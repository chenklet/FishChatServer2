// Code generated by protoc-gen-go.
// source: logic.proto
// DO NOT EDIT!

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LoginReq struct {
	UID        int64  `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	AccessAddr string `protobuf:"bytes,3,opt,name=accessAddr" json:"accessAddr,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LoginReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReq) GetAccessAddr() string {
	if m != nil {
		return m.AccessAddr
	}
	return ""
}

type LoginRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *LoginRes) Reset()                    { *m = LoginRes{} }
func (m *LoginRes) String() string            { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()               {}
func (*LoginRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *LoginRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *LoginRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type PingReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *PingReq) Reset()                    { *m = PingReq{} }
func (m *PingReq) String() string            { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()               {}
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *PingReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type PingRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *PingRes) Reset()                    { *m = PingRes{} }
func (m *PingRes) String() string            { return proto.CompactTextString(m) }
func (*PingRes) ProtoMessage()               {}
func (*PingRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *PingRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *PingRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type SendP2PMsgReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *SendP2PMsgReq) Reset()                    { *m = SendP2PMsgReq{} }
func (m *SendP2PMsgReq) String() string            { return proto.CompactTextString(m) }
func (*SendP2PMsgReq) ProtoMessage()               {}
func (*SendP2PMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SendP2PMsgReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *SendP2PMsgReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *SendP2PMsgReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *SendP2PMsgReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SendP2PMsgRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *SendP2PMsgRes) Reset()                    { *m = SendP2PMsgRes{} }
func (m *SendP2PMsgRes) String() string            { return proto.CompactTextString(m) }
func (*SendP2PMsgRes) ProtoMessage()               {}
func (*SendP2PMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SendP2PMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SendP2PMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

// p2p msg accept ack
type AcceptP2PMsgAckReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,5,opt,name=msgID" json:"msgID,omitempty"`
}

func (m *AcceptP2PMsgAckReq) Reset()                    { *m = AcceptP2PMsgAckReq{} }
func (m *AcceptP2PMsgAckReq) String() string            { return proto.CompactTextString(m) }
func (*AcceptP2PMsgAckReq) ProtoMessage()               {}
func (*AcceptP2PMsgAckReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *AcceptP2PMsgAckReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *AcceptP2PMsgAckReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *AcceptP2PMsgAckReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

type AcceptP2PMsgAckRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *AcceptP2PMsgAckRes) Reset()                    { *m = AcceptP2PMsgAckRes{} }
func (m *AcceptP2PMsgAckRes) String() string            { return proto.CompactTextString(m) }
func (*AcceptP2PMsgAckRes) ProtoMessage()               {}
func (*AcceptP2PMsgAckRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *AcceptP2PMsgAckRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *AcceptP2PMsgAckRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "rpc.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "rpc.LoginRes")
	proto.RegisterType((*PingReq)(nil), "rpc.PingReq")
	proto.RegisterType((*PingRes)(nil), "rpc.PingRes")
	proto.RegisterType((*SendP2PMsgReq)(nil), "rpc.SendP2PMsgReq")
	proto.RegisterType((*SendP2PMsgRes)(nil), "rpc.SendP2PMsgRes")
	proto.RegisterType((*AcceptP2PMsgAckReq)(nil), "rpc.AcceptP2PMsgAckReq")
	proto.RegisterType((*AcceptP2PMsgAckRes)(nil), "rpc.AcceptP2PMsgAckRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LogicRPC service

type LogicRPCClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
	SendP2PMsg(ctx context.Context, in *SendP2PMsgReq, opts ...grpc.CallOption) (*SendP2PMsgRes, error)
}

type logicRPCClient struct {
	cc *grpc.ClientConn
}

func NewLogicRPCClient(cc *grpc.ClientConn) LogicRPCClient {
	return &logicRPCClient{cc}
}

func (c *logicRPCClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicRPCClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicRPCClient) SendP2PMsg(ctx context.Context, in *SendP2PMsgReq, opts ...grpc.CallOption) (*SendP2PMsgRes, error) {
	out := new(SendP2PMsgRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/SendP2PMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogicRPC service

type LogicRPCServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Ping(context.Context, *PingReq) (*PingRes, error)
	SendP2PMsg(context.Context, *SendP2PMsgReq) (*SendP2PMsgRes, error)
}

func RegisterLogicRPCServer(s *grpc.Server, srv LogicRPCServer) {
	s.RegisterService(&_LogicRPC_serviceDesc, srv)
}

func _LogicRPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicRPC_SendP2PMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendP2PMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).SendP2PMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/SendP2PMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).SendP2PMsg(ctx, req.(*SendP2PMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LogicRPC",
	HandlerType: (*LogicRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LogicRPC_Login_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _LogicRPC_Ping_Handler,
		},
		{
			MethodName: "SendP2PMsg",
			Handler:    _LogicRPC_SendP2PMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic.proto",
}

func init() { proto.RegisterFile("logic.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x5d, 0xd7, 0x7d, 0xb8, 0xab, 0x03, 0xb9, 0x88, 0x94, 0x29, 0x32, 0x82, 0xe0, 0x9e, 0xf6,
	0x30, 0xc1, 0x17, 0x7d, 0x29, 0x1b, 0xc2, 0x40, 0xa1, 0x64, 0xf8, 0x03, 0x66, 0x1a, 0x43, 0x99,
	0x6b, 0x6a, 0x12, 0x7f, 0x88, 0xff, 0x58, 0x92, 0xb6, 0x6b, 0x8b, 0x7b, 0x2a, 0xbe, 0xe5, 0x9c,
	0x53, 0xce, 0x3d, 0xbd, 0xf7, 0xc0, 0xe9, 0xa7, 0x14, 0x09, 0x9b, 0x67, 0x4a, 0x1a, 0x89, 0xbe,
	0xca, 0x18, 0xa1, 0x70, 0xf2, 0x22, 0x45, 0x92, 0x52, 0xfe, 0x85, 0xe7, 0xe0, 0xbf, 0xad, 0x57,
	0x81, 0x37, 0xf5, 0x66, 0x3e, 0xb5, 0x4f, 0xbc, 0x80, 0xbe, 0x91, 0x3b, 0x9e, 0x06, 0xdd, 0xa9,
	0x37, 0x1b, 0xd1, 0x1c, 0xe0, 0x0d, 0xc0, 0x96, 0x31, 0xae, 0x75, 0x18, 0xc7, 0x2a, 0xf0, 0x9d,
	0x54, 0x63, 0xc8, 0xd3, 0xc1, 0x53, 0x63, 0x00, 0x43, 0xae, 0xd4, 0x52, 0xc6, 0xdc, 0xf9, 0x8e,
	0x69, 0x09, 0xf1, 0x12, 0x06, 0x5c, 0xa9, 0x8d, 0x51, 0x85, 0x79, 0x81, 0xc8, 0x15, 0x0c, 0xa3,
	0x24, 0x15, 0x47, 0x03, 0x91, 0xc7, 0x52, 0x6c, 0xe3, 0xac, 0x61, 0xbc, 0xe1, 0x69, 0x1c, 0x2d,
	0xa2, 0x57, 0xed, 0xfc, 0xaf, 0x61, 0xa4, 0xe5, 0xb7, 0x62, 0xbc, 0x9a, 0x52, 0x11, 0x56, 0x35,
	0x5b, 0x25, 0xb8, 0xb1, 0x6a, 0x37, 0x57, 0x0f, 0x84, 0x5d, 0xcd, 0x5e, 0x8b, 0xf5, 0xaa, 0xf8,
	0xff, 0x1c, 0xd8, 0xc4, 0x7b, 0x2d, 0x82, 0x9e, 0xe3, 0xec, 0x93, 0x84, 0xcd, 0xa1, 0x6d, 0x72,
	0x7f, 0x00, 0x86, 0x8c, 0xf1, 0xcc, 0xe4, 0x26, 0x21, 0xdb, 0xfd, 0x5b, 0xf8, 0x7e, 0x2d, 0x3c,
	0x79, 0x3e, 0x32, 0xa7, 0x45, 0xde, 0xc5, 0x8f, 0x97, 0x17, 0x80, 0xd1, 0x68, 0x89, 0x77, 0xd0,
	0x77, 0x65, 0xc0, 0xf1, 0x5c, 0x65, 0x6c, 0x5e, 0x96, 0x6d, 0xd2, 0x80, 0x9a, 0x74, 0xf0, 0x16,
	0x7a, 0xf6, 0xb4, 0x78, 0xe6, 0x84, 0xa2, 0x02, 0x93, 0x3a, 0xb2, 0x5f, 0x3d, 0x00, 0x54, 0xeb,
	0x44, 0x74, 0x6a, 0xe3, 0xa8, 0x93, 0xbf, 0x9c, 0x26, 0x9d, 0xf7, 0x81, 0xeb, 0xfc, 0xfd, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x65, 0xd9, 0x83, 0x02, 0x03, 0x00, 0x00,
}
