// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gateway/gateway.proto

package gateway

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	MovieCreate(ctx context.Context, in *GatewayMovieCreateRequest, opts ...grpc.CallOption) (*GatewayMovieCreateResponse, error)
	MovieList(ctx context.Context, in *GatewayMovieListRequest, opts ...grpc.CallOption) (*GatewayMovieListResponse, error)
	MovieUpdate(ctx context.Context, in *GatewayMovieUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MovieDelete(ctx context.Context, in *GatewayMovieDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MovieGetOne(ctx context.Context, in *GatewayMovieGetOneRequest, opts ...grpc.CallOption) (*GatewayMovieGetOneResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) MovieCreate(ctx context.Context, in *GatewayMovieCreateRequest, opts ...grpc.CallOption) (*GatewayMovieCreateResponse, error) {
	out := new(GatewayMovieCreateResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.homework.api.gateway.Gateway/MovieCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) MovieList(ctx context.Context, in *GatewayMovieListRequest, opts ...grpc.CallOption) (*GatewayMovieListResponse, error) {
	out := new(GatewayMovieListResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.homework.api.gateway.Gateway/MovieList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) MovieUpdate(ctx context.Context, in *GatewayMovieUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozon.dev.homework.api.gateway.Gateway/MovieUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) MovieDelete(ctx context.Context, in *GatewayMovieDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozon.dev.homework.api.gateway.Gateway/MovieDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) MovieGetOne(ctx context.Context, in *GatewayMovieGetOneRequest, opts ...grpc.CallOption) (*GatewayMovieGetOneResponse, error) {
	out := new(GatewayMovieGetOneResponse)
	err := c.cc.Invoke(ctx, "/ozon.dev.homework.api.gateway.Gateway/MovieGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	MovieCreate(context.Context, *GatewayMovieCreateRequest) (*GatewayMovieCreateResponse, error)
	MovieList(context.Context, *GatewayMovieListRequest) (*GatewayMovieListResponse, error)
	MovieUpdate(context.Context, *GatewayMovieUpdateRequest) (*emptypb.Empty, error)
	MovieDelete(context.Context, *GatewayMovieDeleteRequest) (*emptypb.Empty, error)
	MovieGetOne(context.Context, *GatewayMovieGetOneRequest) (*GatewayMovieGetOneResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) MovieCreate(context.Context, *GatewayMovieCreateRequest) (*GatewayMovieCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieCreate not implemented")
}
func (UnimplementedGatewayServer) MovieList(context.Context, *GatewayMovieListRequest) (*GatewayMovieListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieList not implemented")
}
func (UnimplementedGatewayServer) MovieUpdate(context.Context, *GatewayMovieUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieUpdate not implemented")
}
func (UnimplementedGatewayServer) MovieDelete(context.Context, *GatewayMovieDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieDelete not implemented")
}
func (UnimplementedGatewayServer) MovieGetOne(context.Context, *GatewayMovieGetOneRequest) (*GatewayMovieGetOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieGetOne not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_MovieCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayMovieCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).MovieCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.homework.api.gateway.Gateway/MovieCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).MovieCreate(ctx, req.(*GatewayMovieCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_MovieList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayMovieListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).MovieList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.homework.api.gateway.Gateway/MovieList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).MovieList(ctx, req.(*GatewayMovieListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_MovieUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayMovieUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).MovieUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.homework.api.gateway.Gateway/MovieUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).MovieUpdate(ctx, req.(*GatewayMovieUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_MovieDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayMovieDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).MovieDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.homework.api.gateway.Gateway/MovieDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).MovieDelete(ctx, req.(*GatewayMovieDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_MovieGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayMovieGetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).MovieGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozon.dev.homework.api.gateway.Gateway/MovieGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).MovieGetOne(ctx, req.(*GatewayMovieGetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozon.dev.homework.api.gateway.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MovieCreate",
			Handler:    _Gateway_MovieCreate_Handler,
		},
		{
			MethodName: "MovieList",
			Handler:    _Gateway_MovieList_Handler,
		},
		{
			MethodName: "MovieUpdate",
			Handler:    _Gateway_MovieUpdate_Handler,
		},
		{
			MethodName: "MovieDelete",
			Handler:    _Gateway_MovieDelete_Handler,
		},
		{
			MethodName: "MovieGetOne",
			Handler:    _Gateway_MovieGetOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/gateway.proto",
}
