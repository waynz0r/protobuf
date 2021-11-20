// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue435.proto

package issue435

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

type Message struct {
	NonnullableOptional  SubMessage   `protobuf:"bytes,1,opt,name=nonnullable_optional,json=nonnullableOptional" json:"nonnullable_optional"`
	NonnullableRepeated  []SubMessage `protobuf:"bytes,2,rep,name=nonnullable_repeated,json=nonnullableRepeated" json:"nonnullable_repeated"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_dba67e55f800df0d, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetNonnullableOptional() SubMessage {
	if m != nil {
		return m.NonnullableOptional
	}
	return SubMessage{}
}

func (m *Message) GetNonnullableRepeated() []SubMessage {
	if m != nil {
		return m.NonnullableRepeated
	}
	return nil
}

type SubMessage struct {
	Value                *int64   `protobuf:"varint,1,opt,name=value,def=7" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubMessage) Reset()         { *m = SubMessage{} }
func (m *SubMessage) String() string { return proto.CompactTextString(m) }
func (*SubMessage) ProtoMessage()    {}
func (*SubMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_dba67e55f800df0d, []int{1}
}
func (m *SubMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubMessage.Unmarshal(m, b)
}
func (m *SubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubMessage.Marshal(b, m, deterministic)
}
func (m *SubMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubMessage.Merge(m, src)
}
func (m *SubMessage) XXX_Size() int {
	return xxx_messageInfo_SubMessage.Size(m)
}
func (m *SubMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubMessage proto.InternalMessageInfo

const Default_SubMessage_Value int64 = 7

func (m *SubMessage) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_SubMessage_Value
}

func init() {
	proto.RegisterType((*Message)(nil), "issue435.Message")
	proto.RegisterType((*SubMessage)(nil), "issue435.SubMessage")
}

func init() { proto.RegisterFile("issue435.proto", fileDescriptor_dba67e55f800df0d) }

var fileDescriptor_dba67e55f800df0d = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0x36, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x0c,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xcb, 0x13, 0x2b, 0xf3, 0xaa,
	0x0c, 0x8a, 0xf4, 0xc1, 0x6a, 0x92, 0x4a, 0xd3, 0xf4, 0xd3, 0xf3, 0xd3, 0xf3, 0xc1, 0x1c, 0x30,
	0x0b, 0xa2, 0x57, 0x69, 0x39, 0x23, 0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90,
	0x2f, 0x97, 0x48, 0x5e, 0x7e, 0x5e, 0x5e, 0x69, 0x4e, 0x4e, 0x62, 0x52, 0x4e, 0x6a, 0x7c, 0x7e,
	0x41, 0x49, 0x66, 0x7e, 0x5e, 0x62, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e,
	0xdc, 0xda, 0xe0, 0xd2, 0x24, 0xa8, 0x1e, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x84, 0x91,
	0xf4, 0xf9, 0x43, 0xb5, 0xa1, 0x1b, 0x57, 0x94, 0x5a, 0x90, 0x9a, 0x58, 0x92, 0x9a, 0x22, 0xc1,
	0xa4, 0xc0, 0x4c, 0x82, 0x71, 0x41, 0x50, 0x6d, 0x4a, 0xaa, 0x5c, 0x5c, 0x08, 0x85, 0x42, 0xe2,
	0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x60, 0xc7, 0x31, 0x5b, 0x31, 0x9a, 0x07, 0x41, 0xf8,
	0x4e, 0x5c, 0x51, 0xf0, 0xe0, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0x95, 0xa2, 0x7b, 0x87, 0x29,
	0x01, 0x00, 0x00,
}
