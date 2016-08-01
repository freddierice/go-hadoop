package rpc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"

	. "github.com/freddierice/go-hadoop/proto"
	"github.com/golang/protobuf/proto"
)

// Conn represents an rpc connection.
type Conn struct {
	net.Conn
	ClientId []byte
	CallId   int
	Context  string
	User     string
}

// Dial initializes an RPC connection for a user within a context to a
// remote host.
func Dial(host, context, user string, sasl bool) (rc *Conn, err error) {

	// hadoop has great protocols.
	callId := -3
	if sasl {
		callId = -33
	}

	rc = &Conn{
		ClientId: []byte(NewUUID()),
		CallId:   callId,
		Context:  context,
		User:     user,
	}

	// connect to
	if rc.Conn, err = net.Dial("tcp", host); err != nil {
		return nil, err
	}

	// conn write the Hadoop header
	if _, err := rc.Conn.Write([]byte("hrpc")); err != nil {
		return nil, err
	}
	num8 := uint8(9)
	if err := binary.Write(rc.Conn, binary.BigEndian, &num8); err != nil {
		return nil, err
	}
	num8 = 0
	if err := binary.Write(rc.Conn, binary.BigEndian, &num8); err != nil {
		return nil, err
	}
	num8 = 0
	if err := binary.Write(rc.Conn, binary.BigEndian, &num8); err != nil {
		return nil, err
	}

	// conn write the RPC header and the context
	ipc := rc.newIpcConnectionContextProto()
	header := rc.newRpcRequestHeaderProto()
	buf, err := RpcPackage(header, ipc)
	if err != nil {
		return nil, err
	}
	if _, err := rc.Write(buf); err != nil {
		return nil, err
	}

	return rc, nil
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
func (rc *Conn) recv(fill proto.Message) (*RpcResponseHeaderProto, error) {
	var recvLen uint32
	binary.Read(rc, binary.BigEndian, &recvLen)

	allRecv := make([]byte, int(recvLen))
	_, err := io.ReadFull(rc, allRecv)
	if err != nil {
		return nil, err
	}

	allReader := bytes.NewReader(allRecv)
	headerBytes, err := VarintUnpackage(allReader)
	if err != nil {
		return nil, err
	}

	resp := &RpcResponseHeaderProto{}
	if err := proto.Unmarshal(headerBytes, resp); err != nil {
		return nil, err
	}

	if resp.GetStatus() != RpcResponseHeaderProto_SUCCESS {
		log.Printf("RpcResponseHeaderProto: %v", resp)
		return nil, errors.New("error response from hadoop")
	}

	messageBytes, err := VarintUnpackage(allReader)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(messageBytes, fill); err != nil {
		return nil, err
	}

	return resp, nil
}

// send makes a request to run method with req.
func (rc *Conn) send(method string, req proto.Message) error {
	rpcRequestHeader := rc.newRpcRequestHeaderProto()
	requestHeader := rc.newRequestHeaderProto(method)

	buf, err := RpcPackage(rpcRequestHeader, requestHeader, req)
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

func (rc *Conn) newRequestHeaderProto(method string) *RequestHeaderProto {
	declaringClassProtocolName := rc.Context
	clientProtocolVersion := uint64(1)
	return &RequestHeaderProto{
		MethodName:                 &method,
		DeclaringClassProtocolName: &declaringClassProtocolName,
		ClientProtocolVersion:      &clientProtocolVersion,
	}
}

func (rc *Conn) newRpcRequestHeaderProto() *RpcRequestHeaderProto {
	callId := int32(rc.CallId)
	rc.incrementCallId()

	rpcKind := RpcKindProto_RPC_PROTOCOL_BUFFER
	rpcOp := RpcRequestHeaderProto_RPC_FINAL_PACKET
	return &RpcRequestHeaderProto{
		CallId:   &callId,
		ClientId: rc.ClientId,
		RpcKind:  &rpcKind,
		RpcOp:    &rpcOp,
	}
}

func (rc *Conn) newIpcConnectionContextProto() *IpcConnectionContextProto {
	effectiveUser := rc.User
	protocolName := rc.Context
	//realUser := "user"
	return &IpcConnectionContextProto{
		UserInfo: &UserInformationProto{
			EffectiveUser: &effectiveUser,
			//RealUser:      &realUser,
		},
		Protocol: &protocolName,
	}
}
