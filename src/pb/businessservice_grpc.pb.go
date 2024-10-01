// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: businessservice.proto

package pb

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
	Business_GetBusiness_FullMethodName    = "/businessservice.Business/GetBusiness"
	Business_CreateBusiness_FullMethodName = "/businessservice.Business/CreateBusiness"
)

// BusinessClient is the client API for Business service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessClient interface {
	GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*GetBusinessResponse, error)
	CreateBusiness(ctx context.Context, in *CreateBusinessRequest, opts ...grpc.CallOption) (*CreateBusinessResponse, error)
}

type businessClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessClient(cc grpc.ClientConnInterface) BusinessClient {
	return &businessClient{cc}
}

func (c *businessClient) GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*GetBusinessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBusinessResponse)
	err := c.cc.Invoke(ctx, Business_GetBusiness_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessClient) CreateBusiness(ctx context.Context, in *CreateBusinessRequest, opts ...grpc.CallOption) (*CreateBusinessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBusinessResponse)
	err := c.cc.Invoke(ctx, Business_CreateBusiness_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServer is the server API for Business service.
// All implementations must embed UnimplementedBusinessServer
// for forward compatibility.
type BusinessServer interface {
	GetBusiness(context.Context, *GetBusinessRequest) (*GetBusinessResponse, error)
	CreateBusiness(context.Context, *CreateBusinessRequest) (*CreateBusinessResponse, error)
	mustEmbedUnimplementedBusinessServer()
}

// UnimplementedBusinessServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBusinessServer struct{}

func (UnimplementedBusinessServer) GetBusiness(context.Context, *GetBusinessRequest) (*GetBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusiness not implemented")
}
func (UnimplementedBusinessServer) CreateBusiness(context.Context, *CreateBusinessRequest) (*CreateBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusiness not implemented")
}
func (UnimplementedBusinessServer) mustEmbedUnimplementedBusinessServer() {}
func (UnimplementedBusinessServer) testEmbeddedByValue()                  {}

// UnsafeBusinessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessServer will
// result in compilation errors.
type UnsafeBusinessServer interface {
	mustEmbedUnimplementedBusinessServer()
}

func RegisterBusinessServer(s grpc.ServiceRegistrar, srv BusinessServer) {
	// If the following call pancis, it indicates UnimplementedBusinessServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Business_ServiceDesc, srv)
}

func _Business_GetBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServer).GetBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_GetBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServer).GetBusiness(ctx, req.(*GetBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Business_CreateBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServer).CreateBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Business_CreateBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServer).CreateBusiness(ctx, req.(*CreateBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Business_ServiceDesc is the grpc.ServiceDesc for Business service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Business_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "businessservice.Business",
	HandlerType: (*BusinessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusiness",
			Handler:    _Business_GetBusiness_Handler,
		},
		{
			MethodName: "CreateBusiness",
			Handler:    _Business_CreateBusiness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "businessservice.proto",
}
