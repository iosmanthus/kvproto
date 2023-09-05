// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schedulingpb.proto

package schedulingpb

import (
	"fmt"
	"io"
	"math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Participant struct {
	// name is the unique name of the scheduling participant.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// id is the unique id of the scheduling participant.
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// listen_urls is the serivce endpoint list in the url format.
	// listen_urls[0] is primary service endpoint.
	ListenUrls           []string `protobuf:"bytes,3,rep,name=listen_urls,json=listenUrls,proto3" json:"listen_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Participant) Reset()         { *m = Participant{} }
func (m *Participant) String() string { return proto.CompactTextString(m) }
func (*Participant) ProtoMessage()    {}
func (*Participant) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4bfd49510230d67, []int{0}
}
func (m *Participant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Participant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Participant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Participant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Participant.Merge(m, src)
}
func (m *Participant) XXX_Size() int {
	return m.Size()
}
func (m *Participant) XXX_DiscardUnknown() {
	xxx_messageInfo_Participant.DiscardUnknown(m)
}

var xxx_messageInfo_Participant proto.InternalMessageInfo

func (m *Participant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Participant) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Participant) GetListenUrls() []string {
	if m != nil {
		return m.ListenUrls
	}
	return nil
}

func init() {
	proto.RegisterType((*Participant)(nil), "schedulingpb.Participant")
}

func init() { proto.RegisterFile("schedulingpb.proto", fileDescriptor_b4bfd49510230d67) }

var fileDescriptor_b4bfd49510230d67 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0xce, 0x48,
	0x4d, 0x29, 0xcd, 0xc9, 0xcc, 0x4b, 0x2f, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xe8, 0x83, 0x58, 0x10, 0x35, 0x52,
	0xfc, 0x45, 0xa5, 0xc5, 0x25, 0x60, 0x26, 0x44, 0x40, 0x29, 0x88, 0x8b, 0x3b, 0x20, 0xb1, 0xa8,
	0x24, 0x33, 0x39, 0xb3, 0x20, 0x31, 0xaf, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x52, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe7, 0xe2, 0xce, 0xc9, 0x2c, 0x2e,
	0x49, 0xcd, 0x8b, 0x2f, 0x2d, 0xca, 0x29, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x0c, 0xe2, 0x82,
	0x08, 0x85, 0x16, 0xe5, 0x14, 0x3b, 0xa9, 0xdd, 0x58, 0xc1, 0xc1, 0x78, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0xc0, 0x25, 0x90, 0x5f,
	0x94, 0xae, 0x57, 0x92, 0x99, 0x5d, 0xa6, 0x97, 0x5d, 0x06, 0xb6, 0x3b, 0x89, 0x0d, 0x4c, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x73, 0xb0, 0x2b, 0xcd, 0x00, 0x00, 0x00,
}

func (m *Participant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Participant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Participant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ListenUrls) > 0 {
		for iNdEx := len(m.ListenUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ListenUrls[iNdEx])
			copy(dAtA[i:], m.ListenUrls[iNdEx])
			i = encodeVarintSchedulingpb(dAtA, i, uint64(len(m.ListenUrls[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Id != 0 {
		i = encodeVarintSchedulingpb(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchedulingpb(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchedulingpb(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchedulingpb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Participant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchedulingpb(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovSchedulingpb(uint64(m.Id))
	}
	if len(m.ListenUrls) > 0 {
		for _, s := range m.ListenUrls {
			l = len(s)
			n += 1 + l + sovSchedulingpb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSchedulingpb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchedulingpb(x uint64) (n int) {
	return sovSchedulingpb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Participant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedulingpb
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
			return fmt.Errorf("proto: Participant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Participant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedulingpb
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
				return ErrInvalidLengthSchedulingpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedulingpb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedulingpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedulingpb
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
				return ErrInvalidLengthSchedulingpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchedulingpb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenUrls = append(m.ListenUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedulingpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchedulingpb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchedulingpb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedulingpb
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
					return 0, ErrIntOverflowSchedulingpb
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
					return 0, ErrIntOverflowSchedulingpb
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
				return 0, ErrInvalidLengthSchedulingpb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchedulingpb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchedulingpb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchedulingpb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedulingpb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchedulingpb = fmt.Errorf("proto: unexpected end of group")
)