// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: em/liquidityprovider/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgMintTokens struct {
	LiquidityProvider string                                   `protobuf:"bytes,1,opt,name=liquidity_provider,json=liquidityProvider,proto3" json:"liquidity_provider,omitempty" yaml:"liquidity_provider"`
	Amount            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *MsgMintTokens) Reset()         { *m = MsgMintTokens{} }
func (m *MsgMintTokens) String() string { return proto.CompactTextString(m) }
func (*MsgMintTokens) ProtoMessage()    {}
func (*MsgMintTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_98978ada1f5f3138, []int{0}
}
func (m *MsgMintTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintTokens.Merge(m, src)
}
func (m *MsgMintTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintTokens proto.InternalMessageInfo

func (m *MsgMintTokens) GetLiquidityProvider() string {
	if m != nil {
		return m.LiquidityProvider
	}
	return ""
}

func (m *MsgMintTokens) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type MsgMintTokensResponse struct {
}

func (m *MsgMintTokensResponse) Reset()         { *m = MsgMintTokensResponse{} }
func (m *MsgMintTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintTokensResponse) ProtoMessage()    {}
func (*MsgMintTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98978ada1f5f3138, []int{1}
}
func (m *MsgMintTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintTokensResponse.Merge(m, src)
}
func (m *MsgMintTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintTokensResponse proto.InternalMessageInfo

type MsgBurnTokens struct {
	LiquidityProvider string                                   `protobuf:"bytes,1,opt,name=liquidity_provider,json=liquidityProvider,proto3" json:"liquidity_provider,omitempty" yaml:"liquidity_provider"`
	Amount            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *MsgBurnTokens) Reset()         { *m = MsgBurnTokens{} }
func (m *MsgBurnTokens) String() string { return proto.CompactTextString(m) }
func (*MsgBurnTokens) ProtoMessage()    {}
func (*MsgBurnTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_98978ada1f5f3138, []int{2}
}
func (m *MsgBurnTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnTokens.Merge(m, src)
}
func (m *MsgBurnTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnTokens proto.InternalMessageInfo

func (m *MsgBurnTokens) GetLiquidityProvider() string {
	if m != nil {
		return m.LiquidityProvider
	}
	return ""
}

func (m *MsgBurnTokens) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type MsgBurnTokensResponse struct {
}

func (m *MsgBurnTokensResponse) Reset()         { *m = MsgBurnTokensResponse{} }
func (m *MsgBurnTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBurnTokensResponse) ProtoMessage()    {}
func (*MsgBurnTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98978ada1f5f3138, []int{3}
}
func (m *MsgBurnTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnTokensResponse.Merge(m, src)
}
func (m *MsgBurnTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnTokensResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgMintTokens)(nil), "em.liquidityprovider.v1.MsgMintTokens")
	proto.RegisterType((*MsgMintTokensResponse)(nil), "em.liquidityprovider.v1.MsgMintTokensResponse")
	proto.RegisterType((*MsgBurnTokens)(nil), "em.liquidityprovider.v1.MsgBurnTokens")
	proto.RegisterType((*MsgBurnTokensResponse)(nil), "em.liquidityprovider.v1.MsgBurnTokensResponse")
}

func init() { proto.RegisterFile("em/liquidityprovider/v1/tx.proto", fileDescriptor_98978ada1f5f3138) }

var fileDescriptor_98978ada1f5f3138 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xcb, 0x4a, 0xfb, 0x40,
	0x14, 0xc6, 0x93, 0x7f, 0xa1, 0xf0, 0x1f, 0xe9, 0xc2, 0xa0, 0xf4, 0x02, 0x26, 0x25, 0x0b, 0xe9,
	0xa6, 0x33, 0xa4, 0x82, 0x0b, 0x77, 0xc6, 0xad, 0x85, 0x52, 0x5c, 0xb9, 0x91, 0xa4, 0x19, 0xe2,
	0xd0, 0xce, 0x4c, 0xcc, 0x4c, 0x42, 0xf3, 0x16, 0x3e, 0x87, 0x4f, 0xd2, 0x65, 0xc5, 0x8d, 0xab,
	0x2a, 0xed, 0x1b, 0xf4, 0x09, 0xa4, 0xb9, 0xf4, 0x42, 0xb5, 0xb8, 0x76, 0x95, 0x70, 0xf2, 0xe5,
	0x7c, 0xe7, 0xf7, 0xcd, 0x1c, 0xd0, 0xc4, 0x14, 0x8d, 0xc8, 0x53, 0x44, 0x3c, 0x22, 0x93, 0x20,
	0xe4, 0x31, 0xf1, 0x70, 0x88, 0x62, 0x0b, 0xc9, 0x31, 0x0c, 0x42, 0x2e, 0xb9, 0x56, 0xc5, 0x14,
	0xee, 0x29, 0x60, 0x6c, 0x35, 0x4e, 0x7c, 0xee, 0xf3, 0x54, 0x83, 0x56, 0x6f, 0x99, 0xbc, 0xa1,
	0x0f, 0xb8, 0xa0, 0x5c, 0x20, 0xd7, 0x11, 0x18, 0xc5, 0x96, 0x8b, 0xa5, 0x63, 0xa1, 0x01, 0x27,
	0x2c, 0xfb, 0x6e, 0xbe, 0xa9, 0xa0, 0xd2, 0x15, 0x7e, 0x97, 0x30, 0x79, 0xc7, 0x87, 0x98, 0x09,
	0xed, 0x16, 0x68, 0xeb, 0xfe, 0x0f, 0x85, 0x41, 0x4d, 0x6d, 0xaa, 0xad, 0xff, 0xf6, 0xd9, 0x72,
	0x66, 0xd4, 0x13, 0x87, 0x8e, 0xae, 0xcc, 0x7d, 0x8d, 0xd9, 0x3f, 0x5e, 0x17, 0x7b, 0x79, 0x4d,
	0x93, 0xa0, 0xec, 0x50, 0x1e, 0x31, 0x59, 0xfb, 0xd7, 0x2c, 0xb5, 0x8e, 0x3a, 0x75, 0x98, 0x0d,
	0x04, 0x57, 0x03, 0xc1, 0x7c, 0x20, 0x78, 0xc3, 0x09, 0xb3, 0xaf, 0x27, 0x33, 0x43, 0x59, 0xce,
	0x8c, 0x4a, 0x66, 0x90, 0xfd, 0x66, 0xbe, 0x7c, 0x18, 0x2d, 0x9f, 0xc8, 0xc7, 0xc8, 0x85, 0x03,
	0x4e, 0x51, 0x8e, 0x93, 0x3d, 0xda, 0xc2, 0x1b, 0x22, 0x99, 0x04, 0x58, 0xa4, 0x1d, 0x44, 0x3f,
	0xf7, 0x32, 0xab, 0xe0, 0x74, 0x07, 0xaa, 0x8f, 0x45, 0xc0, 0x99, 0xc0, 0x05, 0xae, 0x1d, 0x85,
	0xec, 0xcf, 0xe1, 0x6e, 0xa0, 0x0a, 0xdc, 0xce, 0xab, 0x0a, 0x4a, 0x5d, 0xe1, 0x6b, 0x1e, 0x00,
	0x5b, 0x27, 0x7c, 0x0e, 0x7f, 0xb8, 0x43, 0x70, 0x27, 0xb4, 0x06, 0xfc, 0x9d, 0xae, 0x70, 0x5b,
	0xb9, 0x6c, 0x05, 0x7b, 0xd0, 0x65, 0xa3, 0x3b, 0xec, 0xb2, 0xcf, 0x64, 0xf7, 0x26, 0x73, 0x5d,
	0x9d, 0xce, 0x75, 0xf5, 0x73, 0xae, 0xab, 0xcf, 0x0b, 0x5d, 0x99, 0x2e, 0x74, 0xe5, 0x7d, 0xa1,
	0x2b, 0xf7, 0x97, 0x5b, 0xc1, 0xe1, 0x36, 0xe5, 0x0c, 0x27, 0x08, 0xd3, 0xf6, 0x08, 0x7b, 0x3e,
	0x0e, 0xd1, 0xf8, 0x9b, 0xc5, 0x4a, 0xc3, 0x74, 0xcb, 0xe9, 0x2a, 0x5c, 0x7c, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x32, 0x16, 0xea, 0x39, 0x7d, 0x03, 0x00, 0x00,
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
	MintTokens(ctx context.Context, in *MsgMintTokens, opts ...grpc.CallOption) (*MsgMintTokensResponse, error)
	BurnTokens(ctx context.Context, in *MsgBurnTokens, opts ...grpc.CallOption) (*MsgBurnTokensResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MintTokens(ctx context.Context, in *MsgMintTokens, opts ...grpc.CallOption) (*MsgMintTokensResponse, error) {
	out := new(MsgMintTokensResponse)
	err := c.cc.Invoke(ctx, "/em.liquidityprovider.v1.Msg/MintTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnTokens(ctx context.Context, in *MsgBurnTokens, opts ...grpc.CallOption) (*MsgBurnTokensResponse, error) {
	out := new(MsgBurnTokensResponse)
	err := c.cc.Invoke(ctx, "/em.liquidityprovider.v1.Msg/BurnTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	MintTokens(context.Context, *MsgMintTokens) (*MsgMintTokensResponse, error)
	BurnTokens(context.Context, *MsgBurnTokens) (*MsgBurnTokensResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MintTokens(ctx context.Context, req *MsgMintTokens) (*MsgMintTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintTokens not implemented")
}
func (*UnimplementedMsgServer) BurnTokens(ctx context.Context, req *MsgBurnTokens) (*MsgBurnTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnTokens not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MintTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/em.liquidityprovider.v1.Msg/MintTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintTokens(ctx, req.(*MsgMintTokens))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/em.liquidityprovider.v1.Msg/BurnTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnTokens(ctx, req.(*MsgBurnTokens))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "em.liquidityprovider.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintTokens",
			Handler:    _Msg_MintTokens_Handler,
		},
		{
			MethodName: "BurnTokens",
			Handler:    _Msg_BurnTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "em/liquidityprovider/v1/tx.proto",
}

func (m *MsgMintTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.LiquidityProvider) > 0 {
		i -= len(m.LiquidityProvider)
		copy(dAtA[i:], m.LiquidityProvider)
		i = encodeVarintTx(dAtA, i, uint64(len(m.LiquidityProvider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBurnTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.LiquidityProvider) > 0 {
		i -= len(m.LiquidityProvider)
		copy(dAtA[i:], m.LiquidityProvider)
		i = encodeVarintTx(dAtA, i, uint64(len(m.LiquidityProvider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBurnTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgMintTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LiquidityProvider)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgMintTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBurnTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LiquidityProvider)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgBurnTokensResponse) Size() (n int) {
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
func (m *MsgMintTokens) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMintTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityProvider", wireType)
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
			m.LiquidityProvider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgMintTokensResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMintTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgBurnTokens) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBurnTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityProvider", wireType)
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
			m.LiquidityProvider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgBurnTokensResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgBurnTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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