// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: f-taxes.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FTaxes_SubmitTransaction_FullMethodName = "/FTaxesGrpc.FTaxes/SubmitTransaction"
	FTaxes_SubmitTransfer_FullMethodName    = "/FTaxesGrpc.FTaxes/SubmitTransfer"
	FTaxes_SubmitGenericFee_FullMethodName  = "/FTaxesGrpc.FTaxes/SubmitGenericFee"
	FTaxes_ShowJobProgress_FullMethodName   = "/FTaxesGrpc.FTaxes/ShowJobProgress"
)

// FTaxesClient is the client API for FTaxes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FTaxesClient interface {
	SubmitTransaction(ctx context.Context, in *SrcTx, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SubmitTransfer(ctx context.Context, in *SrcTransfer, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SubmitGenericFee(ctx context.Context, in *SrcGenericFee, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ShowJobProgress(ctx context.Context, in *JobProgress, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type fTaxesClient struct {
	cc grpc.ClientConnInterface
}

func NewFTaxesClient(cc grpc.ClientConnInterface) FTaxesClient {
	return &fTaxesClient{cc}
}

func (c *fTaxesClient) SubmitTransaction(ctx context.Context, in *SrcTx, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FTaxes_SubmitTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fTaxesClient) SubmitTransfer(ctx context.Context, in *SrcTransfer, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FTaxes_SubmitTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fTaxesClient) SubmitGenericFee(ctx context.Context, in *SrcGenericFee, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FTaxes_SubmitGenericFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fTaxesClient) ShowJobProgress(ctx context.Context, in *JobProgress, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FTaxes_ShowJobProgress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FTaxesServer is the server API for FTaxes service.
// All implementations must embed UnimplementedFTaxesServer
// for forward compatibility
type FTaxesServer interface {
	SubmitTransaction(context.Context, *SrcTx) (*emptypb.Empty, error)
	SubmitTransfer(context.Context, *SrcTransfer) (*emptypb.Empty, error)
	SubmitGenericFee(context.Context, *SrcGenericFee) (*emptypb.Empty, error)
	ShowJobProgress(context.Context, *JobProgress) (*emptypb.Empty, error)
	mustEmbedUnimplementedFTaxesServer()
}

// UnimplementedFTaxesServer must be embedded to have forward compatible implementations.
type UnimplementedFTaxesServer struct {
}

func (UnimplementedFTaxesServer) SubmitTransaction(context.Context, *SrcTx) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (UnimplementedFTaxesServer) SubmitTransfer(context.Context, *SrcTransfer) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransfer not implemented")
}
func (UnimplementedFTaxesServer) SubmitGenericFee(context.Context, *SrcGenericFee) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitGenericFee not implemented")
}
func (UnimplementedFTaxesServer) ShowJobProgress(context.Context, *JobProgress) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowJobProgress not implemented")
}
func (UnimplementedFTaxesServer) mustEmbedUnimplementedFTaxesServer() {}

// UnsafeFTaxesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FTaxesServer will
// result in compilation errors.
type UnsafeFTaxesServer interface {
	mustEmbedUnimplementedFTaxesServer()
}

func RegisterFTaxesServer(s grpc.ServiceRegistrar, srv FTaxesServer) {
	s.RegisterService(&FTaxes_ServiceDesc, srv)
}

func _FTaxes_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SrcTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTaxesServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FTaxes_SubmitTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTaxesServer).SubmitTransaction(ctx, req.(*SrcTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _FTaxes_SubmitTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SrcTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTaxesServer).SubmitTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FTaxes_SubmitTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTaxesServer).SubmitTransfer(ctx, req.(*SrcTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _FTaxes_SubmitGenericFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SrcGenericFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTaxesServer).SubmitGenericFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FTaxes_SubmitGenericFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTaxesServer).SubmitGenericFee(ctx, req.(*SrcGenericFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _FTaxes_ShowJobProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobProgress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FTaxesServer).ShowJobProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FTaxes_ShowJobProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FTaxesServer).ShowJobProgress(ctx, req.(*JobProgress))
	}
	return interceptor(ctx, in, info, handler)
}

// FTaxes_ServiceDesc is the grpc.ServiceDesc for FTaxes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FTaxes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FTaxesGrpc.FTaxes",
	HandlerType: (*FTaxesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _FTaxes_SubmitTransaction_Handler,
		},
		{
			MethodName: "SubmitTransfer",
			Handler:    _FTaxes_SubmitTransfer_Handler,
		},
		{
			MethodName: "SubmitGenericFee",
			Handler:    _FTaxes_SubmitGenericFee_Handler,
		},
		{
			MethodName: "ShowJobProgress",
			Handler:    _FTaxes_ShowJobProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "f-taxes.proto",
}

const (
	PluginCtl_UpdateTransactions_FullMethodName = "/FTaxesGrpc.PluginCtl/UpdateTransactions"
)

// PluginCtlClient is the client API for PluginCtl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginCtlClient interface {
	UpdateTransactions(ctx context.Context, in *TxUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pluginCtlClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginCtlClient(cc grpc.ClientConnInterface) PluginCtlClient {
	return &pluginCtlClient{cc}
}

func (c *pluginCtlClient) UpdateTransactions(ctx context.Context, in *TxUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PluginCtl_UpdateTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginCtlServer is the server API for PluginCtl service.
// All implementations must embed UnimplementedPluginCtlServer
// for forward compatibility
type PluginCtlServer interface {
	UpdateTransactions(context.Context, *TxUpdate) (*emptypb.Empty, error)
	mustEmbedUnimplementedPluginCtlServer()
}

// UnimplementedPluginCtlServer must be embedded to have forward compatible implementations.
type UnimplementedPluginCtlServer struct {
}

func (UnimplementedPluginCtlServer) UpdateTransactions(context.Context, *TxUpdate) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransactions not implemented")
}
func (UnimplementedPluginCtlServer) mustEmbedUnimplementedPluginCtlServer() {}

// UnsafePluginCtlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginCtlServer will
// result in compilation errors.
type UnsafePluginCtlServer interface {
	mustEmbedUnimplementedPluginCtlServer()
}

func RegisterPluginCtlServer(s grpc.ServiceRegistrar, srv PluginCtlServer) {
	s.RegisterService(&PluginCtl_ServiceDesc, srv)
}

func _PluginCtl_UpdateTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginCtlServer).UpdateTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginCtl_UpdateTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginCtlServer).UpdateTransactions(ctx, req.(*TxUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginCtl_ServiceDesc is the grpc.ServiceDesc for PluginCtl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginCtl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FTaxesGrpc.PluginCtl",
	HandlerType: (*PluginCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateTransactions",
			Handler:    _PluginCtl_UpdateTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "f-taxes.proto",
}