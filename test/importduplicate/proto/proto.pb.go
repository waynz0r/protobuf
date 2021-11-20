// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/proto.proto

package proto

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/waynz0r/protobuf/gogoproto"
	proto "github.com/waynz0r/protobuf/proto"
	math "math"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Subject struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subject) Reset()         { *m = Subject{} }
func (m *Subject) String() string { return proto.CompactTextString(m) }
func (*Subject) ProtoMessage()    {}
func (*Subject) Descriptor() ([]byte, []int) {
	return fileDescriptor_9897619deba694d0, []int{0}
}
func (m *Subject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subject.Unmarshal(m, b)
}
func (m *Subject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subject.Marshal(b, m, deterministic)
}
func (m *Subject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subject.Merge(m, src)
}
func (m *Subject) XXX_Size() int {
	return xxx_messageInfo_Subject.Size(m)
}
func (m *Subject) XXX_DiscardUnknown() {
	xxx_messageInfo_Subject.DiscardUnknown(m)
}

var xxx_messageInfo_Subject proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Subject)(nil), "proto.Subject")
}

func init() { proto.RegisterFile("proto/proto.proto", fileDescriptor_9897619deba694d0) }

var fileDescriptor_9897619deba694d0 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x07, 0x93, 0x7a, 0x60, 0x52, 0x88, 0x15, 0x4c, 0x49, 0x19, 0xa4, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x97, 0x27, 0x56, 0xe6, 0x55, 0x19, 0x14, 0x41, 0x94,
	0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe7, 0xa7, 0xe7, 0x43, 0x74, 0x82, 0x58, 0x10, 0x8d, 0x4a, 0x9c,
	0x5c, 0xec, 0xc1, 0xa5, 0x49, 0x59, 0xa9, 0xc9, 0x25, 0x4e, 0x02, 0x1f, 0x1e, 0xca, 0x31, 0xfe,
	0x78, 0x28, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x8e, 0x47, 0x72, 0x8c, 0x49, 0x6c, 0x60, 0x35,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x3b, 0xa4, 0xbb, 0x71, 0x00, 0x00, 0x00,
}

func (this *Subject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Subject)
	if !ok {
		that2, ok := that.(Subject)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Subject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&proto.Subject{")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProto(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedSubject(r randyProto, easy bool) *Subject {
	this := &Subject{}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedProto(r, 1)
	}
	return this
}

type randyProto interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProto(r randyProto) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringProto(r randyProto) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneProto(r)
	}
	return string(tmps)
}
func randUnrecognizedProto(r randyProto, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldProto(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldProto(dAtA []byte, r randyProto, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateProto(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateProto(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateProto(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateProto(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateProto(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateProto(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateProto(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
