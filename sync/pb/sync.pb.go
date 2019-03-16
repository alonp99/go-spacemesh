// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync/pb/sync.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FetchBlockReq struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchBlockReq) Reset()         { *m = FetchBlockReq{} }
func (m *FetchBlockReq) String() string { return proto.CompactTextString(m) }
func (*FetchBlockReq) ProtoMessage()    {}
func (*FetchBlockReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_37767203ac85219d, []int{0}
}
func (m *FetchBlockReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchBlockReq.Unmarshal(m, b)
}
func (m *FetchBlockReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchBlockReq.Marshal(b, m, deterministic)
}
func (dst *FetchBlockReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockReq.Merge(dst, src)
}
func (m *FetchBlockReq) XXX_Size() int {
	return xxx_messageInfo_FetchBlockReq.Size(m)
}
func (m *FetchBlockReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockReq.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockReq proto.InternalMessageInfo

func (m *FetchBlockReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FetchBlockResp struct {
	Block                *Block   `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchBlockResp) Reset()         { *m = FetchBlockResp{} }
func (m *FetchBlockResp) String() string { return proto.CompactTextString(m) }
func (*FetchBlockResp) ProtoMessage()    {}
func (*FetchBlockResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_37767203ac85219d, []int{1}
}
func (m *FetchBlockResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchBlockResp.Unmarshal(m, b)
}
func (m *FetchBlockResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchBlockResp.Marshal(b, m, deterministic)
}
func (dst *FetchBlockResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockResp.Merge(dst, src)
}
func (m *FetchBlockResp) XXX_Size() int {
	return xxx_messageInfo_FetchBlockResp.Size(m)
}
func (m *FetchBlockResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockResp.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockResp proto.InternalMessageInfo

func (m *FetchBlockResp) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type LayerHashReq struct {
	Layer                uint64   `protobuf:"varint,1,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerHashReq) Reset()         { *m = LayerHashReq{} }
func (m *LayerHashReq) String() string { return proto.CompactTextString(m) }
func (*LayerHashReq) ProtoMessage()    {}
func (*LayerHashReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_37767203ac85219d, []int{2}
}
func (m *LayerHashReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LayerHashReq.Unmarshal(m, b)
}
func (m *LayerHashReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LayerHashReq.Marshal(b, m, deterministic)
}
func (dst *LayerHashReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LayerHashReq.Merge(dst, src)
}
func (m *LayerHashReq) XXX_Size() int {
	return xxx_messageInfo_LayerHashReq.Size(m)
}
func (m *LayerHashReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LayerHashReq.DiscardUnknown(m)
}

var xxx_messageInfo_LayerHashReq proto.InternalMessageInfo

func (m *LayerHashReq) GetLayer() uint64 {
	if m != nil {
		return m.Layer
	}
	return 0
}

type LayerHashResp struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerHashResp) Reset()         { *m = LayerHashResp{} }
func (m *LayerHashResp) String() string { return proto.CompactTextString(m) }
func (*LayerHashResp) ProtoMessage()    {}
func (*LayerHashResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_37767203ac85219d, []int{3}
}
func (m *LayerHashResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LayerHashResp.Unmarshal(m, b)
}
func (m *LayerHashResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LayerHashResp.Marshal(b, m, deterministic)
}
func (dst *LayerHashResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LayerHashResp.Merge(dst, src)
}
func (m *LayerHashResp) XXX_Size() int {
	return xxx_messageInfo_LayerHashResp.Size(m)
}
func (m *LayerHashResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LayerHashResp.DiscardUnknown(m)
}

var xxx_messageInfo_LayerHashResp proto.InternalMessageInfo

func (m *LayerHashResp) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type LayerIdsReq struct {
	Layer                uint64   `protobuf:"varint,1,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerIdsReq) Reset()         { *m = LayerIdsReq{} }
func (m *LayerIdsReq) String() string { return proto.CompactTextString(m) }
func (*LayerIdsReq) ProtoMessage()    {}
func (*LayerIdsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_37767203ac85219d, []int{4}
}
func (m *LayerIdsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LayerIdsReq.Unmarshal(m, b)
}
func (m *LayerIdsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LayerIdsReq.Marshal(b, m, deterministic)
}
func (dst *LayerIdsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LayerIdsReq.Merge(dst, src)
}
func (m *LayerIdsReq) XXX_Size() int {
	return xxx_messageInfo_LayerIdsReq.Size(m)
}
func (m *LayerIdsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LayerIdsReq.DiscardUnknown(m)
}

var xxx_messageInfo_LayerIdsReq proto.InternalMessageInfo

func (m *LayerIdsReq) GetLayer() uint64 {
	if m != nil {
		return m.Layer
	}
	return 0
}

type LayerIdsResp struct {
	Ids                  []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerIdsResp) Reset()         { *m = LayerIdsResp{} }
func (m *LayerIdsResp) String() string { return proto.CompactTextString(m) }
func (*LayerIdsResp) ProtoMessage()    {}
func (*LayerIdsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_37767203ac85219d, []int{5}
}
func (m *LayerIdsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LayerIdsResp.Unmarshal(m, b)
}
func (m *LayerIdsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LayerIdsResp.Marshal(b, m, deterministic)
}
func (dst *LayerIdsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LayerIdsResp.Merge(dst, src)
}
func (m *LayerIdsResp) XXX_Size() int {
	return xxx_messageInfo_LayerIdsResp.Size(m)
}
func (m *LayerIdsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LayerIdsResp.DiscardUnknown(m)
}

var xxx_messageInfo_LayerIdsResp proto.InternalMessageInfo

func (m *LayerIdsResp) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Block struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Layer                uint64   `protobuf:"varint,2,opt,name=layer,proto3" json:"layer,omitempty"`
	VisibleMesh          []uint64 `protobuf:"varint,3,rep,packed,name=VisibleMesh,proto3" json:"VisibleMesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_37767203ac85219d, []int{6}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block) GetLayer() uint64 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *Block) GetVisibleMesh() []uint64 {
	if m != nil {
		return m.VisibleMesh
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchBlockReq)(nil), "sync.FetchBlockReq")
	proto.RegisterType((*FetchBlockResp)(nil), "sync.FetchBlockResp")
	proto.RegisterType((*LayerHashReq)(nil), "sync.LayerHashReq")
	proto.RegisterType((*LayerHashResp)(nil), "sync.LayerHashResp")
	proto.RegisterType((*LayerIdsReq)(nil), "sync.LayerIdsReq")
	proto.RegisterType((*LayerIdsResp)(nil), "sync.LayerIdsResp")
	proto.RegisterType((*Block)(nil), "sync.Block")
}

func init() { proto.RegisterFile("sync/pb/sync.proto", fileDescriptor_sync_37767203ac85219d) }

var fileDescriptor_sync_37767203ac85219d = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3d, 0x6b, 0xc3, 0x30,
	0x10, 0x86, 0xb1, 0x2d, 0x77, 0x38, 0x27, 0xa1, 0x1c, 0x1d, 0xbc, 0x55, 0x55, 0x3a, 0x78, 0x4a,
	0xa0, 0xf9, 0x07, 0x19, 0x4a, 0x0d, 0x2d, 0x05, 0x0d, 0x1d, 0xba, 0x59, 0x96, 0x40, 0xa6, 0xa6,
	0x56, 0x73, 0x59, 0xf2, 0xef, 0x8b, 0xce, 0x86, 0x8a, 0x42, 0x26, 0xbd, 0xaf, 0xee, 0xe1, 0xd1,
	0x07, 0x20, 0x5d, 0xbe, 0xfb, 0x7d, 0x30, 0xfb, 0xb8, 0xee, 0xc2, 0x69, 0x3a, 0x4f, 0x28, 0x62,
	0x56, 0xf7, 0xb0, 0x7e, 0x76, 0xe7, 0xde, 0x1f, 0xc7, 0xa9, 0xff, 0xd2, 0xee, 0x07, 0x37, 0x90,
	0xb7, 0xb6, 0xce, 0x64, 0xd6, 0x08, 0x9d, 0xb7, 0x56, 0x1d, 0x60, 0x93, 0x02, 0x14, 0xf0, 0x01,
	0x4a, 0x13, 0x4b, 0x5d, 0xc8, 0xac, 0xa9, 0x9e, 0xaa, 0x1d, 0x4b, 0xe7, 0xf9, 0x3c, 0x51, 0x8f,
	0xb0, 0x7a, 0xed, 0x2e, 0xee, 0xf4, 0xd2, 0x91, 0x8f, 0xd2, 0x3b, 0x28, 0xc7, 0xd8, 0x17, 0xef,
	0x5c, 0xd4, 0x16, 0xd6, 0x09, 0x45, 0x01, 0x11, 0x84, 0xef, 0xc8, 0x33, 0xb5, 0xd2, 0x9c, 0xd5,
	0x16, 0x2a, 0x86, 0x5a, 0x4b, 0xd7, 0x4d, 0x72, 0x39, 0x8f, 0x21, 0x0a, 0x78, 0x0b, 0xc5, 0x60,
	0xa9, 0xce, 0x64, 0xd1, 0x08, 0x1d, 0xa3, 0x7a, 0x87, 0x92, 0x6f, 0xf8, 0xff, 0x7d, 0x7f, 0xc2,
	0x3c, 0x11, 0xa2, 0x84, 0xea, 0x63, 0xa0, 0xc1, 0x8c, 0xee, 0xcd, 0x91, 0xaf, 0x0b, 0x16, 0xa5,
	0x5b, 0x47, 0xf1, 0x99, 0x07, 0x63, 0x6e, 0xf8, 0x2f, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0x51, 0x46, 0x59, 0x61, 0x01, 0x00, 0x00,
}