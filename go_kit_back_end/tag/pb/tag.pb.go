// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pb/tag.proto

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

type GetTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int32 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
}

func (x *GetTagRequest) Reset() {
	*x = GetTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagRequest) ProtoMessage() {}

func (x *GetTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagRequest.ProtoReflect.Descriptor instead.
func (*GetTagRequest) Descriptor() ([]byte, []int) {
	return file_pb_tag_proto_rawDescGZIP(), []int{0}
}

func (x *GetTagRequest) GetVideoId() int32 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type GetTagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetTagReply) Reset() {
	*x = GetTagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagReply) ProtoMessage() {}

func (x *GetTagReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagReply.ProtoReflect.Descriptor instead.
func (*GetTagReply) Descriptor() ([]byte, []int) {
	return file_pb_tag_proto_rawDescGZIP(), []int{1}
}

func (x *GetTagReply) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_pb_tag_proto protoreflect.FileDescriptor

var file_pb_tag_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x74, 0x61, 0x67, 0x22, 0x29, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x21,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x32, 0x3a, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x67, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x12, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x61,
	0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x61, 0x6e, 0x67,
	0x7a, 0x68, 0x65, 0x6e, 0x5a, 0x68, 0x61, 0x6f, 0x2f, 0x65, 0x61, 0x73, 0x79, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2f, 0x67, 0x6f, 0x5f, 0x6b, 0x69, 0x74, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x65,
	0x6e, 0x64, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_tag_proto_rawDescOnce sync.Once
	file_pb_tag_proto_rawDescData = file_pb_tag_proto_rawDesc
)

func file_pb_tag_proto_rawDescGZIP() []byte {
	file_pb_tag_proto_rawDescOnce.Do(func() {
		file_pb_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_tag_proto_rawDescData)
	})
	return file_pb_tag_proto_rawDescData
}

var file_pb_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_tag_proto_goTypes = []interface{}{
	(*GetTagRequest)(nil), // 0: tag.GetTagRequest
	(*GetTagReply)(nil),   // 1: tag.GetTagReply
}
var file_pb_tag_proto_depIdxs = []int32{
	0, // 0: tag.TagServe.getTag:input_type -> tag.GetTagRequest
	1, // 1: tag.TagServe.getTag:output_type -> tag.GetTagReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_tag_proto_init() }
func file_pb_tag_proto_init() {
	if File_pb_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagRequest); i {
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
		file_pb_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagReply); i {
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
			RawDescriptor: file_pb_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_tag_proto_goTypes,
		DependencyIndexes: file_pb_tag_proto_depIdxs,
		MessageInfos:      file_pb_tag_proto_msgTypes,
	}.Build()
	File_pb_tag_proto = out.File
	file_pb_tag_proto_rawDesc = nil
	file_pb_tag_proto_goTypes = nil
	file_pb_tag_proto_depIdxs = nil
}
