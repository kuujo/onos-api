// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onos/e2t/e2/v1beta1/control.proto

// Package onos.e2t.e2.v1beta1 defines the interior gRPC interfaces for xApps to interact with E2T.

package v1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ControlRequest struct {
	Headers RequestHeaders `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers"`
	Message ControlMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
}

func (m *ControlRequest) Reset()         { *m = ControlRequest{} }
func (m *ControlRequest) String() string { return proto.CompactTextString(m) }
func (*ControlRequest) ProtoMessage()    {}
func (*ControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a705973a48f294ef, []int{0}
}
func (m *ControlRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlRequest.Merge(m, src)
}
func (m *ControlRequest) XXX_Size() int {
	return m.Size()
}
func (m *ControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControlRequest proto.InternalMessageInfo

func (m *ControlRequest) GetHeaders() RequestHeaders {
	if m != nil {
		return m.Headers
	}
	return RequestHeaders{}
}

func (m *ControlRequest) GetMessage() ControlMessage {
	if m != nil {
		return m.Message
	}
	return ControlMessage{}
}

type ControlMessage struct {
	SubscriptionID SubscriptionID `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3,casttype=SubscriptionID" json:"subscription_id,omitempty"`
	E2NodeID       E2NodeID       `protobuf:"bytes,2,opt,name=e2_node_id,json=e2NodeId,proto3,casttype=E2NodeID" json:"e2_node_id,omitempty"`
	ServiceModel   ServiceModel   `protobuf:"bytes,3,opt,name=service_model,json=serviceModel,proto3" json:"service_model"`
	Header         Payload        `protobuf:"bytes,4,opt,name=header,proto3" json:"header"`
	Body           Payload        `protobuf:"bytes,5,opt,name=body,proto3" json:"body"`
}

func (m *ControlMessage) Reset()         { *m = ControlMessage{} }
func (m *ControlMessage) String() string { return proto.CompactTextString(m) }
func (*ControlMessage) ProtoMessage()    {}
func (*ControlMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a705973a48f294ef, []int{1}
}
func (m *ControlMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlMessage.Merge(m, src)
}
func (m *ControlMessage) XXX_Size() int {
	return m.Size()
}
func (m *ControlMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ControlMessage proto.InternalMessageInfo

func (m *ControlMessage) GetSubscriptionID() SubscriptionID {
	if m != nil {
		return m.SubscriptionID
	}
	return ""
}

func (m *ControlMessage) GetE2NodeID() E2NodeID {
	if m != nil {
		return m.E2NodeID
	}
	return ""
}

func (m *ControlMessage) GetServiceModel() ServiceModel {
	if m != nil {
		return m.ServiceModel
	}
	return ServiceModel{}
}

func (m *ControlMessage) GetHeader() Payload {
	if m != nil {
		return m.Header
	}
	return Payload{}
}

func (m *ControlMessage) GetBody() Payload {
	if m != nil {
		return m.Body
	}
	return Payload{}
}

type ControlResponse struct {
	Headers ResponseHeaders `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers"`
	Outcome Payload         `protobuf:"bytes,2,opt,name=outcome,proto3" json:"outcome"`
}

func (m *ControlResponse) Reset()         { *m = ControlResponse{} }
func (m *ControlResponse) String() string { return proto.CompactTextString(m) }
func (*ControlResponse) ProtoMessage()    {}
func (*ControlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a705973a48f294ef, []int{2}
}
func (m *ControlResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlResponse.Merge(m, src)
}
func (m *ControlResponse) XXX_Size() int {
	return m.Size()
}
func (m *ControlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlResponse proto.InternalMessageInfo

func (m *ControlResponse) GetHeaders() ResponseHeaders {
	if m != nil {
		return m.Headers
	}
	return ResponseHeaders{}
}

func (m *ControlResponse) GetOutcome() Payload {
	if m != nil {
		return m.Outcome
	}
	return Payload{}
}

type ControlOutcome struct {
	Body Payload `protobuf:"bytes,1,opt,name=body,proto3" json:"body"`
}

func (m *ControlOutcome) Reset()         { *m = ControlOutcome{} }
func (m *ControlOutcome) String() string { return proto.CompactTextString(m) }
func (*ControlOutcome) ProtoMessage()    {}
func (*ControlOutcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_a705973a48f294ef, []int{3}
}
func (m *ControlOutcome) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlOutcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlOutcome.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlOutcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlOutcome.Merge(m, src)
}
func (m *ControlOutcome) XXX_Size() int {
	return m.Size()
}
func (m *ControlOutcome) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlOutcome.DiscardUnknown(m)
}

var xxx_messageInfo_ControlOutcome proto.InternalMessageInfo

func (m *ControlOutcome) GetBody() Payload {
	if m != nil {
		return m.Body
	}
	return Payload{}
}

func init() {
	proto.RegisterType((*ControlRequest)(nil), "onos.e2t.e2.v1beta1.ControlRequest")
	proto.RegisterType((*ControlMessage)(nil), "onos.e2t.e2.v1beta1.ControlMessage")
	proto.RegisterType((*ControlResponse)(nil), "onos.e2t.e2.v1beta1.ControlResponse")
	proto.RegisterType((*ControlOutcome)(nil), "onos.e2t.e2.v1beta1.ControlOutcome")
}

func init() { proto.RegisterFile("onos/e2t/e2/v1beta1/control.proto", fileDescriptor_a705973a48f294ef) }

var fileDescriptor_a705973a48f294ef = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x52, 0xd6, 0x61, 0xa0, 0x93, 0x0c, 0x87, 0xa8, 0x9a, 0x52, 0x16, 0x7a, 0xe0,
	0x94, 0x68, 0x46, 0xda, 0x01, 0x71, 0xea, 0x86, 0xb4, 0x49, 0x14, 0x50, 0xc6, 0xbd, 0x4a, 0xe2,
	0x47, 0x89, 0xd4, 0xe6, 0x95, 0xd8, 0x9d, 0xb4, 0x6f, 0xc1, 0x85, 0x0b, 0x9f, 0x68, 0xc7, 0x1d,
	0xe1, 0x52, 0xa1, 0xf4, 0x5b, 0x70, 0x42, 0xb1, 0x9d, 0xa8, 0x82, 0xac, 0x2a, 0x37, 0xeb, 0xf9,
	0xf7, 0x7e, 0xf1, 0xf3, 0xdf, 0xa1, 0x47, 0x98, 0xa1, 0x0c, 0x80, 0xab, 0x00, 0x78, 0x70, 0x75,
	0x1c, 0x83, 0x8a, 0x8e, 0x83, 0x04, 0x33, 0x95, 0xe3, 0xcc, 0x5f, 0xe4, 0xa8, 0x90, 0x3d, 0x29,
	0x11, 0x1f, 0xb8, 0xf2, 0x81, 0xfb, 0x16, 0xe9, 0x3f, 0x9d, 0xe2, 0x14, 0xf5, 0x7e, 0x50, 0xae,
	0x0c, 0xda, 0x3f, 0x6c, 0xb2, 0x01, 0x37, 0xbb, 0xde, 0x77, 0x42, 0x7b, 0xa7, 0x46, 0x1d, 0xc2,
	0x97, 0x25, 0x48, 0xc5, 0x4e, 0x69, 0xf7, 0x33, 0x44, 0x02, 0x72, 0xe9, 0x90, 0x67, 0xe4, 0xc5,
	0x43, 0xfe, 0xdc, 0x6f, 0xf8, 0x9a, 0x6f, 0xf1, 0x73, 0x83, 0x8e, 0x3a, 0x37, 0xab, 0x41, 0x2b,
	0xac, 0x3a, 0x4b, 0xc9, 0x1c, 0xa4, 0x8c, 0xa6, 0xe0, 0xb4, 0xb7, 0x48, 0xec, 0xa7, 0xc7, 0x06,
	0xad, 0x24, 0xb6, 0xd3, 0xfb, 0xd9, 0xae, 0x0f, 0x67, 0x09, 0x36, 0xa6, 0x07, 0x72, 0x19, 0xcb,
	0x24, 0x4f, 0x17, 0x2a, 0xc5, 0x6c, 0x92, 0x0a, 0x7d, 0xc8, 0x07, 0xa3, 0x61, 0xb1, 0x1a, 0xf4,
	0x2e, 0x37, 0xb6, 0x2e, 0xce, 0x7e, 0xff, 0x53, 0x09, 0x7b, 0x9b, 0xcd, 0x17, 0x82, 0x9d, 0x50,
	0x0a, 0x7c, 0x92, 0xa1, 0x80, 0xd2, 0xd4, 0xd6, 0x26, 0xa7, 0x58, 0x0d, 0xf6, 0xdf, 0xf0, 0x77,
	0x28, 0x40, 0x3b, 0xea, 0x75, 0xb8, 0x0f, 0x66, 0x25, 0xd8, 0x5b, 0xfa, 0x58, 0x42, 0x7e, 0x95,
	0x26, 0x30, 0x99, 0xa3, 0x80, 0x99, 0x73, 0x4f, 0x0f, 0x79, 0xd4, 0x38, 0xe4, 0xa5, 0x21, 0xc7,
	0x25, 0x68, 0x47, 0x7c, 0x24, 0x37, 0x6a, 0xec, 0x15, 0xdd, 0x33, 0xf7, 0xe6, 0x74, 0xb4, 0xe6,
	0xb0, 0x51, 0xf3, 0x21, 0xba, 0x9e, 0x61, 0x24, 0xac, 0xc1, 0x76, 0xb0, 0x13, 0xda, 0x89, 0x51,
	0x5c, 0x3b, 0xf7, 0x77, 0xee, 0xd4, 0xbc, 0xf7, 0x8d, 0xd0, 0x83, 0x3a, 0x78, 0xb9, 0xc0, 0x4c,
	0x02, 0x3b, 0xfb, 0x3b, 0xf9, 0xe1, 0x1d, 0xc9, 0x1b, 0xfe, 0x8e, 0xe8, 0x5f, 0xd3, 0x2e, 0x2e,
	0x55, 0x82, 0xf3, 0x2a, 0xfa, 0x5d, 0x0e, 0x55, 0xb5, 0x78, 0xe7, 0x75, 0xe4, 0xef, 0x4d, 0xa5,
	0x9e, 0x90, 0xfc, 0xdf, 0x84, 0xfc, 0x53, 0x6d, 0xb2, 0x01, 0xb0, 0x8f, 0xb4, 0x6b, 0x2b, 0x6c,
	0xeb, 0x73, 0xb4, 0x4f, 0xbb, 0x3f, 0xdc, 0x0e, 0x99, 0x5b, 0x18, 0x39, 0x37, 0x85, 0x4b, 0x6e,
	0x0b, 0x97, 0xfc, 0x2a, 0x5c, 0xf2, 0x75, 0xed, 0xb6, 0x6e, 0xd7, 0x6e, 0xeb, 0xc7, 0xda, 0x6d,
	0xc5, 0x7b, 0xfa, 0x1f, 0x7b, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x64, 0x0c, 0x2a, 0xad, 0xd1,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlServiceClient interface {
	Control(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlResponse, error)
}

type controlServiceClient struct {
	cc *grpc.ClientConn
}

func NewControlServiceClient(cc *grpc.ClientConn) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) Control(ctx context.Context, in *ControlRequest, opts ...grpc.CallOption) (*ControlResponse, error) {
	out := new(ControlResponse)
	err := c.cc.Invoke(ctx, "/onos.e2t.e2.v1beta1.ControlService/Control", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
type ControlServiceServer interface {
	Control(context.Context, *ControlRequest) (*ControlResponse, error)
}

// UnimplementedControlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedControlServiceServer struct {
}

func (*UnimplementedControlServiceServer) Control(ctx context.Context, req *ControlRequest) (*ControlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Control not implemented")
}

func RegisterControlServiceServer(s *grpc.Server, srv ControlServiceServer) {
	s.RegisterService(&_ControlService_serviceDesc, srv)
}

func _ControlService_Control_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Control(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/onos.e2t.e2.v1beta1.ControlService/Control",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Control(ctx, req.(*ControlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onos.e2t.e2.v1beta1.ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Control",
			Handler:    _ControlService_Control_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onos/e2t/e2/v1beta1/control.proto",
}

func (m *ControlRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Headers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ControlMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.ServiceModel.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.E2NodeID) > 0 {
		i -= len(m.E2NodeID)
		copy(dAtA[i:], m.E2NodeID)
		i = encodeVarintControl(dAtA, i, uint64(len(m.E2NodeID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SubscriptionID) > 0 {
		i -= len(m.SubscriptionID)
		copy(dAtA[i:], m.SubscriptionID)
		i = encodeVarintControl(dAtA, i, uint64(len(m.SubscriptionID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ControlResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Outcome.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Headers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ControlOutcome) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlOutcome) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlOutcome) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintControl(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintControl(dAtA []byte, offset int, v uint64) int {
	offset -= sovControl(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ControlRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Headers.Size()
	n += 1 + l + sovControl(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovControl(uint64(l))
	return n
}

func (m *ControlMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SubscriptionID)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = len(m.E2NodeID)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	l = m.ServiceModel.Size()
	n += 1 + l + sovControl(uint64(l))
	l = m.Header.Size()
	n += 1 + l + sovControl(uint64(l))
	l = m.Body.Size()
	n += 1 + l + sovControl(uint64(l))
	return n
}

func (m *ControlResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Headers.Size()
	n += 1 + l + sovControl(uint64(l))
	l = m.Outcome.Size()
	n += 1 + l + sovControl(uint64(l))
	return n
}

func (m *ControlOutcome) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Body.Size()
	n += 1 + l + sovControl(uint64(l))
	return n
}

func sovControl(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozControl(x uint64) (n int) {
	return sovControl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ControlRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Headers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubscriptionID = SubscriptionID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field E2NodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.E2NodeID = E2NodeID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceModel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServiceModel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Headers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outcome", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Outcome.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlOutcome) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlOutcome: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlOutcome: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthControl
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipControl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowControl
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowControl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowControl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthControl
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupControl
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthControl
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthControl        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowControl          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupControl = fmt.Errorf("proto: unexpected end of group")
)
