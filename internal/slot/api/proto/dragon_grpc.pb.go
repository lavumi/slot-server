// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: dragon.proto

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

// DragonReelClient is the client API for DragonReel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DragonReelClient interface {
	Spin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DragonReelResult, error)
}

type dragonReelClient struct {
	cc grpc.ClientConnInterface
}

func NewDragonReelClient(cc grpc.ClientConnInterface) DragonReelClient {
	return &dragonReelClient{cc}
}

func (c *dragonReelClient) Spin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DragonReelResult, error) {
	out := new(DragonReelResult)
	err := c.cc.Invoke(ctx, "/slot.DragonReel/Spin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DragonReelServer is the server API for DragonReel service.
// All implementations must embed UnimplementedDragonReelServer
// for forward compatibility
type DragonReelServer interface {
	Spin(context.Context, *Request) (*DragonReelResult, error)
	mustEmbedUnimplementedDragonReelServer()
}

// UnimplementedDragonReelServer must be embedded to have forward compatible implementations.
type UnimplementedDragonReelServer struct {
}

func (UnimplementedDragonReelServer) Spin(context.Context, *Request) (*DragonReelResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Spin not implemented")
}
func (UnimplementedDragonReelServer) mustEmbedUnimplementedDragonReelServer() {}

// UnsafeDragonReelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DragonReelServer will
// result in compilation errors.
type UnsafeDragonReelServer interface {
	mustEmbedUnimplementedDragonReelServer()
}

func RegisterDragonReelServer(s grpc.ServiceRegistrar, srv DragonReelServer) {
	s.RegisterService(&DragonReel_ServiceDesc, srv)
}

func _DragonReel_Spin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DragonReelServer).Spin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slot.DragonReel/Spin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DragonReelServer).Spin(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// DragonReel_ServiceDesc is the grpc.ServiceDesc for DragonReel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DragonReel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slot.DragonReel",
	HandlerType: (*DragonReelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Spin",
			Handler:    _DragonReel_Spin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dragon.proto",
}
