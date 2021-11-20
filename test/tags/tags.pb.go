// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tags.proto

package tags

import (
	fmt "fmt"
	_ "github.com/waynz0r/protobuf/gogoproto"
	proto "github.com/waynz0r/protobuf/proto"
	math "math"
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

type Outside struct {
	*Inside `protobuf:"bytes,1,opt,name=Inside,embedded=Inside" json:""`
	Field2  *string `protobuf:"bytes,2,opt,name=Field2" json:"MyField2" xml:",comment"`
	// Types that are valid to be assigned to Filed:
	//	*Outside_Field3
	Filed                isOutside_Filed `protobuf_oneof:"filed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Outside) Reset()         { *m = Outside{} }
func (m *Outside) String() string { return proto.CompactTextString(m) }
func (*Outside) ProtoMessage()    {}
func (*Outside) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d9cbcae1e528f6, []int{0}
}
func (m *Outside) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outside.Unmarshal(m, b)
}
func (m *Outside) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outside.Marshal(b, m, deterministic)
}
func (m *Outside) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outside.Merge(m, src)
}
func (m *Outside) XXX_Size() int {
	return xxx_messageInfo_Outside.Size(m)
}
func (m *Outside) XXX_DiscardUnknown() {
	xxx_messageInfo_Outside.DiscardUnknown(m)
}

var xxx_messageInfo_Outside proto.InternalMessageInfo

type isOutside_Filed interface {
	isOutside_Filed()
}

type Outside_Field3 struct {
	Field3 string `protobuf:"bytes,3,opt,name=Field3,oneof" json:"MyField3" xml:",comment"`
}

func (*Outside_Field3) isOutside_Filed() {}

func (m *Outside) GetFiled() isOutside_Filed {
	if m != nil {
		return m.Filed
	}
	return nil
}

func (m *Outside) GetField2() string {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return ""
}

func (m *Outside) GetField3() string {
	if x, ok := m.GetFiled().(*Outside_Field3); ok {
		return x.Field3
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Outside) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Outside_Field3)(nil),
	}
}

type Inside struct {
	Field1               *string  `protobuf:"bytes,1,opt,name=Field1" json:"MyField1" xml:",chardata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inside) Reset()         { *m = Inside{} }
func (m *Inside) String() string { return proto.CompactTextString(m) }
func (*Inside) ProtoMessage()    {}
func (*Inside) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d9cbcae1e528f6, []int{1}
}
func (m *Inside) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inside.Unmarshal(m, b)
}
func (m *Inside) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inside.Marshal(b, m, deterministic)
}
func (m *Inside) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inside.Merge(m, src)
}
func (m *Inside) XXX_Size() int {
	return xxx_messageInfo_Inside.Size(m)
}
func (m *Inside) XXX_DiscardUnknown() {
	xxx_messageInfo_Inside.DiscardUnknown(m)
}

var xxx_messageInfo_Inside proto.InternalMessageInfo

func (m *Inside) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

func init() {
	proto.RegisterType((*Outside)(nil), "tags.Outside")
	proto.RegisterType((*Inside)(nil), "tags.Inside")
}

func init() { proto.RegisterFile("tags.proto", fileDescriptor_e7d9cbcae1e528f6) }

var fileDescriptor_e7d9cbcae1e528f6 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x4c, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x0c, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xcb, 0x13, 0x2b, 0xf3, 0xaa, 0x0c, 0x8a, 0xf4, 0xc1,
	0xf2, 0x49, 0xa5, 0x69, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa7, 0xb4,
	0x85, 0x91, 0x8b, 0xdd, 0xbf, 0xb4, 0xa4, 0x38, 0x33, 0x25, 0x55, 0x48, 0x8f, 0x8b, 0xcd, 0x33,
	0x0f, 0xc4, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd1, 0x03, 0x5b, 0x00, 0x11, 0x73,
	0xe2, 0xb8, 0x70, 0x4f, 0x9e, 0xf1, 0xd5, 0x3d, 0x79, 0x86, 0x20, 0xa8, 0x2a, 0x21, 0x33, 0x2e,
	0x36, 0xb7, 0xcc, 0xd4, 0x9c, 0x14, 0x23, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xb9, 0x57,
	0xf7, 0xe4, 0x39, 0x7c, 0x2b, 0x21, 0x62, 0x9f, 0xee, 0xc9, 0xf3, 0x55, 0xe4, 0xe6, 0x58, 0x29,
	0xe9, 0x24, 0xe7, 0xe7, 0xe6, 0xa6, 0xe6, 0x95, 0x28, 0x05, 0x41, 0x55, 0x0b, 0x59, 0x40, 0xf5,
	0x19, 0x4b, 0x30, 0x63, 0xe8, 0x33, 0xc6, 0xd4, 0xe7, 0xc1, 0x00, 0xd5, 0x69, 0xec, 0xc4, 0xce,
	0xc5, 0x9a, 0x96, 0x99, 0x93, 0x9a, 0xa2, 0xe4, 0x08, 0x73, 0xaa, 0x90, 0x39, 0xd4, 0x30, 0x43,
	0xb0, 0xa3, 0x39, 0x9d, 0xe4, 0x91, 0x0c, 0x33, 0xfc, 0x74, 0x4f, 0x9e, 0x1f, 0x6a, 0x58, 0x46,
	0x62, 0x51, 0x4a, 0x62, 0x49, 0x22, 0xcc, 0x15, 0x86, 0x4e, 0x2c, 0x3f, 0x1e, 0xca, 0x31, 0x02,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xec, 0x2f, 0x4f, 0xbf, 0x44, 0x01, 0x00, 0x00,
}

func NewPopulatedOutside(r randyTags, easy bool) *Outside {
	this := &Outside{}
	if r.Intn(5) != 0 {
		this.Inside = NewPopulatedInside(r, easy)
	}
	if r.Intn(5) != 0 {
		v1 := string(randStringTags(r))
		this.Field2 = &v1
	}
	oneofNumber_Filed := []int32{3}[r.Intn(1)]
	switch oneofNumber_Filed {
	case 3:
		this.Filed = NewPopulatedOutside_Field3(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 4)
	}
	return this
}

func NewPopulatedOutside_Field3(r randyTags, easy bool) *Outside_Field3 {
	this := &Outside_Field3{}
	this.Field3 = string(randStringTags(r))
	return this
}
func NewPopulatedInside(r randyTags, easy bool) *Inside {
	this := &Inside{}
	if r.Intn(5) != 0 {
		v2 := string(randStringTags(r))
		this.Field1 = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTags(r, 2)
	}
	return this
}

type randyTags interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTags(r randyTags) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTags(r randyTags) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTags(r)
	}
	return string(tmps)
}
func randUnrecognizedTags(r randyTags, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTags(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTags(dAtA []byte, r randyTags, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTags(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTags(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTags(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
