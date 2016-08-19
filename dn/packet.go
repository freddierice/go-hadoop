package dn

import (
	"encoding/binary"
	"io"

	"github.com/golang/protobuf/proto"
	"gopkg.in/freddierice/go-hadoop.v2/util"
	hhdfs "gopkg.in/freddierice/go-hproto.v1/hdfs"
)

func ReadPacket(r io.Reader) (pkt []byte, last bool, err error) {

	var pktLen uint32
	if err := binary.Read(r, binary.BigEndian, &pktLen); err != nil {
		return nil, false, err
	}
	pktLen -= 4

	packetHeader := &hhdfs.PacketHeaderProto{}
	buf, err := util.ShortUnpackageBytes(r)
	if err != nil {
		return nil, false, err
	}
	if err := proto.Unmarshal(buf, packetHeader); err != nil {
		return nil, false, err
	}
	if packetHeader.GetLastPacketInBlock() {
		return nil, true, nil
	}

	dataLen := uint32(packetHeader.GetDataLen())
	pkt = make([]byte, int(pktLen))

	if _, err := io.ReadFull(r, pkt); err != nil {
		return nil, false, err
	}

	// there is also a checksum
	if dataLen != pktLen {
		pkt = pkt[4:]
	}

	return pkt, false, nil
}
