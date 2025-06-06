// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: report.proto

package grpc_client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Report_SendReport_FullMethodName = "/manager.v1.Report/SendReport"
)

// ReportClient is the client API for Report service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportClient interface {
	SendReport(ctx context.Context, in *SendReportRequest, opts ...grpc.CallOption) (*SendReportResponse, error)
}

type reportClient struct {
	cc grpc.ClientConnInterface
}

func NewReportClient(cc grpc.ClientConnInterface) ReportClient {
	return &reportClient{cc}
}

func (c *reportClient) SendReport(ctx context.Context, in *SendReportRequest, opts ...grpc.CallOption) (*SendReportResponse, error) {
	out := new(SendReportResponse)
	err := c.cc.Invoke(ctx, Report_SendReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServer is the server API for Report service.
// All implementations should embed UnimplementedReportServer
// for forward compatibility
type ReportServer interface {
	SendReport(context.Context, *SendReportRequest) (*SendReportResponse, error)
}

// UnimplementedReportServer should be embedded to have forward compatible implementations.
type UnimplementedReportServer struct {
}

func (UnimplementedReportServer) SendReport(context.Context, *SendReportRequest) (*SendReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReport not implemented")
}

// UnsafeReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportServer will
// result in compilation errors.
type UnsafeReportServer interface {
	mustEmbedUnimplementedReportServer()
}

func RegisterReportServer(s grpc.ServiceRegistrar, srv ReportServer) {
	s.RegisterService(&Report_ServiceDesc, srv)
}

func _Report_SendReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).SendReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Report_SendReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).SendReport(ctx, req.(*SendReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Report_ServiceDesc is the grpc.ServiceDesc for Report service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Report_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.v1.Report",
	HandlerType: (*ReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendReport",
			Handler:    _Report_SendReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "report.proto",
}
