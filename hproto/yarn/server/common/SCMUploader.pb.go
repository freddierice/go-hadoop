// Code generated by protoc-gen-go.
// source: SCMUploader.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	SCMUploader.proto

It has these top-level messages:
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hadoop_yarn3 "."

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

// Client API for SCMUploaderProtocolService service

type SCMUploaderProtocolServiceClient interface {
	Notify(ctx context.Context, in *hadoop_yarn3.SCMUploaderNotifyRequestProto, opts ...grpc.CallOption) (*hadoop_yarn3.SCMUploaderNotifyResponseProto, error)
	CanUpload(ctx context.Context, in *hadoop_yarn3.SCMUploaderCanUploadRequestProto, opts ...grpc.CallOption) (*hadoop_yarn3.SCMUploaderCanUploadResponseProto, error)
}

type sCMUploaderProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewSCMUploaderProtocolServiceClient(cc *grpc.ClientConn) SCMUploaderProtocolServiceClient {
	return &sCMUploaderProtocolServiceClient{cc}
}

func (c *sCMUploaderProtocolServiceClient) Notify(ctx context.Context, in *hadoop_yarn3.SCMUploaderNotifyRequestProto, opts ...grpc.CallOption) (*hadoop_yarn3.SCMUploaderNotifyResponseProto, error) {
	out := new(hadoop_yarn3.SCMUploaderNotifyResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.yarn.SCMUploaderProtocolService/notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sCMUploaderProtocolServiceClient) CanUpload(ctx context.Context, in *hadoop_yarn3.SCMUploaderCanUploadRequestProto, opts ...grpc.CallOption) (*hadoop_yarn3.SCMUploaderCanUploadResponseProto, error) {
	out := new(hadoop_yarn3.SCMUploaderCanUploadResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.yarn.SCMUploaderProtocolService/canUpload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SCMUploaderProtocolService service

type SCMUploaderProtocolServiceServer interface {
	Notify(context.Context, *hadoop_yarn3.SCMUploaderNotifyRequestProto) (*hadoop_yarn3.SCMUploaderNotifyResponseProto, error)
	CanUpload(context.Context, *hadoop_yarn3.SCMUploaderCanUploadRequestProto) (*hadoop_yarn3.SCMUploaderCanUploadResponseProto, error)
}

func RegisterSCMUploaderProtocolServiceServer(s *grpc.Server, srv SCMUploaderProtocolServiceServer) {
	s.RegisterService(&_SCMUploaderProtocolService_serviceDesc, srv)
}

func _SCMUploaderProtocolService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_yarn3.SCMUploaderNotifyRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCMUploaderProtocolServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.yarn.SCMUploaderProtocolService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCMUploaderProtocolServiceServer).Notify(ctx, req.(*hadoop_yarn3.SCMUploaderNotifyRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _SCMUploaderProtocolService_CanUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_yarn3.SCMUploaderCanUploadRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCMUploaderProtocolServiceServer).CanUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.yarn.SCMUploaderProtocolService/CanUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCMUploaderProtocolServiceServer).CanUpload(ctx, req.(*hadoop_yarn3.SCMUploaderCanUploadRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCMUploaderProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hadoop.yarn.SCMUploaderProtocolService",
	HandlerType: (*SCMUploaderProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "notify",
			Handler:    _SCMUploaderProtocolService_Notify_Handler,
		},
		{
			MethodName: "canUpload",
			Handler:    _SCMUploaderProtocolService_CanUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("SCMUploader.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0x76, 0xf6, 0x0d,
	0x2d, 0xc8, 0xc9, 0x4f, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce,
	0x48, 0x4c, 0xc9, 0xcf, 0x2f, 0xd0, 0xab, 0x4c, 0x2c, 0xca, 0x93, 0x52, 0x07, 0x91, 0xf1, 0xc5,
	0xa9, 0x45, 0x65, 0xa9, 0x45, 0xf1, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x10, 0x5e, 0x66, 0x72, 0x6a,
	0x3c, 0x58, 0x79, 0x31, 0x44, 0x97, 0xd1, 0x6b, 0x46, 0x2e, 0x29, 0x24, 0xb3, 0x02, 0x40, 0x82,
	0xc9, 0xf9, 0x39, 0xc1, 0x10, 0xb5, 0x42, 0x89, 0x5c, 0x6c, 0x79, 0xf9, 0x25, 0x99, 0x69, 0x95,
	0x42, 0x5a, 0x7a, 0x48, 0xe6, 0xeb, 0x21, 0x69, 0xf1, 0x03, 0xcb, 0x07, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x80, 0xf5, 0x4b, 0x69, 0x13, 0x52, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0x56,
	0x2c, 0x94, 0xc5, 0xc5, 0x99, 0x9c, 0x98, 0x07, 0x51, 0x21, 0xa4, 0x8b, 0x4b, 0xa7, 0x33, 0x4c,
	0x09, 0x8a, 0x45, 0x7a, 0x44, 0x28, 0x47, 0xb2, 0xcb, 0xc9, 0x92, 0x4b, 0x26, 0xbf, 0x28, 0x5d,
	0x2f, 0xb1, 0x20, 0x31, 0x39, 0x23, 0x15, 0x45, 0x2f, 0x38, 0x34, 0x9c, 0x84, 0xb1, 0x04, 0x45,
	0x07, 0x23, 0xe3, 0x02, 0x46, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0xc4, 0xa5, 0xe7,
	0x72, 0x01, 0x00, 0x00,
}
