// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync.proto

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
	Id                   uint32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchBlockReq) Reset()         { *m = FetchBlockReq{} }
func (m *FetchBlockReq) String() string { return proto.CompactTextString(m) }
func (*FetchBlockReq) ProtoMessage()    {}
func (*FetchBlockReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_857b64371a2b2501, []int{0}
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

func (m *FetchBlockReq) GetId() uint32 {
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
	return fileDescriptor_sync_857b64371a2b2501, []int{1}
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
	Layer                uint32   `protobuf:"varint,1,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerHashReq) Reset()         { *m = LayerHashReq{} }
func (m *LayerHashReq) String() string { return proto.CompactTextString(m) }
func (*LayerHashReq) ProtoMessage()    {}
func (*LayerHashReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_857b64371a2b2501, []int{2}
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

func (m *LayerHashReq) GetLayer() uint32 {
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
	return fileDescriptor_sync_857b64371a2b2501, []int{3}
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
	Layer                uint32   `protobuf:"varint,1,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerIdsReq) Reset()         { *m = LayerIdsReq{} }
func (m *LayerIdsReq) String() string { return proto.CompactTextString(m) }
func (*LayerIdsReq) ProtoMessage()    {}
func (*LayerIdsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_857b64371a2b2501, []int{4}
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

func (m *LayerIdsReq) GetLayer() uint32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

type LayerIdsResp struct {
	Ids                  []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LayerIdsResp) Reset()         { *m = LayerIdsResp{} }
func (m *LayerIdsResp) String() string { return proto.CompactTextString(m) }
func (*LayerIdsResp) ProtoMessage()    {}
func (*LayerIdsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_857b64371a2b2501, []int{5}
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

func (m *LayerIdsResp) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Block struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Layer                uint32   `protobuf:"varint,2,opt,name=layer,proto3" json:"layer,omitempty"`
	VisibleMesh          []uint32 `protobuf:"varint,3,rep,packed,name=VisibleMesh,proto3" json:"VisibleMesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_857b64371a2b2501, []int{6}
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

func (m *Block) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block) GetLayer() uint32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *Block) GetVisibleMesh() []uint32 {
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

func init() { proto.RegisterFile("sync.proto", fileDescriptor_sync_857b64371a2b2501) }

var fileDescriptor_sync_857b64371a2b2501 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3d, 0x4b, 0x04, 0x31,
	0x10, 0x86, 0xd9, 0x2f, 0x8b, 0xc9, 0xe5, 0x90, 0x60, 0x91, 0xce, 0x98, 0xb3, 0xd8, 0xea, 0x0a,
	0xef, 0x1f, 0x5c, 0x21, 0x2e, 0x28, 0x42, 0x0a, 0x0b, 0xbb, 0xcd, 0x26, 0x90, 0xc5, 0xc5, 0x8d,
	0x37, 0xd7, 0xdc, 0xbf, 0x97, 0xcc, 0x2e, 0x18, 0x84, 0xeb, 0xde, 0x37, 0xf3, 0xf0, 0x24, 0x19,
	0x00, 0xbc, 0x7c, 0x0f, 0xfb, 0x78, 0x9a, 0xcf, 0xb3, 0xa8, 0x53, 0xd6, 0xf7, 0xc0, 0x9f, 0xfd,
	0x79, 0x08, 0xc7, 0x69, 0x1e, 0xbe, 0x8c, 0xff, 0x11, 0x5b, 0x28, 0x3b, 0x27, 0x0b, 0x55, 0xb4,
	0xdc, 0x94, 0x9d, 0xd3, 0x07, 0xd8, 0xe6, 0x00, 0x46, 0xf1, 0x00, 0x8d, 0x4d, 0x45, 0x56, 0xaa,
	0x68, 0xd9, 0x13, 0xdb, 0x93, 0x74, 0x99, 0x2f, 0x13, 0xfd, 0x08, 0x9b, 0xd7, 0xfe, 0xe2, 0x4f,
	0x2f, 0x3d, 0x86, 0x24, 0xbd, 0x83, 0x66, 0x4a, 0x7d, 0xf5, 0x2e, 0x45, 0xef, 0x80, 0x67, 0x14,
	0x46, 0x21, 0xa0, 0x0e, 0x3d, 0x06, 0xa2, 0x36, 0x86, 0xb2, 0xde, 0x01, 0x23, 0xa8, 0x73, 0x78,
	0xdd, 0xa4, 0xd6, 0xfb, 0x08, 0xc2, 0x28, 0x6e, 0xa1, 0x1a, 0x1d, 0xca, 0x42, 0x55, 0x2d, 0x37,
	0x29, 0xea, 0x77, 0x68, 0xe8, 0x85, 0xff, 0xff, 0xf7, 0x27, 0x2c, 0x33, 0xa1, 0x50, 0xc0, 0x3e,
	0x46, 0x1c, 0xed, 0xe4, 0xdf, 0x3c, 0x06, 0x59, 0x91, 0x28, 0x3f, 0x3a, 0xd6, 0x9f, 0x65, 0xb4,
	0xf6, 0x86, 0x76, 0x79, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x72, 0xa5, 0xb4, 0xbb, 0x59, 0x01,
	0x00, 0x00,
}
