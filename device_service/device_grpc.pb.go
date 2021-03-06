// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: device.proto

package device_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceServiceClient interface {
	// This method records the provided data by the device's inputs
	// the purpose of this method it's just to save the information each
	// time that is called without execute any calculation
	RecordData(ctx context.Context, in *DeviceRecordData, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// This method contains all the logic to calculate the right output
	// given an input data
	GetDeviceOutput(ctx context.Context, in *DeviceRecordData, opts ...grpc.CallOption) (*DeviceOutputData, error)
	// This method retrieves the information of a certain device given his id
	GetDeviceRecordedData(ctx context.Context, in *DeviceId, opts ...grpc.CallOption) (*DeviceRecordData, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) RecordData(ctx context.Context, in *DeviceRecordData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/device.DeviceService/RecordData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceOutput(ctx context.Context, in *DeviceRecordData, opts ...grpc.CallOption) (*DeviceOutputData, error) {
	out := new(DeviceOutputData)
	err := c.cc.Invoke(ctx, "/device.DeviceService/GetDeviceOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceRecordedData(ctx context.Context, in *DeviceId, opts ...grpc.CallOption) (*DeviceRecordData, error) {
	out := new(DeviceRecordData)
	err := c.cc.Invoke(ctx, "/device.DeviceService/GetDeviceRecordedData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations must embed UnimplementedDeviceServiceServer
// for forward compatibility
type DeviceServiceServer interface {
	// This method records the provided data by the device's inputs
	// the purpose of this method it's just to save the information each
	// time that is called without execute any calculation
	RecordData(context.Context, *DeviceRecordData) (*emptypb.Empty, error)
	// This method contains all the logic to calculate the right output
	// given an input data
	GetDeviceOutput(context.Context, *DeviceRecordData) (*DeviceOutputData, error)
	// This method retrieves the information of a certain device given his id
	GetDeviceRecordedData(context.Context, *DeviceId) (*DeviceRecordData, error)
	mustEmbedUnimplementedDeviceServiceServer()
}

// UnimplementedDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (UnimplementedDeviceServiceServer) RecordData(context.Context, *DeviceRecordData) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordData not implemented")
}
func (UnimplementedDeviceServiceServer) GetDeviceOutput(context.Context, *DeviceRecordData) (*DeviceOutputData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceOutput not implemented")
}
func (UnimplementedDeviceServiceServer) GetDeviceRecordedData(context.Context, *DeviceId) (*DeviceRecordData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceRecordedData not implemented")
}
func (UnimplementedDeviceServiceServer) mustEmbedUnimplementedDeviceServiceServer() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_RecordData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRecordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).RecordData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.DeviceService/RecordData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).RecordData(ctx, req.(*DeviceRecordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRecordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.DeviceService/GetDeviceOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceOutput(ctx, req.(*DeviceRecordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceRecordedData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceRecordedData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.DeviceService/GetDeviceRecordedData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceRecordedData(ctx, req.(*DeviceId))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordData",
			Handler:    _DeviceService_RecordData_Handler,
		},
		{
			MethodName: "GetDeviceOutput",
			Handler:    _DeviceService_GetDeviceOutput_Handler,
		},
		{
			MethodName: "GetDeviceRecordedData",
			Handler:    _DeviceService_GetDeviceRecordedData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}
