// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.0
// - protoc             (unknown)
// source: teleport/userprovisioning/v1/statichostuser_service.proto

package userprovisioningv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StaticHostUsersService_GetStaticHostUser_FullMethodName    = "/teleport.userprovisioning.v1.StaticHostUsersService/GetStaticHostUser"
	StaticHostUsersService_ListStaticHostUsers_FullMethodName  = "/teleport.userprovisioning.v1.StaticHostUsersService/ListStaticHostUsers"
	StaticHostUsersService_CreateStaticHostUser_FullMethodName = "/teleport.userprovisioning.v1.StaticHostUsersService/CreateStaticHostUser"
	StaticHostUsersService_UpdateStaticHostUser_FullMethodName = "/teleport.userprovisioning.v1.StaticHostUsersService/UpdateStaticHostUser"
	StaticHostUsersService_UpsertStaticHostUser_FullMethodName = "/teleport.userprovisioning.v1.StaticHostUsersService/UpsertStaticHostUser"
	StaticHostUsersService_DeleteStaticHostUser_FullMethodName = "/teleport.userprovisioning.v1.StaticHostUsersService/DeleteStaticHostUser"
)

// StaticHostUsersServiceClient is the client API for StaticHostUsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// StaticHostUsersService provides methods to manage static host users.
type StaticHostUsersServiceClient interface {
	// GetStaticHostUser retrieves a static host user resource by name.
	GetStaticHostUser(ctx context.Context, in *GetStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error)
	// ListStaticHostUsers gets all existing static host users.
	ListStaticHostUsers(ctx context.Context, in *ListStaticHostUsersRequest, opts ...grpc.CallOption) (*ListStaticHostUsersResponse, error)
	// CreateStaticHostUser creates a static host user if one does not already exist.
	CreateStaticHostUser(ctx context.Context, in *CreateStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error)
	// UpdateStaticHostUser updates an existing static host user.
	UpdateStaticHostUser(ctx context.Context, in *UpdateStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error)
	// UpsertStaticHostUser creates a new static host user or forcefully updates an existing static host user.
	UpsertStaticHostUser(ctx context.Context, in *UpsertStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error)
	// DeleteStaticHostUser removes an existing static host user resource by name.
	DeleteStaticHostUser(ctx context.Context, in *DeleteStaticHostUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type staticHostUsersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticHostUsersServiceClient(cc grpc.ClientConnInterface) StaticHostUsersServiceClient {
	return &staticHostUsersServiceClient{cc}
}

func (c *staticHostUsersServiceClient) GetStaticHostUser(ctx context.Context, in *GetStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaticHostUser)
	err := c.cc.Invoke(ctx, StaticHostUsersService_GetStaticHostUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticHostUsersServiceClient) ListStaticHostUsers(ctx context.Context, in *ListStaticHostUsersRequest, opts ...grpc.CallOption) (*ListStaticHostUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStaticHostUsersResponse)
	err := c.cc.Invoke(ctx, StaticHostUsersService_ListStaticHostUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticHostUsersServiceClient) CreateStaticHostUser(ctx context.Context, in *CreateStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaticHostUser)
	err := c.cc.Invoke(ctx, StaticHostUsersService_CreateStaticHostUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticHostUsersServiceClient) UpdateStaticHostUser(ctx context.Context, in *UpdateStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaticHostUser)
	err := c.cc.Invoke(ctx, StaticHostUsersService_UpdateStaticHostUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticHostUsersServiceClient) UpsertStaticHostUser(ctx context.Context, in *UpsertStaticHostUserRequest, opts ...grpc.CallOption) (*StaticHostUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StaticHostUser)
	err := c.cc.Invoke(ctx, StaticHostUsersService_UpsertStaticHostUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticHostUsersServiceClient) DeleteStaticHostUser(ctx context.Context, in *DeleteStaticHostUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, StaticHostUsersService_DeleteStaticHostUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticHostUsersServiceServer is the server API for StaticHostUsersService service.
// All implementations must embed UnimplementedStaticHostUsersServiceServer
// for forward compatibility.
//
// StaticHostUsersService provides methods to manage static host users.
type StaticHostUsersServiceServer interface {
	// GetStaticHostUser retrieves a static host user resource by name.
	GetStaticHostUser(context.Context, *GetStaticHostUserRequest) (*StaticHostUser, error)
	// ListStaticHostUsers gets all existing static host users.
	ListStaticHostUsers(context.Context, *ListStaticHostUsersRequest) (*ListStaticHostUsersResponse, error)
	// CreateStaticHostUser creates a static host user if one does not already exist.
	CreateStaticHostUser(context.Context, *CreateStaticHostUserRequest) (*StaticHostUser, error)
	// UpdateStaticHostUser updates an existing static host user.
	UpdateStaticHostUser(context.Context, *UpdateStaticHostUserRequest) (*StaticHostUser, error)
	// UpsertStaticHostUser creates a new static host user or forcefully updates an existing static host user.
	UpsertStaticHostUser(context.Context, *UpsertStaticHostUserRequest) (*StaticHostUser, error)
	// DeleteStaticHostUser removes an existing static host user resource by name.
	DeleteStaticHostUser(context.Context, *DeleteStaticHostUserRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedStaticHostUsersServiceServer()
}

// UnimplementedStaticHostUsersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStaticHostUsersServiceServer struct{}

func (UnimplementedStaticHostUsersServiceServer) GetStaticHostUser(context.Context, *GetStaticHostUserRequest) (*StaticHostUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStaticHostUser not implemented")
}
func (UnimplementedStaticHostUsersServiceServer) ListStaticHostUsers(context.Context, *ListStaticHostUsersRequest) (*ListStaticHostUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStaticHostUsers not implemented")
}
func (UnimplementedStaticHostUsersServiceServer) CreateStaticHostUser(context.Context, *CreateStaticHostUserRequest) (*StaticHostUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaticHostUser not implemented")
}
func (UnimplementedStaticHostUsersServiceServer) UpdateStaticHostUser(context.Context, *UpdateStaticHostUserRequest) (*StaticHostUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaticHostUser not implemented")
}
func (UnimplementedStaticHostUsersServiceServer) UpsertStaticHostUser(context.Context, *UpsertStaticHostUserRequest) (*StaticHostUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertStaticHostUser not implemented")
}
func (UnimplementedStaticHostUsersServiceServer) DeleteStaticHostUser(context.Context, *DeleteStaticHostUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStaticHostUser not implemented")
}
func (UnimplementedStaticHostUsersServiceServer) mustEmbedUnimplementedStaticHostUsersServiceServer() {
}
func (UnimplementedStaticHostUsersServiceServer) testEmbeddedByValue() {}

// UnsafeStaticHostUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticHostUsersServiceServer will
// result in compilation errors.
type UnsafeStaticHostUsersServiceServer interface {
	mustEmbedUnimplementedStaticHostUsersServiceServer()
}

func RegisterStaticHostUsersServiceServer(s grpc.ServiceRegistrar, srv StaticHostUsersServiceServer) {
	// If the following call pancis, it indicates UnimplementedStaticHostUsersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StaticHostUsersService_ServiceDesc, srv)
}

func _StaticHostUsersService_GetStaticHostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaticHostUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticHostUsersServiceServer).GetStaticHostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticHostUsersService_GetStaticHostUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticHostUsersServiceServer).GetStaticHostUser(ctx, req.(*GetStaticHostUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticHostUsersService_ListStaticHostUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaticHostUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticHostUsersServiceServer).ListStaticHostUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticHostUsersService_ListStaticHostUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticHostUsersServiceServer).ListStaticHostUsers(ctx, req.(*ListStaticHostUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticHostUsersService_CreateStaticHostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaticHostUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticHostUsersServiceServer).CreateStaticHostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticHostUsersService_CreateStaticHostUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticHostUsersServiceServer).CreateStaticHostUser(ctx, req.(*CreateStaticHostUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticHostUsersService_UpdateStaticHostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaticHostUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticHostUsersServiceServer).UpdateStaticHostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticHostUsersService_UpdateStaticHostUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticHostUsersServiceServer).UpdateStaticHostUser(ctx, req.(*UpdateStaticHostUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticHostUsersService_UpsertStaticHostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertStaticHostUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticHostUsersServiceServer).UpsertStaticHostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticHostUsersService_UpsertStaticHostUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticHostUsersServiceServer).UpsertStaticHostUser(ctx, req.(*UpsertStaticHostUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticHostUsersService_DeleteStaticHostUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStaticHostUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticHostUsersServiceServer).DeleteStaticHostUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticHostUsersService_DeleteStaticHostUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticHostUsersServiceServer).DeleteStaticHostUser(ctx, req.(*DeleteStaticHostUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaticHostUsersService_ServiceDesc is the grpc.ServiceDesc for StaticHostUsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaticHostUsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.userprovisioning.v1.StaticHostUsersService",
	HandlerType: (*StaticHostUsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStaticHostUser",
			Handler:    _StaticHostUsersService_GetStaticHostUser_Handler,
		},
		{
			MethodName: "ListStaticHostUsers",
			Handler:    _StaticHostUsersService_ListStaticHostUsers_Handler,
		},
		{
			MethodName: "CreateStaticHostUser",
			Handler:    _StaticHostUsersService_CreateStaticHostUser_Handler,
		},
		{
			MethodName: "UpdateStaticHostUser",
			Handler:    _StaticHostUsersService_UpdateStaticHostUser_Handler,
		},
		{
			MethodName: "UpsertStaticHostUser",
			Handler:    _StaticHostUsersService_UpsertStaticHostUser_Handler,
		},
		{
			MethodName: "DeleteStaticHostUser",
			Handler:    _StaticHostUsersService_DeleteStaticHostUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/userprovisioning/v1/statichostuser_service.proto",
}
