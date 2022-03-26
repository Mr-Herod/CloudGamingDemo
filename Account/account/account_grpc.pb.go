// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package account

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterRsp, error)
	UserLogIn(ctx context.Context, in *UserLogInReq, opts ...grpc.CallOption) (*UserLogInRsp, error)
	UserLogOut(ctx context.Context, in *UserLogOutReq, opts ...grpc.CallOption) (*UserLogOutRsp, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterRsp, error) {
	out := new(UserRegisterRsp)
	err := c.cc.Invoke(ctx, "/account.AccountService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UserLogIn(ctx context.Context, in *UserLogInReq, opts ...grpc.CallOption) (*UserLogInRsp, error) {
	out := new(UserLogInRsp)
	err := c.cc.Invoke(ctx, "/account.AccountService/UserLogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UserLogOut(ctx context.Context, in *UserLogOutReq, opts ...grpc.CallOption) (*UserLogOutRsp, error) {
	out := new(UserLogOutRsp)
	err := c.cc.Invoke(ctx, "/account.AccountService/UserLogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterRsp, error)
	UserLogIn(context.Context, *UserLogInReq) (*UserLogInRsp, error)
	UserLogOut(context.Context, *UserLogOutReq) (*UserLogOutRsp, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) UserRegister(context.Context, *UserRegisterReq) (*UserRegisterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedAccountServiceServer) UserLogIn(context.Context, *UserLogInReq) (*UserLogInRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogIn not implemented")
}
func (UnimplementedAccountServiceServer) UserLogOut(context.Context, *UserLogOutReq) (*UserLogOutRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogOut not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UserLogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UserLogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/UserLogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UserLogIn(ctx, req.(*UserLogInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UserLogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogOutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UserLogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/UserLogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UserLogOut(ctx, req.(*UserLogOutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _AccountService_UserRegister_Handler,
		},
		{
			MethodName: "UserLogIn",
			Handler:    _AccountService_UserLogIn_Handler,
		},
		{
			MethodName: "UserLogOut",
			Handler:    _AccountService_UserLogOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account.proto",
}
