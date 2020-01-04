// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room.proto

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

//房间的数据
type RoomData struct {
	RoomId               uint32          `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	MemberCount          uint32          `protobuf:"varint,2,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	MemberMaxCount       uint32          `protobuf:"varint,3,opt,name=memberMaxCount,proto3" json:"memberMaxCount,omitempty"`
	RoomMaster           *PlayerBaseInfo `protobuf:"bytes,4,opt,name=roomMaster,proto3" json:"roomMaster,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RoomData) Reset()         { *m = RoomData{} }
func (m *RoomData) String() string { return proto.CompactTextString(m) }
func (*RoomData) ProtoMessage()    {}
func (*RoomData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{0}
}

func (m *RoomData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomData.Unmarshal(m, b)
}
func (m *RoomData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomData.Marshal(b, m, deterministic)
}
func (m *RoomData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomData.Merge(m, src)
}
func (m *RoomData) XXX_Size() int {
	return xxx_messageInfo_RoomData.Size(m)
}
func (m *RoomData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomData.DiscardUnknown(m)
}

var xxx_messageInfo_RoomData proto.InternalMessageInfo

func (m *RoomData) GetRoomId() uint32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomData) GetMemberCount() uint32 {
	if m != nil {
		return m.MemberCount
	}
	return 0
}

func (m *RoomData) GetMemberMaxCount() uint32 {
	if m != nil {
		return m.MemberMaxCount
	}
	return 0
}

func (m *RoomData) GetRoomMaster() *PlayerBaseInfo {
	if m != nil {
		return m.RoomMaster
	}
	return nil
}

//请求房间
type RoomListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomListReq) Reset()         { *m = RoomListReq{} }
func (m *RoomListReq) String() string { return proto.CompactTextString(m) }
func (*RoomListReq) ProtoMessage()    {}
func (*RoomListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{1}
}

func (m *RoomListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomListReq.Unmarshal(m, b)
}
func (m *RoomListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomListReq.Marshal(b, m, deterministic)
}
func (m *RoomListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomListReq.Merge(m, src)
}
func (m *RoomListReq) XXX_Size() int {
	return xxx_messageInfo_RoomListReq.Size(m)
}
func (m *RoomListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomListReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomListReq proto.InternalMessageInfo

//请求房间列表
type RoomListRsp struct {
	Rooms                []*RoomData `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoomListRsp) Reset()         { *m = RoomListRsp{} }
func (m *RoomListRsp) String() string { return proto.CompactTextString(m) }
func (*RoomListRsp) ProtoMessage()    {}
func (*RoomListRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{2}
}

func (m *RoomListRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomListRsp.Unmarshal(m, b)
}
func (m *RoomListRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomListRsp.Marshal(b, m, deterministic)
}
func (m *RoomListRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomListRsp.Merge(m, src)
}
func (m *RoomListRsp) XXX_Size() int {
	return xxx_messageInfo_RoomListRsp.Size(m)
}
func (m *RoomListRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomListRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RoomListRsp proto.InternalMessageInfo

func (m *RoomListRsp) GetRooms() []*RoomData {
	if m != nil {
		return m.Rooms
	}
	return nil
}

//请求加入房间
type JoinRoomReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomReq) Reset()         { *m = JoinRoomReq{} }
func (m *JoinRoomReq) String() string { return proto.CompactTextString(m) }
func (*JoinRoomReq) ProtoMessage()    {}
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{3}
}

