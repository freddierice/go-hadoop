// Code generated by protoc-gen-go.
// source: IpcConnectionContext.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	IpcConnectionContext.proto

It has these top-level messages:
	UserInformationProto
	IpcConnectionContextProto
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Spec for UserInformationProto is specified in ProtoUtil#makeIpcConnectionContext
type UserInformationProto struct {
	EffectiveUser    *string `protobuf:"bytes,1,opt,name=effectiveUser" json:"effectiveUser,omitempty"`
	RealUser         *string `protobuf:"bytes,2,opt,name=realUser" json:"realUser,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserInformationProto) Reset()                    { *m = UserInformationProto{} }
func (m *UserInformationProto) String() string            { return proto.CompactTextString(m) }
func (*UserInformationProto) ProtoMessage()               {}
func (*UserInformationProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInformationProto) GetEffectiveUser() string {
	if m != nil && m.EffectiveUser != nil {
		return *m.EffectiveUser
	}
	return ""
}

func (m *UserInformationProto) GetRealUser() string {
	if m != nil && m.RealUser != nil {
		return *m.RealUser
	}
	return ""
}

// *
// The connection context is sent as part of the connection establishment.
// It establishes the context for ALL Rpc calls within the connection.
type IpcConnectionContextProto struct {
	// UserInfo beyond what is determined as part of security handshake
	// at connection time (kerberos, tokens etc).
	UserInfo *UserInformationProto `protobuf:"bytes,2,opt,name=userInfo" json:"userInfo,omitempty"`
	// Protocol name for next rpc layer.
	// The client created a proxy with this protocol name
	Protocol         *string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IpcConnectionContextProto) Reset()                    { *m = IpcConnectionContextProto{} }
func (m *IpcConnectionContextProto) String() string            { return proto.CompactTextString(m) }
func (*IpcConnectionContextProto) ProtoMessage()               {}
func (*IpcConnectionContextProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IpcConnectionContextProto) GetUserInfo() *UserInformationProto {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *IpcConnectionContextProto) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInformationProto)(nil), "hadoop.common.UserInformationProto")
	proto.RegisterType((*IpcConnectionContextProto)(nil), "hadoop.common.IpcConnectionContextProto")
}

func init() { proto.RegisterFile("IpcConnectionContext.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xf2, 0x2c, 0x48, 0x76,
	0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x03, 0xb2, 0x4a, 0x52, 0x2b, 0x4a, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x33, 0x12, 0x53, 0xf2, 0xf3, 0x0b, 0xf4, 0x92, 0xf3,
	0x73, 0x73, 0xf3, 0xf3, 0x94, 0x22, 0xb8, 0x44, 0x42, 0x8b, 0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2,
	0x8b, 0x72, 0x13, 0x41, 0xca, 0x03, 0xc0, 0xca, 0x54, 0xb8, 0x78, 0x53, 0xd3, 0xd2, 0x40, 0x06,
	0x94, 0xa5, 0x82, 0x14, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x0a, 0x49, 0x71,
	0x71, 0x14, 0xa5, 0x26, 0xe6, 0x80, 0x15, 0x30, 0x81, 0x15, 0xc0, 0xf9, 0x4a, 0x15, 0x5c, 0x92,
	0xd8, 0x9c, 0x01, 0x31, 0xde, 0x9e, 0x8b, 0xa3, 0x14, 0x6a, 0x2d, 0x58, 0x23, 0xb7, 0x91, 0xb2,
	0x1e, 0x8a, 0xc3, 0xf4, 0xb0, 0xb9, 0x2a, 0x08, 0xae, 0x09, 0x64, 0x33, 0xd8, 0x3f, 0xc9, 0xf9,
	0x39, 0x12, 0xcc, 0x10, 0x9b, 0x61, 0x7c, 0x27, 0x7b, 0x2e, 0xb9, 0xfc, 0xa2, 0x74, 0xbd, 0xc4,
	0x82, 0xc4, 0xe4, 0x8c, 0x54, 0x98, 0xb1, 0x99, 0x05, 0xc9, 0x90, 0x10, 0x48, 0x2a, 0x4d, 0x73,
	0x92, 0xc2, 0xe9, 0xb2, 0xe2, 0x05, 0x8c, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x05,
	0xde, 0x49, 0x40, 0x01, 0x00, 0x00,
}
