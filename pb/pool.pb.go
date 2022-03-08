// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pool.proto

package pb

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

type PublishReuqest struct {
	QueueName            string   `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	ExchangeName         string   `protobuf:"bytes,2,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty"`
	Delay                int32    `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	Payload              string   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishReuqest) Reset()         { *m = PublishReuqest{} }
func (m *PublishReuqest) String() string { return proto.CompactTextString(m) }
func (*PublishReuqest) ProtoMessage()    {}
func (*PublishReuqest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{0}
}

func (m *PublishReuqest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishReuqest.Unmarshal(m, b)
}
func (m *PublishReuqest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishReuqest.Marshal(b, m, deterministic)
}
func (m *PublishReuqest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishReuqest.Merge(m, src)
}
func (m *PublishReuqest) XXX_Size() int {
	return xxx_messageInfo_PublishReuqest.Size(m)
}
func (m *PublishReuqest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishReuqest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishReuqest proto.InternalMessageInfo

func (m *PublishReuqest) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *PublishReuqest) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *PublishReuqest) GetDelay() int32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *PublishReuqest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type PublishResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{1}
}

func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (m *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(m, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

func (m *PublishResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*PublishReuqest)(nil), "pool.PublishReuqest")
	proto.RegisterType((*PublishResponse)(nil), "pool.PublishResponse")
}

func init() { proto.RegisterFile("pool.proto", fileDescriptor_8a14d8612184524f) }

var fileDescriptor_8a14d8612184524f = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcf, 0xcf,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x5a, 0x18, 0xb9, 0xf8, 0x02,
	0x4a, 0x93, 0x72, 0x32, 0x8b, 0x33, 0x82, 0x52, 0x4b, 0x0b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb9,
	0xb8, 0x0a, 0x4b, 0x53, 0x4b, 0x53, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x38, 0xc1, 0x22, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0xca, 0x5c, 0xbc, 0xa9, 0x15, 0xc9,
	0x19, 0x89, 0x79, 0xe9, 0x50, 0x15, 0x4c, 0x60, 0x15, 0x3c, 0x30, 0x41, 0xb0, 0x22, 0x11, 0x2e,
	0xd6, 0x94, 0xd4, 0x9c, 0xc4, 0x4a, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x48,
	0x82, 0x8b, 0xbd, 0x20, 0xb1, 0x32, 0x27, 0x3f, 0x31, 0x45, 0x82, 0x05, 0xac, 0x09, 0xc6, 0x55,
	0xd2, 0xe4, 0xe2, 0x87, 0xbb, 0xa2, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b, 0xad,
	0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x18, 0xec, 0x04, 0xd6, 0x20, 0x28, 0xcf, 0xc8, 0x8d, 0x8b, 0x3b,
	0x20, 0x3f, 0x3f, 0x27, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x9c, 0x8b, 0x23, 0xa0,
	0xb4, 0x38, 0x03, 0xa4, 0x55, 0x48, 0x44, 0x0f, 0xec, 0x3f, 0x54, 0xff, 0x48, 0x89, 0xa2, 0x89,
	0x42, 0xcc, 0x77, 0x62, 0x89, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x07, 0x86, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x77, 0x99, 0x4b, 0x1a, 0x01, 0x00, 0x00,
}