func (m *JoinRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomReq.Unmarshal(m, b)
}
func (m *JoinRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomReq.Marshal(b, m, deterministic)
}
func (m *JoinRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomReq.Merge(m, src)
}
func (m *JoinRoomReq) XXX_Size() int {
	return xxx_messageInfo_JoinRoomReq.Size(m)
}
func (m *JoinRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomReq proto.InternalMessageInfo

//请求加入房间
type JoinRoomRsp struct {
	RspHead              *RspHead          `protobuf:"bytes,1,opt,name=rspHead,proto3" json:"rspHead,omitempty"`
	Player               []*PlayerBaseInfo `protobuf:"bytes,2,rep,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JoinRoomRsp) Reset()         { *m = JoinRoomRsp{} }
func (m *JoinRoomRsp) String() string { return proto.CompactTextString(m) }
func (*JoinRoomRsp) ProtoMessage()    {}
func (*JoinRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{4}
}

func (m *JoinRoomRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomRsp.Unmarshal(m, b)
}
func (m *JoinRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomRsp.Marshal(b, m, deterministic)
}
func (m *JoinRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomRsp.Merge(m, src)
}
func (m *JoinRoomRsp) XXX_Size() int {
	return xxx_messageInfo_JoinRoomRsp.Size(m)
}
func (m *JoinRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomRsp proto.InternalMessageInfo

func (m *JoinRoomRsp) GetRspHead() *RspHead {
	if m != nil {
		return m.RspHead
	}
	return nil
}

func (m *JoinRoomRsp) GetPlayer() []*PlayerBaseInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

//请求离开房间
type LeaveRoomReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomReq) Reset()         { *m = LeaveRoomReq{} }
func (m *LeaveRoomReq) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomReq) ProtoMessage()    {}
func (*LeaveRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{5}
}

func (m *LeaveRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomReq.Unmarshal(m, b)
}
func (m *LeaveRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomReq.Marshal(b, m, deterministic)
}
func (m *LeaveRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomReq.Merge(m, src)
}
func (m *LeaveRoomReq) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomReq.Size(m)
}
func (m *LeaveRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomReq proto.InternalMessageInfo

type LeaveRoomRsp struct {
	RspHead              *RspHead `protobuf:"bytes,1,opt,name=rspHead,proto3" json:"rspHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomRsp) Reset()         { *m = LeaveRoomRsp{} }
func (m *LeaveRoomRsp) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomRsp) ProtoMessage()    {}
func (*LeaveRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{6}
}

func (m *LeaveRoomRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomRsp.Unmarshal(m, b)
}
func (m *LeaveRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomRsp.Marshal(b, m, deterministic)
}
func (m *LeaveRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomRsp.Merge(m, src)
}
func (m *LeaveRoomRsp) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomRsp.Size(m)
}
func (m *LeaveRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomRsp proto.InternalMessageInfo

func (m *LeaveRoomRsp) GetRspHead() *RspHead {
	if m != nil {
		return m.RspHead
	}
	return nil
}

//玩家进入房间
type RoomPlayerJoinNtf struct {
	Player               *PlayerBaseInfo `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RoomPlayerJoinNtf) Reset()         { *m = RoomPlayerJoinNtf{} }
func (m *RoomPlayerJoinNtf) String() string { return proto.CompactTextString(m) }
func (*RoomPlayerJoinNtf) ProtoMessage()    {}
func (*RoomPlayerJoinNtf) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{7}
}

func (m *RoomPlayerJoinNtf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomPlayerJoinNtf.Unmarshal(m, b)
}
func (m *RoomPlayerJoinNtf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomPlayerJoinNtf.Marshal(b, m, deterministic)
}
func (m *RoomPlayerJoinNtf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomPlayerJoinNtf.Merge(m, src)
}
func (m *RoomPlayerJoinNtf) XXX_Size() int {
	return xxx_messageInfo_RoomPlayerJoinNtf.Size(m)
}
func (m *RoomPlayerJoinNtf) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomPlayerJoinNtf.DiscardUnknown(m)
}

var xxx_messageInfo_RoomPlayerJoinNtf proto.InternalMessageInfo

func (m *RoomPlayerJoinNtf) GetPlayer() *PlayerBaseInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

//玩家离开房间
type RoomPlayerLeaveNtf struct {
	PlayerId             uint32   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomPlayerLeaveNtf) Reset()         { *m = RoomPlayerLeaveNtf{} }
func (m *RoomPlayerLeaveNtf) String() string { return proto.CompactTextString(m) }
func (*RoomPlayerLeaveNtf) ProtoMessage()    {}
func (*RoomPlayerLeaveNtf) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{8}
}

func (m *RoomPlayerLeaveNtf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomPlayerLeaveNtf.Unmarshal(m, b)
}
func (m *RoomPlayerLeaveNtf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomPlayerLeaveNtf.Marshal(b, m, deterministic)
}
func (m *RoomPlayerLeaveNtf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomPlayerLeaveNtf.Merge(m, src)
}
func (m *RoomPlayerLeaveNtf) XXX_Size() int {
	return xxx_messageInfo_RoomPlayerLeaveNtf.Size(m)
}
func (m *RoomPlayerLeaveNtf) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomPlayerLeaveNtf.DiscardUnknown(m)
}

var xxx_messageInfo_RoomPlayerLeaveNtf proto.InternalMessageInfo

func (m *RoomPlayerLeaveNtf) GetPlayerId() uint32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

//请求战斗准备
type RoomBattleReadyReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleReadyReq) Reset()         { *m = RoomBattleReadyReq{} }
func (m *RoomBattleReadyReq) String() string { return proto.CompactTextString(m) }
func (*RoomBattleReadyReq) ProtoMessage()    {}
func (*RoomBattleReadyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{9}
}

func (m *RoomBattleReadyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleReadyReq.Unmarshal(m, b)
}
func (m *RoomBattleReadyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleReadyReq.Marshal(b, m, deterministic)
}
func (m *RoomBattleReadyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleReadyReq.Merge(m, src)
}
func (m *RoomBattleReadyReq) XXX_Size() int {
	return xxx_messageInfo_RoomBattleReadyReq.Size(m)
}
func (m *RoomBattleReadyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleReadyReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleReadyReq proto.InternalMessageInfo

type RoomBattleReadyRsp struct {
	RspHead              *RspHead `protobuf:"bytes,1,opt,name=rspHead,proto3" json:"rspHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleReadyRsp) Reset()         { *m = RoomBattleReadyRsp{} }
func (m *RoomBattleReadyRsp) String() string { return proto.CompactTextString(m) }
func (*RoomBattleReadyRsp) ProtoMessage()    {}
func (*RoomBattleReadyRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{10}
}

func (m *RoomBattleReadyRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleReadyRsp.Unmarshal(m, b)
}
func (m *RoomBattleReadyRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleReadyRsp.Marshal(b, m, deterministic)
}
func (m *RoomBattleReadyRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleReadyRsp.Merge(m, src)
}
func (m *RoomBattleReadyRsp) XXX_Size() int {
	return xxx_messageInfo_RoomBattleReadyRsp.Size(m)
}
func (m *RoomBattleReadyRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleReadyRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleReadyRsp proto.InternalMessageInfo

func (m *RoomBattleReadyRsp) GetRspHead() *RspHead {
	if m != nil {
		return m.RspHead
	}
	return nil
}

//战斗准备推送
type RoomBattleReadyNtf struct {
	PlayerId             uint32   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleReadyNtf) Reset()         { *m = RoomBattleReadyNtf{} }
func (m *RoomBattleReadyNtf) String() string { return proto.CompactTextString(m) }
func (*RoomBattleReadyNtf) ProtoMessage()    {}
func (*RoomBattleReadyNtf) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{11}
}

func (m *RoomBattleReadyNtf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleReadyNtf.Unmarshal(m, b)
}
func (m *RoomBattleReadyNtf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleReadyNtf.Marshal(b, m, deterministic)
}
func (m *RoomBattleReadyNtf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleReadyNtf.Merge(m, src)
}
func (m *RoomBattleReadyNtf) XXX_Size() int {
	return xxx_messageInfo_RoomBattleReadyNtf.Size(m)
}
func (m *RoomBattleReadyNtf) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleReadyNtf.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleReadyNtf proto.InternalMessageInfo

func (m *RoomBattleReadyNtf) GetPlayerId() uint32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

//战斗取消准备
type RoomBattleCancleReadyReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleCancleReadyReq) Reset()         { *m = RoomBattleCancleReadyReq{} }
func (m *RoomBattleCancleReadyReq) String() string { return proto.CompactTextString(m) }
func (*RoomBattleCancleReadyReq) ProtoMessage()    {}
func (*RoomBattleCancleReadyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{12}
}

func (m *RoomBattleCancleReadyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleCancleReadyReq.Unmarshal(m, b)
}
func (m *RoomBattleCancleReadyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleCancleReadyReq.Marshal(b, m, deterministic)
}
func (m *RoomBattleCancleReadyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleCancleReadyReq.Merge(m, src)
}
func (m *RoomBattleCancleReadyReq) XXX_Size() int {
	return xxx_messageInfo_RoomBattleCancleReadyReq.Size(m)
}
func (m *RoomBattleCancleReadyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleCancleReadyReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleCancleReadyReq proto.InternalMessageInfo

type RoomBattleCancleReadyRsp struct {
	RspHead              *RspHead `protobuf:"bytes,1,opt,name=rspHead,proto3" json:"rspHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleCancleReadyRsp) Reset()         { *m = RoomBattleCancleReadyRsp{} }
func (m *RoomBattleCancleReadyRsp) String() string { return proto.CompactTextString(m) }
func (*RoomBattleCancleReadyRsp) ProtoMessage()    {}
func (*RoomBattleCancleReadyRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{13}
}

func (m *RoomBattleCancleReadyRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleCancleReadyRsp.Unmarshal(m, b)
}
func (m *RoomBattleCancleReadyRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleCancleReadyRsp.Marshal(b, m, deterministic)
}
func (m *RoomBattleCancleReadyRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleCancleReadyRsp.Merge(m, src)
}
func (m *RoomBattleCancleReadyRsp) XXX_Size() int {
	return xxx_messageInfo_RoomBattleCancleReadyRsp.Size(m)
}
func (m *RoomBattleCancleReadyRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleCancleReadyRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleCancleReadyRsp proto.InternalMessageInfo

func (m *RoomBattleCancleReadyRsp) GetRspHead() *RspHead {
	if m != nil {
		return m.RspHead
	}
	return nil
}

//取消战斗
type RoomBattleCancleReadyNtf struct {
	PlayerId             uint32   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleCancleReadyNtf) Reset()         { *m = RoomBattleCancleReadyNtf{} }
func (m *RoomBattleCancleReadyNtf) String() string { return proto.CompactTextString(m) }
func (*RoomBattleCancleReadyNtf) ProtoMessage()    {}
func (*RoomBattleCancleReadyNtf) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{14}
}

func (m *RoomBattleCancleReadyNtf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleCancleReadyNtf.Unmarshal(m, b)
}
func (m *RoomBattleCancleReadyNtf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleCancleReadyNtf.Marshal(b, m, deterministic)
}
func (m *RoomBattleCancleReadyNtf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleCancleReadyNtf.Merge(m, src)
}
func (m *RoomBattleCancleReadyNtf) XXX_Size() int {
	return xxx_messageInfo_RoomBattleCancleReadyNtf.Size(m)
}
func (m *RoomBattleCancleReadyNtf) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleCancleReadyNtf.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleCancleReadyNtf proto.InternalMessageInfo

func (m *RoomBattleCancleReadyNtf) GetPlayerId() uint32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

//战斗准备
type RoomBattleStartReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleStartReq) Reset()         { *m = RoomBattleStartReq{} }
func (m *RoomBattleStartReq) String() string { return proto.CompactTextString(m) }
func (*RoomBattleStartReq) ProtoMessage()    {}
func (*RoomBattleStartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{15}
}

func (m *RoomBattleStartReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleStartReq.Unmarshal(m, b)
}
func (m *RoomBattleStartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleStartReq.Marshal(b, m, deterministic)
}
func (m *RoomBattleStartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleStartReq.Merge(m, src)
}
func (m *RoomBattleStartReq) XXX_Size() int {
	return xxx_messageInfo_RoomBattleStartReq.Size(m)
}
func (m *RoomBattleStartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleStartReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleStartReq proto.InternalMessageInfo

type RoomBattleStartRsp struct {
	RspHead              *RspHead `protobuf:"bytes,1,opt,name=rspHead,proto3" json:"rspHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleStartRsp) Reset()         { *m = RoomBattleStartRsp{} }
func (m *RoomBattleStartRsp) String() string { return proto.CompactTextString(m) }
func (*RoomBattleStartRsp) ProtoMessage()    {}
func (*RoomBattleStartRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{16}
}

func (m *RoomBattleStartRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleStartRsp.Unmarshal(m, b)
}
func (m *RoomBattleStartRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleStartRsp.Marshal(b, m, deterministic)
}
func (m *RoomBattleStartRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleStartRsp.Merge(m, src)
}
func (m *RoomBattleStartRsp) XXX_Size() int {
	return xxx_messageInfo_RoomBattleStartRsp.Size(m)
}
func (m *RoomBattleStartRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleStartRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleStartRsp proto.InternalMessageInfo

func (m *RoomBattleStartRsp) GetRspHead() *RspHead {
	if m != nil {
		return m.RspHead
	}
	return nil
}

type RoomBattleStartNtf struct {
	RspHead              *RspHead `protobuf:"bytes,1,opt,name=rspHead,proto3" json:"rspHead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomBattleStartNtf) Reset()         { *m = RoomBattleStartNtf{} }
func (m *RoomBattleStartNtf) String() string { return proto.CompactTextString(m) }
func (*RoomBattleStartNtf) ProtoMessage()    {}
func (*RoomBattleStartNtf) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{17}
}

func (m *RoomBattleStartNtf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomBattleStartNtf.Unmarshal(m, b)
}
func (m *RoomBattleStartNtf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomBattleStartNtf.Marshal(b, m, deterministic)
}
func (m *RoomBattleStartNtf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomBattleStartNtf.Merge(m, src)
}
func (m *RoomBattleStartNtf) XXX_Size() int {
	return xxx_messageInfo_RoomBattleStartNtf.Size(m)
}
func (m *RoomBattleStartNtf) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomBattleStartNtf.DiscardUnknown(m)
}

var xxx_messageInfo_RoomBattleStartNtf proto.InternalMessageInfo

func (m *RoomBattleStartNtf) GetRspHead() *RspHead {
	if m != nil {
		return m.RspHead
	}
	return nil
}

func init() {
	proto.RegisterType((*RoomData)(nil), "msg.RoomData")
	proto.RegisterType((*RoomListReq)(nil), "msg.RoomListReq")
	proto.RegisterType((*RoomListRsp)(nil), "msg.RoomListRsp")
	proto.RegisterType((*JoinRoomReq)(nil), "msg.JoinRoomReq")
	proto.RegisterType((*JoinRoomRsp)(nil), "msg.JoinRoomRsp")
	proto.RegisterType((*LeaveRoomReq)(nil), "msg.LeaveRoomReq")
	proto.RegisterType((*LeaveRoomRsp)(nil), "msg.LeaveRoomRsp")
	proto.RegisterType((*RoomPlayerJoinNtf)(nil), "msg.RoomPlayerJoinNtf")
	proto.RegisterType((*RoomPlayerLeaveNtf)(nil), "msg.RoomPlayerLeaveNtf")
	proto.RegisterType((*RoomBattleReadyReq)(nil), "msg.RoomBattleReadyReq")
	proto.RegisterType((*RoomBattleReadyRsp)(nil), "msg.RoomBattleReadyRsp")
	proto.RegisterType((*RoomBattleReadyNtf)(nil), "msg.RoomBattleReadyNtf")
	proto.RegisterType((*RoomBattleCancleReadyReq)(nil), "msg.RoomBattleCancleReadyReq")
	proto.RegisterType((*RoomBattleCancleReadyRsp)(nil), "msg.RoomBattleCancleReadyRsp")
	proto.RegisterType((*RoomBattleCancleReadyNtf)(nil), "msg.RoomBattleCancleReadyNtf")
	proto.RegisterType((*RoomBattleStartReq)(nil), "msg.RoomBattleStartReq")
	proto.RegisterType((*RoomBattleStartRsp)(nil), "msg.RoomBattleStartRsp")
	proto.RegisterType((*RoomBattleStartNtf)(nil), "msg.RoomBattleStartNtf")
}

func init() { proto.RegisterFile("room.proto", fileDescriptor_c5fd27dd97284ef4) }

var fileDescriptor_c5fd27dd97284ef4 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xd1, 0x6a, 0xc2, 0x30,
	0x14, 0x86, 0xa9, 0x6e, 0x4e, 0x4e, 0x55, 0x58, 0x36, 0x46, 0xe9, 0x55, 0xe9, 0x40, 0x84, 0x81,
	0x48, 0x05, 0xaf, 0x76, 0x31, 0x74, 0x17, 0x73, 0xe8, 0x18, 0xd9, 0x13, 0x44, 0x8d, 0x22, 0x98,
	0xa6, 0x4b, 0xb2, 0x31, 0x1f, 0x67, 0x6f, 0x3a, 0x92, 0x58, 0x0d, 0xeb, 0x64, 0xf6, 0xf2, 0xfc,
	0xe7, 0xfc, 0xff, 0xf9, 0x4e, 0x4a, 0x01, 0x04, 0xe7, 0xac, 0x9b, 0x09, 0xae, 0x38, 0xaa, 0x32,
	0xb9, 0x0a, 0xfd, 0x0d, 0x5f, 0xad, 0x53, 0xab, 0x84, 0x8d, 0x39, 0x67, 0x8c, 0xef, 0xaa, 0xf8,
	0xdb, 0x83, 0x3a, 0xe6, 0x9c, 0x3d, 0x12, 0x45, 0xd0, 0x0d, 0xd4, 0xb4, 0x75, 0xbc, 0x08, 0xbc,
	0xc8, 0xeb, 0x34, 0xf1, 0xae, 0x42, 0x11, 0xf8, 0x8c, 0xb2, 0x19, 0x15, 0x23, 0xfe, 0x91, 0xaa,
	0xa0, 0x62, 0x9a, 0xae, 0x84, 0xda, 0xd0, 0xb2, 0xe5, 0x94, 0x7c, 0xd9, 0xa1, 0xaa, 0x19, 0xfa,
	0xa5, 0xa2, 0xbe, 0x85, 0x9b, 0x12, 0xa9, 0xa8, 0x08, 0xce, 0x22, 0xaf, 0xe3, 0x27, 0x57, 0x5d,
	0x26, 0x57, 0xdd, 0xd7, 0x0d, 0xd9, 0x52, 0x31, 0x24, 0x92, 0x8e, 0xd3, 0x25, 0xc7, 0xce, 0x58,
	0xdc, 0x04, 0x5f, 0x23, 0x4e, 0xd6, 0x52, 0x61, 0xfa, 0x1e, 0x27, 0x4e, 0x29, 0x33, 0x74, 0x0b,
	0xe7, 0x7a, 0x56, 0x06, 0x5e, 0x54, 0xed, 0xf8, 0x49, 0xd3, 0xa4, 0xe5, 0x27, 0x61, 0xdb, 0xd3,
	0x11, 0xcf, 0x7c, 0x9d, 0x6a, 0x59, 0x47, 0xcc, 0x9c, 0x52, 0x66, 0xa8, 0x0d, 0x17, 0x42, 0x66,
	0x4f, 0x94, 0xd8, 0xc3, 0xfd, 0xa4, 0x61, 0x43, 0xac, 0x86, 0xf3, 0x26, 0xba, 0x83, 0x5a, 0x66,
	0x30, 0x83, 0x8a, 0xd9, 0xf5, 0x27, 0xf9, 0x6e, 0x24, 0x6e, 0x41, 0x63, 0x42, 0xc9, 0x27, 0xcd,
	0x77, 0x0e, 0xdc, 0xfa, 0xf4, 0xa5, 0xf1, 0x03, 0x5c, 0x6a, 0x8b, 0xdd, 0xa2, 0xa9, 0x5f, 0xd4,
	0xd2, 0x21, 0xf1, 0x8e, 0xbf, 0x61, 0x4e, 0xd2, 0x03, 0x74, 0x48, 0x30, 0x0c, 0x3a, 0x22, 0x84,
	0xba, 0xed, 0xef, 0x3f, 0xf7, 0xbe, 0x8e, 0xaf, 0xad, 0x63, 0x48, 0x94, 0xda, 0x50, 0x4c, 0xc9,
	0x62, 0xab, 0x2f, 0xb8, 0x2f, 0xaa, 0x25, 0xee, 0xe8, 0x15, 0xdc, 0xff, 0x51, 0x84, 0x10, 0x1c,
	0x1c, 0x23, 0x92, 0xce, 0x1d, 0x96, 0xe1, 0xb1, 0x5e, 0x09, 0xa2, 0xc1, 0x91, 0x8c, 0x52, 0xaf,
	0xf3, 0xa6, 0x88, 0x50, 0x85, 0xd7, 0xb1, 0x6a, 0x09, 0x96, 0xa2, 0x5b, 0x53, 0x9c, 0xe8, 0x9e,
	0xd5, 0xcc, 0xcf, 0xdc, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x41, 0x4e, 0x38, 0x65, 0xfa, 0x03,
	0x00, 0x00,
}
