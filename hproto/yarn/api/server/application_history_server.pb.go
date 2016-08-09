// Code generated by protoc-gen-go.
// source: application_history_server.proto
// DO NOT EDIT!

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	application_history_server.proto

It has these top-level messages:
	ApplicationHistoryDataProto
	ApplicationStartDataProto
	ApplicationFinishDataProto
	ApplicationAttemptHistoryDataProto
	ApplicationAttemptStartDataProto
	ApplicationAttemptFinishDataProto
	ContainerHistoryDataProto
	ContainerStartDataProto
	ContainerFinishDataProto
*/
package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hadoop_yarn "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApplicationHistoryDataProto struct {
	ApplicationId          *hadoop_yarn.ApplicationIdProto          `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ApplicationName        *string                                  `protobuf:"bytes,2,opt,name=application_name,json=applicationName" json:"application_name,omitempty"`
	ApplicationType        *string                                  `protobuf:"bytes,3,opt,name=application_type,json=applicationType" json:"application_type,omitempty"`
	User                   *string                                  `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Queue                  *string                                  `protobuf:"bytes,5,opt,name=queue" json:"queue,omitempty"`
	SubmitTime             *int64                                   `protobuf:"varint,6,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	StartTime              *int64                                   `protobuf:"varint,7,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	FinishTime             *int64                                   `protobuf:"varint,8,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo        *string                                  `protobuf:"bytes,9,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus *hadoop_yarn.FinalApplicationStatusProto `protobuf:"varint,10,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	YarnApplicationState   *hadoop_yarn.YarnApplicationStateProto   `protobuf:"varint,11,opt,name=yarn_application_state,json=yarnApplicationState,enum=hadoop.yarn.YarnApplicationStateProto" json:"yarn_application_state,omitempty"`
	XXX_unrecognized       []byte                                   `json:"-"`
}

