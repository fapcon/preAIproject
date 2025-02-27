// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: user.proto

package user

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

// UserServiceGRPCClient is the client API for UserServiceGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceGRPCClient interface {
	Create(ctx context.Context, in *UserCreateIn, opts ...grpc.CallOption) (*UserCreateOut, error)
	Update(ctx context.Context, in *UserUpdateIn, opts ...grpc.CallOption) (*UserUpdateOut, error)
	VerifyEmail(ctx context.Context, in *UserVerifyEmailIn, opts ...grpc.CallOption) (*UserUpdateOut, error)
	ChangePassword(ctx context.Context, in *ChangePasswordIn, opts ...grpc.CallOption) (*ChangePasswordOut, error)
	GetByEmail(ctx context.Context, in *GetByEmailIn, opts ...grpc.CallOption) (*UserOut, error)
	GetByPhone(ctx context.Context, in *GetByPhoneIn, opts ...grpc.CallOption) (*UserOut, error)
	GetByID(ctx context.Context, in *GetByIDIn, opts ...grpc.CallOption) (*UserOut, error)
	GetByIDs(ctx context.Context, in *GetByIDsIn, opts ...grpc.CallOption) (*UsersOut, error)
	BanByID(ctx context.Context, in *BanByIDIn, opts ...grpc.CallOption) (*BanByIDOut, error)
	IsBanned(ctx context.Context, in *IsBannedIn, opts ...grpc.CallOption) (*IsBannedOut, error)
	UnbanByID(ctx context.Context, in *UnbanByIDIn, opts ...grpc.CallOption) (*UnbanByIDOut, error)
}

type userServiceGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceGRPCClient(cc grpc.ClientConnInterface) UserServiceGRPCClient {
	return &userServiceGRPCClient{cc}
}

