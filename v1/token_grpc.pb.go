// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: oauthapp/v1/token.proto

package oauthappv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TokeService_GetToken_FullMethodName     = "/oauthapp.v1.TokeService/GetToken"
	TokeService_GetTokenInfo_FullMethodName = "/oauthapp.v1.TokeService/GetTokenInfo"
)

// TokeServiceClient is the client API for TokeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokeServiceClient interface {
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	GetTokenInfo(ctx context.Context, in *GetTokenInfoRequest, opts ...grpc.CallOption) (*GetTokenInfoResponse, error)
}

type tokeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokeServiceClient(cc grpc.ClientConnInterface) TokeServiceClient {
	return &tokeServiceClient{cc}
}

func (c *tokeServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, TokeService_GetToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokeServiceClient) GetTokenInfo(ctx context.Context, in *GetTokenInfoRequest, opts ...grpc.CallOption) (*GetTokenInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTokenInfoResponse)
	err := c.cc.Invoke(ctx, TokeService_GetTokenInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokeServiceServer is the server API for TokeService service.
// All implementations must embed UnimplementedTokeServiceServer
// for forward compatibility.
type TokeServiceServer interface {
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	GetTokenInfo(context.Context, *GetTokenInfoRequest) (*GetTokenInfoResponse, error)
	mustEmbedUnimplementedTokeServiceServer()
}

// UnimplementedTokeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTokeServiceServer struct{}

func (UnimplementedTokeServiceServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedTokeServiceServer) GetTokenInfo(context.Context, *GetTokenInfoRequest) (*GetTokenInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenInfo not implemented")
}
func (UnimplementedTokeServiceServer) mustEmbedUnimplementedTokeServiceServer() {}
func (UnimplementedTokeServiceServer) testEmbeddedByValue()                     {}

// UnsafeTokeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokeServiceServer will
// result in compilation errors.
type UnsafeTokeServiceServer interface {
	mustEmbedUnimplementedTokeServiceServer()
}

func RegisterTokeServiceServer(s grpc.ServiceRegistrar, srv TokeServiceServer) {
	// If the following call pancis, it indicates UnimplementedTokeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TokeService_ServiceDesc, srv)
}

func _TokeService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokeServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokeService_GetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokeServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokeService_GetTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokeServiceServer).GetTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokeService_GetTokenInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokeServiceServer).GetTokenInfo(ctx, req.(*GetTokenInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokeService_ServiceDesc is the grpc.ServiceDesc for TokeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oauthapp.v1.TokeService",
	HandlerType: (*TokeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _TokeService_GetToken_Handler,
		},
		{
			MethodName: "GetTokenInfo",
			Handler:    _TokeService_GetTokenInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauthapp/v1/token.proto",
}