// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/vendor.proto

package vendor

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
	VendorService_GetVendor_FullMethodName    = "/vendor.VendorService/GetVendor"
	VendorService_CreateVendor_FullMethodName = "/vendor.VendorService/CreateVendor"
)

// VendorServiceClient is the client API for VendorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Vendor service definition
type VendorServiceClient interface {
	GetVendor(ctx context.Context, in *VendorRequest, opts ...grpc.CallOption) (*VendorResponse, error)
	CreateVendor(ctx context.Context, in *CreateVendorRequest, opts ...grpc.CallOption) (*CreateVendorResponse, error)
}

type vendorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVendorServiceClient(cc grpc.ClientConnInterface) VendorServiceClient {
	return &vendorServiceClient{cc}
}

func (c *vendorServiceClient) GetVendor(ctx context.Context, in *VendorRequest, opts ...grpc.CallOption) (*VendorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VendorResponse)
	err := c.cc.Invoke(ctx, VendorService_GetVendor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vendorServiceClient) CreateVendor(ctx context.Context, in *CreateVendorRequest, opts ...grpc.CallOption) (*CreateVendorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVendorResponse)
	err := c.cc.Invoke(ctx, VendorService_CreateVendor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VendorServiceServer is the server API for VendorService service.
// All implementations must embed UnimplementedVendorServiceServer
// for forward compatibility.
//
// Vendor service definition
type VendorServiceServer interface {
	GetVendor(context.Context, *VendorRequest) (*VendorResponse, error)
	CreateVendor(context.Context, *CreateVendorRequest) (*CreateVendorResponse, error)
	mustEmbedUnimplementedVendorServiceServer()
}

// UnimplementedVendorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVendorServiceServer struct{}

func (UnimplementedVendorServiceServer) GetVendor(context.Context, *VendorRequest) (*VendorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVendor not implemented")
}
func (UnimplementedVendorServiceServer) CreateVendor(context.Context, *CreateVendorRequest) (*CreateVendorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVendor not implemented")
}
func (UnimplementedVendorServiceServer) mustEmbedUnimplementedVendorServiceServer() {}
func (UnimplementedVendorServiceServer) testEmbeddedByValue()                       {}

// UnsafeVendorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendorServiceServer will
// result in compilation errors.
type UnsafeVendorServiceServer interface {
	mustEmbedUnimplementedVendorServiceServer()
}

func RegisterVendorServiceServer(s grpc.ServiceRegistrar, srv VendorServiceServer) {
	// If the following call pancis, it indicates UnimplementedVendorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VendorService_ServiceDesc, srv)
}

func _VendorService_GetVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VendorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorServiceServer).GetVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorService_GetVendor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorServiceServer).GetVendor(ctx, req.(*VendorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VendorService_CreateVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVendorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorServiceServer).CreateVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorService_CreateVendor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorServiceServer).CreateVendor(ctx, req.(*CreateVendorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VendorService_ServiceDesc is the grpc.ServiceDesc for VendorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vendor.VendorService",
	HandlerType: (*VendorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVendor",
			Handler:    _VendorService_GetVendor_Handler,
		},
		{
			MethodName: "CreateVendor",
			Handler:    _VendorService_CreateVendor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vendor.proto",
}