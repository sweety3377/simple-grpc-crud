// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/client.proto

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

// ClientGreeterClient is the client API for ClientGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientGreeterClient interface {
	Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error)
	Read(ctx context.Context, in *ReadClientRequest, opts ...grpc.CallOption) (*ReadClientResponse, error)
	Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*UpdateClientResponse, error)
	Delete(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error)
}

type clientGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewClientGreeterClient(cc grpc.ClientConnInterface) ClientGreeterClient {
	return &clientGreeterClient{cc}
}

func (c *clientGreeterClient) Create(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResponse, error) {
	out := new(CreateClientResponse)
	err := c.cc.Invoke(ctx, "/simple_crud_example.ClientGreeter/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientGreeterClient) Read(ctx context.Context, in *ReadClientRequest, opts ...grpc.CallOption) (*ReadClientResponse, error) {
	out := new(ReadClientResponse)
	err := c.cc.Invoke(ctx, "/simple_crud_example.ClientGreeter/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientGreeterClient) Update(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*UpdateClientResponse, error) {
	out := new(UpdateClientResponse)
	err := c.cc.Invoke(ctx, "/simple_crud_example.ClientGreeter/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientGreeterClient) Delete(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error) {
	out := new(DeleteClientResponse)
	err := c.cc.Invoke(ctx, "/simple_crud_example.ClientGreeter/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientGreeterServer is the server API for ClientGreeter service.
// All implementations must embed UnimplementedClientGreeterServer
// for forward compatibility
type ClientGreeterServer interface {
	Create(context.Context, *CreateClientRequest) (*CreateClientResponse, error)
	Read(context.Context, *ReadClientRequest) (*ReadClientResponse, error)
	Update(context.Context, *UpdateClientRequest) (*UpdateClientResponse, error)
	Delete(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)
	mustEmbedUnimplementedClientGreeterServer()
}

// UnimplementedClientGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedClientGreeterServer struct {
}

func (UnimplementedClientGreeterServer) Create(context.Context, *CreateClientRequest) (*CreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClientGreeterServer) Read(context.Context, *ReadClientRequest) (*ReadClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedClientGreeterServer) Update(context.Context, *UpdateClientRequest) (*UpdateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClientGreeterServer) Delete(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClientGreeterServer) mustEmbedUnimplementedClientGreeterServer() {}

// UnsafeClientGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientGreeterServer will
// result in compilation errors.
type UnsafeClientGreeterServer interface {
	mustEmbedUnimplementedClientGreeterServer()
}

func RegisterClientGreeterServer(s grpc.ServiceRegistrar, srv ClientGreeterServer) {
	s.RegisterService(&ClientGreeter_ServiceDesc, srv)
}

func _ClientGreeter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientGreeterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple_crud_example.ClientGreeter/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientGreeterServer).Create(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientGreeter_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientGreeterServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple_crud_example.ClientGreeter/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientGreeterServer).Read(ctx, req.(*ReadClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientGreeter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientGreeterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple_crud_example.ClientGreeter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientGreeterServer).Update(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientGreeter_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientGreeterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple_crud_example.ClientGreeter/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientGreeterServer).Delete(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientGreeter_ServiceDesc is the grpc.ServiceDesc for ClientGreeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientGreeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simple_crud_example.ClientGreeter",
	HandlerType: (*ClientGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientGreeter_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ClientGreeter_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientGreeter_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClientGreeter_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/client.proto",
}
