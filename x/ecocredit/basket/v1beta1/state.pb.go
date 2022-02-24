// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/basket/v1beta1/state.proto

package v1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
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

// BasketBalance stores the amount of credits from a batch in a basket
type BasketBalance struct {
	// basket_denom is the denom of the basket
	BasketDenom string `protobuf:"bytes,1,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// batch_id is the id of the credit batch
	BatchId uint64 `protobuf:"varint,2,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	// balance is the amount of ecocredits held in the basket
	Balance string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *BasketBalance) Reset()         { *m = BasketBalance{} }
func (m *BasketBalance) String() string { return proto.CompactTextString(m) }
func (*BasketBalance) ProtoMessage()    {}
func (*BasketBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_87a9a9e13a9680aa, []int{0}
}
func (m *BasketBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasketBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasketBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasketBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketBalance.Merge(m, src)
}
func (m *BasketBalance) XXX_Size() int {
	return m.Size()
}
func (m *BasketBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketBalance.DiscardUnknown(m)
}

var xxx_messageInfo_BasketBalance proto.InternalMessageInfo

func (m *BasketBalance) GetBasketDenom() string {
	if m != nil {
		return m.BasketDenom
	}
	return ""
}

func (m *BasketBalance) GetBatchId() uint64 {
	if m != nil {
		return m.BatchId
	}
	return 0
}

func (m *BasketBalance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func init() {
	proto.RegisterType((*BasketBalance)(nil), "regen.ecocredit.basket.v1beta1.BasketBalance")
}

func init() {
	proto.RegisterFile("regen/ecocredit/basket/v1beta1/state.proto", fileDescriptor_87a9a9e13a9680aa)
}

var fileDescriptor_87a9a9e13a9680aa = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x4a, 0x4d, 0x4f,
	0xcd, 0xd3, 0x4f, 0x4d, 0xce, 0x4f, 0x2e, 0x4a, 0x4d, 0xc9, 0x2c, 0xd1, 0x4f, 0x4a, 0x2c, 0xce,
	0x4e, 0x2d, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2e, 0x49, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x03, 0xab, 0xd5, 0x83, 0xab, 0xd5, 0x83, 0xa8,
	0xd5, 0x83, 0xaa, 0x95, 0x92, 0x4d, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0xcf, 0x2f, 0xca, 0xd5,
	0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x04, 0x71, 0x20, 0xda, 0x95, 0xba, 0x18, 0xb9,
	0x78, 0x9d, 0xc0, 0x3a, 0x9c, 0x12, 0x73, 0x12, 0xf3, 0x92, 0x53, 0x85, 0x14, 0xb9, 0x78, 0x20,
	0x46, 0xc4, 0xa7, 0xa4, 0xe6, 0xe5, 0xe7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x43,
	0xc4, 0x5c, 0x40, 0x42, 0x42, 0x92, 0x5c, 0x1c, 0x49, 0x89, 0x25, 0xc9, 0x19, 0xf1, 0x99, 0x29,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0xec, 0x60, 0xbe, 0x67, 0x8a, 0x90, 0x04, 0x17, 0x7b,
	0x12, 0xc4, 0x20, 0x09, 0x66, 0xb0, 0x46, 0x18, 0xd7, 0x4a, 0xf1, 0xd3, 0xbc, 0xcb, 0x7d, 0xcc,
	0xd2, 0x5c, 0xe2, 0x5c, 0xa2, 0xc8, 0xe6, 0xeb, 0xc0, 0x4d, 0x62, 0x74, 0x8a, 0x3a, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x87, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0x7d, 0xb0, 0x87, 0x75, 0xf3, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xa1, 0xbc,
	0x9c, 0xd4, 0x94, 0xf4, 0xd4, 0x22, 0xfd, 0x0a, 0x9c, 0x61, 0x96, 0xc4, 0x06, 0xf6, 0xaf, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x92, 0xd5, 0xce, 0x5c, 0x01, 0x00, 0x00,
}

func (m *BasketBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasketBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasketBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintState(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BatchId != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.BatchId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BasketDenom) > 0 {
		i -= len(m.BasketDenom)
		copy(dAtA[i:], m.BasketDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.BasketDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasketBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BasketDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.BatchId != 0 {
		n += 1 + sovState(uint64(m.BatchId))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasketBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: BasketBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasketBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
			}
			m.BatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)