// Code generated by protoc-gen-go.
// source: MRClientProtocol.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hadoop_common "."

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MRClientProtocolService service

type MRClientProtocolServiceClient interface {
	GetJobReport(ctx context.Context, in *GetJobReportRequestProto, opts ...grpc.CallOption) (*GetJobReportResponseProto, error)
	GetTaskReport(ctx context.Context, in *GetTaskReportRequestProto, opts ...grpc.CallOption) (*GetTaskReportResponseProto, error)
	GetTaskAttemptReport(ctx context.Context, in *GetTaskAttemptReportRequestProto, opts ...grpc.CallOption) (*GetTaskAttemptReportResponseProto, error)
	GetCounters(ctx context.Context, in *GetCountersRequestProto, opts ...grpc.CallOption) (*GetCountersResponseProto, error)
	GetTaskAttemptCompletionEvents(ctx context.Context, in *GetTaskAttemptCompletionEventsRequestProto, opts ...grpc.CallOption) (*GetTaskAttemptCompletionEventsResponseProto, error)
	GetTaskReports(ctx context.Context, in *GetTaskReportsRequestProto, opts ...grpc.CallOption) (*GetTaskReportsResponseProto, error)
	GetDiagnostics(ctx context.Context, in *GetDiagnosticsRequestProto, opts ...grpc.CallOption) (*GetDiagnosticsResponseProto, error)
	GetDelegationToken(ctx context.Context, in *hadoop_common.GetDelegationTokenRequestProto, opts ...grpc.CallOption) (*hadoop_common.GetDelegationTokenResponseProto, error)
	KillJob(ctx context.Context, in *KillJobRequestProto, opts ...grpc.CallOption) (*KillJobResponseProto, error)
	KillTask(ctx context.Context, in *KillTaskRequestProto, opts ...grpc.CallOption) (*KillTaskResponseProto, error)
	KillTaskAttempt(ctx context.Context, in *KillTaskAttemptRequestProto, opts ...grpc.CallOption) (*KillTaskAttemptResponseProto, error)
	FailTaskAttempt(ctx context.Context, in *FailTaskAttemptRequestProto, opts ...grpc.CallOption) (*FailTaskAttemptResponseProto, error)
	RenewDelegationToken(ctx context.Context, in *hadoop_common.RenewDelegationTokenRequestProto, opts ...grpc.CallOption) (*hadoop_common.RenewDelegationTokenResponseProto, error)
	CancelDelegationToken(ctx context.Context, in *hadoop_common.CancelDelegationTokenRequestProto, opts ...grpc.CallOption) (*hadoop_common.CancelDelegationTokenResponseProto, error)
}

type mRClientProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewMRClientProtocolServiceClient(cc *grpc.ClientConn) MRClientProtocolServiceClient {
	return &mRClientProtocolServiceClient{cc}
}

