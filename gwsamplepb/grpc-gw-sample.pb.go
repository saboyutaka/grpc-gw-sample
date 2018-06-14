// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gwsamplepb/grpc-gw-sample.proto

package gwsamplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMessage) Reset()         { *m = RequestMessage{} }
func (m *RequestMessage) String() string { return proto.CompactTextString(m) }
func (*RequestMessage) ProtoMessage()    {}
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_gw_sample_31fe2b5c7d0c3f5b, []int{0}
}
func (m *RequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMessage.Unmarshal(m, b)
}
func (m *RequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMessage.Marshal(b, m, deterministic)
}
func (dst *RequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMessage.Merge(dst, src)
}
func (m *RequestMessage) XXX_Size() int {
	return xxx_messageInfo_RequestMessage.Size(m)
}
func (m *RequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMessage proto.InternalMessageInfo

func (m *RequestMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ResponseMessage struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMessage) Reset()         { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()    {}
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_gw_sample_31fe2b5c7d0c3f5b, []int{1}
}
func (m *ResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMessage.Unmarshal(m, b)
}
func (m *ResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMessage.Marshal(b, m, deterministic)
}
func (dst *ResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMessage.Merge(dst, src)
}
func (m *ResponseMessage) XXX_Size() int {
	return xxx_messageInfo_ResponseMessage.Size(m)
}
func (m *ResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMessage proto.InternalMessageInfo

func (m *ResponseMessage) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestMessage)(nil), "gwsamplepb.RequestMessage")
	proto.RegisterType((*ResponseMessage)(nil), "gwsamplepb.ResponseMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewaySampleClient is the client API for GatewaySample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewaySampleClient interface {
	Echo(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type gatewaySampleClient struct {
	cc *grpc.ClientConn
}

func NewGatewaySampleClient(cc *grpc.ClientConn) GatewaySampleClient {
	return &gatewaySampleClient{cc}
}

func (c *gatewaySampleClient) Echo(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/gwsamplepb.GatewaySample/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewaySampleServer is the server API for GatewaySample service.
type GatewaySampleServer interface {
	Echo(context.Context, *RequestMessage) (*ResponseMessage, error)
}

func RegisterGatewaySampleServer(s *grpc.Server, srv GatewaySampleServer) {
	s.RegisterService(&_GatewaySample_serviceDesc, srv)
}

func _GatewaySample_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySampleServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwsamplepb.GatewaySample/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySampleServer).Echo(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewaySample_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gwsamplepb.GatewaySample",
	HandlerType: (*GatewaySampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _GatewaySample_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gwsamplepb/grpc-gw-sample.proto",
}

func init() {
	proto.RegisterFile("gwsamplepb/grpc-gw-sample.proto", fileDescriptor_grpc_gw_sample_31fe2b5c7d0c3f5b)
}

var fileDescriptor_grpc_gw_sample_31fe2b5c7d0c3f5b = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2f, 0x2f, 0x4e,
	0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0x48, 0xd2, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4d, 0x2f, 0xd7, 0x85,
	0x08, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x21, 0x14, 0x48, 0xc9, 0xa4, 0xe7, 0xe7,
	0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64,
	0xe6, 0xe7, 0x15, 0x43, 0x54, 0x2a, 0xa9, 0x71, 0xf1, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0xf8, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x9a, 0x5c, 0xfc, 0x41, 0xa9, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x30, 0x85, 0x62, 0x5c, 0x6c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25,
	0x50, 0x95, 0x50, 0x9e, 0x51, 0x2e, 0x17, 0xaf, 0x7b, 0x62, 0x49, 0x6a, 0x79, 0x62, 0x65, 0x30,
	0xd8, 0x0d, 0x42, 0x31, 0x5c, 0x2c, 0xae, 0xc9, 0x19, 0xf9, 0x42, 0x52, 0x7a, 0x08, 0x67, 0xe9,
	0xa1, 0xda, 0x2a, 0x25, 0x8d, 0x2a, 0x87, 0x62, 0x93, 0x92, 0x74, 0xd3, 0xe5, 0x27, 0x93, 0x99,
	0x44, 0x95, 0x04, 0xf4, 0xcb, 0x0c, 0xf5, 0x53, 0x2b, 0xc0, 0xca, 0xf4, 0x53, 0x93, 0x33, 0xf2,
	0xad, 0x18, 0xb5, 0x92, 0xd8, 0xc0, 0x1e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x43,
	0x82, 0xa7, 0x15, 0x01, 0x00, 0x00,
}