// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: culc.proto

package culc

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Calculate(ctx context.Context, in *CalculateReq, opts ...grpc.CallOption) (*CalculateRes, error)
	NewClient(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*ClientRes, error)
	StreamServerStatuses(ctx context.Context, in *StreamServerStatusesRequest, opts ...grpc.CallOption) (Auth_StreamServerStatusesClient, error)
	GetHistoryEx(ctx context.Context, in *HistoryReq, opts ...grpc.CallOption) (*HistoryRes, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, "/auth.Auth/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Calculate(ctx context.Context, in *CalculateReq, opts ...grpc.CallOption) (*CalculateRes, error) {
	out := new(CalculateRes)
	err := c.cc.Invoke(ctx, "/auth.Auth/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) NewClient(ctx context.Context, in *ClientReq, opts ...grpc.CallOption) (*ClientRes, error) {
	out := new(ClientRes)
	err := c.cc.Invoke(ctx, "/auth.Auth/NewClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) StreamServerStatuses(ctx context.Context, in *StreamServerStatusesRequest, opts ...grpc.CallOption) (Auth_StreamServerStatusesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Auth_ServiceDesc.Streams[0], "/auth.Auth/StreamServerStatuses", opts...)
	if err != nil {
		return nil, err
	}
	x := &authStreamServerStatusesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Auth_StreamServerStatusesClient interface {
	Recv() (*StreamServerStatusesResponse, error)
	grpc.ClientStream
}

type authStreamServerStatusesClient struct {
	grpc.ClientStream
}

func (x *authStreamServerStatusesClient) Recv() (*StreamServerStatusesResponse, error) {
	m := new(StreamServerStatusesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authClient) GetHistoryEx(ctx context.Context, in *HistoryReq, opts ...grpc.CallOption) (*HistoryRes, error) {
	out := new(HistoryRes)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetHistoryEx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Calculate(context.Context, *CalculateReq) (*CalculateRes, error)
	NewClient(context.Context, *ClientReq) (*ClientRes, error)
	StreamServerStatuses(*StreamServerStatusesRequest, Auth_StreamServerStatusesServer) error
	GetHistoryEx(context.Context, *HistoryReq) (*HistoryRes, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Register(context.Context, *RegisterReq) (*RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) Calculate(context.Context, *CalculateReq) (*CalculateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedAuthServer) NewClient(context.Context, *ClientReq) (*ClientRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewClient not implemented")
}
func (UnimplementedAuthServer) StreamServerStatuses(*StreamServerStatusesRequest, Auth_StreamServerStatusesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamServerStatuses not implemented")
}
func (UnimplementedAuthServer) GetHistoryEx(context.Context, *HistoryReq) (*HistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoryEx not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Calculate(ctx, req.(*CalculateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_NewClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).NewClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/NewClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).NewClient(ctx, req.(*ClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_StreamServerStatuses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamServerStatusesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthServer).StreamServerStatuses(m, &authStreamServerStatusesServer{stream})
}

type Auth_StreamServerStatusesServer interface {
	Send(*StreamServerStatusesResponse) error
	grpc.ServerStream
}

type authStreamServerStatusesServer struct {
	grpc.ServerStream
}

func (x *authStreamServerStatusesServer) Send(m *StreamServerStatusesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Auth_GetHistoryEx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetHistoryEx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetHistoryEx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetHistoryEx(ctx, req.(*HistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Calculate",
			Handler:    _Auth_Calculate_Handler,
		},
		{
			MethodName: "NewClient",
			Handler:    _Auth_NewClient_Handler,
		},
		{
			MethodName: "GetHistoryEx",
			Handler:    _Auth_GetHistoryEx_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamServerStatuses",
			Handler:       _Auth_StreamServerStatuses_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "culc.proto",
}
