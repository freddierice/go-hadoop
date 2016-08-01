package rpc

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

func IntPackageBytes(b []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	num32 := uint32(len(b))

	binary.Write(buf, binary.BigEndian, &num32)
	buf.Write(b)

	return buf.Bytes(), nil
}

func VarintPackageBytes(b []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	bLen := proto.EncodeVarint(uint64(len(b)))

	buf.Write(bLen)
	buf.Write(b)

	return buf.Bytes(), nil
}

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

func VarintPackage(msg proto.Message) ([]byte, error) {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return VarintPackageBytes(msgBytes)
}

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
