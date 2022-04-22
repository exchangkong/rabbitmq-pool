// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pool.proto

package pb

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

// RabbitmqPoolServiceClient is the client API for RabbitmqPoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RabbitmqPoolServiceClient interface {
	Pushlish(ctx context.Context, in *PublishReuqest, opts ...grpc.CallOption) (*PublishResponse, error)
}

type rabbitmqPoolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRabbitmqPoolServiceClient(cc grpc.ClientConnInterface) RabbitmqPoolServiceClient {
	return &rabbitmqPoolServiceClient{cc}
}

func (c *rabbitmqPoolServiceClient) Pushlish(ctx context.Context, in *PublishReuqest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/RabbitmqPoolService/Pushlish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RabbitmqPoolServiceServer is the server API for RabbitmqPoolService service.
// All implementations must embed UnimplementedRabbitmqPoolServiceServer
// for forward compatibility
type RabbitmqPoolServiceServer interface {
	Pushlish(context.Context, *PublishReuqest) (*PublishResponse, error)
	mustEmbedUnimplementedRabbitmqPoolServiceServer()
}

// UnimplementedRabbitmqPoolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRabbitmqPoolServiceServer struct {
}

func (UnimplementedRabbitmqPoolServiceServer) Pushlish(context.Context, *PublishReuqest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pushlish not implemented")
}
func (UnimplementedRabbitmqPoolServiceServer) mustEmbedUnimplementedRabbitmqPoolServiceServer() {}

// UnsafeRabbitmqPoolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RabbitmqPoolServiceServer will
// result in compilation errors.
type UnsafeRabbitmqPoolServiceServer interface {
	mustEmbedUnimplementedRabbitmqPoolServiceServer()
}

func RegisterRabbitmqPoolServiceServer(s grpc.ServiceRegistrar, srv RabbitmqPoolServiceServer) {
	s.RegisterService(&RabbitmqPoolService_ServiceDesc, srv)
}

func _RabbitmqPoolService_Pushlish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishReuqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RabbitmqPoolServiceServer).Pushlish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RabbitmqPoolService/Pushlish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RabbitmqPoolServiceServer).Pushlish(ctx, req.(*PublishReuqest))
	}
	return interceptor(ctx, in, info, handler)
}

// RabbitmqPoolService_ServiceDesc is the grpc.ServiceDesc for RabbitmqPoolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RabbitmqPoolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RabbitmqPoolService",
	HandlerType: (*RabbitmqPoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pushlish",
			Handler:    _RabbitmqPoolService_Pushlish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pool.proto",
}
