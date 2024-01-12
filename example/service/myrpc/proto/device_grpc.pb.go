// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: device.proto

package proto

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
	Device_DeviceReg_FullMethodName = "/proto.Device/deviceReg"
	Device_DeviceGet_FullMethodName = "/proto.Device/deviceGet"
)

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceClient interface {
	// 设备注册
	DeviceReg(ctx context.Context, in *DeviceRegReq, opts ...grpc.CallOption) (*DeviceResp, error)
	// 设备列表
	DeviceGet(ctx context.Context, in *DeviceGetReq, opts ...grpc.CallOption) (*DeviceGetResp, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) DeviceReg(ctx context.Context, in *DeviceRegReq, opts ...grpc.CallOption) (*DeviceResp, error) {
	out := new(DeviceResp)
	err := c.cc.Invoke(ctx, Device_DeviceReg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) DeviceGet(ctx context.Context, in *DeviceGetReq, opts ...grpc.CallOption) (*DeviceGetResp, error) {
	out := new(DeviceGetResp)
	err := c.cc.Invoke(ctx, Device_DeviceGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
// All implementations must embed UnimplementedDeviceServer
// for forward compatibility
type DeviceServer interface {
	// 设备注册
	DeviceReg(context.Context, *DeviceRegReq) (*DeviceResp, error)
	// 设备列表
	DeviceGet(context.Context, *DeviceGetReq) (*DeviceGetResp, error)
	mustEmbedUnimplementedDeviceServer()
}

// UnimplementedDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServer struct {
}

func (UnimplementedDeviceServer) DeviceReg(context.Context, *DeviceRegReq) (*DeviceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceReg not implemented")
}
func (UnimplementedDeviceServer) DeviceGet(context.Context, *DeviceGetReq) (*DeviceGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceGet not implemented")
}
func (UnimplementedDeviceServer) mustEmbedUnimplementedDeviceServer() {}

// UnsafeDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServer will
// result in compilation errors.
type UnsafeDeviceServer interface {
	mustEmbedUnimplementedDeviceServer()
}

func RegisterDeviceServer(s grpc.ServiceRegistrar, srv DeviceServer) {
	s.RegisterService(&Device_ServiceDesc, srv)
}

func _Device_DeviceReg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRegReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeviceReg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Device_DeviceReg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeviceReg(ctx, req.(*DeviceRegReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_DeviceGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).DeviceGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Device_DeviceGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).DeviceGet(ctx, req.(*DeviceGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Device_ServiceDesc is the grpc.ServiceDesc for Device service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Device_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "deviceReg",
			Handler:    _Device_DeviceReg_Handler,
		},
		{
			MethodName: "deviceGet",
			Handler:    _Device_DeviceGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}
