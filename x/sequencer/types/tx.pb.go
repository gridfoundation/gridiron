// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gridiron/sequencer/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// MsgCreateSequencer defines a SDK message for creating a new sequencer.
type MsgCreateSequencer struct {
	// creator is the bech32-encoded address of the sequencer account which is the account that the message was sent from.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// pubkey is the public key of the sequencers' furyint client, as a Protobuf Any.
	FuryintPubKey *types.Any `protobuf:"bytes,2,opt,name=furyintPubKey,proto3" json:"furyintPubKey,omitempty"`
	// rollappId defines the rollapp to which the sequencer belongs.
	RollappId string `protobuf:"bytes,3,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// description defines the descriptive terms for the sequencer.
	Description Description `protobuf:"bytes,4,opt,name=description,proto3" json:"description"`
}

func (m *MsgCreateSequencer) Reset()         { *m = MsgCreateSequencer{} }
func (m *MsgCreateSequencer) String() string { return proto.CompactTextString(m) }
func (*MsgCreateSequencer) ProtoMessage()    {}
func (*MsgCreateSequencer) Descriptor() ([]byte, []int) {
	return fileDescriptor_26d679aa996065f1, []int{0}
}
func (m *MsgCreateSequencer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateSequencer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateSequencer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateSequencer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateSequencer.Merge(m, src)
}
func (m *MsgCreateSequencer) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateSequencer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateSequencer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateSequencer proto.InternalMessageInfo

func (m *MsgCreateSequencer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateSequencer) GetFuryintPubKey() *types.Any {
	if m != nil {
		return m.FuryintPubKey
	}
	return nil
}

func (m *MsgCreateSequencer) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *MsgCreateSequencer) GetDescription() Description {
	if m != nil {
		return m.Description
	}
	return Description{}
}

type MsgCreateSequencerResponse struct {
}

func (m *MsgCreateSequencerResponse) Reset()         { *m = MsgCreateSequencerResponse{} }
func (m *MsgCreateSequencerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateSequencerResponse) ProtoMessage()    {}
func (*MsgCreateSequencerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_26d679aa996065f1, []int{1}
}
func (m *MsgCreateSequencerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateSequencerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateSequencerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateSequencerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateSequencerResponse.Merge(m, src)
}
func (m *MsgCreateSequencerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateSequencerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateSequencerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateSequencerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateSequencer)(nil), "gridfoundation.gridiron.sequencer.MsgCreateSequencer")
	proto.RegisterType((*MsgCreateSequencerResponse)(nil), "gridfoundation.gridiron.sequencer.MsgCreateSequencerResponse")
}

func init() { proto.RegisterFile("gridiron/sequencer/tx.proto", fileDescriptor_26d679aa996065f1) }

var fileDescriptor_26d679aa996065f1 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3f, 0x4b, 0xc3, 0x40,
	0x14, 0xcf, 0xd9, 0xa2, 0xf4, 0x2a, 0x08, 0xa1, 0x43, 0x0c, 0x25, 0x96, 0x82, 0xd0, 0xa5, 0x17,
	0x68, 0x1d, 0x5d, 0xac, 0x2e, 0x22, 0x05, 0x89, 0xb8, 0xb8, 0x48, 0x93, 0x9c, 0x31, 0xd0, 0xe4,
	0x9d, 0x77, 0x17, 0xe8, 0xb9, 0x3b, 0x0a, 0x7e, 0x18, 0x3f, 0x44, 0x71, 0xea, 0xe8, 0x24, 0xd2,
	0x7e, 0x08, 0x57, 0x69, 0xd2, 0xb4, 0xd1, 0x0a, 0x05, 0xb7, 0xfc, 0x7e, 0xef, 0xfd, 0xfe, 0xe4,
	0x25, 0xb8, 0xee, 0xab, 0x88, 0xc6, 0x22, 0x84, 0xd8, 0x16, 0xf4, 0x21, 0xa1, 0xb1, 0x47, 0xb9,
	0x2d, 0x47, 0x84, 0x71, 0x90, 0xa0, 0x37, 0x96, 0xd3, 0x91, 0x7a, 0x24, 0x4b, 0x40, 0x96, 0xab,
	0xe6, 0xe1, 0x5f, 0x7a, 0x9f, 0x0a, 0x8f, 0x87, 0x4c, 0xce, 0x57, 0x53, 0x23, 0x73, 0x3f, 0x00,
	0x08, 0x86, 0xd4, 0x4e, 0x91, 0x9b, 0xdc, 0xd9, 0x83, 0x58, 0xe5, 0x23, 0x0f, 0x44, 0x04, 0xe2,
	0x36, 0x45, 0x76, 0x06, 0x16, 0xa3, 0x5a, 0x00, 0x01, 0x64, 0xfc, 0xfc, 0x29, 0x63, 0x9b, 0x5f,
	0x08, 0xeb, 0x7d, 0x11, 0x9c, 0x72, 0x3a, 0x90, 0xf4, 0x2a, 0x0f, 0xd5, 0x0d, 0xbc, 0xe3, 0xcd,
	0x29, 0xe0, 0x06, 0x6a, 0xa0, 0x56, 0xc5, 0xc9, 0xa1, 0xee, 0xe0, 0x5d, 0x5f, 0x45, 0x61, 0x2c,
	0x2f, 0x13, 0xf7, 0x82, 0x2a, 0x63, 0xab, 0x81, 0x5a, 0xd5, 0x4e, 0x8d, 0x64, 0x9d, 0x48, 0xde,
	0x89, 0x9c, 0xc4, 0xaa, 0x67, 0xbc, 0xbd, 0xb6, 0x6b, 0x8b, 0x12, 0x1e, 0x57, 0x4c, 0x02, 0xc9,
	0x54, 0xce, 0x0f, 0x0f, 0xbd, 0x8e, 0x2b, 0x1c, 0x86, 0xc3, 0x01, 0x63, 0xe7, 0xbe, 0x51, 0x4a,
	0xf3, 0x56, 0x84, 0x7e, 0x8d, 0xab, 0x85, 0x1b, 0x18, 0xe5, 0x34, 0xb0, 0x4d, 0x36, 0x5d, 0x93,
	0x9c, 0xad, 0x44, 0xbd, 0xf2, 0xf8, 0xe3, 0x40, 0x73, 0x8a, 0x3e, 0xcd, 0x3a, 0x36, 0xd7, 0x5f,
	0xdc, 0xa1, 0x82, 0x41, 0x2c, 0x68, 0xe7, 0x19, 0xe1, 0x52, 0x5f, 0x04, 0xfa, 0x13, 0xc2, 0x7b,
	0xbf, 0x8f, 0x73, 0xb4, 0x39, 0x7b, 0xdd, 0xd9, 0x3c, 0xfe, 0x8f, 0x2a, 0xef, 0xd3, 0xeb, 0x8f,
	0xa7, 0x16, 0x9a, 0x4c, 0x2d, 0xf4, 0x39, 0xb5, 0xd0, 0xcb, 0xcc, 0xd2, 0x26, 0x33, 0x4b, 0x7b,
	0x9f, 0x59, 0xda, 0x4d, 0x37, 0x08, 0xe5, 0x7d, 0xe2, 0x12, 0x0f, 0x22, 0xbb, 0x98, 0xb0, 0x02,
	0xf6, 0xa8, 0xf8, 0x3b, 0x2a, 0x46, 0x85, 0xbb, 0x9d, 0x7e, 0xa7, 0xee, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x74, 0xc3, 0x3a, 0xb2, 0x02, 0x00, 0x00,
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
	CreateSequencer(ctx context.Context, in *MsgCreateSequencer, opts ...grpc.CallOption) (*MsgCreateSequencerResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateSequencer(ctx context.Context, in *MsgCreateSequencer, opts ...grpc.CallOption) (*MsgCreateSequencerResponse, error) {
	out := new(MsgCreateSequencerResponse)
	err := c.cc.Invoke(ctx, "/gridfoundation.gridiron.sequencer.Msg/CreateSequencer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateSequencer(context.Context, *MsgCreateSequencer) (*MsgCreateSequencerResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateSequencer(ctx context.Context, req *MsgCreateSequencer) (*MsgCreateSequencerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSequencer not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateSequencer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateSequencer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateSequencer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gridfoundation.gridiron.sequencer.Msg/CreateSequencer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateSequencer(ctx, req.(*MsgCreateSequencer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gridfoundation.gridiron.sequencer.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSequencer",
			Handler:    _Msg_CreateSequencer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gridiron/sequencer/tx.proto",
}

func (m *MsgCreateSequencer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateSequencer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateSequencer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.FuryintPubKey != nil {
		{
			size, err := m.FuryintPubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateSequencerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateSequencerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateSequencerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateSequencer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.FuryintPubKey != nil {
		l = m.FuryintPubKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateSequencerResponse) Size() (n int) {
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
func (m *MsgCreateSequencer) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateSequencer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateSequencer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuryintPubKey", wireType)
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
			if m.FuryintPubKey == nil {
				m.FuryintPubKey = &types.Any{}
			}
			if err := m.FuryintPubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
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
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateSequencerResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateSequencerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateSequencerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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