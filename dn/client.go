package dn

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"net"

	"github.com/golang/protobuf/proto"

	"gopkg.in/freddierice/go-hadoop.v2/util"
	hcommon "gopkg.in/freddierice/go-hproto.v1/common"
	hhdfs "gopkg.in/freddierice/go-hproto.v1/hdfs"
	"gopkg.in/freddierice/go-sasl.v2"
)

type Client struct {
	// conn is the underlying connection
	conn net.Conn

	// encOn lets the library know if the connection is encrypted over the
	// stream reader/writers.
	encOn     bool
	encReader io.Reader
	encWriter io.Writer

	Username string
	Hostname string
	Service  string
	Password string
}

// Dial attempts to connect to a remote datanode through SASL. On success, a
// Client is returned that can be used to do other operations.
func Dial(token, password, host string) (*Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	cl := &Client{
		conn:     conn,
		Username: token,
		Service:  "hdfs",
		Password: password,
	}

	if err := cl.Authenticate(); err != nil {
		return nil, err
	}

	return cl, nil
}

// authenticate attempts to authenticate with another datanode over a conn.
func (cl *Client) Authenticate() error {
	// send hello to the server
	hello := []byte{0xde, 0xad, 0xbe, 0xef}
	if _, err := cl.Write(hello); err != nil {
		return err
	}

	// create SASL client
	saslClient, err := sasl.NewClient(cl.Service, "0", &sasl.Config{
		Authname: cl.Username + " ",
		Username: cl.Username,
		Password: cl.Password,
		MaxSsf:   65535,
		MinSsf:   1,
	})
	if err != nil {
		return err
	}

	// set the mechanism. server only uses DIGEST-MD5.
	_, response, _, err := saslClient.Start([]string{"DIGEST-MD5"})
	if err != nil {
		return err
	}

	statusSuccess := hhdfs.DataTransferEncryptorMessageProto_SUCCESS
	prot := &hhdfs.DataTransferEncryptorMessageProto{
		Status:  &statusSuccess,
		Payload: response,
	}

	buf, err := util.VarintPackage(prot)
	if err != nil {
		return err
	}
	if _, err := cl.Write(buf); err != nil {
		return err
	}

	// unpackage server response
	buf, err = util.VarintUnpackage(cl)
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(buf, prot); err != nil {
		return err
	}

	// take a step with the info we got from the server
	response, _, err = saslClient.Step(prot.Payload)
	if err != nil {
		return err
	}

	// package up client response along with a choice to aes encrypt
	suite := hhdfs.CipherSuiteProto_AES_CTR_NOPADDING
	prot = &hhdfs.DataTransferEncryptorMessageProto{
		Status:  &statusSuccess,
		Payload: response,
		CipherOption: []*hhdfs.CipherOptionProto{&hhdfs.CipherOptionProto{
			Suite: &suite,
		}}}
	buf, err = util.VarintPackage(prot)
	if err != nil {
		return err
	}
	if _, err := cl.Write(buf); err != nil {
		return err
	}

	// unpackage server response
	buf, err = util.VarintUnpackage(cl)
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(buf, prot); err != nil {
		return err
	}

	// the last client step. after this, the handshake is done.
	response, _, err = saslClient.Step(prot.Payload)
	if err != nil {
		return err
	}

	cipherOption := prot.CipherOption[0]
	inKeyEncoded, err := util.IntPackageBytes(cipherOption.InKey)
	if err != nil {
		return err
	}
	outKeyEncoded, err := util.IntPackageBytes(cipherOption.OutKey)
	if err != nil {
		return err
	}

	inKey, err := saslClient.Decode(inKeyEncoded)
	if err != nil {
		return err
	}

	outKey, err := saslClient.Decode(outKeyEncoded)
	if err != nil {
		return err
	}

	blockIn, err := aes.NewCipher(inKey)
	if err != nil {
		panic(err)
	}
	blockOut, err := aes.NewCipher(outKey)
	if err != nil {
		panic(err)
	}

	// use out key for reading and in key for writing because
	// we are the client.
	cl.encReader = &cipher.StreamReader{
		S: cipher.NewCTR(blockOut, cipherOption.OutIv),
		R: cl.conn,
	}

	cl.encWriter = &cipher.StreamWriter{
		S: cipher.NewCTR(blockIn, cipherOption.InIv),
		W: cl.conn,
	}

	cl.encOn = true

	return nil
}

// Read implements io.Reader by reading len(buf) bytes from the underlying
// connection.
func (cl *Client) Read(buf []byte) (int, error) {
	if cl.encOn {
		return cl.encReader.Read(buf)
	}
	return cl.conn.Read(buf)
}

// Write implements io.Writer by writing len(buf) bytes to the underlying
// connection.
func (cl *Client) Write(buf []byte) (int, error) {
	if cl.encOn {
		return cl.encWriter.Write(buf)
	}
	return cl.conn.Write(buf)
}

// ReadBlock reads a block from the datanode.
func (cl *Client) ReadBlock(blockToken *hcommon.TokenProto,
	block *hhdfs.ExtendedBlockProto) (dat []byte, err error) {

	if err = WriteVersion(cl); err != nil {
		return nil, err
	}
	if err = READ_BLOCK.Write(cl); err != nil {
		return nil, err
	}

	sendChecksums := false
	offset := uint64(0)
	readLen := uint64(10)
	clientName := "myName"
	readOp := &hhdfs.OpReadBlockProto{
		Header: &hhdfs.ClientOperationHeaderProto{
			ClientName: &clientName,
			BaseHeader: &hhdfs.BaseHeaderProto{
				Block: block,
				Token: blockToken,
			},
		},
		Offset:        &offset,
		Len:           &readLen,
		SendChecksums: &sendChecksums,
	}

	buf, err := util.VarintPackage(readOp)
	if err != nil {
		return nil, err
	}
	if _, err := cl.Write(buf); err != nil {
		return nil, err
	}

	resp := &hhdfs.BlockOpResponseProto{}
	buf, err = util.VarintUnpackage(cl)
	if err != nil {
		return nil, err
	}
	if err = proto.Unmarshal(buf, resp); err != nil {
		return nil, err
	}

	// read all the packets
	bbuf := &bytes.Buffer{}
	for last, pkt := false, []byte(nil); !last && err == nil; pkt, last,
		err = ReadPacket(cl) {
		bbuf.Write(pkt)
	}

	return bbuf.Bytes(), err
}

func (cl *Client) WriteBlock(blockToken *hcommon.TokenProto) {

}
