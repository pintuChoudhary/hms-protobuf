// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gateway/registration.proto

package apigateway

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
	ServiceRegistry_Register_FullMethodName   = "/auditlogging.ServiceRegistry/Register"
	ServiceRegistry_Deregister_FullMethodName = "/auditlogging.ServiceRegistry/Deregister"
)

// ServiceRegistryClient is the client API for ServiceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceRegistryClient interface {
	Register(ctx context.Context, in *ServiceRegistrationRequest, opts ...grpc.CallOption) (*ServiceRegistrationResponse, error)
	Deregister(ctx context.Context, in *ServiceDeregistrationRequest, opts ...grpc.CallOption) (*ServiceDeregistrationResponse, error)
}

type serviceRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceRegistryClient(cc grpc.ClientConnInterface) ServiceRegistryClient {
	return &serviceRegistryClient{cc}
}

func (c *serviceRegistryClient) Register(ctx context.Context, in *ServiceRegistrationRequest, opts ...grpc.CallOption) (*ServiceRegistrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceRegistrationResponse)
	err := c.cc.Invoke(ctx, ServiceRegistry_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRegistryClient) Deregister(ctx context.Context, in *ServiceDeregistrationRequest, opts ...grpc.CallOption) (*ServiceDeregistrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceDeregistrationResponse)
	err := c.cc.Invoke(ctx, ServiceRegistry_Deregister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceRegistryServer is the server API for ServiceRegistry service.
// All implementations must embed UnimplementedServiceRegistryServer
// for forward compatibility.
type ServiceRegistryServer interface {
	Register(context.Context, *ServiceRegistrationRequest) (*ServiceRegistrationResponse, error)
	Deregister(context.Context, *ServiceDeregistrationRequest) (*ServiceDeregistrationResponse, error)
	mustEmbedUnimplementedServiceRegistryServer()
}

// UnimplementedServiceRegistryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceRegistryServer struct{}

func (UnimplementedServiceRegistryServer) Register(context.Context, *ServiceRegistrationRequest) (*ServiceRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedServiceRegistryServer) Deregister(context.Context, *ServiceDeregistrationRequest) (*ServiceDeregistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedServiceRegistryServer) mustEmbedUnimplementedServiceRegistryServer() {}
func (UnimplementedServiceRegistryServer) testEmbeddedByValue()                         {}

// UnsafeServiceRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceRegistryServer will
// result in compilation errors.
type UnsafeServiceRegistryServer interface {
	mustEmbedUnimplementedServiceRegistryServer()
}

func RegisterServiceRegistryServer(s grpc.ServiceRegistrar, srv ServiceRegistryServer) {
	// If the following call pancis, it indicates UnimplementedServiceRegistryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServiceRegistry_ServiceDesc, srv)
}

func _ServiceRegistry_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRegistryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceRegistry_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRegistryServer).Register(ctx, req.(*ServiceRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRegistry_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceDeregistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRegistryServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceRegistry_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRegistryServer).Deregister(ctx, req.(*ServiceDeregistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceRegistry_ServiceDesc is the grpc.ServiceDesc for ServiceRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auditlogging.ServiceRegistry",
	HandlerType: (*ServiceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ServiceRegistry_Register_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _ServiceRegistry_Deregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/registration.proto",
}
