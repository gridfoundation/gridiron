// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gridiron/sequencer/scheduler.proto

package types

import (
	fmt "fmt"
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

// Scheduler defines the operating status of a sequencer
type Scheduler struct {
	// sequencerAddress is the bech32-encoded address of the sequencer account, identifying the sequencer
	SequencerAddress string `protobuf:"bytes,1,opt,name=sequencerAddress,proto3" json:"sequencerAddress,omitempty"`
	// status is the operating status of this sequencer
	Status OperatingStatus `protobuf:"varint,2,opt,name=status,proto3,enum=gridfoundation.gridiron.sequencer.OperatingStatus" json:"status,omitempty"`
}

func (m *Scheduler) Reset()         { *m = Scheduler{} }
func (m *Scheduler) String() string { return proto.CompactTextString(m) }
func (*Scheduler) ProtoMessage()    {}
func (*Scheduler) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e4ce79da53e96a, []int{0}
}
func (m *Scheduler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scheduler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scheduler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scheduler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scheduler.Merge(m, src)
}
func (m *Scheduler) XXX_Size() int {
	return m.Size()
}
func (m *Scheduler) XXX_DiscardUnknown() {
	xxx_messageInfo_Scheduler.DiscardUnknown(m)
}

var xxx_messageInfo_Scheduler proto.InternalMessageInfo

func (m *Scheduler) GetSequencerAddress() string {
	if m != nil {
		return m.SequencerAddress
	}
	return ""
}

func (m *Scheduler) GetStatus() OperatingStatus {
	if m != nil {
		return m.Status
	}
	return Unspecified
}

func init() {
	proto.RegisterType((*Scheduler)(nil), "gridfoundation.gridiron.sequencer.Scheduler")
}

func init() {
	proto.RegisterFile("gridiron/sequencer/scheduler.proto", fileDescriptor_68e4ce79da53e96a)
}

var fileDescriptor_68e4ce79da53e96a = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x4e, 0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0x4e, 0x2d, 0xd2,
	0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x80, 0x2b, 0xaa, 0xa8, 0xac, 0xd2, 0x83, 0x73, 0xf4, 0xe0, 0x3a, 0xa4, 0xb4, 0xb0, 0x19,
	0x93, 0x5f, 0x90, 0x5a, 0x94, 0x58, 0x92, 0x99, 0x97, 0x1e, 0x5f, 0x5c, 0x92, 0x58, 0x52, 0x5a,
	0x0c, 0x31, 0x4d, 0xa9, 0x89, 0x91, 0x8b, 0x33, 0x18, 0x66, 0x83, 0x90, 0x16, 0x97, 0x00, 0x5c,
	0x87, 0x63, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x86,
	0xb8, 0x90, 0x27, 0x17, 0x1b, 0xc4, 0x24, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x43, 0x3d,
	0x42, 0x0e, 0xd3, 0xf3, 0x87, 0xb9, 0x21, 0x18, 0xac, 0x31, 0x08, 0x6a, 0x80, 0x93, 0xef, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0x1b, 0x8f, 0xe0, 0xe8, 0x57, 0x20, 0x79, 0xb2, 0xa4, 0xb2,
	0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x35, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x80,
	0x27, 0x35, 0x4f, 0x01, 0x00, 0x00,
}

func (m *Scheduler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scheduler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scheduler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintScheduler(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SequencerAddress) > 0 {
		i -= len(m.SequencerAddress)
		copy(dAtA[i:], m.SequencerAddress)
		i = encodeVarintScheduler(dAtA, i, uint64(len(m.SequencerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScheduler(dAtA []byte, offset int, v uint64) int {
	offset -= sovScheduler(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Scheduler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SequencerAddress)
	if l > 0 {
		n += 1 + l + sovScheduler(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovScheduler(uint64(m.Status))
	}
	return n
}

func sovScheduler(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScheduler(x uint64) (n int) {
	return sovScheduler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Scheduler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduler
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
			return fmt.Errorf("proto: Scheduler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scheduler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduler
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
				return ErrInvalidLengthScheduler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OperatingStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipScheduler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScheduler
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
func skipScheduler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScheduler
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
					return 0, ErrIntOverflowScheduler
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
					return 0, ErrIntOverflowScheduler
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
				return 0, ErrInvalidLengthScheduler
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScheduler
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScheduler
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScheduler        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScheduler          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScheduler = fmt.Errorf("proto: unexpected end of group")
)
