package util

import (
	"bytes"
	"encoding/binary"
	"io"
)

// WriteVInt writes an integer just as the hadoop library does.
func WriteVInt(w io.Writer, i int) {
	// code follows almost exact logic from the actual implementation. it was
	// not written to look pretty or be particularly readable.
	if i >= -120 && i <= 127 {
		out := int8(i)
		binary.Write(w, binary.BigEndian, &out)
		return
	}

	blen := -112
	if i < 0 {
		i ^= -1
		blen = -120
	}

	tmp := i
	for tmp != 0 {
		tmp = tmp >> 8
		blen--
	}

	out := int8(blen)
	binary.Write(w, binary.BigEndian, &out)

	if blen < -120 {
		blen = -(blen + 120)
	} else {
		blen = -(blen + 112)
	}

	for idx := blen; idx != 0; idx-- {
		shiftbits := (idx - 1) * 8
		mask := 0xFF << uint(shiftbits)
		out := int8((i & mask) >> uint(shiftbits))
		binary.Write(w, binary.BigEndian, &out)
	}
}

// WriteString writes a string to the writer as hadoop does. It assumes that
// the string is UTF-8 encoded. If s is not UTF-8 encoded, hadoop might fail
// unexpectedly.
func WriteString(w io.Writer, s string) {
	if len(s) == 0 {
		out := int32(-1)
		binary.Write(w, binary.BigEndian, &out)
		return
	}

	b := []byte(s)
	blen := int32(len(b))

	binary.Write(w, binary.BigEndian, &blen)
	w.Write(b)
}

// ByteToSignedByte takes ina a byte and returns the signed version. This is to
// keep the compatibility between java/golang symantics for the signedness of a
// byte.
func ByteToSignedByte(v byte) int8 {
	var num int8
	binary.Read(bytes.NewBuffer([]byte{v}), binary.BigEndian, &num)
	return num
}

// DecodeVIntSize is a utility function in hadoop that parses the first byte
// of a vint/vlong to determine the following number of bytes (1 to 9).
func DecodeVIntSize(v int8) int {
	vint := int(v)
	if vint >= -112 {
		return 1
	} else if vint < -120 {
		return -119 - vint
	} else {
		return -111 - vint
	}
}

// IsNegativeVInt determines whether v represents a negative value.
func IsNegativeVInt(v int8) bool {
	return v < -120 || (v >= -112 && v < 0)
}

// ReadString reads a string from the io.Reader as hadoop does. It assumes that
// the string is UTF-8 encoded.
func ReadString(r io.Reader) (string, error) {
	var inlen int32
	binary.Read(r, binary.BigEndian, &inlen)

	if inlen == 0 {
		return "", nil
	}

	b := make([]byte, int(inlen))
	_, err := io.ReadFull(r, b)

	return string(b), err
}

// ReadVInt as it is implemented in hadoop. This code is not meant to be
// readable, but instead mimic the style in the actual code base.
func ReadVInt(r io.Reader) (int, error) {
	buf := make([]byte, 1)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	firstByte := ByteToSignedByte(buf[0])

	blen := DecodeVIntSize(firstByte)
	if blen == 1 {
		return int(firstByte), nil
	}

	i := 0
	for idx := 0; idx < blen-1; idx++ {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		i = i << 8
		i = i | (int(buf[0]) & 0xFF)
	}

	if IsNegativeVInt(firstByte) {
		return i ^ -1, nil
	}

	return i, nil
}
