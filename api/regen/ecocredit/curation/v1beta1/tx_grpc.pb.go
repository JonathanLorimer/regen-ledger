// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: regen/ecocredit/curation/v1beta1/tx.proto

package curationv1beta1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	DefineTag(ctx context.Context, in *MsgDefineTag, opts ...grpc.CallOption) (*MsgDefineTagResponse, error)
	DefineNumericAtt(ctx context.Context, in *MsgDefineNumericAttr, opts ...grpc.CallOption) (*MsgDefineNumericAttrResponse, error)
	Tag(ctx context.Context, in *MsgTag, opts ...grpc.CallOption) (*MsgTagResponse, error)
	SetNumericAttr(ctx context.Context, in *MsgSetNumericAttr, opts ...grpc.CallOption) (*MsgSetNumericAttrResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DefineTag(ctx context.Context, in *MsgDefineTag, opts ...grpc.CallOption) (*MsgDefineTagResponse, error) {
	out := new(MsgDefineTagResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.curation.v1beta1.Msg/DefineTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DefineNumericAtt(ctx context.Context, in *MsgDefineNumericAttr, opts ...grpc.CallOption) (*MsgDefineNumericAttrResponse, error) {
	out := new(MsgDefineNumericAttrResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.curation.v1beta1.Msg/DefineNumericAtt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Tag(ctx context.Context, in *MsgTag, opts ...grpc.CallOption) (*MsgTagResponse, error) {
	out := new(MsgTagResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.curation.v1beta1.Msg/Tag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetNumericAttr(ctx context.Context, in *MsgSetNumericAttr, opts ...grpc.CallOption) (*MsgSetNumericAttrResponse, error) {
	out := new(MsgSetNumericAttrResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.curation.v1beta1.Msg/SetNumericAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	DefineTag(context.Context, *MsgDefineTag) (*MsgDefineTagResponse, error)
	DefineNumericAtt(context.Context, *MsgDefineNumericAttr) (*MsgDefineNumericAttrResponse, error)
	Tag(context.Context, *MsgTag) (*MsgTagResponse, error)
	SetNumericAttr(context.Context, *MsgSetNumericAttr) (*MsgSetNumericAttrResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) DefineTag(context.Context, *MsgDefineTag) (*MsgDefineTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefineTag not implemented")
}
func (UnimplementedMsgServer) DefineNumericAtt(context.Context, *MsgDefineNumericAttr) (*MsgDefineNumericAttrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefineNumericAtt not implemented")
}
func (UnimplementedMsgServer) Tag(context.Context, *MsgTag) (*MsgTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tag not implemented")
}
func (UnimplementedMsgServer) SetNumericAttr(context.Context, *MsgSetNumericAttr) (*MsgSetNumericAttrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNumericAttr not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_DefineTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDefineTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DefineTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.curation.v1beta1.Msg/DefineTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DefineTag(ctx, req.(*MsgDefineTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DefineNumericAtt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDefineNumericAttr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DefineNumericAtt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.curation.v1beta1.Msg/DefineNumericAtt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DefineNumericAtt(ctx, req.(*MsgDefineNumericAttr))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Tag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Tag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.curation.v1beta1.Msg/Tag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Tag(ctx, req.(*MsgTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetNumericAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetNumericAttr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetNumericAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.curation.v1beta1.Msg/SetNumericAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetNumericAttr(ctx, req.(*MsgSetNumericAttr))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.curation.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DefineTag",
			Handler:    _Msg_DefineTag_Handler,
		},
		{
			MethodName: "DefineNumericAtt",
			Handler:    _Msg_DefineNumericAtt_Handler,
		},
		{
			MethodName: "Tag",
			Handler:    _Msg_Tag_Handler,
		},
		{
			MethodName: "SetNumericAttr",
			Handler:    _Msg_SetNumericAttr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/curation/v1beta1/tx.proto",
}