func (m *ApplicationHistoryDataProto) Reset()                    { *m = ApplicationHistoryDataProto{} }
func (m *ApplicationHistoryDataProto) String() string            { return proto.CompactTextString(m) }
func (*ApplicationHistoryDataProto) ProtoMessage()               {}
func (*ApplicationHistoryDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApplicationHistoryDataProto) GetApplicationId() *hadoop_yarn.ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *ApplicationHistoryDataProto) GetApplicationName() string {
	if m != nil && m.ApplicationName != nil {
		return *m.ApplicationName
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetApplicationType() string {
	if m != nil && m.ApplicationType != nil {
		return *m.ApplicationType
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetQueue() string {
	if m != nil && m.Queue != nil {
		return *m.Queue
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetSubmitTime() int64 {
	if m != nil && m.SubmitTime != nil {
		return *m.SubmitTime
	}
	return 0
}

func (m *ApplicationHistoryDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *ApplicationHistoryDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ApplicationHistoryDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetFinalApplicationStatus() hadoop_yarn.FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return hadoop_yarn.FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationHistoryDataProto) GetYarnApplicationState() hadoop_yarn.YarnApplicationStateProto {
	if m != nil && m.YarnApplicationState != nil {
		return *m.YarnApplicationState
	}
	return hadoop_yarn.YarnApplicationStateProto_NEW
}

type ApplicationStartDataProto struct {
	ApplicationId    *hadoop_yarn.ApplicationIdProto `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ApplicationName  *string                         `protobuf:"bytes,2,opt,name=application_name,json=applicationName" json:"application_name,omitempty"`
	ApplicationType  *string                         `protobuf:"bytes,3,opt,name=application_type,json=applicationType" json:"application_type,omitempty"`
	User             *string                         `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Queue            *string                         `protobuf:"bytes,5,opt,name=queue" json:"queue,omitempty"`
	SubmitTime       *int64                          `protobuf:"varint,6,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	StartTime        *int64                          `protobuf:"varint,7,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ApplicationStartDataProto) Reset()                    { *m = ApplicationStartDataProto{} }
func (m *ApplicationStartDataProto) String() string            { return proto.CompactTextString(m) }
func (*ApplicationStartDataProto) ProtoMessage()               {}
func (*ApplicationStartDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApplicationStartDataProto) GetApplicationId() *hadoop_yarn.ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *ApplicationStartDataProto) GetApplicationName() string {
	if m != nil && m.ApplicationName != nil {
		return *m.ApplicationName
	}
	return ""
}

func (m *ApplicationStartDataProto) GetApplicationType() string {
	if m != nil && m.ApplicationType != nil {
		return *m.ApplicationType
	}
	return ""
}

func (m *ApplicationStartDataProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *ApplicationStartDataProto) GetQueue() string {
	if m != nil && m.Queue != nil {
		return *m.Queue
	}
	return ""
}

func (m *ApplicationStartDataProto) GetSubmitTime() int64 {
	if m != nil && m.SubmitTime != nil {
		return *m.SubmitTime
	}
	return 0
}

func (m *ApplicationStartDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

type ApplicationFinishDataProto struct {
	ApplicationId          *hadoop_yarn.ApplicationIdProto          `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	FinishTime             *int64                                   `protobuf:"varint,2,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo        *string                                  `protobuf:"bytes,3,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus *hadoop_yarn.FinalApplicationStatusProto `protobuf:"varint,4,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	YarnApplicationState   *hadoop_yarn.YarnApplicationStateProto   `protobuf:"varint,5,opt,name=yarn_application_state,json=yarnApplicationState,enum=hadoop.yarn.YarnApplicationStateProto" json:"yarn_application_state,omitempty"`
	XXX_unrecognized       []byte                                   `json:"-"`
}

func (m *ApplicationFinishDataProto) Reset()                    { *m = ApplicationFinishDataProto{} }
func (m *ApplicationFinishDataProto) String() string            { return proto.CompactTextString(m) }
func (*ApplicationFinishDataProto) ProtoMessage()               {}
func (*ApplicationFinishDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApplicationFinishDataProto) GetApplicationId() *hadoop_yarn.ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *ApplicationFinishDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ApplicationFinishDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationFinishDataProto) GetFinalApplicationStatus() hadoop_yarn.FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return hadoop_yarn.FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationFinishDataProto) GetYarnApplicationState() hadoop_yarn.YarnApplicationStateProto {
	if m != nil && m.YarnApplicationState != nil {
		return *m.YarnApplicationState
	}
	return hadoop_yarn.YarnApplicationStateProto_NEW
}

type ApplicationAttemptHistoryDataProto struct {
	ApplicationAttemptId        *hadoop_yarn.ApplicationAttemptIdProto        `protobuf:"bytes,1,opt,name=application_attempt_id,json=applicationAttemptId" json:"application_attempt_id,omitempty"`
	Host                        *string                                       `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	RpcPort                     *int32                                        `protobuf:"varint,3,opt,name=rpc_port,json=rpcPort" json:"rpc_port,omitempty"`
	TrackingUrl                 *string                                       `protobuf:"bytes,4,opt,name=tracking_url,json=trackingUrl" json:"tracking_url,omitempty"`
	DiagnosticsInfo             *string                                       `protobuf:"bytes,5,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus      *hadoop_yarn.FinalApplicationStatusProto      `protobuf:"varint,6,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	MasterContainerId           *hadoop_yarn.ContainerIdProto                 `protobuf:"bytes,7,opt,name=master_container_id,json=masterContainerId" json:"master_container_id,omitempty"`
	YarnApplicationAttemptState *hadoop_yarn.YarnApplicationAttemptStateProto `protobuf:"varint,8,opt,name=yarn_application_attempt_state,json=yarnApplicationAttemptState,enum=hadoop.yarn.YarnApplicationAttemptStateProto" json:"yarn_application_attempt_state,omitempty"`
	XXX_unrecognized            []byte                                        `json:"-"`
}

func (m *ApplicationAttemptHistoryDataProto) Reset()         { *m = ApplicationAttemptHistoryDataProto{} }
func (m *ApplicationAttemptHistoryDataProto) String() string { return proto.CompactTextString(m) }
func (*ApplicationAttemptHistoryDataProto) ProtoMessage()    {}
func (*ApplicationAttemptHistoryDataProto) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

func (m *ApplicationAttemptHistoryDataProto) GetApplicationAttemptId() *hadoop_yarn.ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

func (m *ApplicationAttemptHistoryDataProto) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *ApplicationAttemptHistoryDataProto) GetRpcPort() int32 {
	if m != nil && m.RpcPort != nil {
		return *m.RpcPort
	}
	return 0
}

func (m *ApplicationAttemptHistoryDataProto) GetTrackingUrl() string {
	if m != nil && m.TrackingUrl != nil {
		return *m.TrackingUrl
	}
	return ""
}

func (m *ApplicationAttemptHistoryDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationAttemptHistoryDataProto) GetFinalApplicationStatus() hadoop_yarn.FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return hadoop_yarn.FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationAttemptHistoryDataProto) GetMasterContainerId() *hadoop_yarn.ContainerIdProto {
	if m != nil {
		return m.MasterContainerId
	}
	return nil
}

func (m *ApplicationAttemptHistoryDataProto) GetYarnApplicationAttemptState() hadoop_yarn.YarnApplicationAttemptStateProto {
	if m != nil && m.YarnApplicationAttemptState != nil {
		return *m.YarnApplicationAttemptState
	}
	return hadoop_yarn.YarnApplicationAttemptStateProto_APP_ATTEMPT_NEW
}

type ApplicationAttemptStartDataProto struct {
	ApplicationAttemptId *hadoop_yarn.ApplicationAttemptIdProto `protobuf:"bytes,1,opt,name=application_attempt_id,json=applicationAttemptId" json:"application_attempt_id,omitempty"`
	Host                 *string                                `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	RpcPort              *int32                                 `protobuf:"varint,3,opt,name=rpc_port,json=rpcPort" json:"rpc_port,omitempty"`
	MasterContainerId    *hadoop_yarn.ContainerIdProto          `protobuf:"bytes,4,opt,name=master_container_id,json=masterContainerId" json:"master_container_id,omitempty"`
	XXX_unrecognized     []byte                                 `json:"-"`
}

func (m *ApplicationAttemptStartDataProto) Reset()         { *m = ApplicationAttemptStartDataProto{} }
func (m *ApplicationAttemptStartDataProto) String() string { return proto.CompactTextString(m) }
func (*ApplicationAttemptStartDataProto) ProtoMessage()    {}
func (*ApplicationAttemptStartDataProto) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4}
}

func (m *ApplicationAttemptStartDataProto) GetApplicationAttemptId() *hadoop_yarn.ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

func (m *ApplicationAttemptStartDataProto) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *ApplicationAttemptStartDataProto) GetRpcPort() int32 {
	if m != nil && m.RpcPort != nil {
		return *m.RpcPort
	}
	return 0
}

func (m *ApplicationAttemptStartDataProto) GetMasterContainerId() *hadoop_yarn.ContainerIdProto {
	if m != nil {
		return m.MasterContainerId
	}
	return nil
}

type ApplicationAttemptFinishDataProto struct {
	ApplicationAttemptId        *hadoop_yarn.ApplicationAttemptIdProto        `protobuf:"bytes,1,opt,name=application_attempt_id,json=applicationAttemptId" json:"application_attempt_id,omitempty"`
	TrackingUrl                 *string                                       `protobuf:"bytes,2,opt,name=tracking_url,json=trackingUrl" json:"tracking_url,omitempty"`
	DiagnosticsInfo             *string                                       `protobuf:"bytes,3,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus      *hadoop_yarn.FinalApplicationStatusProto      `protobuf:"varint,4,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	YarnApplicationAttemptState *hadoop_yarn.YarnApplicationAttemptStateProto `protobuf:"varint,5,opt,name=yarn_application_attempt_state,json=yarnApplicationAttemptState,enum=hadoop.yarn.YarnApplicationAttemptStateProto" json:"yarn_application_attempt_state,omitempty"`
	XXX_unrecognized            []byte                                        `json:"-"`
}

func (m *ApplicationAttemptFinishDataProto) Reset()         { *m = ApplicationAttemptFinishDataProto{} }
func (m *ApplicationAttemptFinishDataProto) String() string { return proto.CompactTextString(m) }
func (*ApplicationAttemptFinishDataProto) ProtoMessage()    {}
func (*ApplicationAttemptFinishDataProto) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *ApplicationAttemptFinishDataProto) GetApplicationAttemptId() *hadoop_yarn.ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

func (m *ApplicationAttemptFinishDataProto) GetTrackingUrl() string {
	if m != nil && m.TrackingUrl != nil {
		return *m.TrackingUrl
	}
	return ""
}

func (m *ApplicationAttemptFinishDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationAttemptFinishDataProto) GetFinalApplicationStatus() hadoop_yarn.FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return hadoop_yarn.FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationAttemptFinishDataProto) GetYarnApplicationAttemptState() hadoop_yarn.YarnApplicationAttemptStateProto {
	if m != nil && m.YarnApplicationAttemptState != nil {
		return *m.YarnApplicationAttemptState
	}
	return hadoop_yarn.YarnApplicationAttemptStateProto_APP_ATTEMPT_NEW
}

type ContainerHistoryDataProto struct {
	ContainerId         *hadoop_yarn.ContainerIdProto    `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	AllocatedResource   *hadoop_yarn.ResourceProto       `protobuf:"bytes,2,opt,name=allocated_resource,json=allocatedResource" json:"allocated_resource,omitempty"`
	AssignedNodeId      *hadoop_yarn.NodeIdProto         `protobuf:"bytes,3,opt,name=assigned_node_id,json=assignedNodeId" json:"assigned_node_id,omitempty"`
	Priority            *hadoop_yarn.PriorityProto       `protobuf:"bytes,4,opt,name=priority" json:"priority,omitempty"`
	StartTime           *int64                           `protobuf:"varint,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	FinishTime          *int64                           `protobuf:"varint,6,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo     *string                          `protobuf:"bytes,7,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	ContainerExitStatus *int32                           `protobuf:"varint,8,opt,name=container_exit_status,json=containerExitStatus" json:"container_exit_status,omitempty"`
	ContainerState      *hadoop_yarn.ContainerStateProto `protobuf:"varint,9,opt,name=container_state,json=containerState,enum=hadoop.yarn.ContainerStateProto" json:"container_state,omitempty"`
	XXX_unrecognized    []byte                           `json:"-"`
}

func (m *ContainerHistoryDataProto) Reset()                    { *m = ContainerHistoryDataProto{} }
func (m *ContainerHistoryDataProto) String() string            { return proto.CompactTextString(m) }
func (*ContainerHistoryDataProto) ProtoMessage()               {}
func (*ContainerHistoryDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ContainerHistoryDataProto) GetContainerId() *hadoop_yarn.ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetAllocatedResource() *hadoop_yarn.ResourceProto {
	if m != nil {
		return m.AllocatedResource
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetAssignedNodeId() *hadoop_yarn.NodeIdProto {
	if m != nil {
		return m.AssignedNodeId
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetPriority() *hadoop_yarn.PriorityProto {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *ContainerHistoryDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ContainerHistoryDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ContainerHistoryDataProto) GetContainerExitStatus() int32 {
	if m != nil && m.ContainerExitStatus != nil {
		return *m.ContainerExitStatus
	}
	return 0
}

func (m *ContainerHistoryDataProto) GetContainerState() hadoop_yarn.ContainerStateProto {
	if m != nil && m.ContainerState != nil {
		return *m.ContainerState
	}
	return hadoop_yarn.ContainerStateProto_C_NEW
}

type ContainerStartDataProto struct {
	ContainerId       *hadoop_yarn.ContainerIdProto `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	AllocatedResource *hadoop_yarn.ResourceProto    `protobuf:"bytes,2,opt,name=allocated_resource,json=allocatedResource" json:"allocated_resource,omitempty"`
	AssignedNodeId    *hadoop_yarn.NodeIdProto      `protobuf:"bytes,3,opt,name=assigned_node_id,json=assignedNodeId" json:"assigned_node_id,omitempty"`
	Priority          *hadoop_yarn.PriorityProto    `protobuf:"bytes,4,opt,name=priority" json:"priority,omitempty"`
	StartTime         *int64                        `protobuf:"varint,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	XXX_unrecognized  []byte                        `json:"-"`
}

func (m *ContainerStartDataProto) Reset()                    { *m = ContainerStartDataProto{} }
func (m *ContainerStartDataProto) String() string            { return proto.CompactTextString(m) }
func (*ContainerStartDataProto) ProtoMessage()               {}
func (*ContainerStartDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ContainerStartDataProto) GetContainerId() *hadoop_yarn.ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *ContainerStartDataProto) GetAllocatedResource() *hadoop_yarn.ResourceProto {
	if m != nil {
		return m.AllocatedResource
	}
	return nil
}

func (m *ContainerStartDataProto) GetAssignedNodeId() *hadoop_yarn.NodeIdProto {
	if m != nil {
		return m.AssignedNodeId
	}
	return nil
}

func (m *ContainerStartDataProto) GetPriority() *hadoop_yarn.PriorityProto {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *ContainerStartDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

type ContainerFinishDataProto struct {
	ContainerId         *hadoop_yarn.ContainerIdProto    `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	FinishTime          *int64                           `protobuf:"varint,2,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo     *string                          `protobuf:"bytes,3,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	ContainerExitStatus *int32                           `protobuf:"varint,4,opt,name=container_exit_status,json=containerExitStatus" json:"container_exit_status,omitempty"`
	ContainerState      *hadoop_yarn.ContainerStateProto `protobuf:"varint,5,opt,name=container_state,json=containerState,enum=hadoop.yarn.ContainerStateProto" json:"container_state,omitempty"`
	XXX_unrecognized    []byte                           `json:"-"`
}

func (m *ContainerFinishDataProto) Reset()                    { *m = ContainerFinishDataProto{} }
func (m *ContainerFinishDataProto) String() string            { return proto.CompactTextString(m) }
func (*ContainerFinishDataProto) ProtoMessage()               {}
func (*ContainerFinishDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ContainerFinishDataProto) GetContainerId() *hadoop_yarn.ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *ContainerFinishDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ContainerFinishDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ContainerFinishDataProto) GetContainerExitStatus() int32 {
	if m != nil && m.ContainerExitStatus != nil {
		return *m.ContainerExitStatus
	}
	return 0
}

func (m *ContainerFinishDataProto) GetContainerState() hadoop_yarn.ContainerStateProto {
	if m != nil && m.ContainerState != nil {
		return *m.ContainerState
	}
	return hadoop_yarn.ContainerStateProto_C_NEW
}

func init() {
	proto.RegisterType((*ApplicationHistoryDataProto)(nil), "hadoop.yarn.ApplicationHistoryDataProto")
	proto.RegisterType((*ApplicationStartDataProto)(nil), "hadoop.yarn.ApplicationStartDataProto")
	proto.RegisterType((*ApplicationFinishDataProto)(nil), "hadoop.yarn.ApplicationFinishDataProto")
	proto.RegisterType((*ApplicationAttemptHistoryDataProto)(nil), "hadoop.yarn.ApplicationAttemptHistoryDataProto")
	proto.RegisterType((*ApplicationAttemptStartDataProto)(nil), "hadoop.yarn.ApplicationAttemptStartDataProto")
	proto.RegisterType((*ApplicationAttemptFinishDataProto)(nil), "hadoop.yarn.ApplicationAttemptFinishDataProto")
	proto.RegisterType((*ContainerHistoryDataProto)(nil), "hadoop.yarn.ContainerHistoryDataProto")
	proto.RegisterType((*ContainerStartDataProto)(nil), "hadoop.yarn.ContainerStartDataProto")
	proto.RegisterType((*ContainerFinishDataProto)(nil), "hadoop.yarn.ContainerFinishDataProto")
}

func init() { proto.RegisterFile("application_history_server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x97, 0xc1, 0x4e, 0xdb, 0x48,
	0x18, 0xc7, 0xe5, 0xc4, 0x26, 0xc9, 0x17, 0x36, 0x2c, 0x86, 0x65, 0x0d, 0x2c, 0x10, 0x7c, 0x40,
	0xec, 0x61, 0x73, 0xc8, 0x61, 0xcf, 0x0b, 0xcb, 0xa2, 0xcd, 0xa1, 0x08, 0x05, 0x7a, 0xa8, 0x54,
	0xc9, 0x1a, 0x9c, 0x21, 0x19, 0x35, 0xf1, 0xb8, 0xe3, 0x49, 0xd5, 0xbc, 0x41, 0x4f, 0x7d, 0x86,
	0x4a, 0x95, 0xca, 0x03, 0xf4, 0x09, 0xfa, 0x4a, 0x7d, 0x80, 0xaa, 0x93, 0x19, 0x3b, 0x19, 0x3b,
	0x49, 0x09, 0x34, 0xa8, 0x3d, 0x70, 0x89, 0xe2, 0xff, 0xff, 0xfb, 0xbe, 0x19, 0x7f, 0xdf, 0x2f,
	0xe3, 0x18, 0xaa, 0x28, 0x0c, 0xbb, 0xc4, 0x47, 0x9c, 0xd0, 0xc0, 0xeb, 0x90, 0x88, 0x53, 0x36,
	0xf0, 0x22, 0xcc, 0x5e, 0x61, 0x56, 0x0b, 0x19, 0xe5, 0xd4, 0x2e, 0x77, 0x50, 0x8b, 0xd2, 0xb0,
	0x36, 0x40, 0x2c, 0xd8, 0x5a, 0x1d, 0x7e, 0x7a, 0xd2, 0x88, 0x94, 0xef, 0x7e, 0x34, 0x61, 0xfb,
	0x68, 0x5c, 0xe4, 0x7f, 0x55, 0xe3, 0x04, 0x71, 0x74, 0x2e, 0xf3, 0x4f, 0xa1, 0xa2, 0xaf, 0x41,
	0x5a, 0x8e, 0x51, 0x35, 0x0e, 0xcb, 0xf5, 0xbd, 0x9a, 0x56, 0xb8, 0xa6, 0x55, 0x68, 0xb4, 0x64,
	0x62, 0xf3, 0x17, 0xa4, 0x6b, 0xf6, 0x9f, 0xf0, 0xab, 0x5e, 0x27, 0x40, 0x3d, 0xec, 0xe4, 0x44,
	0xa5, 0x52, 0x73, 0x45, 0xd3, 0xcf, 0x84, 0x9c, 0x0d, 0xe5, 0x83, 0x10, 0x3b, 0xf9, 0x89, 0xd0,
	0x4b, 0x21, 0xdb, 0x36, 0x98, 0x7d, 0x71, 0xbb, 0x8e, 0x29, 0x6d, 0xf9, 0xdd, 0x5e, 0x07, 0xeb,
	0x65, 0x1f, 0xf7, 0xb1, 0x63, 0x49, 0x51, 0x5d, 0xd8, 0x7b, 0x50, 0x8e, 0xfa, 0x57, 0x3d, 0xc2,
	0x3d, 0x4e, 0xc4, 0xd2, 0x4b, 0xc2, 0xcb, 0x37, 0x41, 0x49, 0x97, 0x42, 0xb1, 0x77, 0x00, 0x22,
	0x8e, 0x58, 0xec, 0x17, 0xa4, 0x5f, 0x92, 0x8a, 0xb4, 0x45, 0xfe, 0x35, 0x09, 0x48, 0xd4, 0x51,
	0x7e, 0x51, 0xe5, 0x2b, 0x49, 0x06, 0x88, 0x5d, 0xb7, 0x08, 0x6a, 0x07, 0x34, 0xe2, 0xc4, 0x8f,
	0x3c, 0x12, 0x5c, 0x53, 0xa7, 0xa4, 0x76, 0xad, 0xe9, 0x0d, 0x21, 0xdb, 0x57, 0xe0, 0x88, 0x44,
	0xd4, 0xf5, 0xf4, 0xdb, 0x14, 0x4b, 0xf1, 0x7e, 0xe4, 0x80, 0x48, 0xa9, 0xd4, 0x0f, 0x53, 0xdd,
	0x3d, 0x1d, 0x06, 0x6b, 0x2d, 0xbe, 0x90, 0xa1, 0xaa, 0xcd, 0x1b, 0xd7, 0x53, 0x4d, 0xfb, 0x39,
	0x6c, 0xc8, 0x61, 0x67, 0x97, 0xc0, 0x4e, 0x59, 0xae, 0x70, 0x90, 0x5a, 0xe1, 0x99, 0xf8, 0xc8,
	0xd4, 0xc0, 0xaa, 0xfe, 0xfa, 0x60, 0x8a, 0xe5, 0xbe, 0xcf, 0xc1, 0x66, 0x5a, 0x64, 0xfc, 0x91,
	0x99, 0x0c, 0x33, 0xee, 0xe7, 0x1c, 0x6c, 0x69, 0x77, 0x79, 0x2a, 0x61, 0x59, 0x7c, 0x9b, 0x32,
	0x68, 0xe6, 0xe6, 0x42, 0x33, 0x7f, 0x77, 0x34, 0xcd, 0x07, 0x47, 0xd3, 0x5a, 0x00, 0x9a, 0x1f,
	0x4c, 0x70, 0x35, 0xf1, 0x88, 0x73, 0xdc, 0x0b, 0xf9, 0xc4, 0xb9, 0x26, 0x36, 0xa1, 0xaf, 0x8f,
	0x54, 0xd8, 0x78, 0x08, 0x07, 0xb3, 0x86, 0x10, 0x17, 0x4c, 0x66, 0xb1, 0x8e, 0xa6, 0x58, 0x43,
	0xc6, 0x3a, 0xa2, 0xab, 0x31, 0xad, 0xf2, 0xbb, 0xbd, 0x09, 0x45, 0x16, 0xfa, 0x5e, 0x48, 0x19,
	0x97, 0xdd, 0xb7, 0x9a, 0x05, 0x71, 0x7d, 0x2e, 0x2e, 0xed, 0x7d, 0x58, 0xe6, 0x0c, 0xf9, 0x2f,
	0x48, 0xd0, 0xf6, 0xfa, 0xac, 0x1b, 0xa3, 0x59, 0x4e, 0xb4, 0xa7, 0xac, 0x3b, 0x75, 0x86, 0xd6,
	0xdd, 0x67, 0xb8, 0xb4, 0xa0, 0x19, 0x3e, 0x81, 0xb5, 0x1e, 0x8a, 0x38, 0x66, 0x9e, 0x4f, 0x03,
	0x8e, 0x48, 0x20, 0xbe, 0x89, 0xde, 0x15, 0x64, 0xef, 0x76, 0x52, 0xe5, 0xff, 0x4d, 0x02, 0x92,
	0x96, 0xad, 0xaa, 0x4c, 0x4d, 0xb7, 0x19, 0xec, 0x4e, 0x20, 0x91, 0x8c, 0x44, 0xa1, 0x51, 0x94,
	0x1b, 0xff, 0xeb, 0x5b, 0x68, 0xc4, 0xed, 0xd7, 0x08, 0xd9, 0x1e, 0xcc, 0x8e, 0x70, 0xbf, 0x18,
	0x50, 0x9d, 0xea, 0xe9, 0x47, 0xd9, 0x4f, 0x85, 0xc9, 0x8c, 0xa6, 0x9b, 0xf7, 0x6b, 0xba, 0x7b,
	0x93, 0x87, 0xfd, 0xc9, 0x1d, 0x67, 0x4f, 0xa9, 0x87, 0xed, 0x40, 0x96, 0xfc, 0xdc, 0x7c, 0xe4,
	0xff, 0xc0, 0xd3, 0xeb, 0x76, 0x54, 0xad, 0x85, 0xa3, 0xfa, 0xd6, 0x84, 0xcd, 0xd1, 0xe4, 0x26,
	0x8e, 0xb2, 0x7f, 0x60, 0x39, 0xc5, 0x83, 0x31, 0x0f, 0x0f, 0x65, 0x5f, 0xfb, 0xf9, 0x35, 0xc0,
	0x46, 0xdd, 0x2e, 0x15, 0x2b, 0xe3, 0x96, 0xc7, 0x70, 0x44, 0xfb, 0xcc, 0x57, 0x0f, 0x92, 0x72,
	0x7d, 0x2b, 0x55, 0xa7, 0x19, 0x9b, 0x31, 0x54, 0xa3, 0xac, 0x44, 0xb7, 0x8f, 0xc5, 0x83, 0x38,
	0x8a, 0x48, 0x3b, 0x10, 0x95, 0x02, 0xda, 0xc2, 0xc3, 0x0d, 0xe5, 0x65, 0x21, 0x27, 0x55, 0xe8,
	0x4c, 0x78, 0xc9, 0x5e, 0x2a, 0x49, 0x86, 0x12, 0xed, 0xbf, 0xa1, 0x18, 0x32, 0x42, 0x19, 0xe1,
	0x83, 0x18, 0xee, 0xf4, 0x26, 0xce, 0x63, 0x53, 0x65, 0x8f, 0x62, 0x33, 0x8f, 0x63, 0xeb, 0x96,
	0xbf, 0x70, 0x4b, 0x73, 0x3d, 0x27, 0x0b, 0xd3, 0x49, 0xab, 0xc3, 0x6f, 0xe3, 0x9e, 0xe3, 0xd7,
	0x84, 0x27, 0x98, 0x15, 0xe5, 0x4f, 0x76, 0x6d, 0x64, 0xfe, 0x27, 0xbc, 0x98, 0x9c, 0x06, 0xac,
	0x8c, 0x73, 0x14, 0x2a, 0x25, 0x89, 0x4a, 0x75, 0xfa, 0xa8, 0x34, 0x3a, 0x2a, 0x7e, 0x4a, 0x74,
	0x3f, 0xe5, 0xe0, 0x77, 0x3d, 0x4e, 0x3f, 0xb2, 0x1e, 0x71, 0x98, 0x0b, 0x07, 0xf7, 0x26, 0x07,
	0xce, 0xa8, 0x0f, 0xd9, 0x53, 0xef, 0xfb, 0x9b, 0xb8, 0xc8, 0x7f, 0x65, 0x33, 0x69, 0x33, 0xef,
	0x44, 0x9b, 0x75, 0x3f, 0xda, 0x8e, 0x4f, 0xe0, 0x0f, 0xca, 0xda, 0x35, 0x14, 0x22, 0xbf, 0x83,
	0x53, 0xd9, 0xf2, 0x1d, 0xf2, 0x78, 0x77, 0xf2, 0x05, 0xf2, 0x42, 0xbe, 0x83, 0xca, 0x72, 0xd1,
	0x1b, 0xc3, 0x78, 0x67, 0x18, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0xe8, 0xd2, 0x5d, 0xac,
	0x0e, 0x00, 0x00,
}
