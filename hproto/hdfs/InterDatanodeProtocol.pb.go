// Code generated by protoc-gen-go.
// source: InterDatanodeProtocol.proto
// DO NOT EDIT!

package hdfs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hadoop_hdfs "."

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
func (*InitReplicaRecoveryRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

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
	State            *ReplicaStateProto      `protobuf:"varint,2,opt,name=state,enum=hadoop.hdfs.ReplicaStateProto" json:"state,omitempty"`
	Block            *hadoop_hdfs.BlockProto `protobuf:"bytes,3,opt,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *InitReplicaRecoveryResponseProto) Reset()         { *m = InitReplicaRecoveryResponseProto{} }
func (m *InitReplicaRecoveryResponseProto) String() string { return proto.CompactTextString(m) }
func (*InitReplicaRecoveryResponseProto) ProtoMessage()    {}
func (*InitReplicaRecoveryResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{1}
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

func (m *InitReplicaRecoveryResponseProto) GetBlock() *hadoop_hdfs.BlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

// *
// Update replica with new generation stamp and length
type UpdateReplicaUnderRecoveryRequestProto struct {
	Block      *hadoop_hdfs.ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	RecoveryId *uint64                         `protobuf:"varint,2,req,name=recoveryId" json:"recoveryId,omitempty"`
	NewLength  *uint64                         `protobuf:"varint,3,req,name=newLength" json:"newLength,omitempty"`
	// New blockId for copy (truncate), default is 0.
	NewBlockId       *uint64 `protobuf:"varint,4,opt,name=newBlockId,def=0" json:"newBlockId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateReplicaUnderRecoveryRequestProto) Reset() {
	*m = UpdateReplicaUnderRecoveryRequestProto{}
}
func (m *UpdateReplicaUnderRecoveryRequestProto) String() string { return proto.CompactTextString(m) }
func (*UpdateReplicaUnderRecoveryRequestProto) ProtoMessage()    {}
func (*UpdateReplicaUnderRecoveryRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{2}
}

const Default_UpdateReplicaUnderRecoveryRequestProto_NewBlockId uint64 = 0

