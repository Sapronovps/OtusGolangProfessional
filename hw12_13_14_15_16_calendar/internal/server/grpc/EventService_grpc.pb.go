// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: EventService.proto

package internalgrpc

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
	EventService_CreateEvent_FullMethodName = "/event.EventService/CreateEvent"
	EventService_GetEvent_FullMethodName    = "/event.EventService/GetEvent"
	EventService_UpdateEvent_FullMethodName = "/event.EventService/UpdateEvent"
	EventService_DeleteEvent_FullMethodName = "/event.EventService/DeleteEvent"
	EventService_ListByDay_FullMethodName   = "/event.EventService/ListByDay"
	EventService_ListByWeek_FullMethodName  = "/event.EventService/ListByWeek"
	EventService_ListByMonth_FullMethodName = "/event.EventService/ListByMonth"
)

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	CreateEvent(ctx context.Context, in *RequestCreateEvent, opts ...grpc.CallOption) (*Event, error)
	GetEvent(ctx context.Context, in *RequestGetEvent, opts ...grpc.CallOption) (*Event, error)
	UpdateEvent(ctx context.Context, in *RequestUpdateEvent, opts ...grpc.CallOption) (*Event, error)
	DeleteEvent(ctx context.Context, in *RequestDeleteEvent, opts ...grpc.CallOption) (*Event, error)
	ListByDay(ctx context.Context, in *RequestListByDate, opts ...grpc.CallOption) (*ResponseListByDate, error)
	ListByWeek(ctx context.Context, in *RequestListByDate, opts ...grpc.CallOption) (*ResponseListByDate, error)
	ListByMonth(ctx context.Context, in *RequestListByDate, opts ...grpc.CallOption) (*ResponseListByDate, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) CreateEvent(ctx context.Context, in *RequestCreateEvent, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventService_CreateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEvent(ctx context.Context, in *RequestGetEvent, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventService_GetEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) UpdateEvent(ctx context.Context, in *RequestUpdateEvent, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventService_UpdateEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) DeleteEvent(ctx context.Context, in *RequestDeleteEvent, opts ...grpc.CallOption) (*Event, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Event)
	err := c.cc.Invoke(ctx, EventService_DeleteEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) ListByDay(ctx context.Context, in *RequestListByDate, opts ...grpc.CallOption) (*ResponseListByDate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseListByDate)
	err := c.cc.Invoke(ctx, EventService_ListByDay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) ListByWeek(ctx context.Context, in *RequestListByDate, opts ...grpc.CallOption) (*ResponseListByDate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseListByDate)
	err := c.cc.Invoke(ctx, EventService_ListByWeek_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) ListByMonth(ctx context.Context, in *RequestListByDate, opts ...grpc.CallOption) (*ResponseListByDate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseListByDate)
	err := c.cc.Invoke(ctx, EventService_ListByMonth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility.
type EventServiceServer interface {
	CreateEvent(context.Context, *RequestCreateEvent) (*Event, error)
	GetEvent(context.Context, *RequestGetEvent) (*Event, error)
	UpdateEvent(context.Context, *RequestUpdateEvent) (*Event, error)
	DeleteEvent(context.Context, *RequestDeleteEvent) (*Event, error)
	ListByDay(context.Context, *RequestListByDate) (*ResponseListByDate, error)
	ListByWeek(context.Context, *RequestListByDate) (*ResponseListByDate, error)
	ListByMonth(context.Context, *RequestListByDate) (*ResponseListByDate, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEventServiceServer struct{}

func (UnimplementedEventServiceServer) CreateEvent(context.Context, *RequestCreateEvent) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedEventServiceServer) GetEvent(context.Context, *RequestGetEvent) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventServiceServer) UpdateEvent(context.Context, *RequestUpdateEvent) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventServiceServer) DeleteEvent(context.Context, *RequestDeleteEvent) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedEventServiceServer) ListByDay(context.Context, *RequestListByDate) (*ResponseListByDate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByDay not implemented")
}
func (UnimplementedEventServiceServer) ListByWeek(context.Context, *RequestListByDate) (*ResponseListByDate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByWeek not implemented")
}
func (UnimplementedEventServiceServer) ListByMonth(context.Context, *RequestListByDate) (*ResponseListByDate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByMonth not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}
func (UnimplementedEventServiceServer) testEmbeddedByValue()                      {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	// If the following call pancis, it indicates UnimplementedEventServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateEvent(ctx, req.(*RequestCreateEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEvent(ctx, req.(*RequestGetEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).UpdateEvent(ctx, req.(*RequestUpdateEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).DeleteEvent(ctx, req.(*RequestDeleteEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_ListByDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestListByDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).ListByDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_ListByDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).ListByDay(ctx, req.(*RequestListByDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_ListByWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestListByDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).ListByWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_ListByWeek_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).ListByWeek(ctx, req.(*RequestListByDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_ListByMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestListByDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).ListByMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_ListByMonth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).ListByMonth(ctx, req.(*RequestListByDate))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventService_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _EventService_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _EventService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _EventService_DeleteEvent_Handler,
		},
		{
			MethodName: "ListByDay",
			Handler:    _EventService_ListByDay_Handler,
		},
		{
			MethodName: "ListByWeek",
			Handler:    _EventService_ListByWeek_Handler,
		},
		{
			MethodName: "ListByMonth",
			Handler:    _EventService_ListByMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "EventService.proto",
}
