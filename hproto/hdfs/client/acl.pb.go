// Code generated by protoc-gen-go.
// source: acl.proto
// DO NOT EDIT!

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	acl.proto

It has these top-level messages:
	AclEntryProto
	AclStatusProto
	ModifyAclEntriesRequestProto
	ModifyAclEntriesResponseProto
	RemoveAclRequestProto
	RemoveAclResponseProto
	RemoveAclEntriesRequestProto
	RemoveAclEntriesResponseProto
	RemoveDefaultAclRequestProto
	RemoveDefaultAclResponseProto
	SetAclRequestProto
	SetAclResponseProto
	GetAclStatusRequestProto
	GetAclStatusResponseProto
*/
package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hadoop_hdfs "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AclEntryProto_AclEntryScopeProto int32

const (
	AclEntryProto_ACCESS  AclEntryProto_AclEntryScopeProto = 0
	AclEntryProto_DEFAULT AclEntryProto_AclEntryScopeProto = 1
)

var AclEntryProto_AclEntryScopeProto_name = map[int32]string{
	0: "ACCESS",
	1: "DEFAULT",
}
var AclEntryProto_AclEntryScopeProto_value = map[string]int32{
	"ACCESS":  0,
	"DEFAULT": 1,
}

func (x AclEntryProto_AclEntryScopeProto) Enum() *AclEntryProto_AclEntryScopeProto {
	p := new(AclEntryProto_AclEntryScopeProto)
	*p = x
	return p
}
func (x AclEntryProto_AclEntryScopeProto) String() string {
	return proto.EnumName(AclEntryProto_AclEntryScopeProto_name, int32(x))
}
func (x *AclEntryProto_AclEntryScopeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_AclEntryScopeProto_value, data, "AclEntryProto_AclEntryScopeProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_AclEntryScopeProto(value)
	return nil
}
func (AclEntryProto_AclEntryScopeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type AclEntryProto_AclEntryTypeProto int32

const (
	AclEntryProto_USER  AclEntryProto_AclEntryTypeProto = 0
	AclEntryProto_GROUP AclEntryProto_AclEntryTypeProto = 1
	AclEntryProto_MASK  AclEntryProto_AclEntryTypeProto = 2
	AclEntryProto_OTHER AclEntryProto_AclEntryTypeProto = 3
)

var AclEntryProto_AclEntryTypeProto_name = map[int32]string{
	0: "USER",
	1: "GROUP",
	2: "MASK",
	3: "OTHER",
}
var AclEntryProto_AclEntryTypeProto_value = map[string]int32{
	"USER":  0,
	"GROUP": 1,
	"MASK":  2,
	"OTHER": 3,
}

func (x AclEntryProto_AclEntryTypeProto) Enum() *AclEntryProto_AclEntryTypeProto {
	p := new(AclEntryProto_AclEntryTypeProto)
	*p = x
	return p
}
func (x AclEntryProto_AclEntryTypeProto) String() string {
	return proto.EnumName(AclEntryProto_AclEntryTypeProto_name, int32(x))
}
func (x *AclEntryProto_AclEntryTypeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_AclEntryTypeProto_value, data, "AclEntryProto_AclEntryTypeProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_AclEntryTypeProto(value)
	return nil
}
func (AclEntryProto_AclEntryTypeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

type AclEntryProto_FsActionProto int32

const (
	AclEntryProto_NONE          AclEntryProto_FsActionProto = 0
	AclEntryProto_EXECUTE       AclEntryProto_FsActionProto = 1
	AclEntryProto_WRITE         AclEntryProto_FsActionProto = 2
	AclEntryProto_WRITE_EXECUTE AclEntryProto_FsActionProto = 3
	AclEntryProto_READ          AclEntryProto_FsActionProto = 4
	AclEntryProto_READ_EXECUTE  AclEntryProto_FsActionProto = 5
	AclEntryProto_READ_WRITE    AclEntryProto_FsActionProto = 6
	AclEntryProto_PERM_ALL      AclEntryProto_FsActionProto = 7
)

var AclEntryProto_FsActionProto_name = map[int32]string{
	0: "NONE",
	1: "EXECUTE",
	2: "WRITE",
	3: "WRITE_EXECUTE",
	4: "READ",
	5: "READ_EXECUTE",
	6: "READ_WRITE",
	7: "PERM_ALL",
}
var AclEntryProto_FsActionProto_value = map[string]int32{
	"NONE":          0,
	"EXECUTE":       1,
	"WRITE":         2,
	"WRITE_EXECUTE": 3,
	"READ":          4,
	"READ_EXECUTE":  5,
	"READ_WRITE":    6,
	"PERM_ALL":      7,
}

func (x AclEntryProto_FsActionProto) Enum() *AclEntryProto_FsActionProto {
	p := new(AclEntryProto_FsActionProto)
	*p = x
	return p
}
func (x AclEntryProto_FsActionProto) String() string {
	return proto.EnumName(AclEntryProto_FsActionProto_name, int32(x))
}
func (x *AclEntryProto_FsActionProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_FsActionProto_value, data, "AclEntryProto_FsActionProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_FsActionProto(value)
	return nil
}
func (AclEntryProto_FsActionProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2}
}

type AclEntryProto struct {
	Type             *AclEntryProto_AclEntryTypeProto  `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.AclEntryProto_AclEntryTypeProto" json:"type,omitempty"`
	Scope            *AclEntryProto_AclEntryScopeProto `protobuf:"varint,2,req,name=scope,enum=hadoop.hdfs.AclEntryProto_AclEntryScopeProto" json:"scope,omitempty"`
	Permissions      *AclEntryProto_FsActionProto      `protobuf:"varint,3,req,name=permissions,enum=hadoop.hdfs.AclEntryProto_FsActionProto" json:"permissions,omitempty"`
	Name             *string                           `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *AclEntryProto) Reset()                    { *m = AclEntryProto{} }
func (m *AclEntryProto) String() string            { return proto.CompactTextString(m) }
func (*AclEntryProto) ProtoMessage()               {}
func (*AclEntryProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AclEntryProto) GetType() AclEntryProto_AclEntryTypeProto {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return AclEntryProto_USER
}

func (m *AclEntryProto) GetScope() AclEntryProto_AclEntryScopeProto {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return AclEntryProto_ACCESS
}

func (m *AclEntryProto) GetPermissions() AclEntryProto_FsActionProto {
	if m != nil && m.Permissions != nil {
		return *m.Permissions
	}
	return AclEntryProto_NONE
}

func (m *AclEntryProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type AclStatusProto struct {
	Owner            *string                        `protobuf:"bytes,1,req,name=owner" json:"owner,omitempty"`
	Group            *string                        `protobuf:"bytes,2,req,name=group" json:"group,omitempty"`
	Sticky           *bool                          `protobuf:"varint,3,req,name=sticky" json:"sticky,omitempty"`
	Entries          []*AclEntryProto               `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	Permission       *hadoop_hdfs.FsPermissionProto `protobuf:"bytes,5,opt,name=permission" json:"permission,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *AclStatusProto) Reset()                    { *m = AclStatusProto{} }
func (m *AclStatusProto) String() string            { return proto.CompactTextString(m) }
func (*AclStatusProto) ProtoMessage()               {}
func (*AclStatusProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AclStatusProto) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *AclStatusProto) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *AclStatusProto) GetSticky() bool {
	if m != nil && m.Sticky != nil {
		return *m.Sticky
	}
	return false
}

func (m *AclStatusProto) GetEntries() []*AclEntryProto {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AclStatusProto) GetPermission() *hadoop_hdfs.FsPermissionProto {
	if m != nil {
		return m.Permission
	}
	return nil
}

type ModifyAclEntriesRequestProto struct {
	Src              *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec          []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ModifyAclEntriesRequestProto) Reset()                    { *m = ModifyAclEntriesRequestProto{} }
func (m *ModifyAclEntriesRequestProto) String() string            { return proto.CompactTextString(m) }
func (*ModifyAclEntriesRequestProto) ProtoMessage()               {}
func (*ModifyAclEntriesRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ModifyAclEntriesRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *ModifyAclEntriesRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type ModifyAclEntriesResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ModifyAclEntriesResponseProto) Reset()                    { *m = ModifyAclEntriesResponseProto{} }
func (m *ModifyAclEntriesResponseProto) String() string            { return proto.CompactTextString(m) }
func (*ModifyAclEntriesResponseProto) ProtoMessage()               {}
func (*ModifyAclEntriesResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RemoveAclRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RemoveAclRequestProto) Reset()                    { *m = RemoveAclRequestProto{} }
func (m *RemoveAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclRequestProto) ProtoMessage()               {}
func (*RemoveAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type RemoveAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveAclResponseProto) Reset()                    { *m = RemoveAclResponseProto{} }
func (m *RemoveAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclResponseProto) ProtoMessage()               {}
func (*RemoveAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RemoveAclEntriesRequestProto struct {
	Src              *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec          []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *RemoveAclEntriesRequestProto) Reset()                    { *m = RemoveAclEntriesRequestProto{} }
func (m *RemoveAclEntriesRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclEntriesRequestProto) ProtoMessage()               {}
func (*RemoveAclEntriesRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RemoveAclEntriesRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *RemoveAclEntriesRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type RemoveAclEntriesResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveAclEntriesResponseProto) Reset()                    { *m = RemoveAclEntriesResponseProto{} }
func (m *RemoveAclEntriesResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveAclEntriesResponseProto) ProtoMessage()               {}
func (*RemoveAclEntriesResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type RemoveDefaultAclRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RemoveDefaultAclRequestProto) Reset()                    { *m = RemoveDefaultAclRequestProto{} }
func (m *RemoveDefaultAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveDefaultAclRequestProto) ProtoMessage()               {}
func (*RemoveDefaultAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RemoveDefaultAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type RemoveDefaultAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveDefaultAclResponseProto) Reset()                    { *m = RemoveDefaultAclResponseProto{} }
func (m *RemoveDefaultAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RemoveDefaultAclResponseProto) ProtoMessage()               {}
func (*RemoveDefaultAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SetAclRequestProto struct {
	Src              *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec          []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SetAclRequestProto) Reset()                    { *m = SetAclRequestProto{} }
func (m *SetAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*SetAclRequestProto) ProtoMessage()               {}
func (*SetAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *SetAclRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type SetAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SetAclResponseProto) Reset()                    { *m = SetAclResponseProto{} }
func (m *SetAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*SetAclResponseProto) ProtoMessage()               {}
func (*SetAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type GetAclStatusRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetAclStatusRequestProto) Reset()                    { *m = GetAclStatusRequestProto{} }
func (m *GetAclStatusRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetAclStatusRequestProto) ProtoMessage()               {}
func (*GetAclStatusRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetAclStatusRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type GetAclStatusResponseProto struct {
	Result           *AclStatusProto `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetAclStatusResponseProto) Reset()                    { *m = GetAclStatusResponseProto{} }
func (m *GetAclStatusResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetAclStatusResponseProto) ProtoMessage()               {}
func (*GetAclStatusResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetAclStatusResponseProto) GetResult() *AclStatusProto {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*AclEntryProto)(nil), "hadoop.hdfs.AclEntryProto")
	proto.RegisterType((*AclStatusProto)(nil), "hadoop.hdfs.AclStatusProto")
	proto.RegisterType((*ModifyAclEntriesRequestProto)(nil), "hadoop.hdfs.ModifyAclEntriesRequestProto")
	proto.RegisterType((*ModifyAclEntriesResponseProto)(nil), "hadoop.hdfs.ModifyAclEntriesResponseProto")
	proto.RegisterType((*RemoveAclRequestProto)(nil), "hadoop.hdfs.RemoveAclRequestProto")
	proto.RegisterType((*RemoveAclResponseProto)(nil), "hadoop.hdfs.RemoveAclResponseProto")
	proto.RegisterType((*RemoveAclEntriesRequestProto)(nil), "hadoop.hdfs.RemoveAclEntriesRequestProto")
	proto.RegisterType((*RemoveAclEntriesResponseProto)(nil), "hadoop.hdfs.RemoveAclEntriesResponseProto")
	proto.RegisterType((*RemoveDefaultAclRequestProto)(nil), "hadoop.hdfs.RemoveDefaultAclRequestProto")
	proto.RegisterType((*RemoveDefaultAclResponseProto)(nil), "hadoop.hdfs.RemoveDefaultAclResponseProto")
	proto.RegisterType((*SetAclRequestProto)(nil), "hadoop.hdfs.SetAclRequestProto")
	proto.RegisterType((*SetAclResponseProto)(nil), "hadoop.hdfs.SetAclResponseProto")
	proto.RegisterType((*GetAclStatusRequestProto)(nil), "hadoop.hdfs.GetAclStatusRequestProto")
	proto.RegisterType((*GetAclStatusResponseProto)(nil), "hadoop.hdfs.GetAclStatusResponseProto")
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryScopeProto", AclEntryProto_AclEntryScopeProto_name, AclEntryProto_AclEntryScopeProto_value)
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryTypeProto", AclEntryProto_AclEntryTypeProto_name, AclEntryProto_AclEntryTypeProto_value)
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_FsActionProto", AclEntryProto_FsActionProto_name, AclEntryProto_FsActionProto_value)
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x8e, 0xd2, 0x40,
	0x14, 0xb6, 0x94, 0xdf, 0xd3, 0x85, 0x74, 0x47, 0x77, 0x53, 0xd7, 0xbf, 0x4d, 0x13, 0x13, 0x4c,
	0x76, 0x89, 0x41, 0xbd, 0xd4, 0xd8, 0x85, 0xb2, 0xfe, 0xc0, 0x42, 0xa6, 0x10, 0xbd, 0x30, 0xd9,
	0x34, 0x65, 0x58, 0x88, 0xc0, 0xd4, 0x4e, 0xd1, 0x70, 0xe3, 0xb3, 0xf8, 0x3c, 0x3e, 0x87, 0x0f,
	0xe2, 0x74, 0xa6, 0xe0, 0x20, 0x71, 0x31, 0x26, 0xde, 0x34, 0x67, 0xce, 0xf9, 0xbe, 0xef, 0x9c,
	0xaf, 0xa7, 0x1d, 0x28, 0xf9, 0xc1, 0xb4, 0x16, 0x46, 0x34, 0xa6, 0xc8, 0x18, 0xfb, 0x43, 0x4a,
	0xc3, 0xda, 0x78, 0x38, 0x62, 0x47, 0x90, 0x3c, 0x65, 0xc1, 0xfe, 0xa1, 0x43, 0xd9, 0x09, 0xa6,
	0xee, 0x3c, 0x8e, 0x96, 0x3d, 0x01, 0x7d, 0x09, 0xd9, 0x78, 0x19, 0x12, 0x4b, 0x3b, 0xce, 0x54,
	0x2b, 0xf5, 0x93, 0x9a, 0xc2, 0xac, 0x6d, 0x20, 0xd7, 0xa7, 0x3e, 0x87, 0x8b, 0x0c, 0x16, 0x4c,
	0xd4, 0x80, 0x1c, 0x0b, 0x28, 0x97, 0xc8, 0x08, 0x89, 0xd3, 0xbf, 0x90, 0xf0, 0x12, 0xbc, 0xd4,
	0x90, 0x5c, 0xf4, 0x06, 0x8c, 0x90, 0x44, 0xb3, 0x09, 0x63, 0x13, 0x3a, 0x67, 0x96, 0x2e, 0xa4,
	0xaa, 0xd7, 0x48, 0xb5, 0x98, 0x13, 0xc4, 0x1c, 0x2b, 0x55, 0x54, 0x32, 0x42, 0x90, 0x9d, 0xfb,
	0x33, 0x62, 0x65, 0x8f, 0xb5, 0x6a, 0x09, 0x8b, 0xd8, 0x3e, 0x05, 0xb4, 0xdd, 0x1c, 0x01, 0xe4,
	0x9d, 0x46, 0xc3, 0xf5, 0x3c, 0xf3, 0x06, 0x32, 0xa0, 0xd0, 0x74, 0x5b, 0xce, 0xa0, 0xdd, 0x37,
	0x35, 0xfb, 0x39, 0xec, 0x6f, 0xd9, 0x45, 0x45, 0xc8, 0x0e, 0x3c, 0x17, 0x73, 0x6c, 0x09, 0x72,
	0xe7, 0xb8, 0x3b, 0xe8, 0x99, 0x5a, 0x92, 0xec, 0x38, 0xde, 0x5b, 0x33, 0x93, 0x24, 0xbb, 0xfd,
	0x57, 0xbc, 0xae, 0xdb, 0x5f, 0xa1, 0xbc, 0x31, 0x5f, 0x82, 0xba, 0xe8, 0x5e, 0xb8, 0xb2, 0x8d,
	0xfb, 0xde, 0x6d, 0x0c, 0xfa, 0x2e, 0x27, 0x73, 0xca, 0x3b, 0xfc, 0x9a, 0x87, 0x19, 0xb4, 0x0f,
	0x65, 0x11, 0x5e, 0xae, 0xaa, 0x7a, 0x42, 0xc2, 0xae, 0xd3, 0x34, 0xb3, 0xc8, 0x84, 0xbd, 0x24,
	0x5a, 0xd7, 0x72, 0xa8, 0x02, 0x20, 0x32, 0x92, 0x9e, 0x47, 0x7b, 0x50, 0xec, 0xb9, 0xb8, 0x73,
	0xe9, 0xb4, 0xdb, 0x66, 0xc1, 0xfe, 0xae, 0x41, 0x85, 0xcf, 0xef, 0xc5, 0x7e, 0xbc, 0x60, 0x72,
	0x82, 0x5b, 0x90, 0xa3, 0x5f, 0xe6, 0x24, 0x12, 0x8b, 0x2e, 0x61, 0x79, 0x48, 0xb2, 0x57, 0x11,
	0x5d, 0x84, 0x62, 0x77, 0x3c, 0x2b, 0x0e, 0xe8, 0x10, 0xf2, 0x2c, 0x9e, 0x04, 0x1f, 0x97, 0x62,
	0x0f, 0x45, 0x9c, 0x9e, 0xd0, 0x53, 0x28, 0x10, 0xfe, 0x4a, 0x26, 0x84, 0xf1, 0x77, 0xab, 0x57,
	0x8d, 0xfa, 0xd1, 0x9f, 0x17, 0x84, 0x57, 0x50, 0xf4, 0x02, 0xe0, 0xd7, 0x76, 0xac, 0x1c, 0x5f,
	0x8a, 0x51, 0xbf, 0xbf, 0x41, 0x6c, 0xb1, 0xde, 0x1a, 0x20, 0xc9, 0x0a, 0xc3, 0x1e, 0xc1, 0xdd,
	0x0e, 0x1d, 0x4e, 0x46, 0xcb, 0x54, 0x9f, 0x6b, 0x62, 0xf2, 0x69, 0x41, 0x58, 0x2c, 0x9d, 0x99,
	0xa0, 0xb3, 0x28, 0x48, 0x7d, 0x25, 0x61, 0x32, 0x27, 0xff, 0x17, 0xbc, 0x90, 0x04, 0xdc, 0xd7,
	0xce, 0x39, 0x53, 0xa8, 0xfd, 0x00, 0xee, 0x6d, 0xf7, 0x61, 0x21, 0xff, 0xa2, 0xe4, 0xfe, 0xed,
	0x47, 0x70, 0x80, 0xc9, 0x8c, 0x7e, 0x26, 0x1c, 0x70, 0xfd, 0x04, 0xb6, 0x05, 0x87, 0x0a, 0x54,
	0x15, 0xe1, 0x6e, 0xd6, 0x95, 0xff, 0xec, 0x66, 0xbb, 0x8f, 0x3a, 0xc8, 0xe3, 0xd5, 0x20, 0x4d,
	0x32, 0xf2, 0x17, 0xd3, 0x78, 0xb7, 0xa9, 0xb5, 0xa4, 0xca, 0x50, 0x25, 0x3f, 0x00, 0xf2, 0xc8,
	0x6e, 0xa1, 0x7f, 0x74, 0x74, 0x00, 0x37, 0x57, 0xea, 0x6a, 0xd3, 0x13, 0xb0, 0xce, 0x45, 0x5a,
	0x7e, 0xed, 0x3b, 0x3c, 0xf4, 0xe0, 0xf6, 0x26, 0x5a, 0x91, 0x42, 0x4f, 0x20, 0x1f, 0x11, 0xc6,
	0x9d, 0x09, 0x86, 0x51, 0xbf, 0xf3, 0xfb, 0x58, 0xca, 0x0f, 0x85, 0x53, 0xe8, 0xd9, 0x33, 0x78,
	0x48, 0xa3, 0xab, 0x9a, 0x1f, 0xfa, 0xc1, 0x98, 0x6c, 0x10, 0xc4, 0x95, 0x1b, 0xd0, 0xf4, 0x52,
	0x3e, 0x2b, 0x71, 0x01, 0x41, 0x65, 0xdf, 0x34, 0xed, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a,
	0xf8, 0x89, 0x4e, 0xae, 0x05, 0x00, 0x00,
}
