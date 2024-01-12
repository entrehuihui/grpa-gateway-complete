// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: login.proto

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
	Login_LoginWeb_FullMethodName      = "/proto.Login/loginWeb"
	Login_LoginPwd_FullMethodName      = "/proto.Login/loginPwd"
	Login_LoginPwdLogin_FullMethodName = "/proto.Login/loginPwdLogin"
	Login_LoginInfo_FullMethodName     = "/proto.Login/loginInfo"
	Login_LoginUser_FullMethodName     = "/proto.Login/loginUser"
	Login_LoginUserCode_FullMethodName = "/proto.Login/loginUserCode"
)

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	// 登录
	LoginWeb(ctx context.Context, in *LoginWebReq, opts ...grpc.CallOption) (*Loginresp, error)
	// 修改密码(登陆前)
	LoginPwd(ctx context.Context, in *LoginPwdReq, opts ...grpc.CallOption) (*Loginresp, error)
	// 修改密码(登陆后)
	LoginPwdLogin(ctx context.Context, in *LoginPwdLoginReq, opts ...grpc.CallOption) (*Loginresp, error)
	// 修改邮箱/昵称
	LoginInfo(ctx context.Context, in *LoginInfoReq, opts ...grpc.CallOption) (*Loginresp, error)
	// 创建账号
	LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*Loginresp, error)
	// 发送创建账户验证码
	LoginUserCode(ctx context.Context, in *LoginUserCodeReq, opts ...grpc.CallOption) (*Loginresp, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) LoginWeb(ctx context.Context, in *LoginWebReq, opts ...grpc.CallOption) (*Loginresp, error) {
	out := new(Loginresp)
	err := c.cc.Invoke(ctx, Login_LoginWeb_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginPwd(ctx context.Context, in *LoginPwdReq, opts ...grpc.CallOption) (*Loginresp, error) {
	out := new(Loginresp)
	err := c.cc.Invoke(ctx, Login_LoginPwd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginPwdLogin(ctx context.Context, in *LoginPwdLoginReq, opts ...grpc.CallOption) (*Loginresp, error) {
	out := new(Loginresp)
	err := c.cc.Invoke(ctx, Login_LoginPwdLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginInfo(ctx context.Context, in *LoginInfoReq, opts ...grpc.CallOption) (*Loginresp, error) {
	out := new(Loginresp)
	err := c.cc.Invoke(ctx, Login_LoginInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*Loginresp, error) {
	out := new(Loginresp)
	err := c.cc.Invoke(ctx, Login_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) LoginUserCode(ctx context.Context, in *LoginUserCodeReq, opts ...grpc.CallOption) (*Loginresp, error) {
	out := new(Loginresp)
	err := c.cc.Invoke(ctx, Login_LoginUserCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	// 登录
	LoginWeb(context.Context, *LoginWebReq) (*Loginresp, error)
	// 修改密码(登陆前)
	LoginPwd(context.Context, *LoginPwdReq) (*Loginresp, error)
	// 修改密码(登陆后)
	LoginPwdLogin(context.Context, *LoginPwdLoginReq) (*Loginresp, error)
	// 修改邮箱/昵称
	LoginInfo(context.Context, *LoginInfoReq) (*Loginresp, error)
	// 创建账号
	LoginUser(context.Context, *LoginUserReq) (*Loginresp, error)
	// 发送创建账户验证码
	LoginUserCode(context.Context, *LoginUserCodeReq) (*Loginresp, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) LoginWeb(context.Context, *LoginWebReq) (*Loginresp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWeb not implemented")
}
func (UnimplementedLoginServer) LoginPwd(context.Context, *LoginPwdReq) (*Loginresp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginPwd not implemented")
}
func (UnimplementedLoginServer) LoginPwdLogin(context.Context, *LoginPwdLoginReq) (*Loginresp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginPwdLogin not implemented")
}
func (UnimplementedLoginServer) LoginInfo(context.Context, *LoginInfoReq) (*Loginresp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginInfo not implemented")
}
func (UnimplementedLoginServer) LoginUser(context.Context, *LoginUserReq) (*Loginresp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedLoginServer) LoginUserCode(context.Context, *LoginUserCodeReq) (*Loginresp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUserCode not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_LoginWeb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWebReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginWeb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginWeb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginWeb(ctx, req.(*LoginWebReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPwdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginPwd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginPwd(ctx, req.(*LoginPwdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginPwdLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPwdLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginPwdLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginPwdLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginPwdLogin(ctx, req.(*LoginPwdLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginInfo(ctx, req.(*LoginInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginUser(ctx, req.(*LoginUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_LoginUserCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginUserCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginUserCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginUserCode(ctx, req.(*LoginUserCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc.ServiceDesc for Login service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "loginWeb",
			Handler:    _Login_LoginWeb_Handler,
		},
		{
			MethodName: "loginPwd",
			Handler:    _Login_LoginPwd_Handler,
		},
		{
			MethodName: "loginPwdLogin",
			Handler:    _Login_LoginPwdLogin_Handler,
		},
		{
			MethodName: "loginInfo",
			Handler:    _Login_LoginInfo_Handler,
		},
		{
			MethodName: "loginUser",
			Handler:    _Login_LoginUser_Handler,
		},
		{
			MethodName: "loginUserCode",
			Handler:    _Login_LoginUserCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
