// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: swim.proto

package swim

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
	FailureDetector_Ping_FullMethodName             = "/swim.FailureDetector/Ping"
	FailureDetector_IndirectPing_FullMethodName     = "/swim.FailureDetector/IndirectPing"
	FailureDetector_UpdateMembership_FullMethodName = "/swim.FailureDetector/UpdateMembership"
)

// FailureDetectorClient is the client API for FailureDetector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Existing FailureDetector service plus the new UpdateMembership RPC.
type FailureDetectorClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	IndirectPing(ctx context.Context, in *IndirectPingRequest, opts ...grpc.CallOption) (*IndirectPingResponse, error)
	UpdateMembership(ctx context.Context, in *MembershipUpdateRequest, opts ...grpc.CallOption) (*MembershipUpdateResponse, error)
}

type failureDetectorClient struct {
	cc grpc.ClientConnInterface
}

func NewFailureDetectorClient(cc grpc.ClientConnInterface) FailureDetectorClient {
	return &failureDetectorClient{cc}
}

func (c *failureDetectorClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, FailureDetector_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *failureDetectorClient) IndirectPing(ctx context.Context, in *IndirectPingRequest, opts ...grpc.CallOption) (*IndirectPingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IndirectPingResponse)
	err := c.cc.Invoke(ctx, FailureDetector_IndirectPing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *failureDetectorClient) UpdateMembership(ctx context.Context, in *MembershipUpdateRequest, opts ...grpc.CallOption) (*MembershipUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MembershipUpdateResponse)
	err := c.cc.Invoke(ctx, FailureDetector_UpdateMembership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FailureDetectorServer is the server API for FailureDetector service.
// All implementations must embed UnimplementedFailureDetectorServer
// for forward compatibility.
//
// Existing FailureDetector service plus the new UpdateMembership RPC.
type FailureDetectorServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	IndirectPing(context.Context, *IndirectPingRequest) (*IndirectPingResponse, error)
	UpdateMembership(context.Context, *MembershipUpdateRequest) (*MembershipUpdateResponse, error)
	mustEmbedUnimplementedFailureDetectorServer()
}

// UnimplementedFailureDetectorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFailureDetectorServer struct{}

func (UnimplementedFailureDetectorServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedFailureDetectorServer) IndirectPing(context.Context, *IndirectPingRequest) (*IndirectPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndirectPing not implemented")
}
func (UnimplementedFailureDetectorServer) UpdateMembership(context.Context, *MembershipUpdateRequest) (*MembershipUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMembership not implemented")
}
func (UnimplementedFailureDetectorServer) mustEmbedUnimplementedFailureDetectorServer() {}
func (UnimplementedFailureDetectorServer) testEmbeddedByValue()                         {}

// UnsafeFailureDetectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FailureDetectorServer will
// result in compilation errors.
type UnsafeFailureDetectorServer interface {
	mustEmbedUnimplementedFailureDetectorServer()
}

func RegisterFailureDetectorServer(s grpc.ServiceRegistrar, srv FailureDetectorServer) {
	// If the following call pancis, it indicates UnimplementedFailureDetectorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FailureDetector_ServiceDesc, srv)
}

func _FailureDetector_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailureDetectorServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FailureDetector_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailureDetectorServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FailureDetector_IndirectPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndirectPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailureDetectorServer).IndirectPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FailureDetector_IndirectPing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailureDetectorServer).IndirectPing(ctx, req.(*IndirectPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FailureDetector_UpdateMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MembershipUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailureDetectorServer).UpdateMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FailureDetector_UpdateMembership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailureDetectorServer).UpdateMembership(ctx, req.(*MembershipUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FailureDetector_ServiceDesc is the grpc.ServiceDesc for FailureDetector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FailureDetector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "swim.FailureDetector",
	HandlerType: (*FailureDetectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _FailureDetector_Ping_Handler,
		},
		{
			MethodName: "IndirectPing",
			Handler:    _FailureDetector_IndirectPing_Handler,
		},
		{
			MethodName: "UpdateMembership",
			Handler:    _FailureDetector_UpdateMembership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swim.proto",
}

const (
	Dissemination_Disseminate_FullMethodName = "/swim.Dissemination/Disseminate"
	Dissemination_Join_FullMethodName        = "/swim.Dissemination/Join"
)

// DisseminationClient is the client API for Dissemination service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DisseminationClient interface {
	Disseminate(ctx context.Context, in *DisseminationRequest, opts ...grpc.CallOption) (*DisseminationResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
}

type disseminationClient struct {
	cc grpc.ClientConnInterface
}

func NewDisseminationClient(cc grpc.ClientConnInterface) DisseminationClient {
	return &disseminationClient{cc}
}

func (c *disseminationClient) Disseminate(ctx context.Context, in *DisseminationRequest, opts ...grpc.CallOption) (*DisseminationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisseminationResponse)
	err := c.cc.Invoke(ctx, Dissemination_Disseminate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disseminationClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, Dissemination_Join_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisseminationServer is the server API for Dissemination service.
// All implementations must embed UnimplementedDisseminationServer
// for forward compatibility.
type DisseminationServer interface {
	Disseminate(context.Context, *DisseminationRequest) (*DisseminationResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	mustEmbedUnimplementedDisseminationServer()
}

// UnimplementedDisseminationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDisseminationServer struct{}

func (UnimplementedDisseminationServer) Disseminate(context.Context, *DisseminationRequest) (*DisseminationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disseminate not implemented")
}
func (UnimplementedDisseminationServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedDisseminationServer) mustEmbedUnimplementedDisseminationServer() {}
func (UnimplementedDisseminationServer) testEmbeddedByValue()                       {}

// UnsafeDisseminationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisseminationServer will
// result in compilation errors.
type UnsafeDisseminationServer interface {
	mustEmbedUnimplementedDisseminationServer()
}

func RegisterDisseminationServer(s grpc.ServiceRegistrar, srv DisseminationServer) {
	// If the following call pancis, it indicates UnimplementedDisseminationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Dissemination_ServiceDesc, srv)
}

func _Dissemination_Disseminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisseminationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisseminationServer).Disseminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dissemination_Disseminate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisseminationServer).Disseminate(ctx, req.(*DisseminationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dissemination_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisseminationServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dissemination_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisseminationServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dissemination_ServiceDesc is the grpc.ServiceDesc for Dissemination service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dissemination_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "swim.Dissemination",
	HandlerType: (*DisseminationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Disseminate",
			Handler:    _Dissemination_Disseminate_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Dissemination_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swim.proto",
}
