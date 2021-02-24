// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: em/authority/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type GenesisState struct {
	AuthorityKey     string                                      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
	RestrictedDenoms RestrictedDenoms                            `protobuf:"bytes,2,rep,name=restricted_denoms,json=restrictedDenoms,proto3,castrepeated=RestrictedDenoms" json:"restricted_denoms" yaml:"restricted_denoms"`
	MinGasPrices     github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,3,rep,name=min_gas_prices,json=minGasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"min_gas_prices" yaml:"min_gas_prices"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d82b0301f14f469, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAuthorityKey() string {
	if m != nil {
		return m.AuthorityKey
	}
	return ""
}

func (m *GenesisState) GetRestrictedDenoms() RestrictedDenoms {
	if m != nil {
		return m.RestrictedDenoms
	}
	return nil
}

func (m *GenesisState) GetMinGasPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.MinGasPrices
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "em.authority.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("em/authority/v1beta1/genesis.proto", fileDescriptor_8d82b0301f14f469)
}

var fileDescriptor_8d82b0301f14f469 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x6a, 0xdb, 0x40,
	0x14, 0xc7, 0x25, 0x1b, 0x0a, 0x55, 0x4d, 0x71, 0x85, 0x0b, 0xc2, 0x94, 0x91, 0x11, 0x2d, 0x18,
	0x8a, 0x67, 0xea, 0xb6, 0xab, 0xee, 0xaa, 0x18, 0x1c, 0x48, 0x16, 0x41, 0xd9, 0x65, 0x63, 0xf4,
	0xf1, 0x90, 0x07, 0x7b, 0x34, 0x46, 0x33, 0x0e, 0xd1, 0x0d, 0xb2, 0xf4, 0x39, 0x72, 0x12, 0x2f,
	0xbd, 0xcc, 0x4a, 0x09, 0xf2, 0x0d, 0x7c, 0x82, 0xa0, 0x0f, 0x2b, 0x8e, 0xe3, 0x95, 0xc4, 0xcc,
	0xef, 0xfd, 0xfe, 0x6f, 0xde, 0xd3, 0x2c, 0x60, 0xc4, 0x5d, 0xca, 0x29, 0x8f, 0xa9, 0x4c, 0xc8,
	0xed, 0xd0, 0x03, 0xe9, 0x0e, 0x49, 0x08, 0x11, 0x08, 0x2a, 0xf0, 0x22, 0xe6, 0x92, 0xeb, 0x1d,
	0x60, 0xb8, 0x66, 0x70, 0xc5, 0x74, 0x3b, 0x21, 0x0f, 0x79, 0x01, 0x90, 0xfc, 0xaf, 0x64, 0xbb,
	0xc8, 0xe7, 0x82, 0x71, 0x41, 0x3c, 0x57, 0x40, 0xad, 0xf3, 0x39, 0x8d, 0xaa, 0xfb, 0xef, 0x27,
	0xf3, 0x5e, 0xed, 0x05, 0x65, 0xa5, 0x0d, 0xad, 0x35, 0x2e, 0x7b, 0xb8, 0x96, 0xae, 0x04, 0xfd,
	0x97, 0xd6, 0x9c, 0x41, 0x62, 0xa8, 0x3d, 0xb5, 0xff, 0xd1, 0x46, 0x59, 0x6a, 0xb6, 0xfe, 0xef,
	0x4b, 0x2e, 0x20, 0xd9, 0xa5, 0xa6, 0x96, 0xb8, 0x6c, 0xfe, 0xcf, 0x9a, 0x41, 0x62, 0x39, 0x39,
	0xaa, 0xdf, 0xab, 0xda, 0x97, 0x18, 0x84, 0x8c, 0xa9, 0x2f, 0x21, 0x98, 0x04, 0x10, 0x71, 0x26,
	0x8c, 0x46, 0xaf, 0xd9, 0xff, 0xf4, 0xfb, 0x07, 0x3e, 0xf5, 0x22, 0xec, 0xd4, 0xf8, 0x28, 0xa7,
	0xed, 0xbf, 0xeb, 0xd4, 0x54, 0x76, 0xa9, 0x69, 0x94, 0xee, 0x77, 0x36, 0xeb, 0xe1, 0xc9, 0x6c,
	0x1f, 0x15, 0x09, 0xa7, 0x1d, 0x1f, 0x9d, 0xe8, 0x2b, 0x55, 0xfb, 0xcc, 0x68, 0x34, 0x09, 0x5d,
	0x31, 0x59, 0xc4, 0xd4, 0x07, 0x61, 0x34, 0x8b, 0x3e, 0xbe, 0xe1, 0x72, 0x5a, 0x38, 0x9f, 0x56,
	0xdd, 0xc6, 0x08, 0xfc, 0x33, 0x4e, 0x23, 0xfb, 0xb2, 0x8a, 0xff, 0x5a, 0xc6, 0xbf, 0x35, 0xe4,
	0xd9, 0x3f, 0x43, 0x2a, 0xa7, 0x4b, 0x0f, 0xfb, 0x9c, 0x91, 0x6a, 0xec, 0xe5, 0x67, 0x20, 0x82,
	0x19, 0x91, 0xc9, 0x02, 0xc4, 0x5e, 0x26, 0x9c, 0x16, 0xa3, 0xd1, 0xd8, 0x15, 0x57, 0x45, 0xb5,
	0x7d, 0xbe, 0xce, 0x90, 0xba, 0xc9, 0x90, 0xfa, 0x9c, 0x21, 0x75, 0xb5, 0x45, 0xca, 0x66, 0x8b,
	0x94, 0xc7, 0x2d, 0x52, 0x6e, 0xf0, 0x81, 0x14, 0x06, 0x8c, 0x47, 0x90, 0x10, 0x60, 0x83, 0x39,
	0x04, 0x21, 0xc4, 0xe4, 0xee, 0x60, 0x79, 0x45, 0x80, 0xf7, 0xa1, 0xd8, 0xd8, 0x9f, 0x97, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xea, 0xf6, 0x9a, 0x40, 0x49, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinGasPrices) > 0 {
		for iNdEx := len(m.MinGasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinGasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RestrictedDenoms) > 0 {
		for iNdEx := len(m.RestrictedDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RestrictedDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AuthorityKey) > 0 {
		i -= len(m.AuthorityKey)
		copy(dAtA[i:], m.AuthorityKey)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AuthorityKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AuthorityKey)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.RestrictedDenoms) > 0 {
		for _, e := range m.RestrictedDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MinGasPrices) > 0 {
		for _, e := range m.MinGasPrices {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorityKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestrictedDenoms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RestrictedDenoms = append(m.RestrictedDenoms, RestrictedDenom{})
			if err := m.RestrictedDenoms[len(m.RestrictedDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinGasPrices = append(m.MinGasPrices, types.DecCoin{})
			if err := m.MinGasPrices[len(m.MinGasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)