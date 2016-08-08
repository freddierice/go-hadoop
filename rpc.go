package hadoop

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
	"gopkg.in/freddierice/go-hadoop.v1/hproto"
	"gopkg.in/freddierice/go-sasl.v1"
)

// Auth is the type of authentication that should be used over RPC.
type Auth uint8

const (
	// AuthNone uses no authentication (also refered to as SIMPLE
	// authentication).
	AuthNone Auth = 0x0

	// AuthSasl uses SASL (and an underlying mechanism) to authenticate
	// with the host.
	AuthSasl Auth = 0xDF
)

var (
	// AuthenticationMethods is the set of authentication methods that the
	// client will accept.
	AllowedMechanisms = map[string]bool{"GSSAPI": true}
)

// Conn represents an rpc connection.
type Conn struct {
	// net.Conn is the underlying TCP connection.
	net.Conn

	// ClientId is the sequence of bytes that are used to uniquely identify this
	// client from others that may be connecting to the server.
	ClientId []byte

	// CallId is an id that identifies the count of the next . Before the
	// connection is established, it determines whether SASL or SIMPLE
	// authentication is in progress.
	CallId int

	// Context is the namespace where the functions (to be called) lie. The
	// context cannot change without creating a new connection.
	Context string

	// Username is the name the user wants to authenticate as.
	Username string

	// Hostname is the network location of the server (as a FQDN).
	Hostname string

	// Service is a field for use with Kerberos.
	Service string
}

// Dial initializes an RPC connection for a user within a context to a
// remote host. Service is an optional argument if kerberos is not used.
func Dial(username, host, context, service string, auth Auth) (rc *Conn,
	err error) {

	// hadoop has great protocols.
	rc = &Conn{
		ClientId: []byte(NewUUID()),
		CallId:   -3,
		Context:  context,
		Username: username,
		Hostname: strings.Split(host, ":")[0],
		Service:  service,
	}

	if rc.Conn, err = net.Dial("tcp", host); err != nil {
		return nil, err
	}

	if err := rc.writeHeader(auth); err != nil {
		return nil, err
	}

	if err := rc.authenticate(auth); err != nil {
		return nil, err
	}

	// conn write the RPC header and the context
	ipc := rc.newIpcConnectionContextProto()
	header := rc.newRpcRequestHeaderProto()
	buf, err := rpcPackage(header, ipc)
	if err != nil {
		return nil, err
	}
	if _, err := rc.Write(buf); err != nil {
		return nil, err
	}

	return rc, err
}

// authenticate authenticates a user over an rpc connection, through a
// particular authentication method.
func (rc *Conn) authenticate(auth Auth) (err error) {

	switch auth {
	case AuthNone:
		return nil
	case AuthSasl:
		return rc.authenticateSasl()
	default:
		return fmt.Errorf("unknown authentication method: %v", auth)
	}
}

// authenticateSasl authenticates a user through SASL/Kerberos. If the server
// does not allow for the GSSAPI mechanism, then the function will fail.
func (rc *Conn) authenticateSasl() (err error) {

	// create a new sasl client
	saslClient, err := sasl.NewClient(rc.Service, rc.Hostname, &sasl.Config{
		Username: rc.Username,
	})
	if err != nil {
		return err
	}

	// 1. Initialize negotiation with the rpc server
	state := hproto.RpcSaslProto_NEGOTIATE
	negotiate := &hproto.RpcSaslProto{
		State: &state,
	}
	saslRes, err := rc.sendSasl(negotiate)
	if err != nil {
		return err
	}

	for {
		switch *saslRes.State {
		case hproto.RpcSaslProto_NEGOTIATE:
			// build a mechanism list that work for both the client
			// and the server.
			mechlist := make([]string, len(saslRes.Auths))
			i := 0
			for _, auth := range saslRes.Auths {
				if _, ok := AllowedMechanisms[*auth.Mechanism]; ok {
					mechlist[i] = *auth.Mechanism
					i++
				}
			}
			if i == 0 {
				return fmt.Errorf("client and server could not agree on a " +
					"mechanism")
			}
			mechlist = mechlist[0:i]

			// make the first step with an agreed upon mechanism list.
			mech, resp, err := saslClient.Start(mechlist)
			if err != nil {
				return fmt.Errorf("could not negotiate sasl: %v", err)
			}

			var chosenAuth *hproto.RpcSaslProto_SaslAuth
			for _, serverAuth := range saslRes.Auths {
				if *serverAuth.Mechanism == mech {
					chosenAuth = serverAuth
					break
				}
			}

			// 2. initiate the challenge/response with a chosen algorithm.
			state = hproto.RpcSaslProto_INITIATE
			initiate := &hproto.RpcSaslProto{
				State: &state,
				Token: []byte(resp),
				Auths: []*hproto.RpcSaslProto_SaslAuth{
					chosenAuth,
				},
			}
			saslRes, err = rc.sendSasl(initiate)
			if err != nil {
				return err
			}
		case hproto.RpcSaslProto_CHALLENGE:
			// 3. do the challenge to prove that we are who we say we are.
			token, err := saslClient.Step(string(saslRes.Token))
			if err != nil {
				return err
			}

			// 4. construct a response to send to the server
			state = hproto.RpcSaslProto_RESPONSE
			response := &hproto.RpcSaslProto{
				State: &state,
				Token: []byte(token),
			}
			if saslRes, err = rc.sendSasl(response); err != nil {
				return err
			}
		case hproto.RpcSaslProto_SUCCESS:
			return nil
		default:
			return fmt.Errorf("recieved unexpected saslRes.State from server: %v\n", *saslRes.State)
		}
	}

	return nil
}