func (c *mRClientProtocolServiceClient) GetJobReport(ctx context.Context, in *GetJobReportRequestProto, opts ...grpc.CallOption) (*GetJobReportResponseProto, error) {
	out := new(GetJobReportResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getJobReport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) GetTaskReport(ctx context.Context, in *GetTaskReportRequestProto, opts ...grpc.CallOption) (*GetTaskReportResponseProto, error) {
	out := new(GetTaskReportResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getTaskReport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) GetTaskAttemptReport(ctx context.Context, in *GetTaskAttemptReportRequestProto, opts ...grpc.CallOption) (*GetTaskAttemptReportResponseProto, error) {
	out := new(GetTaskAttemptReportResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getTaskAttemptReport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) GetCounters(ctx context.Context, in *GetCountersRequestProto, opts ...grpc.CallOption) (*GetCountersResponseProto, error) {
	out := new(GetCountersResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getCounters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) GetTaskAttemptCompletionEvents(ctx context.Context, in *GetTaskAttemptCompletionEventsRequestProto, opts ...grpc.CallOption) (*GetTaskAttemptCompletionEventsResponseProto, error) {
	out := new(GetTaskAttemptCompletionEventsResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getTaskAttemptCompletionEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) GetTaskReports(ctx context.Context, in *GetTaskReportsRequestProto, opts ...grpc.CallOption) (*GetTaskReportsResponseProto, error) {
	out := new(GetTaskReportsResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getTaskReports", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) GetDiagnostics(ctx context.Context, in *GetDiagnosticsRequestProto, opts ...grpc.CallOption) (*GetDiagnosticsResponseProto, error) {
	out := new(GetDiagnosticsResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getDiagnostics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) GetDelegationToken(ctx context.Context, in *hadoop_common.GetDelegationTokenRequestProto, opts ...grpc.CallOption) (*hadoop_common.GetDelegationTokenResponseProto, error) {
	out := new(hadoop_common.GetDelegationTokenResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/getDelegationToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) KillJob(ctx context.Context, in *KillJobRequestProto, opts ...grpc.CallOption) (*KillJobResponseProto, error) {
	out := new(KillJobResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/killJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) KillTask(ctx context.Context, in *KillTaskRequestProto, opts ...grpc.CallOption) (*KillTaskResponseProto, error) {
	out := new(KillTaskResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/killTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) KillTaskAttempt(ctx context.Context, in *KillTaskAttemptRequestProto, opts ...grpc.CallOption) (*KillTaskAttemptResponseProto, error) {
	out := new(KillTaskAttemptResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/killTaskAttempt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) FailTaskAttempt(ctx context.Context, in *FailTaskAttemptRequestProto, opts ...grpc.CallOption) (*FailTaskAttemptResponseProto, error) {
	out := new(FailTaskAttemptResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/failTaskAttempt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) RenewDelegationToken(ctx context.Context, in *hadoop_common.RenewDelegationTokenRequestProto, opts ...grpc.CallOption) (*hadoop_common.RenewDelegationTokenResponseProto, error) {
	out := new(hadoop_common.RenewDelegationTokenResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/renewDelegationToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mRClientProtocolServiceClient) CancelDelegationToken(ctx context.Context, in *hadoop_common.CancelDelegationTokenRequestProto, opts ...grpc.CallOption) (*hadoop_common.CancelDelegationTokenResponseProto, error) {
	out := new(hadoop_common.CancelDelegationTokenResponseProto)
	err := grpc.Invoke(ctx, "/hadoop.mapreduce.MRClientProtocolService/cancelDelegationToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MRClientProtocolService service

type MRClientProtocolServiceServer interface {
	GetJobReport(context.Context, *GetJobReportRequestProto) (*GetJobReportResponseProto, error)
	GetTaskReport(context.Context, *GetTaskReportRequestProto) (*GetTaskReportResponseProto, error)
	GetTaskAttemptReport(context.Context, *GetTaskAttemptReportRequestProto) (*GetTaskAttemptReportResponseProto, error)
	GetCounters(context.Context, *GetCountersRequestProto) (*GetCountersResponseProto, error)
	GetTaskAttemptCompletionEvents(context.Context, *GetTaskAttemptCompletionEventsRequestProto) (*GetTaskAttemptCompletionEventsResponseProto, error)
	GetTaskReports(context.Context, *GetTaskReportsRequestProto) (*GetTaskReportsResponseProto, error)
	GetDiagnostics(context.Context, *GetDiagnosticsRequestProto) (*GetDiagnosticsResponseProto, error)
	GetDelegationToken(context.Context, *hadoop_common.GetDelegationTokenRequestProto) (*hadoop_common.GetDelegationTokenResponseProto, error)
	KillJob(context.Context, *KillJobRequestProto) (*KillJobResponseProto, error)
	KillTask(context.Context, *KillTaskRequestProto) (*KillTaskResponseProto, error)
	KillTaskAttempt(context.Context, *KillTaskAttemptRequestProto) (*KillTaskAttemptResponseProto, error)
	FailTaskAttempt(context.Context, *FailTaskAttemptRequestProto) (*FailTaskAttemptResponseProto, error)
	RenewDelegationToken(context.Context, *hadoop_common.RenewDelegationTokenRequestProto) (*hadoop_common.RenewDelegationTokenResponseProto, error)
	CancelDelegationToken(context.Context, *hadoop_common.CancelDelegationTokenRequestProto) (*hadoop_common.CancelDelegationTokenResponseProto, error)
}

func RegisterMRClientProtocolServiceServer(s *grpc.Server, srv MRClientProtocolServiceServer) {
	s.RegisterService(&_MRClientProtocolService_serviceDesc, srv)
}

func _MRClientProtocolService_GetJobReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobReportRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetJobReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetJobReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetJobReport(ctx, req.(*GetJobReportRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_GetTaskReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskReportRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetTaskReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetTaskReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetTaskReport(ctx, req.(*GetTaskReportRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_GetTaskAttemptReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskAttemptReportRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetTaskAttemptReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetTaskAttemptReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetTaskAttemptReport(ctx, req.(*GetTaskAttemptReportRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_GetCounters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountersRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetCounters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetCounters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetCounters(ctx, req.(*GetCountersRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_GetTaskAttemptCompletionEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskAttemptCompletionEventsRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetTaskAttemptCompletionEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetTaskAttemptCompletionEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetTaskAttemptCompletionEvents(ctx, req.(*GetTaskAttemptCompletionEventsRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_GetTaskReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskReportsRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetTaskReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetTaskReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetTaskReports(ctx, req.(*GetTaskReportsRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_GetDiagnostics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiagnosticsRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetDiagnostics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetDiagnostics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetDiagnostics(ctx, req.(*GetDiagnosticsRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_GetDelegationToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_common.GetDelegationTokenRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).GetDelegationToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/GetDelegationToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).GetDelegationToken(ctx, req.(*hadoop_common.GetDelegationTokenRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_KillJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillJobRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).KillJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/KillJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).KillJob(ctx, req.(*KillJobRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_KillTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillTaskRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).KillTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/KillTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).KillTask(ctx, req.(*KillTaskRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_KillTaskAttempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillTaskAttemptRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).KillTaskAttempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/KillTaskAttempt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).KillTaskAttempt(ctx, req.(*KillTaskAttemptRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_FailTaskAttempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailTaskAttemptRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).FailTaskAttempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/FailTaskAttempt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).FailTaskAttempt(ctx, req.(*FailTaskAttemptRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_RenewDelegationToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_common.RenewDelegationTokenRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).RenewDelegationToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/RenewDelegationToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).RenewDelegationToken(ctx, req.(*hadoop_common.RenewDelegationTokenRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _MRClientProtocolService_CancelDelegationToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(hadoop_common.CancelDelegationTokenRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MRClientProtocolServiceServer).CancelDelegationToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hadoop.mapreduce.MRClientProtocolService/CancelDelegationToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MRClientProtocolServiceServer).CancelDelegationToken(ctx, req.(*hadoop_common.CancelDelegationTokenRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _MRClientProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hadoop.mapreduce.MRClientProtocolService",
	HandlerType: (*MRClientProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getJobReport",
			Handler:    _MRClientProtocolService_GetJobReport_Handler,
		},
		{
			MethodName: "getTaskReport",
			Handler:    _MRClientProtocolService_GetTaskReport_Handler,
		},
		{
			MethodName: "getTaskAttemptReport",
			Handler:    _MRClientProtocolService_GetTaskAttemptReport_Handler,
		},
		{
			MethodName: "getCounters",
			Handler:    _MRClientProtocolService_GetCounters_Handler,
		},
		{
			MethodName: "getTaskAttemptCompletionEvents",
			Handler:    _MRClientProtocolService_GetTaskAttemptCompletionEvents_Handler,
		},
		{
			MethodName: "getTaskReports",
			Handler:    _MRClientProtocolService_GetTaskReports_Handler,
		},
		{
			MethodName: "getDiagnostics",
			Handler:    _MRClientProtocolService_GetDiagnostics_Handler,
		},
		{
			MethodName: "getDelegationToken",
			Handler:    _MRClientProtocolService_GetDelegationToken_Handler,
		},
		{
			MethodName: "killJob",
			Handler:    _MRClientProtocolService_KillJob_Handler,
		},
		{
			MethodName: "killTask",
			Handler:    _MRClientProtocolService_KillTask_Handler,
		},
		{
			MethodName: "killTaskAttempt",
			Handler:    _MRClientProtocolService_KillTaskAttempt_Handler,
		},
		{
			MethodName: "failTaskAttempt",
			Handler:    _MRClientProtocolService_FailTaskAttempt_Handler,
		},
		{
			MethodName: "renewDelegationToken",
			Handler:    _MRClientProtocolService_RenewDelegationToken_Handler,
		},
		{
			MethodName: "cancelDelegationToken",
			Handler:    _MRClientProtocolService_CancelDelegationToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("MRClientProtocol.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xd5, 0x13, 0xc8, 0x8c, 0x31, 0x59, 0x83, 0x49, 0x15, 0xe2, 0x06, 0x88, 0xc1, 0x02,
	0x6c, 0x57, 0x38, 0xb0, 0xf0, 0x47, 0x02, 0x21, 0xa1, 0x6c, 0x07, 0x24, 0x0e, 0x93, 0x71, 0x5f,
	0x32, 0xd3, 0xc4, 0xaf, 0xb1, 0x9d, 0xa1, 0x4a, 0x48, 0x5c, 0xf9, 0x12, 0x7c, 0x33, 0x3e, 0x0c,
	0x49, 0x93, 0x0e, 0xdb, 0x4d, 0xea, 0x70, 0xab, 0xea, 0xdf, 0xf3, 0xfc, 0xde, 0xd8, 0x89, 0x4c,
	0x6e, 0xbd, 0xcf, 0xd2, 0x42, 0x80, 0xb4, 0x1f, 0x34, 0x5a, 0xe4, 0x58, 0x24, 0xaa, 0xf9, 0x41,
	0x77, 0xce, 0xd9, 0x0c, 0x51, 0x25, 0x25, 0x53, 0x1a, 0x66, 0x15, 0x87, 0xe9, 0xf6, 0x09, 0xf0,
	0x4a, 0x0b, 0xbb, 0x68, 0x89, 0xe9, 0x5e, 0xa9, 0xcf, 0x0c, 0xe8, 0x0b, 0xc1, 0xe1, 0x6c, 0xf9,
	0x8f, 0x69, 0x17, 0x0e, 0xff, 0x6c, 0x91, 0xbd, 0xb0, 0xf5, 0xa4, 0x05, 0x69, 0x4e, 0xb6, 0x72,
	0xb0, 0x6f, 0xf1, 0x73, 0x06, 0x0a, 0xb5, 0xa5, 0xfb, 0x49, 0xe8, 0x49, 0xde, 0x38, 0xeb, 0x19,
	0x7c, 0xab, 0xc0, 0xb4, 0x2d, 0xd3, 0x87, 0x31, 0xd6, 0x28, 0x94, 0x06, 0x96, 0x30, 0xfd, 0x4a,
	0xae, 0xd7, 0xa2, 0x53, 0x66, 0xe6, 0x9d, 0xa9, 0x3f, 0xfd, 0x0f, 0xf0, 0x54, 0x8f, 0xa2, 0xb0,
	0xeb, 0xfa, 0x49, 0x76, 0x3b, 0xd7, 0x0b, 0x6b, 0xa1, 0x54, 0xb6, 0x53, 0x1e, 0x0e, 0xb6, 0x78,
	0x9c, 0x67, 0x3e, 0x1a, 0x9b, 0x71, 0x07, 0x98, 0x91, 0x6b, 0xf5, 0x00, 0x29, 0x56, 0xd2, 0x82,
	0x36, 0xf4, 0x41, 0x6f, 0xc7, 0x6a, 0xd9, 0xd3, 0xed, 0x47, 0x50, 0xd7, 0xf2, 0x7b, 0x42, 0xee,
	0xf8, 0xcf, 0x99, 0x62, 0xa9, 0x0a, 0xb0, 0x02, 0xe5, 0xab, 0x8b, 0xfa, 0xb0, 0x0d, 0x7d, 0x16,
	0x9b, 0x3e, 0x4c, 0x78, 0xc3, 0x3c, 0xff, 0xff, 0xb4, 0x3b, 0x5f, 0x49, 0xb6, 0xbd, 0x23, 0x37,
	0x34, 0x76, 0x8c, 0xbe, 0xfe, 0x20, 0x4e, 0xaf, 0xeb, 0x5e, 0x0a, 0x96, 0x4b, 0x34, 0x56, 0xf0,
	0x21, 0x9d, 0x43, 0x8c, 0xd0, 0x79, 0xb4, 0xab, 0x33, 0x84, 0x36, 0x3a, 0x28, 0x20, 0x67, 0xcd,
	0x0e, 0x9c, 0xe2, 0x1c, 0x24, 0xbd, 0x2c, 0xe1, 0x58, 0x96, 0x28, 0x97, 0x0d, 0x3e, 0xe2, 0x39,
	0x93, 0x11, 0xb8, 0x2b, 0xfd, 0x48, 0xae, 0xcc, 0x45, 0x51, 0xd4, 0xdf, 0x18, 0xbd, 0xbb, 0x3e,
	0xee, 0xbb, 0x76, 0xc9, 0x33, 0xdc, 0xdb, 0x80, 0xb9, 0xcd, 0x9f, 0xc8, 0xd5, 0xa6, 0xb9, 0xd9,
	0x5d, 0x3a, 0x90, 0x69, 0x77, 0xde, 0xe9, 0xbe, 0xbf, 0x89, 0x73, 0xcb, 0x15, 0xb9, 0xb1, 0x2a,
	0xef, 0xde, 0x1c, 0x7a, 0x30, 0x9c, 0xbd, 0xfc, 0xb0, 0xfa, 0x36, 0x6a, 0x03, 0x1e, 0x18, 0xbf,
	0x30, 0x11, 0x33, 0xbe, 0xf6, 0x91, 0x98, 0x71, 0x0d, 0x77, 0x8d, 0x0b, 0xb2, 0xab, 0x41, 0xc2,
	0xf7, 0xf0, 0x8d, 0x78, 0x1c, 0x1c, 0x71, 0xd6, 0x03, 0x79, 0xe2, 0x27, 0xa3, 0x02, 0xae, 0xfa,
	0x07, 0xb9, 0xc9, 0x99, 0xe4, 0x50, 0x84, 0xee, 0xb0, 0x2a, 0xed, 0xa3, 0x3c, 0xf9, 0xd3, 0x71,
	0x09, 0xc7, 0x7e, 0x7c, 0x44, 0x6e, 0xa3, 0xce, 0x13, 0xa6, 0x18, 0x3f, 0x87, 0x55, 0x7c, 0xc1,
	0xb4, 0x6c, 0xaf, 0x9f, 0xe3, 0x9d, 0xf0, 0xee, 0xf9, 0x35, 0x99, 0xfc, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xf8, 0xe4, 0xad, 0x05, 0xe7, 0x06, 0x00, 0x00,
}