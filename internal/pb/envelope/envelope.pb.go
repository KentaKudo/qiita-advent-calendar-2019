// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envelope.proto

package envelope

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	Id        string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp *types.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Payload   *types.Any       `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Event) Reset()      { *m = Event{} }
func (*Event) ProtoMessage() {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee266e8c558e9dc5, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetPayload() *types.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor_ee266e8c558e9dc5) }

var fileDescriptor_ee266e8c558e9dc5 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x2b, 0x4b,
	0xcd, 0xc9, 0x2f, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12,
	0x73, 0x0b, 0xa0, 0x0a, 0x24, 0xd1, 0x15, 0x24, 0xe6, 0x55, 0x42, 0xa4, 0x94, 0x1a, 0x19, 0xb9,
	0x58, 0x5d, 0xcb, 0x52, 0xf3, 0x4a, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x2c, 0xb8, 0x38, 0xe1, 0xe6, 0x48, 0x30, 0x29, 0x30,
	0x6a, 0x70, 0x1b, 0x49, 0xe9, 0x41, 0x0c, 0xd2, 0x83, 0x19, 0xa4, 0x17, 0x02, 0x53, 0x11, 0x84,
	0x50, 0x2c, 0xa4, 0xc7, 0xc5, 0x5e, 0x90, 0x58, 0x99, 0x93, 0x9f, 0x98, 0x22, 0xc1, 0x0c, 0xd6,
	0x27, 0x82, 0xa1, 0xcf, 0x31, 0xaf, 0x32, 0x08, 0xa6, 0xc8, 0xc9, 0xe4, 0xc2, 0x43, 0x39, 0x86,
	0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c,
	0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8, 0xe2, 0x91,
	0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x28,
	0xd4, 0x38, 0x0e, 0x01, 0x00, 0x00,
}

func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	if !this.Payload.Equal(that1.Payload) {
		return false
	}
	return true
}
func (this *Event) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&envelope.Event{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	if this.Timestamp != nil {
		s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	}
	if this.Payload != nil {
		s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEnvelope(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEnvelope(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Timestamp != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEnvelope(dAtA, i, uint64(m.Timestamp.Size()))
		n1, err1 := m.Timestamp.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.Payload != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEnvelope(dAtA, i, uint64(m.Payload.Size()))
		n2, err2 := m.Payload.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	return i, nil
}

func encodeVarintEnvelope(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovEnvelope(uint64(l))
	}
	return n
}

func sovEnvelope(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEnvelope(x uint64) (n int) {
	return sovEnvelope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Timestamp:` + strings.Replace(fmt.Sprintf("%v", this.Timestamp), "Timestamp", "types.Timestamp", 1) + `,`,
		`Payload:` + strings.Replace(fmt.Sprintf("%v", this.Payload), "Any", "types.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEnvelope(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvelope
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
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
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnvelope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
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
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEnvelope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
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
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEnvelope
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &types.Any{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvelope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvelope
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEnvelope
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
func skipEnvelope(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnvelope
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
					return 0, ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnvelope
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
				return 0, ErrInvalidLengthEnvelope
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthEnvelope
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEnvelope
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEnvelope(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthEnvelope
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEnvelope = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnvelope   = fmt.Errorf("proto: integer overflow")
)
