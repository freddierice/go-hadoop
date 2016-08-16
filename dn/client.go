package dn

import (
	"net"

	"github.com/golang/protobuf/proto"

	"gopkg.in/freddierice/go-hadoop.v2"
	hhdfs "gopkg.in/freddierice/go-hproto.v1/hdfs"
)

type Client struct {
	// conn is the underlying connection
	conn net.Conn
}

// Dial attempts to connect to a remote datanode through SASL. On success, a
// Client is returned that can be used to do other operations.
func Dial(host string) (*Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	if err := authenticate(conn); err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

// authenticate attempts to authenticate with another datanode over a conn.
func authenticate(conn net.Conn) error {
	hello := []byte{0xde, 0xad, 0xbe, 0xef}

	status := hhdfs.DataTransferEncryptorMessageProto_SUCCESS
	prot := &hhdfs.DataTransferEncryptorMessageProto{
		Status:  &status,
		Payload: []byte{},
	}

	buf, err := proto.Marshal(prot)
	if err != nil {
		return err
	}

	conn.Write(hello)
	buf, err = hadoop.varintUnpackage(conn)

	return nil
}

// Read implements io.Reader by reading len(buf) bytes from the underlying
// connection.
func (c *Client) Read(buf []byte) (int, error) {
	return c.conn.Read(buf)
}

// Write implements io.Writer by writing len(buf) bytes to the underlying
// connection.
func (c *Client) Write(buf []byte) (int, error) {
	return c.conn.Write(buf)
}
