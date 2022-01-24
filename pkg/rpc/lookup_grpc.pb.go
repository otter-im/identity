// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

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

// LookupServiceClient is the client API for LookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LookupServiceClient interface {
	Authorize(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error)
}

type lookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLookupServiceClient(cc grpc.ClientConnInterface) LookupServiceClient {
	return &lookupServiceClient{cc}
}

func (c *lookupServiceClient) Authorize(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error) {
	out := new(AuthorizationResponse)
	err := c.cc.Invoke(ctx, "/service.LookupService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LookupServiceServer is the server API for LookupService service.
// All implementations must embed UnimplementedLookupServiceServer
// for forward compatibility
type LookupServiceServer interface {
	Authorize(context.Context, *AuthorizationRequest) (*AuthorizationResponse, error)
	mustEmbedUnimplementedLookupServiceServer()
}

// UnimplementedLookupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLookupServiceServer struct {
}

func (UnimplementedLookupServiceServer) Authorize(context.Context, *AuthorizationRequest) (*AuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedLookupServiceServer) mustEmbedUnimplementedLookupServiceServer() {}

// UnsafeLookupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LookupServiceServer will
// result in compilation errors.
type UnsafeLookupServiceServer interface {
	mustEmbedUnimplementedLookupServiceServer()
}

func RegisterLookupServiceServer(s grpc.ServiceRegistrar, srv LookupServiceServer) {
	s.RegisterService(&LookupService_ServiceDesc, srv)
}

func _LookupService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.LookupService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServiceServer).Authorize(ctx, req.(*AuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LookupService_ServiceDesc is the grpc.ServiceDesc for LookupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LookupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.LookupService",
	HandlerType: (*LookupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _LookupService_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/rpc/lookup.proto",
}
