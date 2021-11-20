// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: timefail.proto

package timefail

import (
	fmt "fmt"
	_ "github.com/waynz0r/protobuf/gogoproto"
	proto "github.com/waynz0r/protobuf/proto"
	_ "github.com/waynz0r/protobuf/types"
	github_com_waynz0r_protobuf_types "github.com/waynz0r/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TimeFail struct {
	TimeTest *time.Time `protobuf:"bytes,1,opt,name=time_test,json=timeTest,proto3,stdtime" json:"time_test,omitempty"`
}

func (m *TimeFail) Reset()      { *m = TimeFail{} }
func (*TimeFail) ProtoMessage() {}
func (*TimeFail) Descriptor() ([]byte, []int) {
	return fileDescriptor_395e61815f86626a, []int{0}
}
func (m *TimeFail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeFail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeFail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeFail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeFail.Merge(m, src)
}
func (m *TimeFail) XXX_Size() int {
	return m.Size()
}
func (m *TimeFail) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeFail.DiscardUnknown(m)
}

var xxx_messageInfo_TimeFail proto.InternalMessageInfo

func (m *TimeFail) GetTimeTest() *time.Time {
	if m != nil {
		return m.TimeTest
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeFail)(nil), "timefail.TimeFail")
}

func init() { proto.RegisterFile("timefail.proto", fileDescriptor_395e61815f86626a) }

var fileDescriptor_395e61815f86626a = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc9, 0xcc, 0x4d,
	0x4d, 0x4b, 0xcc, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x0c,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xcb, 0x13, 0x2b, 0xf3, 0xaa,
	0x0c, 0x8a, 0xf4, 0xc1, 0x6a, 0x92, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30,
	0x0b, 0xa2, 0x57, 0x4a, 0x3e, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x15, 0xa1, 0x0a, 0x64, 0x56, 0x71,
	0x49, 0x62, 0x6e, 0x01, 0x44, 0x81, 0x92, 0x27, 0x17, 0x47, 0x48, 0x66, 0x6e, 0xaa, 0x5b, 0x62,
	0x66, 0x8e, 0x90, 0x2d, 0x17, 0x27, 0x48, 0x3a, 0xbe, 0x24, 0xb5, 0xb8, 0x44, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x4a, 0x0f, 0x62, 0x80, 0x1e, 0xcc, 0x00, 0xbd, 0x10, 0x98, 0x01, 0x4e,
	0x2c, 0x13, 0xee, 0xcb, 0x33, 0x06, 0x81, 0x5d, 0x17, 0x92, 0x5a, 0x5c, 0xe2, 0x64, 0x72, 0xe1,
	0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c,
	0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x6c, 0xb2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0x5a, 0xfa, 0xd4, 0xf6, 0x00, 0x00, 0x00,
}

func (this *TimeFail) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TimeFail)
	if !ok {
		that2, ok := that.(TimeFail)
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
	if that1.TimeTest == nil {
		if this.TimeTest != nil {
			return false
		}
	} else if !this.TimeTest.Equal(*that1.TimeTest) {
		return false
	}
	return true
}
func (this *TimeFail) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&timefail.TimeFail{")
	s = append(s, "TimeTest: "+fmt.Sprintf("%#v", this.TimeTest)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTimefail(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TimeFail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeFail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeFail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeTest != nil {
		n1, err1 := github_com_waynz0r_protobuf_types.StdTimeMarshalTo(*m.TimeTest, dAtA[i-github_com_waynz0r_protobuf_types.SizeOfStdTime(*m.TimeTest):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTimefail(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimefail(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimefail(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimeFail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeTest != nil {
		l = github_com_waynz0r_protobuf_types.SizeOfStdTime(*m.TimeTest)
		n += 1 + l + sovTimefail(uint64(l))
	}
	return n
}

func sovTimefail(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimefail(x uint64) (n int) {
	return sovTimefail(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TimeFail) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TimeFail{`,
		`TimeTest:` + strings.Replace(fmt.Sprintf("%v", this.TimeTest), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTimefail(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TimeFail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimefail
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
			return fmt.Errorf("proto: TimeFail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeFail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeTest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimefail
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
				return ErrInvalidLengthTimefail
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTimefail
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TimeTest == nil {
				m.TimeTest = new(time.Time)
			}
			if err := github_com_waynz0r_protobuf_types.StdTimeUnmarshal(m.TimeTest, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimefail(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimefail
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
func skipTimefail(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimefail
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
					return 0, ErrIntOverflowTimefail
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
					return 0, ErrIntOverflowTimefail
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
				return 0, ErrInvalidLengthTimefail
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimefail
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimefail
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimefail        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimefail          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimefail = fmt.Errorf("proto: unexpected end of group")
)
