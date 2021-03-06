// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: cmd/protoc-gen-go-ext/extension/tag/tag.proto

package tag

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StructTag string `protobuf:"bytes,1,opt,name=struct_tag,json=structTag,proto3" json:"struct_tag,omitempty"`
}

func (x *FieldTag) Reset() {
	*x = FieldTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTag) ProtoMessage() {}

func (x *FieldTag) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTag.ProtoReflect.Descriptor instead.
func (*FieldTag) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTag) GetStructTag() string {
	if x != nil {
		return x.StructTag
	}
	return ""
}

var file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldTag)(nil),
		Field:         65000,
		Name:          "google.protobuf.field_tag",
		Tag:           "bytes,65000,opt,name=field_tag",
		Filename:      "cmd/protoc-gen-go-ext/extension/tag/tag.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional google.protobuf.FieldTag field_tag = 65000;
	E_FieldTag = &file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_extTypes[0]
)

var File_cmd_protoc_gen_go_ext_extension_tag_tag_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x29, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x61, 0x67, 0x3a, 0x57, 0x0a,
	0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe8, 0xfb, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x52, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x42, 0x8c, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x09,
	0x54, 0x61, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x6c, 0x61,
	0x62, 0x2f, 0x6d, 0x69, 0x6e, 0x67, 0x6f, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x3b, 0x74, 0x61, 0x67, 0xf8,
	0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescData = file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDesc
)

func file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDescData
}

var file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_goTypes = []interface{}{
	(*FieldTag)(nil),                  // 0: google.protobuf.FieldTag
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_depIdxs = []int32{
	1, // 0: google.protobuf.field_tag:extendee -> google.protobuf.FieldOptions
	0, // 1: google.protobuf.field_tag:type_name -> google.protobuf.FieldTag
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_init() }
func file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_init() {
	if File_cmd_protoc_gen_go_ext_extension_tag_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldTag); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_msgTypes,
		ExtensionInfos:    file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_extTypes,
	}.Build()
	File_cmd_protoc_gen_go_ext_extension_tag_tag_proto = out.File
	file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_rawDesc = nil
	file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_goTypes = nil
	file_cmd_protoc_gen_go_ext_extension_tag_tag_proto_depIdxs = nil
}
