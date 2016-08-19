package util

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"

	"github.com/golang/protobuf/proto"
)

// IntPackageBytes takes in a byte slice, then packages it as
// -----------------------------------
// |     big endian uint32 len(b)    |
// -----------------------------------
// |        provided bytes (b)       |
// -----------------------------------
func IntPackageBytes(b []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	num32 := uint32(len(b))

	binary.Write(buf, binary.BigEndian, &num32)
	buf.Write(b)

	return buf.Bytes(), nil
}

// IntUnpackageBytes takes in an io.Reader, then returns b on input:
// -----------------------------------
// |     big endian uint32 len(b)    |
// -----------------------------------
// |        provided bytes (b)       |
// -----------------------------------
func IntUnpackageBytes(r io.Reader) ([]byte, error) {
	var num32 uint32
	if err := binary.Read(r, binary.BigEndian, &num32); err != nil {
		return nil, err
	}

	log.Print(num32)
	buf := make([]byte, int(num32))
	_, err := io.ReadFull(r, buf)
	return buf, err
}

// ShortUnpackageBytes takes in an io.Reader, then returns b on input:
// -----------------------------------
// |     big endian uint16 len(b)    |
// -----------------------------------
// |        provided bytes (b)       |
// -----------------------------------
func ShortUnpackageBytes(r io.Reader) ([]byte, error) {
	var num16 uint16
	if err := binary.Read(r, binary.BigEndian, &num16); err != nil {
		return nil, err
	}

	buf := make([]byte, int(num16))
	_, err := io.ReadFull(r, buf)
	return buf, err
}

// VarintPackageBytes takes in a byteslice then packes it as
// -----------------------------------
// |      varint encoded len(b)      |
// -----------------------------------
// |        provided bytes (b)       |
// -----------------------------------
func VarintPackageBytes(b []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	bLen := proto.EncodeVarint(uint64(len(b)))

	buf.Write(bLen)
	buf.Write(b)

	return buf.Bytes(), nil
}

// VarintUnpackage takes an io.Reader, reads a varint, then returns that many
// bytes from the io.Reader. An error is returned if the reader cannot read
// a sufficient number of bytes (i.e. if there is malformed data).
func VarintUnpackage(r io.Reader) ([]byte, error) {
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

// VarintPackage packages a protobuf structure as
// --------------------------------------
// | varint encoded len(serialized msg) |
// --------------------------------------
// |         (serialized msg)           |
// --------------------------------------
func VarintPackage(msg proto.Message) ([]byte, error) {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return VarintPackageBytes(msgBytes)
}

// RpcPackage takes in a slice of protobufs, and packages them for use with
// the hadoop rpc server.
func RpcPackage(msgs ...proto.Message) ([]byte, error) {
	buf := &bytes.Buffer{}
	for _, msg := range msgs {
		msgBytes, err := VarintPackage(msg)
		if err != nil {
			return nil, err
		}

		buf.Write(msgBytes)
	}

	return IntPackageBytes(buf.Bytes())
}
