package dn

import (
	"encoding/binary"
	"io"
)

// Op is an operation for use with a datanode.
type Op byte

const (
	WRITE_BLOCK Op = iota + 80
	READ_BLOCK
	READ_METADATA
	REPLACE_BLOCK
	COPY_BLOCK
	BLOCK_CHECKSUM
	TRANSFER_BLOCK
	REQUEST_SHORT_CIRCUIT_FDS
	RELEASE_SHORT_CIRCUIT_FDS
	REQUEST_SHORT_CIRCUIT_SHM
	BLOCK_GROUP_CHECKSUM

	CUSTOM Op = 127
)

// DATA_TRANSFER_VERSION is the version of data transfer supported by
// this library.
const DATA_TRANSFER_VERSION = uint16(28)

// WriteVersion writes the current version number to the server.
func WriteVersion(w io.Writer) (err error) {
	version := DATA_TRANSFER_VERSION
	return binary.Write(w, binary.BigEndian, &version)
}

// Write writes op to an io.Writer.
func (op Op) Write(w io.Writer) (err error) {
	return binary.Write(w, binary.BigEndian, &op)
}
