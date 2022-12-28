// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: firstfile.proto

package proto

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FirstServiceClient is the client API for FirstService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FirstServiceClient interface {
}

type firstServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFirstServiceClient(cc grpc.ClientConnInterface) FirstServiceClient {
	return &firstServiceClient{cc}
}

// FirstServiceServer is the server API for FirstService service.
// All implementations must embed UnimplementedFirstServiceServer
// for forward compatibility
type FirstServiceServer interface {
	mustEmbedUnimplementedFirstServiceServer()
}

// UnimplementedFirstServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFirstServiceServer struct {
}

func (UnimplementedFirstServiceServer) mustEmbedUnimplementedFirstServiceServer() {}

// UnsafeFirstServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FirstServiceServer will
// result in compilation errors.
type UnsafeFirstServiceServer interface {
	mustEmbedUnimplementedFirstServiceServer()
}

func RegisterFirstServiceServer(s grpc.ServiceRegistrar, srv FirstServiceServer) {
	s.RegisterService(&FirstService_ServiceDesc, srv)
}

// FirstService_ServiceDesc is the grpc.ServiceDesc for FirstService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FirstService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.FirstService",
	HandlerType: (*FirstServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "firstfile.proto",
}
