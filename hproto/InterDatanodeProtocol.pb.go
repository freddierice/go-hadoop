// Code generated by protoc-gen-go.
// source: InterDatanodeProtocol.proto
// DO NOT EDIT!

package hproto

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
// Block with location information and new generation stamp
// to be used for recovery.
type InitReplicaRecoveryRequestProto struct {
	Block            *RecoveringBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *InitReplicaRecoveryRequestProto) Reset()                    { *m = InitReplicaRecoveryRequestProto{} }
func (m *InitReplicaRecoveryRequestProto) String() string            { return proto.CompactTextString(m) }
func (*InitReplicaRecoveryRequestProto) ProtoMessage()               {}
func (*InitReplicaRecoveryRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *InitReplicaRecoveryRequestProto) GetBlock() *RecoveringBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

// *
// Repica recovery information
type InitReplicaRecoveryResponseProto struct {
	ReplicaFound *bool `protobuf:"varint,1,req,name=replicaFound" json:"replicaFound,omitempty"`
	// The following entries are not set if there was no replica found.
	State            *ReplicaStateProto `protobuf:"varint,2,opt,name=state,enum=ReplicaStateProto" json:"state,omitempty"`
	Block            *BlockProto        `protobuf:"bytes,3,opt,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *InitReplicaRecoveryResponseProto) Reset()         { *m = InitReplicaRecoveryResponseProto{} }
func (m *InitReplicaRecoveryResponseProto) String() string { return proto.CompactTextString(m) }
func (*InitReplicaRecoveryResponseProto) ProtoMessage()    {}
func (*InitReplicaRecoveryResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{1}
}

func (m *InitReplicaRecoveryResponseProto) GetReplicaFound() bool {
	if m != nil && m.ReplicaFound != nil {
		return *m.ReplicaFound
	}
	return false
}

func (m *InitReplicaRecoveryResponseProto) GetState() ReplicaStateProto {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ReplicaStateProto_FINALIZED
}

func (m *InitReplicaRecoveryResponseProto) GetBlock() *BlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

// *
// Update replica with new generation stamp and length
type UpdateReplicaUnderRecoveryRequestProto struct {
	Block            *ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	RecoveryId       *uint64             `protobuf:"varint,2,req,name=recoveryId" json:"recoveryId,omitempty"`
	NewLength        *uint64             `protobuf:"varint,3,req,name=newLength" json:"newLength,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *UpdateReplicaUnderRecoveryRequestProto) Reset() {
	*m = UpdateReplicaUnderRecoveryRequestProto{}
}
func (m *UpdateReplicaUnderRecoveryRequestProto) String() string { return proto.CompactTextString(m) }
func (*UpdateReplicaUnderRecoveryRequestProto) ProtoMessage()    {}
func (*UpdateReplicaUnderRecoveryRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{2}
}

func (m *UpdateReplicaUnderRecoveryRequestProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *UpdateReplicaUnderRecoveryRequestProto) GetRecoveryId() uint64 {
	if m != nil && m.RecoveryId != nil {
		return *m.RecoveryId
	}
	return 0
}

func (m *UpdateReplicaUnderRecoveryRequestProto) GetNewLength() uint64 {
	if m != nil && m.NewLength != nil {
		return *m.NewLength
	}
	return 0
}

