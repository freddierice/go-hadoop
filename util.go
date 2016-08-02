package hadoop

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"
)

// NewUUID creates a new UUID for usage with rpc.
func NewUUID() string {
	return string(uuid.NewV4().String()[0:16])
}

// intPackageBytes takes in a byte slice, then packages it as
// -----------------------------------
// |     big endian uint32 len(b)    |
// -----------------------------------
// |        provided bytes (b)       |
// -----------------------------------
func intPackageBytes(b []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	num32 := uint32(len(b))

	binary.Write(buf, binary.BigEndian, &num32)
	buf.Write(b)

	return buf.Bytes(), nil
}

// varintPackageBytes takes in a byteslice then packes it as
// -----------------------------------
// |      varint encoded len(b)      |
// -----------------------------------
// |        provided bytes (b)       |
// -----------------------------------
func varintPackageBytes(b []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	bLen := proto.EncodeVarint(uint64(len(b)))

	buf.Write(bLen)
	buf.Write(b)

	return buf.Bytes(), nil
}

// varintUnpackage takes an io.Reader, reads a varint, then returns that many
// bytes from the io.Reader. An error is returned if the reader cannot read
// a sufficient number of bytes (i.e. if there is malformed data).
func varintUnpackage(r io.Reader) ([]byte, error) {
	packageLen := uint64(0)
	n := 0
	var buf []byte

	// try reading until we know the length
	tmpBuf := make([]byte, 1)
	for n == 0 {
		if _, err := io.ReadFull(r, tmpBuf); err != nil {
			return nil, err
		}

		buf = append(buf, tmpBuf[0])
		packageLen, n = proto.DecodeVarint(buf)
	}

	// read the rest of the backage
	buf = make([]byte, int(packageLen))
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// varintPackage packages a protobuf structure as
// --------------------------------------
// | varint encoded len(serialized msg) |
// --------------------------------------
// |         (serialized msg)           |
// --------------------------------------
func varintPackage(msg proto.Message) ([]byte, error) {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return varintPackageBytes(msgBytes)
}

// rpcPackage takes in a slice of protobufs, and packages them for use with
// the hadoop rpc server.
func rpcPackage(msgs ...proto.Message) ([]byte, error) {
	buf := &bytes.Buffer{}
	for _, msg := range msgs {
		msgBytes, err := varintPackage(msg)
		if err != nil {
			return nil, err
		}

		buf.Write(msgBytes)
	}

	return intPackageBytes(buf.Bytes())
}
