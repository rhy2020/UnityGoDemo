// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RspHead struct {
	ErrorId              int32    `protobuf:"varint,1,opt,name=errorId,proto3" json:"errorId,omitempty"`
	ErrorString          string   `protobuf:"bytes,2,opt,name=errorString,proto3" json:"errorString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspHead) Reset()         { *m = RspHead{} }
func (m *RspHead) String() string { return proto.CompactTextString(m) }
func (*RspHead) ProtoMessage()    {}
func (*RspHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *RspHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspHead.Unmarshal(m, b)
}
func (m *RspHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspHead.Marshal(b, m, deterministic)
}
func (m *RspHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspHead.Merge(m, src)
}
func (m *RspHead) XXX_Size() int {
	return xxx_messageInfo_RspHead.Size(m)
}
func (m *RspHead) XXX_DiscardUnknown() {
	xxx_messageInfo_RspHead.DiscardUnknown(m)
}

var xxx_messageInfo_RspHead proto.InternalMessageInfo

func (m *RspHead) GetErrorId() int32 {
	if m != nil {
		return m.ErrorId
	}
	return 0
}

func (m *RspHead) GetErrorString() string {
	if m != nil {
		return m.ErrorString
	}
	return ""
}

//心跳
type HeartReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartReq) Reset()         { *m = HeartReq{} }
func (m *HeartReq) String() string { return proto.CompactTextString(m) }
func (*HeartReq) ProtoMessage()    {}
func (*HeartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *HeartReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartReq.Unmarshal(m, b)
}
func (m *HeartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartReq.Marshal(b, m, deterministic)
}
func (m *HeartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartReq.Merge(m, src)
}
func (m *HeartReq) XXX_Size() int {
	return xxx_messageInfo_HeartReq.Size(m)
}
func (m *HeartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartReq proto.InternalMessageInfo

type HeartRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartRsp) Reset()         { *m = HeartRsp{} }
func (m *HeartRsp) String() string { return proto.CompactTextString(m) }
func (*HeartRsp) ProtoMessage()    {}
func (*HeartRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *HeartRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartRsp.Unmarshal(m, b)
}
func (m *HeartRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartRsp.Marshal(b, m, deterministic)
}
func (m *HeartRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartRsp.Merge(m, src)
}
func (m *HeartRsp) XXX_Size() int {
	return xxx_messageInfo_HeartRsp.Size(m)
}
func (m *HeartRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HeartRsp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RspHead)(nil), "msg.RspHead")
	proto.RegisterType((*HeartReq)(nil), "msg.HeartReq")
	proto.RegisterType((*HeartRsp)(nil), "msg.HeartRsp")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0x72, 0xe5,
	0x62, 0x0f, 0x2a, 0x2e, 0xf0, 0x48, 0x4d, 0x4c, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x2d, 0x2a, 0xca,
	0x2f, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71, 0x85, 0x14, 0xb8, 0xb8,
	0xc1, 0xcc, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x64,
	0x21, 0x25, 0x2e, 0x2e, 0x0e, 0x8f, 0xd4, 0xc4, 0xa2, 0x92, 0xa0, 0xd4, 0x42, 0x04, 0xbb, 0xb8,
	0x20, 0x89, 0x0d, 0x6c, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x76, 0x94, 0x9e, 0xf0, 0x7a,
	0x00, 0x00, 0x00,
}