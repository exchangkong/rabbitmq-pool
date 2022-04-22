// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pool.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishReuqest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueName       string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	DeclareExchange string `protobuf:"bytes,2,opt,name=declare_exchange,json=declareExchange,proto3" json:"declare_exchange,omitempty"`
	Delay           int32  `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	Payload         string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PublishReuqest) Reset() {
	*x = PublishReuqest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReuqest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReuqest) ProtoMessage() {}

func (x *PublishReuqest) ProtoReflect() protoreflect.Message {
	mi := &file_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReuqest.ProtoReflect.Descriptor instead.
func (*PublishReuqest) Descriptor() ([]byte, []int) {
	return file_pool_proto_rawDescGZIP(), []int{0}
}

func (x *PublishReuqest) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *PublishReuqest) GetDeclareExchange() string {
	if x != nil {
		return x.DeclareExchange
	}
	return ""
}

func (x *PublishReuqest) GetDelay() int32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *PublishReuqest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_pool_proto_rawDescGZIP(), []int{1}
}

func (x *PublishResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PublishResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pool_proto protoreflect.FileDescriptor

var file_pool_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x75, 0x71, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72,
	0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x43, 0x0a, 0x0f, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x44,
	0x0a, 0x13, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x0f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x75, 0x71, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pool_proto_rawDescOnce sync.Once
	file_pool_proto_rawDescData = file_pool_proto_rawDesc
)

func file_pool_proto_rawDescGZIP() []byte {
	file_pool_proto_rawDescOnce.Do(func() {
		file_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_pool_proto_rawDescData)
	})
	return file_pool_proto_rawDescData
}

var file_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pool_proto_goTypes = []interface{}{
	(*PublishReuqest)(nil),  // 0: PublishReuqest
	(*PublishResponse)(nil), // 1: PublishResponse
}
var file_pool_proto_depIdxs = []int32{
	0, // 0: RabbitmqPoolService.Pushlish:input_type -> PublishReuqest
	1, // 1: RabbitmqPoolService.Pushlish:output_type -> PublishResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pool_proto_init() }
func file_pool_proto_init() {
	if File_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReuqest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pool_proto_goTypes,
		DependencyIndexes: file_pool_proto_depIdxs,
		MessageInfos:      file_pool_proto_msgTypes,
	}.Build()
	File_pool_proto = out.File
	file_pool_proto_rawDesc = nil
	file_pool_proto_goTypes = nil
	file_pool_proto_depIdxs = nil
}
