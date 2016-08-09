// Code generated by protoc-gen-go.
// source: encryption.proto
// DO NOT EDIT!

package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateEncryptionZoneRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	KeyName          *string `protobuf:"bytes,2,opt,name=keyName" json:"keyName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateEncryptionZoneRequestProto) Reset()         { *m = CreateEncryptionZoneRequestProto{} }
func (m *CreateEncryptionZoneRequestProto) String() string { return proto.CompactTextString(m) }
func (*CreateEncryptionZoneRequestProto) ProtoMessage()    {}
func (*CreateEncryptionZoneRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0}
}

func (m *CreateEncryptionZoneRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *CreateEncryptionZoneRequestProto) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

type CreateEncryptionZoneResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CreateEncryptionZoneResponseProto) Reset()         { *m = CreateEncryptionZoneResponseProto{} }
func (m *CreateEncryptionZoneResponseProto) String() string { return proto.CompactTextString(m) }
func (*CreateEncryptionZoneResponseProto) ProtoMessage()    {}
func (*CreateEncryptionZoneResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{1}
}

type ListEncryptionZonesRequestProto struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListEncryptionZonesRequestProto) Reset()                    { *m = ListEncryptionZonesRequestProto{} }
func (m *ListEncryptionZonesRequestProto) String() string            { return proto.CompactTextString(m) }
func (*ListEncryptionZonesRequestProto) ProtoMessage()               {}
func (*ListEncryptionZonesRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *ListEncryptionZonesRequestProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type EncryptionZoneProto struct {
	Id                    *int64                      `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Path                  *string                     `protobuf:"bytes,2,req,name=path" json:"path,omitempty"`
	Suite                 *CipherSuiteProto           `protobuf:"varint,3,req,name=suite,enum=hadoop.hdfs.CipherSuiteProto" json:"suite,omitempty"`
	CryptoProtocolVersion *CryptoProtocolVersionProto `protobuf:"varint,4,req,name=cryptoProtocolVersion,enum=hadoop.hdfs.CryptoProtocolVersionProto" json:"cryptoProtocolVersion,omitempty"`
	KeyName               *string                     `protobuf:"bytes,5,req,name=keyName" json:"keyName,omitempty"`
	XXX_unrecognized      []byte                      `json:"-"`
}

func (m *EncryptionZoneProto) Reset()                    { *m = EncryptionZoneProto{} }
func (m *EncryptionZoneProto) String() string            { return proto.CompactTextString(m) }
func (*EncryptionZoneProto) ProtoMessage()               {}
func (*EncryptionZoneProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *EncryptionZoneProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *EncryptionZoneProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *EncryptionZoneProto) GetSuite() CipherSuiteProto {
	if m != nil && m.Suite != nil {
		return *m.Suite
	}
	return CipherSuiteProto_UNKNOWN
}

func (m *EncryptionZoneProto) GetCryptoProtocolVersion() CryptoProtocolVersionProto {
	if m != nil && m.CryptoProtocolVersion != nil {
		return *m.CryptoProtocolVersion
	}
	return CryptoProtocolVersionProto_UNKNOWN_PROTOCOL_VERSION
}

func (m *EncryptionZoneProto) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

