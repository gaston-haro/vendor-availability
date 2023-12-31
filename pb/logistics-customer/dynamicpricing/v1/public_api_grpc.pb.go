// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: dynamicpricing/v1/public_api.proto

package v1

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

const (
	PublicAPI_GetMultipleVendorsFee_FullMethodName = "/dynamicpricing.v1.PublicAPI/GetMultipleVendorsFee"
	PublicAPI_GetSingleVendorFee_FullMethodName    = "/dynamicpricing.v1.PublicAPI/GetSingleVendorFee"
)

// PublicAPIClient is the client API for PublicAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicAPIClient interface {
	// Retrieve fee information for multiple vendors.
	GetMultipleVendorsFee(ctx context.Context, in *GetMultipleVendorsFeeRequest, opts ...grpc.CallOption) (*GetMultipleVendorsFeeResponse, error)
	// Retrieve fee information for a single vendor.
	GetSingleVendorFee(ctx context.Context, in *GetSingleVendorFeeRequest, opts ...grpc.CallOption) (*GetSingleVendorFeeResponse, error)
}

type publicAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicAPIClient(cc grpc.ClientConnInterface) PublicAPIClient {
	return &publicAPIClient{cc}
}

func (c *publicAPIClient) GetMultipleVendorsFee(ctx context.Context, in *GetMultipleVendorsFeeRequest, opts ...grpc.CallOption) (*GetMultipleVendorsFeeResponse, error) {
	out := new(GetMultipleVendorsFeeResponse)
	err := c.cc.Invoke(ctx, PublicAPI_GetMultipleVendorsFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicAPIClient) GetSingleVendorFee(ctx context.Context, in *GetSingleVendorFeeRequest, opts ...grpc.CallOption) (*GetSingleVendorFeeResponse, error) {
	out := new(GetSingleVendorFeeResponse)
	err := c.cc.Invoke(ctx, PublicAPI_GetSingleVendorFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicAPIServer is the server API for PublicAPI service.
// All implementations must embed UnimplementedPublicAPIServer
// for forward compatibility
type PublicAPIServer interface {
	// Retrieve fee information for multiple vendors.
	GetMultipleVendorsFee(context.Context, *GetMultipleVendorsFeeRequest) (*GetMultipleVendorsFeeResponse, error)
	// Retrieve fee information for a single vendor.
	GetSingleVendorFee(context.Context, *GetSingleVendorFeeRequest) (*GetSingleVendorFeeResponse, error)
	mustEmbedUnimplementedPublicAPIServer()
}

// UnimplementedPublicAPIServer must be embedded to have forward compatible implementations.
type UnimplementedPublicAPIServer struct {
}

func (UnimplementedPublicAPIServer) GetMultipleVendorsFee(context.Context, *GetMultipleVendorsFeeRequest) (*GetMultipleVendorsFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipleVendorsFee not implemented")
}
func (UnimplementedPublicAPIServer) GetSingleVendorFee(context.Context, *GetSingleVendorFeeRequest) (*GetSingleVendorFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleVendorFee not implemented")
}
func (UnimplementedPublicAPIServer) mustEmbedUnimplementedPublicAPIServer() {}

// UnsafePublicAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicAPIServer will
// result in compilation errors.
type UnsafePublicAPIServer interface {
	mustEmbedUnimplementedPublicAPIServer()
}

func RegisterPublicAPIServer(s grpc.ServiceRegistrar, srv PublicAPIServer) {
	s.RegisterService(&PublicAPI_ServiceDesc, srv)
}

func _PublicAPI_GetMultipleVendorsFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipleVendorsFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicAPIServer).GetMultipleVendorsFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicAPI_GetMultipleVendorsFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicAPIServer).GetMultipleVendorsFee(ctx, req.(*GetMultipleVendorsFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicAPI_GetSingleVendorFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleVendorFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicAPIServer).GetSingleVendorFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicAPI_GetSingleVendorFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicAPIServer).GetSingleVendorFee(ctx, req.(*GetSingleVendorFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicAPI_ServiceDesc is the grpc.ServiceDesc for PublicAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dynamicpricing.v1.PublicAPI",
	HandlerType: (*PublicAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMultipleVendorsFee",
			Handler:    _PublicAPI_GetMultipleVendorsFee_Handler,
		},
		{
			MethodName: "GetSingleVendorFee",
			Handler:    _PublicAPI_GetSingleVendorFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dynamicpricing/v1/public_api.proto",
}