func (c *userServiceGRPCClient) Create(ctx context.Context, in *UserCreateIn, opts ...grpc.CallOption) (*UserCreateOut, error) {
	out := new(UserCreateOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) Update(ctx context.Context, in *UserUpdateIn, opts ...grpc.CallOption) (*UserUpdateOut, error) {
	out := new(UserUpdateOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) VerifyEmail(ctx context.Context, in *UserVerifyEmailIn, opts ...grpc.CallOption) (*UserUpdateOut, error) {
	out := new(UserUpdateOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/VerifyEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) ChangePassword(ctx context.Context, in *ChangePasswordIn, opts ...grpc.CallOption) (*ChangePasswordOut, error) {
	out := new(ChangePasswordOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) GetByEmail(ctx context.Context, in *GetByEmailIn, opts ...grpc.CallOption) (*UserOut, error) {
	out := new(UserOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/GetByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) GetByPhone(ctx context.Context, in *GetByPhoneIn, opts ...grpc.CallOption) (*UserOut, error) {
	out := new(UserOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/GetByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) GetByID(ctx context.Context, in *GetByIDIn, opts ...grpc.CallOption) (*UserOut, error) {
	out := new(UserOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) GetByIDs(ctx context.Context, in *GetByIDsIn, opts ...grpc.CallOption) (*UsersOut, error) {
	out := new(UsersOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/GetByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) BanByID(ctx context.Context, in *BanByIDIn, opts ...grpc.CallOption) (*BanByIDOut, error) {
	out := new(BanByIDOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/BanByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) IsBanned(ctx context.Context, in *IsBannedIn, opts ...grpc.CallOption) (*IsBannedOut, error) {
	out := new(IsBannedOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/IsBanned", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceGRPCClient) UnbanByID(ctx context.Context, in *UnbanByIDIn, opts ...grpc.CallOption) (*UnbanByIDOut, error) {
	out := new(UnbanByIDOut)
	err := c.cc.Invoke(ctx, "/user.UserServiceGRPC/UnbanByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceGRPCServer is the server API for UserServiceGRPC service.
// All implementations must embed UnimplementedUserServiceGRPCServer
// for forward compatibility
type UserServiceGRPCServer interface {
	Create(context.Context, *UserCreateIn) (*UserCreateOut, error)
	Update(context.Context, *UserUpdateIn) (*UserUpdateOut, error)
	VerifyEmail(context.Context, *UserVerifyEmailIn) (*UserUpdateOut, error)
	ChangePassword(context.Context, *ChangePasswordIn) (*ChangePasswordOut, error)
	GetByEmail(context.Context, *GetByEmailIn) (*UserOut, error)
	GetByPhone(context.Context, *GetByPhoneIn) (*UserOut, error)
	GetByID(context.Context, *GetByIDIn) (*UserOut, error)
	GetByIDs(context.Context, *GetByIDsIn) (*UsersOut, error)
	BanByID(context.Context, *BanByIDIn) (*BanByIDOut, error)
	IsBanned(context.Context, *IsBannedIn) (*IsBannedOut, error)
	UnbanByID(context.Context, *UnbanByIDIn) (*UnbanByIDOut, error)
	mustEmbedUnimplementedUserServiceGRPCServer()
}

// UnimplementedUserServiceGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceGRPCServer struct {
}

func (UnimplementedUserServiceGRPCServer) Create(context.Context, *UserCreateIn) (*UserCreateOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceGRPCServer) Update(context.Context, *UserUpdateIn) (*UserUpdateOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceGRPCServer) VerifyEmail(context.Context, *UserVerifyEmailIn) (*UserUpdateOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedUserServiceGRPCServer) ChangePassword(context.Context, *ChangePasswordIn) (*ChangePasswordOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceGRPCServer) GetByEmail(context.Context, *GetByEmailIn) (*UserOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (UnimplementedUserServiceGRPCServer) GetByPhone(context.Context, *GetByPhoneIn) (*UserOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPhone not implemented")
}
func (UnimplementedUserServiceGRPCServer) GetByID(context.Context, *GetByIDIn) (*UserOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserServiceGRPCServer) GetByIDs(context.Context, *GetByIDsIn) (*UsersOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIDs not implemented")
}
func (UnimplementedUserServiceGRPCServer) BanByID(context.Context, *BanByIDIn) (*BanByIDOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanByID not implemented")
}
func (UnimplementedUserServiceGRPCServer) IsBanned(context.Context, *IsBannedIn) (*IsBannedOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBanned not implemented")
}
func (UnimplementedUserServiceGRPCServer) UnbanByID(context.Context, *UnbanByIDIn) (*UnbanByIDOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbanByID not implemented")
}
func (UnimplementedUserServiceGRPCServer) mustEmbedUnimplementedUserServiceGRPCServer() {}

// UnsafeUserServiceGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceGRPCServer will
// result in compilation errors.
type UnsafeUserServiceGRPCServer interface {
	mustEmbedUnimplementedUserServiceGRPCServer()
}

func RegisterUserServiceGRPCServer(s grpc.ServiceRegistrar, srv UserServiceGRPCServer) {
	s.RegisterService(&UserServiceGRPC_ServiceDesc, srv)
}

func _UserServiceGRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).Create(ctx, req.(*UserCreateIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).Update(ctx, req.(*UserUpdateIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVerifyEmailIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/VerifyEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).VerifyEmail(ctx, req.(*UserVerifyEmailIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).ChangePassword(ctx, req.(*ChangePasswordIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByEmailIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/GetByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).GetByEmail(ctx, req.(*GetByEmailIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_GetByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByPhoneIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).GetByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/GetByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).GetByPhone(ctx, req.(*GetByPhoneIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).GetByID(ctx, req.(*GetByIDIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_GetByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).GetByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/GetByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).GetByIDs(ctx, req.(*GetByIDsIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_BanByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanByIDIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).BanByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/BanByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).BanByID(ctx, req.(*BanByIDIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_IsBanned_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBannedIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).IsBanned(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/IsBanned",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).IsBanned(ctx, req.(*IsBannedIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceGRPC_UnbanByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbanByIDIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceGRPCServer).UnbanByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserServiceGRPC/UnbanByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceGRPCServer).UnbanByID(ctx, req.(*UnbanByIDIn))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServiceGRPC_ServiceDesc is the grpc.ServiceDesc for UserServiceGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServiceGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserServiceGRPC",
	HandlerType: (*UserServiceGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserServiceGRPC_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserServiceGRPC_Update_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _UserServiceGRPC_VerifyEmail_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserServiceGRPC_ChangePassword_Handler,
		},
		{
			MethodName: "GetByEmail",
			Handler:    _UserServiceGRPC_GetByEmail_Handler,
		},
		{
			MethodName: "GetByPhone",
			Handler:    _UserServiceGRPC_GetByPhone_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _UserServiceGRPC_GetByID_Handler,
		},
		{
			MethodName: "GetByIDs",
			Handler:    _UserServiceGRPC_GetByIDs_Handler,
		},
		{
			MethodName: "BanByID",
			Handler:    _UserServiceGRPC_BanByID_Handler,
		},
		{
			MethodName: "IsBanned",
			Handler:    _UserServiceGRPC_IsBanned_Handler,
		},
		{
			MethodName: "UnbanByID",
			Handler:    _UserServiceGRPC_UnbanByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
