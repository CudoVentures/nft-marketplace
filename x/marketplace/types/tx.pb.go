// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: marketplace/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
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

type MsgListNFT struct {
	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price  types.Coin `protobuf:"bytes,2,opt,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"price"`
	Sender string     `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgListNFT) Reset()         { *m = MsgListNFT{} }
func (m *MsgListNFT) String() string { return proto.CompactTextString(m) }
func (*MsgListNFT) ProtoMessage()    {}
func (*MsgListNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73de1bb07ce97a8, []int{0}
}
func (m *MsgListNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgListNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgListNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgListNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgListNFT.Merge(m, src)
}
func (m *MsgListNFT) XXX_Size() int {
	return m.Size()
}
func (m *MsgListNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgListNFT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgListNFT proto.InternalMessageInfo

type MsgListNFTResponse struct {
}

func (m *MsgListNFTResponse) Reset()         { *m = MsgListNFTResponse{} }
func (m *MsgListNFTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgListNFTResponse) ProtoMessage()    {}
func (*MsgListNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73de1bb07ce97a8, []int{1}
}
func (m *MsgListNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgListNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgListNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgListNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgListNFTResponse.Merge(m, src)
}
func (m *MsgListNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgListNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgListNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgListNFTResponse proto.InternalMessageInfo

type MsgDeListNFT struct {
	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgDeListNFT) Reset()         { *m = MsgDeListNFT{} }
func (m *MsgDeListNFT) String() string { return proto.CompactTextString(m) }
func (*MsgDeListNFT) ProtoMessage()    {}
func (*MsgDeListNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73de1bb07ce97a8, []int{2}
}
func (m *MsgDeListNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeListNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeListNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeListNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeListNFT.Merge(m, src)
}
func (m *MsgDeListNFT) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeListNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeListNFT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeListNFT proto.InternalMessageInfo

type MsgDeListNFTResponse struct {
}

func (m *MsgDeListNFTResponse) Reset()         { *m = MsgDeListNFTResponse{} }
func (m *MsgDeListNFTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeListNFTResponse) ProtoMessage()    {}
func (*MsgDeListNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73de1bb07ce97a8, []int{3}
}
func (m *MsgDeListNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeListNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeListNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeListNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeListNFTResponse.Merge(m, src)
}
func (m *MsgDeListNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeListNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeListNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeListNFTResponse proto.InternalMessageInfo

type MsgBuyNFT struct {
	Id     string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price  types.Coin `protobuf:"bytes,2,opt,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"price"`
	Sender string     `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgBuyNFT) Reset()         { *m = MsgBuyNFT{} }
func (m *MsgBuyNFT) String() string { return proto.CompactTextString(m) }
func (*MsgBuyNFT) ProtoMessage()    {}
func (*MsgBuyNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73de1bb07ce97a8, []int{4}
}
func (m *MsgBuyNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBuyNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBuyNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBuyNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBuyNFT.Merge(m, src)
}
func (m *MsgBuyNFT) XXX_Size() int {
	return m.Size()
}
func (m *MsgBuyNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBuyNFT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBuyNFT proto.InternalMessageInfo

type MsgBuyNFTResponse struct {
}

func (m *MsgBuyNFTResponse) Reset()         { *m = MsgBuyNFTResponse{} }
func (m *MsgBuyNFTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBuyNFTResponse) ProtoMessage()    {}
func (*MsgBuyNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73de1bb07ce97a8, []int{5}
}
func (m *MsgBuyNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBuyNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBuyNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBuyNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBuyNFTResponse.Merge(m, src)
}
func (m *MsgBuyNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBuyNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBuyNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBuyNFTResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgListNFT)(nil), "OmniFlix.marketplace.v1beta1.MsgListNFT")
	proto.RegisterType((*MsgListNFTResponse)(nil), "OmniFlix.marketplace.v1beta1.MsgListNFTResponse")
	proto.RegisterType((*MsgDeListNFT)(nil), "OmniFlix.marketplace.v1beta1.MsgDeListNFT")
	proto.RegisterType((*MsgDeListNFTResponse)(nil), "OmniFlix.marketplace.v1beta1.MsgDeListNFTResponse")
	proto.RegisterType((*MsgBuyNFT)(nil), "OmniFlix.marketplace.v1beta1.MsgBuyNFT")
	proto.RegisterType((*MsgBuyNFTResponse)(nil), "OmniFlix.marketplace.v1beta1.MsgBuyNFTResponse")
}

func init() { proto.RegisterFile("marketplace/v1beta1/tx.proto", fileDescriptor_f73de1bb07ce97a8) }

var fileDescriptor_f73de1bb07ce97a8 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xcb, 0x4e, 0xea, 0x40,
	0x18, 0xee, 0x94, 0x1c, 0x4e, 0x98, 0x73, 0x72, 0x92, 0x53, 0x09, 0xc1, 0x86, 0x0c, 0x84, 0x0d,
	0xc4, 0xc4, 0x19, 0x81, 0xc4, 0x07, 0x40, 0xc3, 0xca, 0x6a, 0x24, 0xae, 0x5c, 0xd9, 0xcb, 0xa4,
	0x4e, 0xa0, 0x9d, 0xa6, 0x53, 0x0c, 0xbc, 0x85, 0x1b, 0x7d, 0x08, 0xdf, 0xc1, 0x3d, 0x4b, 0x96,
	0xae, 0xbc, 0xc0, 0x8b, 0x18, 0x7a, 0x67, 0xa1, 0xb0, 0x74, 0xd5, 0x99, 0xe6, 0xbb, 0xfd, 0xff,
	0x3f, 0x3f, 0xac, 0x39, 0xba, 0x3f, 0xa2, 0x81, 0x37, 0xd6, 0x4d, 0x4a, 0xee, 0x3a, 0x06, 0x0d,
	0xf4, 0x0e, 0x09, 0xa6, 0xd8, 0xf3, 0x79, 0xc0, 0x95, 0xda, 0x85, 0xe3, 0xb2, 0xc1, 0x98, 0x4d,
	0x71, 0x0e, 0x86, 0x63, 0x98, 0x8a, 0x4c, 0x2e, 0x1c, 0x2e, 0x88, 0xa1, 0x8b, 0x8c, 0x6b, 0x72,
	0xe6, 0x46, 0x6c, 0xb5, 0x6c, 0x73, 0x9b, 0x87, 0x47, 0xb2, 0x3e, 0x45, 0x7f, 0x9b, 0x8f, 0x00,
	0x42, 0x4d, 0xd8, 0x67, 0x4c, 0x04, 0xe7, 0x83, 0x2b, 0xe5, 0x1f, 0x94, 0x99, 0x55, 0x05, 0x0d,
	0xd0, 0x2e, 0x0d, 0x65, 0x66, 0x29, 0x37, 0xf0, 0x97, 0xe7, 0x33, 0x93, 0x56, 0xe5, 0x06, 0x68,
	0xff, 0xe9, 0xee, 0xe3, 0xc8, 0x04, 0xaf, 0x4d, 0x12, 0x67, 0x7c, 0xc2, 0x99, 0xdb, 0x27, 0xf3,
	0xd7, 0xba, 0xf4, 0xf4, 0x56, 0x6f, 0xd9, 0x2c, 0xb8, 0x9d, 0x18, 0xd8, 0xe4, 0x0e, 0x89, 0x13,
	0x45, 0x9f, 0x43, 0x61, 0x8d, 0x48, 0x30, 0xf3, 0xa8, 0x08, 0x09, 0xc3, 0x48, 0x58, 0xa9, 0xc0,
	0xa2, 0xa0, 0xae, 0x45, 0xfd, 0x6a, 0x21, 0x74, 0x8d, 0x6f, 0xcd, 0x32, 0x54, 0xb2, 0x5c, 0x43,
	0x2a, 0x3c, 0xee, 0x0a, 0xda, 0x3c, 0x86, 0x7f, 0x35, 0x61, 0x9f, 0xd2, 0xaf, 0xf2, 0x66, 0x6a,
	0xf2, 0x86, 0x5a, 0x05, 0x96, 0xf3, 0xbc, 0x54, 0xef, 0x01, 0xc0, 0x92, 0x26, 0xec, 0xfe, 0x64,
	0xf6, 0xb3, 0xaa, 0xdf, 0x83, 0xff, 0xd3, 0x58, 0x49, 0xd8, 0xee, 0xb3, 0x0c, 0x0b, 0x9a, 0xb0,
	0x15, 0x0a, 0x7f, 0x27, 0xf5, 0xb7, 0xf1, 0x77, 0x6f, 0x02, 0x67, 0x1d, 0x54, 0x8f, 0x76, 0x45,
	0x26, 0x76, 0xca, 0x08, 0x96, 0xb2, 0x46, 0x1f, 0x6c, 0xa5, 0xa7, 0x58, 0xb5, 0xbb, 0x3b, 0x36,
	0x35, 0x33, 0x60, 0x31, 0x1e, 0x42, 0x6b, 0x2b, 0x3b, 0x02, 0xaa, 0x64, 0x47, 0x60, 0xe2, 0xd1,
	0xbf, 0x9c, 0x7f, 0x20, 0x69, 0xbe, 0x44, 0x60, 0xb1, 0x44, 0xe0, 0x7d, 0x89, 0xc0, 0xfd, 0x0a,
	0x49, 0x8b, 0x15, 0x92, 0x5e, 0x56, 0x48, 0xba, 0xee, 0xe5, 0x46, 0x97, 0x08, 0x93, 0xfc, 0x3e,
	0x6e, 0xde, 0xc2, 0x59, 0x1a, 0xc5, 0x70, 0x8b, 0x7a, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x49,
	0x8a, 0x2b, 0xf2, 0xb9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	ListNFT(ctx context.Context, in *MsgListNFT, opts ...grpc.CallOption) (*MsgListNFTResponse, error)
	DeListNFT(ctx context.Context, in *MsgDeListNFT, opts ...grpc.CallOption) (*MsgDeListNFTResponse, error)
	BuyNFT(ctx context.Context, in *MsgBuyNFT, opts ...grpc.CallOption) (*MsgBuyNFTResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ListNFT(ctx context.Context, in *MsgListNFT, opts ...grpc.CallOption) (*MsgListNFTResponse, error) {
	out := new(MsgListNFTResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.marketplace.v1beta1.Msg/ListNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeListNFT(ctx context.Context, in *MsgDeListNFT, opts ...grpc.CallOption) (*MsgDeListNFTResponse, error) {
	out := new(MsgDeListNFTResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.marketplace.v1beta1.Msg/DeListNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BuyNFT(ctx context.Context, in *MsgBuyNFT, opts ...grpc.CallOption) (*MsgBuyNFTResponse, error) {
	out := new(MsgBuyNFTResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.marketplace.v1beta1.Msg/BuyNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ListNFT(context.Context, *MsgListNFT) (*MsgListNFTResponse, error)
	DeListNFT(context.Context, *MsgDeListNFT) (*MsgDeListNFTResponse, error)
	BuyNFT(context.Context, *MsgBuyNFT) (*MsgBuyNFTResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ListNFT(ctx context.Context, req *MsgListNFT) (*MsgListNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNFT not implemented")
}
func (*UnimplementedMsgServer) DeListNFT(ctx context.Context, req *MsgDeListNFT) (*MsgDeListNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeListNFT not implemented")
}
func (*UnimplementedMsgServer) BuyNFT(ctx context.Context, req *MsgBuyNFT) (*MsgBuyNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyNFT not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ListNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgListNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ListNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.marketplace.v1beta1.Msg/ListNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ListNFT(ctx, req.(*MsgListNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeListNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeListNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeListNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.marketplace.v1beta1.Msg/DeListNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeListNFT(ctx, req.(*MsgDeListNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BuyNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBuyNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BuyNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.marketplace.v1beta1.Msg/BuyNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BuyNFT(ctx, req.(*MsgBuyNFT))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OmniFlix.marketplace.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNFT",
			Handler:    _Msg_ListNFT_Handler,
		},
		{
			MethodName: "DeListNFT",
			Handler:    _Msg_DeListNFT_Handler,
		},
		{
			MethodName: "BuyNFT",
			Handler:    _Msg_BuyNFT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace/v1beta1/tx.proto",
}

func (m *MsgListNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgListNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgListNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgListNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgListNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgListNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeListNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeListNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeListNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeListNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeListNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeListNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBuyNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBuyNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBuyNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBuyNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBuyNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBuyNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgListNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgListNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeListNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeListNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBuyNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgBuyNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgListNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgListNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgListNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgListNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgListNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgListNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeListNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeListNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeListNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeListNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeListNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeListNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBuyNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBuyNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBuyNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBuyNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBuyNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBuyNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
