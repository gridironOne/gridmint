// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: p2p/events.proto

package pb

import (
	fmt "fmt"
	gridmint "github.com/gridironOne/gridmint/types/pb/gridmint"
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

type GossipedBlock struct {
	Block  *gridmint.Block  `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Commit *gridmint.Commit `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (m *GossipedBlock) Reset()         { *m = GossipedBlock{} }
func (m *GossipedBlock) String() string { return proto.CompactTextString(m) }
func (*GossipedBlock) ProtoMessage()    {}
func (*GossipedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7b4e9fa812f238a, []int{0}
}
func (m *GossipedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GossipedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GossipedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GossipedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipedBlock.Merge(m, src)
}
func (m *GossipedBlock) XXX_Size() int {
	return m.Size()
}
func (m *GossipedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_GossipedBlock proto.InternalMessageInfo

func (m *GossipedBlock) GetBlock() *gridmint.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GossipedBlock) GetCommit() *gridmint.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func init() {
	proto.RegisterType((*GossipedBlock)(nil), "p2p.events.GossipedBlock")
}

func init() { proto.RegisterFile("p2p/events.proto", fileDescriptor_c7b4e9fa812f238a) }

var fileDescriptor_c7b4e9fa812f238a = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x30, 0x2a, 0xd0,
	0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x30,
	0x2a, 0xd0, 0x83, 0x88, 0x48, 0x49, 0x96, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xa7, 0x54, 0xe6, 0x66,
	0xe6, 0x95, 0x40, 0x29, 0x88, 0x32, 0xa5, 0x18, 0x2e, 0x5e, 0xf7, 0xfc, 0xe2, 0xe2, 0xcc, 0x82,
	0xd4, 0x14, 0xa7, 0x9c, 0xfc, 0xe4, 0x6c, 0x21, 0x65, 0x2e, 0xd6, 0x24, 0x10, 0x43, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xdb, 0x88, 0x57, 0x0f, 0xaa, 0x1c, 0x2c, 0x1b, 0x04, 0x91, 0x13, 0x52, 0xe3,
	0x62, 0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0x2c, 0x91, 0x60, 0x02, 0xab, 0xe2, 0x83, 0xa9, 0x72, 0x06,
	0x8b, 0x06, 0x41, 0x65, 0x9d, 0xec, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0x4a, 0x35, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x17, 0xe4, 0xa0, 0xd4, 0xbc,
	0xe2, 0xcc, 0xfc, 0xbc, 0x8a, 0xca, 0x2a, 0x98, 0x23, 0x41, 0xfe, 0x29, 0x48, 0x4a, 0x62, 0x03,
	0xbb, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xf7, 0x70, 0xf2, 0xe0, 0x00, 0x00, 0x00,
}

func (m *GossipedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GossipedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GossipedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Commit != nil {
		{
			size, err := m.Commit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GossipedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Commit != nil {
		l = m.Commit.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GossipedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: GossipedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GossipedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &gridmint.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Commit == nil {
				m.Commit = &gridmint.Commit{}
			}
			if err := m.Commit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
