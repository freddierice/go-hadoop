// Code generated by protoc-gen-go.
// source: RefreshAuthorizationPolicyProtocol.proto
// DO NOT EDIT!

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

// *
//  Refresh service acl request.
type RefreshServiceAclRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshServiceAclRequestProto) Reset()                    { *m = RefreshServiceAclRequestProto{} }
func (m *RefreshServiceAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RefreshServiceAclRequestProto) ProtoMessage()               {}
func (*RefreshServiceAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// *
// void response
type RefreshServiceAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshServiceAclResponseProto) Reset()                    { *m = RefreshServiceAclResponseProto{} }
func (m *RefreshServiceAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RefreshServiceAclResponseProto) ProtoMessage()               {}
func (*RefreshServiceAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func init() {
	proto.RegisterType((*RefreshServiceAclRequestProto)(nil), "hadoop.common.RefreshServiceAclRequestProto")
	proto.RegisterType((*RefreshServiceAclResponseProto)(nil), "hadoop.common.RefreshServiceAclResponseProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RefreshAuthorizationPolicyProtocolService service

type RefreshAuthorizationPolicyProtocolServiceClient interface {
	// *
	// Refresh the service-level authorization policy in-effect.
	RefreshServiceAcl(ctx context.Context, in *RefreshServiceAclRequestProto, opts ...grpc.CallOption) (*RefreshServiceAclResponseProto, error)
}

type refreshAuthorizationPolicyProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewRefreshAuthorizationPolicyProtocolServiceClient(cc *grpc.ClientConn) RefreshAuthorizationPolicyProtocolServiceClient {
	return &refreshAuthorizationPolicyProtocolServiceClient{cc}
}

func (c *refreshAuthorizationPolicyProtocolServiceClient) RefreshServiceAcl(ctx context.Context, in *RefreshServiceAclRequestProto, opts ...grpc.CallOption) (*RefreshServiceAclResponseProto, error) {
	out := new(RefreshServiceAclResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.common.RefreshAuthorizationPolicyProtocolService/refreshServiceAcl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RefreshAuthorizationPolicyProtocolService service

type RefreshAuthorizationPolicyProtocolServiceServer interface {
	// *
	// Refresh the service-level authorization policy in-effect.
	RefreshServiceAcl(context.Context, *RefreshServiceAclRequestProto) (*RefreshServiceAclResponseProto, error)
}

func RegisterRefreshAuthorizationPolicyProtocolServiceServer(s *grpc.Server, srv RefreshAuthorizationPolicyProtocolServiceServer) {
	s.RegisterService(&_RefreshAuthorizationPolicyProtocolService_serviceDesc, srv)
}

func _RefreshAuthorizationPolicyProtocolService_RefreshServiceAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshServiceAclRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefreshAuthorizationPolicyProtocolServiceServer).RefreshServiceAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.common.RefreshAuthorizationPolicyProtocolService/RefreshServiceAcl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefreshAuthorizationPolicyProtocolServiceServer).RefreshServiceAcl(ctx, req.(*RefreshServiceAclRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _RefreshAuthorizationPolicyProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hadoop.common.RefreshAuthorizationPolicyProtocolService",
	HandlerType: (*RefreshAuthorizationPolicyProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "refreshServiceAcl",
			Handler:    _RefreshAuthorizationPolicyProtocolService_RefreshServiceAcl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor6,
}

func init() { proto.RegisterFile("RefreshAuthorizationPolicyProtocol.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x08, 0x4a, 0x4d, 0x2b,
	0x4a, 0x2d, 0xce, 0x70, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0xac, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf,
	0x0b, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0x0c, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2b,
	0x00, 0x31, 0x84, 0x78, 0x33, 0x12, 0x53, 0xf2, 0xf3, 0x0b, 0xf4, 0x92, 0xf3, 0x73, 0x73, 0xf3,
	0xf3, 0x94, 0xe4, 0xb9, 0x64, 0xa1, 0x5a, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x1d, 0x93,
	0x73, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xc0, 0x1a, 0x95, 0x14, 0xb8, 0xe4, 0xb0, 0x28,
	0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05, 0xab, 0x30, 0x9a, 0xcb, 0xc8, 0xa5, 0x49, 0xd8, 0x7a,
	0xa8, 0x6e, 0xa1, 0x02, 0x2e, 0xc1, 0x22, 0x74, 0xf3, 0x84, 0x74, 0xf4, 0x50, 0x5c, 0xa5, 0x87,
	0xd7, 0x49, 0x52, 0xba, 0x84, 0x55, 0x23, 0xb9, 0xcf, 0x29, 0x88, 0x4b, 0x21, 0xbf, 0x28, 0x5d,
	0x2f, 0xb1, 0x20, 0x31, 0x39, 0x23, 0x15, 0xa6, 0xb5, 0x38, 0x35, 0xb9, 0xb4, 0x28, 0xb3, 0xa4,
	0x12, 0x12, 0x2a, 0x4e, 0x44, 0x84, 0x1f, 0x98, 0x2e, 0xee, 0x60, 0x64, 0x5c, 0xc0, 0xc8, 0x08,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xf2, 0xf6, 0x15, 0x70, 0x01, 0x00, 0x00,
}