type ListEncryptionZonesResponseProto struct {
	Zones            []*EncryptionZoneProto `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
	HasMore          *bool                  `protobuf:"varint,2,req,name=hasMore" json:"hasMore,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListEncryptionZonesResponseProto) Reset()         { *m = ListEncryptionZonesResponseProto{} }
func (m *ListEncryptionZonesResponseProto) String() string { return proto.CompactTextString(m) }
func (*ListEncryptionZonesResponseProto) ProtoMessage()    {}
func (*ListEncryptionZonesResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{4}
}

func (m *ListEncryptionZonesResponseProto) GetZones() []*EncryptionZoneProto {
	if m != nil {
		return m.Zones
	}
	return nil
}

func (m *ListEncryptionZonesResponseProto) GetHasMore() bool {
	if m != nil && m.HasMore != nil {
		return *m.HasMore
	}
	return false
}

type GetEZForPathRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetEZForPathRequestProto) Reset()                    { *m = GetEZForPathRequestProto{} }
func (m *GetEZForPathRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetEZForPathRequestProto) ProtoMessage()               {}
func (*GetEZForPathRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *GetEZForPathRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type GetEZForPathResponseProto struct {
	Zone             *EncryptionZoneProto `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *GetEZForPathResponseProto) Reset()                    { *m = GetEZForPathResponseProto{} }
func (m *GetEZForPathResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetEZForPathResponseProto) ProtoMessage()               {}
func (*GetEZForPathResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *GetEZForPathResponseProto) GetZone() *EncryptionZoneProto {
	if m != nil {
		return m.Zone
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateEncryptionZoneRequestProto)(nil), "hadoop.hdfs.CreateEncryptionZoneRequestProto")
	proto.RegisterType((*CreateEncryptionZoneResponseProto)(nil), "hadoop.hdfs.CreateEncryptionZoneResponseProto")
	proto.RegisterType((*ListEncryptionZonesRequestProto)(nil), "hadoop.hdfs.ListEncryptionZonesRequestProto")
	proto.RegisterType((*EncryptionZoneProto)(nil), "hadoop.hdfs.EncryptionZoneProto")
	proto.RegisterType((*ListEncryptionZonesResponseProto)(nil), "hadoop.hdfs.ListEncryptionZonesResponseProto")
	proto.RegisterType((*GetEZForPathRequestProto)(nil), "hadoop.hdfs.GetEZForPathRequestProto")
	proto.RegisterType((*GetEZForPathResponseProto)(nil), "hadoop.hdfs.GetEZForPathResponseProto")
}

func init() { proto.RegisterFile("encryption.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xe9, 0xba, 0xa1, 0xde, 0xc1, 0x18, 0x91, 0x41, 0x15, 0xc4, 0x5a, 0x11, 0xf7, 0x20,
	0x05, 0xa7, 0xf8, 0xee, 0xc6, 0xf4, 0x45, 0xc7, 0xac, 0xe0, 0xc3, 0xc0, 0x87, 0xd0, 0x5d, 0x6d,
	0x51, 0x9b, 0x98, 0x64, 0x0f, 0xf3, 0xd7, 0xf8, 0xef, 0xfc, 0x1b, 0x26, 0xa9, 0x1b, 0x8d, 0x14,
	0xf1, 0x65, 0xdc, 0xdc, 0x9c, 0x9c, 0x73, 0xf6, 0x15, 0xba, 0x58, 0xa4, 0x62, 0xc9, 0x55, 0xce,
	0x8a, 0x98, 0x0b, 0xa6, 0x18, 0x69, 0x67, 0x74, 0xce, 0x18, 0x8f, 0xb3, 0xf9, 0x93, 0xdc, 0x05,
	0xf3, 0x5b, 0x5e, 0x44, 0x13, 0x08, 0x47, 0x02, 0xa9, 0xc2, 0xf1, 0xfa, 0xc9, 0x8c, 0x15, 0x98,
	0xe0, 0xfb, 0x02, 0xa5, 0x9a, 0xda, 0xc7, 0x5d, 0xf0, 0xa5, 0x48, 0x03, 0x2f, 0x6c, 0xf4, 0xb7,
	0x12, 0x33, 0x92, 0x00, 0x36, 0x5e, 0x70, 0x39, 0xa1, 0x6f, 0x18, 0x34, 0x42, 0x4f, 0x6f, 0x57,
	0xc7, 0xe8, 0x10, 0x0e, 0xea, 0xfd, 0x24, 0x67, 0x85, 0x44, 0x6b, 0x18, 0x9d, 0xc2, 0xfe, 0x4d,
	0x2e, 0x95, 0x2b, 0x91, 0x4e, 0x66, 0x07, 0x1a, 0xf9, 0xdc, 0x46, 0xfa, 0x89, 0x9e, 0xa2, 0x2f,
	0x0f, 0xb6, 0x5d, 0x7d, 0xad, 0x8e, 0x10, 0x68, 0x72, 0xaa, 0x32, 0x5d, 0xcb, 0x94, 0xb5, 0x33,
	0x39, 0x83, 0x96, 0x5c, 0xe4, 0x0a, 0x03, 0x5f, 0x2f, 0x3b, 0x83, 0xbd, 0xb8, 0x02, 0x23, 0x1e,
	0xe5, 0x3c, 0x43, 0x71, 0x6f, 0xee, 0xad, 0x63, 0x52, 0x6a, 0xc9, 0x23, 0xf4, 0x6c, 0x1a, 0xb3,
	0xdb, 0x94, 0xbd, 0x3e, 0xa0, 0x90, 0x3a, 0x3a, 0x68, 0x5a, 0x93, 0x63, 0xd7, 0xa4, 0x4e, 0x59,
	0xda, 0xd5, 0xbb, 0x54, 0x09, 0xb6, 0x6c, 0xd5, 0x35, 0x41, 0x05, 0x61, 0x2d, 0x9c, 0x0a, 0x40,
	0x72, 0x01, 0xad, 0x0f, 0xb3, 0xd5, 0x7f, 0xdc, 0xef, 0xb7, 0x07, 0xa1, 0x53, 0xa6, 0x06, 0x53,
	0x52, 0xca, 0x4d, 0x6a, 0x46, 0xe5, 0x2d, 0x13, 0x68, 0x01, 0x6d, 0x26, 0xab, 0x63, 0x74, 0x02,
	0xc1, 0x35, 0xaa, 0xf1, 0xec, 0x8a, 0x89, 0xa9, 0x66, 0xf6, 0xf7, 0xf7, 0x8f, 0xee, 0x60, 0xc7,
	0x55, 0x57, 0xcb, 0x9d, 0x43, 0xd3, 0xa4, 0x69, 0xbd, 0xf7, 0xaf, 0x6e, 0x56, 0x3d, 0xbc, 0x84,
	0x23, 0x26, 0x9e, 0x63, 0xca, 0x69, 0x9a, 0xa1, 0xf3, 0x86, 0xff, 0xa0, 0x2b, 0x87, 0x61, 0xef,
	0x17, 0x19, 0x6b, 0x22, 0x3f, 0x3d, 0xef, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xb2, 0xc6, 0x30,
	0xf7, 0x02, 0x00, 0x00,
}
