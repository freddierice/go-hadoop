package dn

import (
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"

	"gopkg.in/freddierice/go-hadoop.v2/util"
	hhdfs "gopkg.in/freddierice/go-hproto.v1/hdfs"
	"gopkg.in/freddierice/go-sasl.v2"
)

type Client struct {
	// conn is the underlying connection
	conn     net.Conn
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
	// create SASL client
	log.Print("creating client")
	saslClient, err := sasl.NewClient(cl.Service, "0", &sasl.Config{
		Authname: cl.Username + " ",
		Username: cl.Username,
		Password: cl.Password,
	})
	if err != nil {
		return err
	}

	// server only uses DIGEST-MD5
	log.Print("selecting DIGEST-MD5")
	mech, response, done, err := saslClient.Start([]string{"DIGEST-MD5"})
	if err != nil {
		return err
	}
	log.Printf("sasl mech: %v", mech)
	log.Printf("sasl done: %v", done)

	// create first response
	log.Print("creating hello")
	hello := []byte{0xde, 0xad, 0xbe, 0xef}
	statusSuccess := hhdfs.DataTransferEncryptorMessageProto_SUCCESS
	prot := &hhdfs.DataTransferEncryptorMessageProto{
		Status:  &statusSuccess,
		Payload: response,
	}

	// send hello and first prot
	log.Print("writing hello")
	cl.Write(hello)
	buf, err := util.VarintPackage(prot)
	if err != nil {
		return err
	}
	log.Printf("writing: %v", string(response))
	cl.Write(buf)

	// unpackage response
	log.Print("unpackaging 1")
	buf, err = util.VarintUnpackage(cl)
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(buf, prot); err != nil {
		return err
	}
	fmt.Println(prot)

	// take a step with the info we got from the server
	for {
		log.Print("sasl step")
		response, done, err := saslClient.Step(prot.Payload)
		if err != nil {
			return err
		}
		/*
			if done {
				break
			}
		*/
		log.Printf("sasl says: %v", done)
		_ = done

		// package up client response
		suite := hhdfs.CipherSuiteProto_AES_CTR_NOPADDING
		prot = &hhdfs.DataTransferEncryptorMessageProto{
			Status:  &statusSuccess,
			Payload: response,
			CipherOption: []*hhdfs.CipherOptionProto{&hhdfs.CipherOptionProto{
				Suite: &suite,
			}}}
		buf, err := util.VarintPackage(prot)
		if err != nil {
			return err
		}
		cl.Write(buf)
		log.Printf("and responds: %v", prot)

		// unpackage server response
		buf, err = util.VarintUnpackage(cl)
		if err != nil {
			return err
		}
		if err := proto.Unmarshal(buf, prot); err != nil {
			return err
		}
		fmt.Println(prot)

	}

	fmt.Println("done!")

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