// *
// Response returns updated block information
type UpdateReplicaUnderRecoveryResponseProto struct {
	StorageUuid      *string `protobuf:"bytes,1,opt,name=storageUuid" json:"storageUuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateReplicaUnderRecoveryResponseProto) Reset() {
	*m = UpdateReplicaUnderRecoveryResponseProto{}
}
func (m *UpdateReplicaUnderRecoveryResponseProto) String() string { return proto.CompactTextString(m) }
func (*UpdateReplicaUnderRecoveryResponseProto) ProtoMessage()    {}
func (*UpdateReplicaUnderRecoveryResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{3}
}

func (m *UpdateReplicaUnderRecoveryResponseProto) GetStorageUuid() string {
	if m != nil && m.StorageUuid != nil {
		return *m.StorageUuid
	}
	return ""
}

func init() {
	proto.RegisterType((*InitReplicaRecoveryRequestProto)(nil), "hadoop.hdfs.InitReplicaRecoveryRequestProto")
	proto.RegisterType((*InitReplicaRecoveryResponseProto)(nil), "hadoop.hdfs.InitReplicaRecoveryResponseProto")
	proto.RegisterType((*UpdateReplicaUnderRecoveryRequestProto)(nil), "hadoop.hdfs.UpdateReplicaUnderRecoveryRequestProto")
	proto.RegisterType((*UpdateReplicaUnderRecoveryResponseProto)(nil), "hadoop.hdfs.UpdateReplicaUnderRecoveryResponseProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for InterDatanodeProtocolService service

type InterDatanodeProtocolServiceClient interface {
	// *
	// Initialize recovery of a replica
	InitReplicaRecovery(ctx context.Context, in *InitReplicaRecoveryRequestProto, opts ...grpc.CallOption) (*InitReplicaRecoveryResponseProto, error)
	// *
	// Update a replica with new generation stamp and length
	UpdateReplicaUnderRecovery(ctx context.Context, in *UpdateReplicaUnderRecoveryRequestProto, opts ...grpc.CallOption) (*UpdateReplicaUnderRecoveryResponseProto, error)
}

type interDatanodeProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewInterDatanodeProtocolServiceClient(cc *grpc.ClientConn) InterDatanodeProtocolServiceClient {
	return &interDatanodeProtocolServiceClient{cc}
}

func (c *interDatanodeProtocolServiceClient) InitReplicaRecovery(ctx context.Context, in *InitReplicaRecoveryRequestProto, opts ...grpc.CallOption) (*InitReplicaRecoveryResponseProto, error) {
	out := new(InitReplicaRecoveryResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.hdfs.InterDatanodeProtocolService/initReplicaRecovery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interDatanodeProtocolServiceClient) UpdateReplicaUnderRecovery(ctx context.Context, in *UpdateReplicaUnderRecoveryRequestProto, opts ...grpc.CallOption) (*UpdateReplicaUnderRecoveryResponseProto, error) {
	out := new(UpdateReplicaUnderRecoveryResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.hdfs.InterDatanodeProtocolService/updateReplicaUnderRecovery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InterDatanodeProtocolService service

type InterDatanodeProtocolServiceServer interface {
	// *
	// Initialize recovery of a replica
	InitReplicaRecovery(context.Context, *InitReplicaRecoveryRequestProto) (*InitReplicaRecoveryResponseProto, error)
	// *
	// Update a replica with new generation stamp and length
	UpdateReplicaUnderRecovery(context.Context, *UpdateReplicaUnderRecoveryRequestProto) (*UpdateReplicaUnderRecoveryResponseProto, error)
}

func RegisterInterDatanodeProtocolServiceServer(s *grpc.Server, srv InterDatanodeProtocolServiceServer) {
	s.RegisterService(&_InterDatanodeProtocolService_serviceDesc, srv)
}

func _InterDatanodeProtocolService_InitReplicaRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReplicaRecoveryRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterDatanodeProtocolServiceServer).InitReplicaRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.hdfs.InterDatanodeProtocolService/InitReplicaRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterDatanodeProtocolServiceServer).InitReplicaRecovery(ctx, req.(*InitReplicaRecoveryRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _InterDatanodeProtocolService_UpdateReplicaUnderRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReplicaUnderRecoveryRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterDatanodeProtocolServiceServer).UpdateReplicaUnderRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.hdfs.InterDatanodeProtocolService/UpdateReplicaUnderRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterDatanodeProtocolServiceServer).UpdateReplicaUnderRecovery(ctx, req.(*UpdateReplicaUnderRecoveryRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _InterDatanodeProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hadoop.hdfs.InterDatanodeProtocolService",
	HandlerType: (*InterDatanodeProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "initReplicaRecovery",
			Handler:    _InterDatanodeProtocolService_InitReplicaRecovery_Handler,
		},
		{
			MethodName: "updateReplicaUnderRecovery",
			Handler:    _InterDatanodeProtocolService_UpdateReplicaUnderRecovery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor5,
}

func init() { proto.RegisterFile("InterDatanodeProtocol.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xe5, 0x40, 0x25, 0x7a, 0x82, 0xb8, 0x70, 0x85, 0x54, 0x85, 0x0a, 0x42, 0x24, 0x20,
	0x08, 0xf0, 0x45, 0xe0, 0x09, 0x2a, 0x40, 0x2a, 0x20, 0x34, 0xa5, 0xea, 0x03, 0x78, 0xf1, 0x59,
	0x12, 0xad, 0xb2, 0x33, 0xc7, 0xe9, 0xb6, 0x27, 0xd8, 0x6e, 0xa7, 0xdd, 0xee, 0x62, 0x8f, 0xba,
	0xfc, 0x9b, 0x96, 0xa8, 0xd9, 0xda, 0x9b, 0x24, 0xfa, 0xce, 0x97, 0x9f, 0xbf, 0xe3, 0x63, 0xc3,
	0x9b, 0x85, 0x34, 0xa8, 0x7f, 0x72, 0xc3, 0xa5, 0x12, 0x78, 0xa0, 0x95, 0x51, 0x91, 0x5a, 0xb3,
	0xac, 0xfa, 0xa0, 0x76, 0xc2, 0x85, 0x52, 0x19, 0x4b, 0xc4, 0x51, 0xee, 0x40, 0xf5, 0x6c, 0x0a,
	0xde, 0x7f, 0x78, 0xb7, 0x90, 0xa9, 0x09, 0x31, 0x5b, 0xa7, 0x11, 0x0f, 0x31, 0x52, 0x1b, 0xd4,
	0xe7, 0x21, 0x9e, 0x14, 0x98, 0x9b, 0x1a, 0x42, 0xbf, 0xc0, 0xe8, 0x70, 0xad, 0xa2, 0xe3, 0x29,
	0x71, 0x2d, 0xdf, 0x0e, 0x5e, 0xb3, 0xd6, 0x95, 0xca, 0x78, 0x5e, 0xe9, 0xb5, 0x2b, 0x6c, 0x3c,
	0xde, 0x35, 0x01, 0x77, 0x10, 0x98, 0x67, 0x4a, 0xe6, 0x4d, 0x2c, 0xea, 0xc1, 0x4b, 0xdd, 0xd4,
	0x7f, 0xab, 0x42, 0x8a, 0x1a, 0xfc, 0x22, 0xec, 0x69, 0xd4, 0x87, 0x51, 0x6e, 0xb8, 0xc1, 0xa9,
	0xe5, 0x12, 0xff, 0x55, 0x40, 0x59, 0x4b, 0x5c, 0x56, 0x62, 0xbb, 0x64, 0x6d, 0xa0, 0xef, 0xef,
	0xf3, 0x3d, 0x2b, 0x9d, 0x76, 0x60, 0xb3, 0xed, 0x54, 0x57, 0x04, 0x3e, 0xae, 0x32, 0x51, 0xba,
	0x5b, 0xca, 0x4a, 0x0a, 0xd4, 0x83, 0xdd, 0x7e, 0xee, 0x77, 0x3b, 0x61, 0xbf, 0xce, 0x0c, 0x96,
	0x6e, 0xb1, 0x45, 0xa5, 0x6f, 0x01, 0x74, 0x8b, 0x58, 0x88, 0x32, 0xa7, 0xe5, 0x3f, 0x0f, 0x3b,
	0x0a, 0x9d, 0xc1, 0x58, 0xe2, 0xe9, 0x3f, 0x94, 0xb1, 0x49, 0xca, 0x70, 0x55, 0xf9, 0x41, 0xf0,
	0xfe, 0xc2, 0xa7, 0xa7, 0x22, 0x75, 0xf7, 0xcb, 0x05, 0x3b, 0x37, 0x4a, 0xf3, 0x18, 0x57, 0x45,
	0x5a, 0x6d, 0x17, 0xf1, 0xc7, 0x61, 0x57, 0x0a, 0x6e, 0x2c, 0x98, 0x0d, 0xce, 0x7f, 0x89, 0x7a,
	0x93, 0x46, 0x48, 0x35, 0x4c, 0xd2, 0xed, 0xb1, 0xd0, 0xaf, 0xac, 0x73, 0x30, 0xd8, 0x8e, 0x93,
	0xe0, 0x7c, 0xdb, 0xed, 0xee, 0xc6, 0xbe, 0x20, 0xe0, 0x14, 0x8f, 0xb6, 0x48, 0xbf, 0xf7, 0x68,
	0xfb, 0x8d, 0xc7, 0xf9, 0xb1, 0xf7, 0x4f, 0x9d, 0x24, 0xf3, 0x3f, 0xf0, 0x41, 0xe9, 0x98, 0xf1,
	0x8c, 0x47, 0x09, 0xf6, 0x08, 0x59, 0xef, 0x9e, 0xcc, 0x87, 0x2f, 0x51, 0xfd, 0xce, 0x2f, 0x09,
	0xb9, 0x25, 0xe4, 0x2e, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x48, 0x8c, 0x62, 0x68, 0x03, 0x00, 0x00,
}