// sendSasl ships a sasl step to the server and returns the server's response.
func (rc *Conn) sendSasl(msg *hproto.RpcSaslProto) (*hproto.RpcSaslProto,
	error) {
	resSasl := &hproto.RpcSaslProto{}

	header := rc.newSaslRpcRequestHeaderProto()
	buf, err := rpcPackage(header, msg)
	if err != nil {
		return nil, err
	}
	if _, err := rc.Write(buf); err != nil {
		return nil, err
	}
	if _, err := rc.recv(resSasl); err != nil {
		return nil, err
	}

	return resSasl, nil
}

// writeHeader writes the rpc header to the connection.
func (rc *Conn) writeHeader(auth Auth) error {

	// header
	if _, err := rc.Conn.Write([]byte("hrpc")); err != nil {
		return err
	}

	// version
	num8 := uint8(9)
	if err := binary.Write(rc.Conn, binary.BigEndian, &num8); err != nil {
		return err
	}

	// RPC service class
	num8 = 0
	if err := binary.Write(rc.Conn, binary.BigEndian, &num8); err != nil {
		return err
	}

	// authentication type
	num8 = uint8(auth)
	if err := binary.Write(rc.Conn, binary.BigEndian, &num8); err != nil {
		return err
	}

	return nil
}

// Call makes a request to run rfunc with input req and output resp. Req and
// resp must be the right type, or this function will have undefined behavior.
func (rc *Conn) Call(rfunc string, req, resp proto.Message) error {
	if err := rc.send(rfunc, req); err != nil {
		return err
	}

	_, err := rc.recv(resp)
	return err
}

// recv recieves a message from the server in response to a send. fill must be
// the right type, or this function will have undefined behavior.
func (rc *Conn) recv(fill proto.Message) (*hproto.RpcResponseHeaderProto, error) {
	var recvLen uint32
	binary.Read(rc, binary.BigEndian, &recvLen)

	allRecv := make([]byte, int(recvLen))
	_, err := io.ReadFull(rc, allRecv)
	if err != nil {
		return nil, err
	}

	allReader := bytes.NewReader(allRecv)
	headerBytes, err := varintUnpackage(allReader)
	if err != nil {
		return nil, err
	}

	resp := &hproto.RpcResponseHeaderProto{}
	if err := proto.Unmarshal(headerBytes, resp); err != nil {
		return nil, err
	}

	if resp.GetStatus() != hproto.RpcResponseHeaderProto_SUCCESS {
		log.Print(resp)
		return resp, errors.New("error response from hadoop")
	}

	messageBytes, err := varintUnpackage(allReader)
	if err != nil {
		return resp, err
	}

	if err := proto.Unmarshal(messageBytes, fill); err != nil {
		return resp, err
	}

	return resp, nil
}

// send makes a request to run method with req.
func (rc *Conn) send(method string, req proto.Message) error {
	rpcRequestHeader := rc.newRpcRequestHeaderProto()
	requestHeader := rc.newRequestHeaderProto(method)

	buf, err := rpcPackage(rpcRequestHeader, requestHeader, req)
	if err != nil {
		return err
	}

	_, err = rc.Conn.Write(buf)
	return err
}

// Read implements io.Reader by forwarding Read requests to the underlying
// connection.
func (rc *Conn) Read(b []byte) (int, error) {
	return rc.Conn.Read(b)
}

// Write implements io.Writer by forwarding Write requests to the underlying
// connection.
func (rc *Conn) Write(b []byte) (int, error) {
	return rc.Conn.Write(b)
}

// incrementCallId increments the call id for every request. The CallId starts
// at -3 for the initial connection, 0 for the first call, then increments by
// one for each call thereafter.
func (rc *Conn) incrementCallId() {
	if rc.CallId < 0 {
		rc.CallId = 0
	} else {
		rc.CallId = rc.CallId + 1
	}
}

func (rc *Conn) newRequestHeaderProto(method string) *hproto.RequestHeaderProto {
	declaringClassProtocolName := rc.Context
	clientProtocolVersion := uint64(1)
	return &hproto.RequestHeaderProto{
		MethodName:                 &method,
		DeclaringClassProtocolName: &declaringClassProtocolName,
		ClientProtocolVersion:      &clientProtocolVersion,
	}
}

func (rc *Conn) newSaslRpcRequestHeaderProto() *hproto.RpcRequestHeaderProto {
	callId := int32(-33) // actually part of the spec...

	rpcKind := hproto.RpcKindProto_RPC_PROTOCOL_BUFFER
	rpcOp := hproto.RpcRequestHeaderProto_RPC_FINAL_PACKET
	return &hproto.RpcRequestHeaderProto{
		CallId:   &callId,
		ClientId: rc.ClientId,
		RpcKind:  &rpcKind,
		RpcOp:    &rpcOp,
	}
}

func (rc *Conn) newRpcRequestHeaderProto() *hproto.RpcRequestHeaderProto {
	callId := int32(rc.CallId)
	rc.incrementCallId()

	rpcKind := hproto.RpcKindProto_RPC_PROTOCOL_BUFFER
	rpcOp := hproto.RpcRequestHeaderProto_RPC_FINAL_PACKET
	return &hproto.RpcRequestHeaderProto{
		CallId:   &callId,
		ClientId: rc.ClientId,
		RpcKind:  &rpcKind,
		RpcOp:    &rpcOp,
	}
}

func (rc *Conn) newIpcConnectionContextProto() *hproto.IpcConnectionContextProto {
	effectiveUser := rc.Username
	protocolName := rc.Context
	//realUser := "user"
	return &hproto.IpcConnectionContextProto{
		UserInfo: &hproto.UserInformationProto{
			//RealUser:      &realUser,
			RealUser: &effectiveUser,
		},
		Protocol: &protocolName,
	}
}
