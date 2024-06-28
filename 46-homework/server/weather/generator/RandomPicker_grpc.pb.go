// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: RandomPicker.proto

package prooto

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

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	ReportWeatherCondition(ctx context.Context, in *ReportWeatherRequest, opts ...grpc.CallOption) (*ReportWeatherReponse, error)
	GetCurrentWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error)
	GetWeatherForecast(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherForecastResponse, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) ReportWeatherCondition(ctx context.Context, in *ReportWeatherRequest, opts ...grpc.CallOption) (*ReportWeatherReponse, error) {
	out := new(ReportWeatherReponse)
	err := c.cc.Invoke(ctx, "/proto.WeatherService/ReportWeatherCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetCurrentWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error) {
	out := new(WeatherResponse)
	err := c.cc.Invoke(ctx, "/proto.WeatherService/GetCurrentWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetWeatherForecast(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherForecastResponse, error) {
	out := new(WeatherForecastResponse)
	err := c.cc.Invoke(ctx, "/proto.WeatherService/GetWeatherForecast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	ReportWeatherCondition(context.Context, *ReportWeatherRequest) (*ReportWeatherReponse, error)
	GetCurrentWeather(context.Context, *WeatherRequest) (*WeatherResponse, error)
	GetWeatherForecast(context.Context, *WeatherRequest) (*WeatherForecastResponse, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) ReportWeatherCondition(context.Context, *ReportWeatherRequest) (*ReportWeatherReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportWeatherCondition not implemented")
}
func (UnimplementedWeatherServiceServer) GetCurrentWeather(context.Context, *WeatherRequest) (*WeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentWeather not implemented")
}
func (UnimplementedWeatherServiceServer) GetWeatherForecast(context.Context, *WeatherRequest) (*WeatherForecastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeatherForecast not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_ReportWeatherCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).ReportWeatherCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WeatherService/ReportWeatherCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).ReportWeatherCondition(ctx, req.(*ReportWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetCurrentWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetCurrentWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WeatherService/GetCurrentWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetCurrentWeather(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetWeatherForecast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetWeatherForecast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WeatherService/GetWeatherForecast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetWeatherForecast(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportWeatherCondition",
			Handler:    _WeatherService_ReportWeatherCondition_Handler,
		},
		{
			MethodName: "GetCurrentWeather",
			Handler:    _WeatherService_GetCurrentWeather_Handler,
		},
		{
			MethodName: "GetWeatherForecast",
			Handler:    _WeatherService_GetWeatherForecast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RandomPicker.proto",
}