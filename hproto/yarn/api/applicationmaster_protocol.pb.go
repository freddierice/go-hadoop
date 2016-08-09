// Code generated by protoc-gen-go.
// source: applicationmaster_protocol.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	applicationmaster_protocol.proto

It has these top-level messages:
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hadoop_yarn1 "."

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ApplicationMasterProtocolService service

type ApplicationMasterProtocolServiceClient interface {
	RegisterApplicationMaster(ctx context.Context, in *hadoop_yarn1.RegisterApplicationMasterRequestProto, opts ...grpc.CallOption) (*hadoop_yarn1.RegisterApplicationMasterResponseProto, error)
	FinishApplicationMaster(ctx context.Context, in *hadoop_yarn1.FinishApplicationMasterRequestProto, opts ...grpc.CallOption) (*hadoop_yarn1.FinishApplicationMasterResponseProto, error)
	Allocate(ctx context.Context, in *hadoop_yarn1.AllocateRequestProto, opts ...grpc.CallOption) (*hadoop_yarn1.AllocateResponseProto, error)
}

type applicationMasterProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationMasterProtocolServiceClient(cc *grpc.ClientConn) ApplicationMasterProtocolServiceClient {
	return &applicationMasterProtocolServiceClient{cc}
}

func (c *applicationMasterProtocolServiceClient) RegisterApplicationMaster(ctx context.Context, in *hadoop_yarn1.RegisterApplicationMasterRequestProto, opts ...grpc.CallOption) (*hadoop_yarn1.RegisterApplicationMasterResponseProto, error) {
	out := new(hadoop_yarn1.RegisterApplicationMasterResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.yarn.ApplicationMasterProtocolService/registerApplicationMaster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationMasterProtocolServiceClient) FinishApplicationMaster(ctx context.Context, in *hadoop_yarn1.FinishApplicationMasterRequestProto, opts ...grpc.CallOption) (*hadoop_yarn1.FinishApplicationMasterResponseProto, error) {
	out := new(hadoop_yarn1.FinishApplicationMasterResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.yarn.ApplicationMasterProtocolService/finishApplicationMaster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationMasterProtocolServiceClient) Allocate(ctx context.Context, in *hadoop_yarn1.AllocateRequestProto, opts ...grpc.CallOption) (*hadoop_yarn1.AllocateResponseProto, error) {
	out := new(hadoop_yarn1.AllocateResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.yarn.ApplicationMasterProtocolService/allocate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationMasterProtocolService service

type ApplicationMasterProtocolServiceServer interface {
	RegisterApplicationMaster(context.Context, *hadoop_yarn1.RegisterApplicationMasterRequestProto) (*hadoop_yarn1.RegisterApplicationMasterResponseProto, error)
	FinishApplicationMaster(context.Context, *hadoop_yarn1.FinishApplicationMasterRequestProto) (*hadoop_yarn1.FinishApplicationMasterResponseProto, error)
	Allocate(context.Context, *hadoop_yarn1.AllocateRequestProto) (*hadoop_yarn1.AllocateResponseProto, error)
}

func RegisterApplicationMasterProtocolServiceServer(s *grpc.Server, srv ApplicationMasterProtocolServiceServer) {
	s.RegisterService(&_ApplicationMasterProtocolService_serviceDesc, srv)
}

func _ApplicationMasterProtocolService_RegisterApplicationMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_yarn1.RegisterApplicationMasterRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationMasterProtocolServiceServer).RegisterApplicationMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.yarn.ApplicationMasterProtocolService/RegisterApplicationMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationMasterProtocolServiceServer).RegisterApplicationMaster(ctx, req.(*hadoop_yarn1.RegisterApplicationMasterRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationMasterProtocolService_FinishApplicationMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_yarn1.FinishApplicationMasterRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationMasterProtocolServiceServer).FinishApplicationMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.yarn.ApplicationMasterProtocolService/FinishApplicationMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationMasterProtocolServiceServer).FinishApplicationMaster(ctx, req.(*hadoop_yarn1.FinishApplicationMasterRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationMasterProtocolService_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_yarn1.AllocateRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationMasterProtocolServiceServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.yarn.ApplicationMasterProtocolService/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationMasterProtocolServiceServer).Allocate(ctx, req.(*hadoop_yarn1.AllocateRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationMasterProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hadoop.yarn.ApplicationMasterProtocolService",
	HandlerType: (*ApplicationMasterProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerApplicationMaster",
			Handler:    _ApplicationMasterProtocolService_RegisterApplicationMaster_Handler,
		},
		{
			MethodName: "finishApplicationMaster",
			Handler:    _ApplicationMasterProtocolService_FinishApplicationMaster_Handler,
		},
		{
			MethodName: "allocate",
			Handler:    _ApplicationMasterProtocolService_Allocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("applicationmaster_protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2c, 0x28, 0xc8,
	0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xcb, 0x4d, 0x2c, 0x2e, 0x49, 0x2d, 0x8a, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x03, 0x33, 0x84, 0xb8, 0x33, 0x12, 0x53, 0xf2, 0xf3,
	0x0b, 0xf4, 0x2a, 0x13, 0x8b, 0xf2, 0xa4, 0x24, 0x41, 0x64, 0x7c, 0x71, 0x6a, 0x51, 0x59, 0x66,
	0x72, 0x2a, 0x44, 0x65, 0x31, 0x44, 0x9d, 0xd1, 0x1f, 0x26, 0x2e, 0x05, 0x47, 0x84, 0x61, 0xbe,
	0x60, 0xc3, 0x02, 0xa0, 0x66, 0x05, 0x43, 0x74, 0x08, 0xb5, 0x30, 0x72, 0x49, 0x16, 0xa5, 0xa6,
	0x67, 0x82, 0xe4, 0x30, 0x14, 0x0b, 0x19, 0xe9, 0x21, 0xd9, 0xa5, 0x17, 0x84, 0x4b, 0x5d, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0xd8, 0x6c, 0x29, 0x63, 0x62, 0xf5, 0x14, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x82, 0x35, 0x09, 0xd5, 0x71, 0x89, 0xa7, 0x65, 0xe6, 0x65, 0x16, 0x67, 0x60, 0xba,
	0xc1, 0x00, 0xc5, 0x3c, 0x37, 0xec, 0xaa, 0x50, 0x5c, 0x60, 0x48, 0x9c, 0x0e, 0x64, 0xfb, 0x03,
	0xb9, 0x38, 0x12, 0x73, 0x72, 0xf2, 0x81, 0x0a, 0x52, 0x85, 0x14, 0x51, 0xb4, 0x3b, 0x42, 0x85,
	0x51, 0x6c, 0x50, 0xc2, 0xa1, 0x04, 0xc9, 0x48, 0x27, 0x7b, 0x2e, 0x99, 0xfc, 0xa2, 0x74, 0xbd,
	0xc4, 0x82, 0xc4, 0xe4, 0x8c, 0x54, 0x14, 0xf5, 0xe0, 0xe8, 0x71, 0x92, 0xc4, 0x19, 0x37, 0x1d,
	0x8c, 0x8c, 0x0b, 0x18, 0x19, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x6f, 0x1f, 0x35, 0x0a,
	0x02, 0x00, 0x00,
}