func (m *UpdateReplicaUnderRecoveryRequestProto) GetBlock() *hadoop_hdfs.ExtendedBlockProto {
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

func (m *UpdateReplicaUnderRecoveryRequestProto) GetNewBlockId() uint64 {
	if m != nil && m.NewBlockId != nil {
		return *m.NewBlockId
	}
	return Default_UpdateReplicaUnderRecoveryRequestProto_NewBlockId
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
	return fileDescriptor6, []int{3}
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
	Metadata: fileDescriptor6,
}

func init() { proto.RegisterFile("InterDatanodeProtocol.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x95, 0x43, 0x47, 0x62, 0x6e, 0x10, 0x42, 0x66, 0xc1, 0x28, 0x8c, 0x66, 0x32, 0x91, 0x80,
	0x2c, 0x68, 0x84, 0x42, 0x11, 0x12, 0xcb, 0x0a, 0x10, 0x05, 0x16, 0xc8, 0x55, 0x37, 0xec, 0x4c,
	0x7c, 0x49, 0x22, 0x2a, 0x3b, 0x38, 0x0e, 0x8f, 0x2f, 0x80, 0x8f, 0x60, 0xc1, 0x9e, 0x9f, 0xe0,
	0xd3, 0x70, 0x1e, 0x88, 0x98, 0x06, 0xda, 0x4d, 0xeb, 0x9c, 0x7b, 0xce, 0xf1, 0xb9, 0xd7, 0x36,
	0xdc, 0x5c, 0x49, 0x83, 0xfa, 0x31, 0x37, 0x5c, 0x2a, 0x81, 0xaf, 0xb4, 0x32, 0x2a, 0x53, 0xdb,
	0xa4, 0x6a, 0x17, 0xd4, 0x2f, 0xb8, 0x50, 0xaa, 0x4a, 0x0a, 0xf1, 0xb6, 0x0e, 0xa0, 0xfd, 0xed,
	0x0b, 0xc1, 0xb5, 0x67, 0x76, 0xbd, 0x46, 0xfd, 0x01, 0x75, 0x8f, 0x44, 0xaf, 0xe1, 0x7c, 0x25,
	0x4b, 0xc3, 0xb0, 0xda, 0x96, 0x19, 0x67, 0x98, 0x29, 0x5b, 0xfb, 0xcc, 0xf0, 0x7d, 0x83, 0xb5,
	0xe9, 0x6c, 0xe9, 0x43, 0x38, 0x7a, 0xb3, 0x55, 0xd9, 0xbb, 0x13, 0x12, 0x7a, 0xb1, 0x9f, 0x5e,
	0x24, 0x23, 0xf7, 0x64, 0x50, 0x94, 0x32, 0x5f, 0xb6, 0x9c, 0x4e, 0xc1, 0x7a, 0x7e, 0xf4, 0x83,
	0x40, 0x38, 0x69, 0x5e, 0x57, 0x4a, 0xd6, 0x7d, 0x68, 0x1a, 0xc1, 0x15, 0xdd, 0xd7, 0x9f, 0xaa,
	0x46, 0x8a, 0x6e, 0x93, 0xcb, 0xcc, 0xc1, 0xe8, 0x02, 0x8e, 0x6a, 0xc3, 0x0d, 0x9e, 0x78, 0x21,
	0x89, 0xaf, 0xa6, 0x67, 0x7f, 0x25, 0xe8, 0x98, 0xeb, 0x96, 0x30, 0x6c, 0xdf, 0x91, 0xe9, 0xfc,
	0x77, 0xee, 0x4b, 0x56, 0xe5, 0xa7, 0x37, 0x1c, 0xd5, 0x6e, 0xda, 0x9f, 0x04, 0x6e, 0x6f, 0x2a,
	0x61, 0x95, 0x83, 0xe3, 0x46, 0x0a, 0xd4, 0x93, 0x13, 0x79, 0xe0, 0x4e, 0xe4, 0xdc, 0x71, 0x7e,
	0xf2, 0xc9, 0xa0, 0x55, 0x8a, 0x9d, 0x1d, 0xe8, 0x19, 0x80, 0x1e, 0xec, 0x56, 0xc2, 0xf6, 0xe2,
	0xc5, 0x33, 0x36, 0x42, 0xe8, 0x29, 0x1c, 0x4b, 0xfc, 0xf8, 0x12, 0x65, 0x6e, 0x0a, 0x1b, 0xba,
	0x2d, 0xff, 0x01, 0xe8, 0x05, 0x80, 0xfd, 0xe8, 0x5c, 0xad, 0x7a, 0x66, 0x7b, 0x9a, 0x3d, 0x22,
	0xf7, 0xd8, 0x08, 0x8c, 0x5e, 0xc0, 0x9d, 0xff, 0x75, 0x30, 0x1e, 0x7b, 0x08, 0x7e, 0x6d, 0x94,
	0xe6, 0x39, 0x6e, 0x9a, 0xb2, 0x9d, 0x3a, 0x89, 0x8f, 0xd9, 0x18, 0x4a, 0xbf, 0x79, 0x70, 0x3a,
	0x79, 0xc9, 0xda, 0xfb, 0x53, 0x66, 0x48, 0x35, 0x5c, 0x2f, 0x77, 0x4f, 0x97, 0xde, 0x75, 0xa6,
	0xb1, 0xe7, 0x72, 0x05, 0xf3, 0xfd, 0xec, 0x71, 0xec, 0x2f, 0x04, 0x82, 0xe6, 0x9f, 0x2d, 0xd2,
	0xfb, 0x8e, 0xdb, 0x61, 0xa7, 0x19, 0x2c, 0x0e, 0x16, 0x8d, 0x92, 0x2c, 0x9f, 0xc3, 0x2d, 0xa5,
	0xf3, 0x84, 0x57, 0x3c, 0x2b, 0xd0, 0x71, 0xa8, 0x9c, 0xc7, 0xb8, 0x9c, 0x7e, 0xa9, 0xdd, 0x7f,
	0xfd, 0x95, 0x90, 0xef, 0x84, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xec, 0x9a, 0x56, 0xc8, 0xcd,
	0x03, 0x00, 0x00,
}
