// Code generated by protoc-gen-go.
// source: ZKFCProtocol.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	ZKFCProtocol.proto

It has these top-level messages:
	CedeActiveRequestProto
	CedeActiveResponseProto
	GracefulFailoverRequestProto
	GracefulFailoverResponseProto
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CedeActiveRequestProto struct {
	MillisToCede     *uint32 `protobuf:"varint,1,req,name=millisToCede" json:"millisToCede,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CedeActiveRequestProto) Reset()                    { *m = CedeActiveRequestProto{} }
func (m *CedeActiveRequestProto) String() string            { return proto.CompactTextString(m) }
func (*CedeActiveRequestProto) ProtoMessage()               {}
func (*CedeActiveRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CedeActiveRequestProto) GetMillisToCede() uint32 {
	if m != nil && m.MillisToCede != nil {
		return *m.MillisToCede
	}
	return 0
}

type CedeActiveResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CedeActiveResponseProto) Reset()                    { *m = CedeActiveResponseProto{} }
func (m *CedeActiveResponseProto) String() string            { return proto.CompactTextString(m) }
func (*CedeActiveResponseProto) ProtoMessage()               {}
func (*CedeActiveResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GracefulFailoverRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GracefulFailoverRequestProto) Reset()                    { *m = GracefulFailoverRequestProto{} }
func (m *GracefulFailoverRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GracefulFailoverRequestProto) ProtoMessage()               {}
func (*GracefulFailoverRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GracefulFailoverResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GracefulFailoverResponseProto) Reset()                    { *m = GracefulFailoverResponseProto{} }
func (m *GracefulFailoverResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GracefulFailoverResponseProto) ProtoMessage()               {}
func (*GracefulFailoverResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*CedeActiveRequestProto)(nil), "hadoop.common.CedeActiveRequestProto")
	proto.RegisterType((*CedeActiveResponseProto)(nil), "hadoop.common.CedeActiveResponseProto")
	proto.RegisterType((*GracefulFailoverRequestProto)(nil), "hadoop.common.GracefulFailoverRequestProto")
	proto.RegisterType((*GracefulFailoverResponseProto)(nil), "hadoop.common.GracefulFailoverResponseProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ZKFCProtocolService service

type ZKFCProtocolServiceClient interface {
	// *
	// Request that the service cede its active state, and quit the election
	// for some amount of time
	CedeActive(ctx context.Context, in *CedeActiveRequestProto, opts ...grpc.CallOption) (*CedeActiveResponseProto, error)
	GracefulFailover(ctx context.Context, in *GracefulFailoverRequestProto, opts ...grpc.CallOption) (*GracefulFailoverResponseProto, error)
}

type zKFCProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewZKFCProtocolServiceClient(cc *grpc.ClientConn) ZKFCProtocolServiceClient {
	return &zKFCProtocolServiceClient{cc}
}

func (c *zKFCProtocolServiceClient) CedeActive(ctx context.Context, in *CedeActiveRequestProto, opts ...grpc.CallOption) (*CedeActiveResponseProto, error) {
	out := new(CedeActiveResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.common.ZKFCProtocolService/cedeActive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zKFCProtocolServiceClient) GracefulFailover(ctx context.Context, in *GracefulFailoverRequestProto, opts ...grpc.CallOption) (*GracefulFailoverResponseProto, error) {
	out := new(GracefulFailoverResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.common.ZKFCProtocolService/gracefulFailover", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ZKFCProtocolService service

type ZKFCProtocolServiceServer interface {
	// *
	// Request that the service cede its active state, and quit the election
	// for some amount of time
	CedeActive(context.Context, *CedeActiveRequestProto) (*CedeActiveResponseProto, error)
	GracefulFailover(context.Context, *GracefulFailoverRequestProto) (*GracefulFailoverResponseProto, error)
}

func RegisterZKFCProtocolServiceServer(s *grpc.Server, srv ZKFCProtocolServiceServer) {
	s.RegisterService(&_ZKFCProtocolService_serviceDesc, srv)
}

func _ZKFCProtocolService_CedeActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CedeActiveRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKFCProtocolServiceServer).CedeActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.common.ZKFCProtocolService/CedeActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKFCProtocolServiceServer).CedeActive(ctx, req.(*CedeActiveRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZKFCProtocolService_GracefulFailover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GracefulFailoverRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKFCProtocolServiceServer).GracefulFailover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.common.ZKFCProtocolService/GracefulFailover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKFCProtocolServiceServer).GracefulFailover(ctx, req.(*GracefulFailoverRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _ZKFCProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hadoop.common.ZKFCProtocolService",
	HandlerType: (*ZKFCProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "cedeActive",
			Handler:    _ZKFCProtocolService_CedeActive_Handler,
		},
		{
			MethodName: "gracefulFailover",
			Handler:    _ZKFCProtocolService_GracefulFailover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("ZKFCProtocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x8a, 0xf2, 0x76, 0x73,
	0x0e, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2b, 0x00, 0x31, 0x84, 0x78, 0x33, 0x12,
	0x53, 0xf2, 0xf3, 0x0b, 0xf4, 0x92, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0x94, 0x6c, 0xb8, 0xc4, 0x9c,
	0x53, 0x53, 0x52, 0x1d, 0x93, 0x4b, 0x32, 0xcb, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0xc0, 0x3a, 0x84, 0x94, 0xb8, 0x78, 0x72, 0x33, 0x73, 0x72, 0x32, 0x8b, 0x43, 0xf2, 0x41, 0x2a,
	0x24, 0x18, 0x15, 0x98, 0x34, 0x78, 0x83, 0x50, 0xc4, 0x94, 0x24, 0xb9, 0xc4, 0x91, 0x75, 0x17,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82, 0xb5, 0x2b, 0xc9, 0x71, 0xc9, 0xb8, 0x17, 0x25, 0x26, 0xa7,
	0xa6, 0x95, 0xe6, 0xb8, 0x25, 0x66, 0xe6, 0xe4, 0x97, 0xa5, 0x16, 0x21, 0x1b, 0xaf, 0x24, 0xcf,
	0x25, 0x8b, 0x29, 0x8f, 0x64, 0x80, 0xd1, 0x43, 0x46, 0x2e, 0x61, 0x64, 0xf7, 0x07, 0xa7, 0x16,
	0x95, 0x65, 0x26, 0xa7, 0x0a, 0x45, 0x73, 0x71, 0x25, 0xc3, 0xed, 0x14, 0x52, 0xd5, 0x43, 0xf1,
	0x8f, 0x1e, 0x76, 0xcf, 0x48, 0xa9, 0xe1, 0x51, 0x86, 0x64, 0xa9, 0x50, 0x2e, 0x97, 0x40, 0x3a,
	0x9a, 0xab, 0x84, 0xb4, 0xd1, 0xf4, 0xe2, 0xf3, 0x96, 0x94, 0x0e, 0x41, 0xc5, 0x48, 0xd6, 0x39,
	0x99, 0x71, 0x49, 0xe5, 0x17, 0xa5, 0xeb, 0x25, 0x16, 0x24, 0x26, 0x67, 0xa4, 0xc2, 0x74, 0x66,
	0x24, 0x42, 0xa2, 0xca, 0x09, 0x25, 0xfa, 0xc0, 0x74, 0x71, 0x07, 0x23, 0xe3, 0x02, 0x46, 0x46,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x73, 0xcc, 0xde, 0xd9, 0x01, 0x00, 0x00,
}
