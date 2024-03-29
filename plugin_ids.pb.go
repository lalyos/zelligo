// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: plugin_ids.proto

package zelligo

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

type PluginIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId   int32  `protobuf:"varint,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	ZellijPid  int32  `protobuf:"varint,2,opt,name=zellij_pid,json=zellijPid,proto3" json:"zellij_pid,omitempty"`
	InitialCwd string `protobuf:"bytes,3,opt,name=initial_cwd,json=initialCwd,proto3" json:"initial_cwd,omitempty"`
}

func (x *PluginIds) Reset() {
	*x = PluginIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_ids_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginIds) ProtoMessage() {}

func (x *PluginIds) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_ids_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginIds.ProtoReflect.Descriptor instead.
func (*PluginIds) Descriptor() ([]byte, []int) {
	return file_plugin_ids_proto_rawDescGZIP(), []int{0}
}

func (x *PluginIds) GetPluginId() int32 {
	if x != nil {
		return x.PluginId
	}
	return 0
}

func (x *PluginIds) GetZellijPid() int32 {
	if x != nil {
		return x.ZellijPid
	}
	return 0
}

func (x *PluginIds) GetInitialCwd() string {
	if x != nil {
		return x.InitialCwd
	}
	return ""
}

type ZellijVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ZellijVersion) Reset() {
	*x = ZellijVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_ids_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZellijVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZellijVersion) ProtoMessage() {}

func (x *ZellijVersion) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_ids_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZellijVersion.ProtoReflect.Descriptor instead.
func (*ZellijVersion) Descriptor() ([]byte, []int) {
	return file_plugin_ids_proto_rawDescGZIP(), []int{1}
}

func (x *ZellijVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_plugin_ids_proto protoreflect.FileDescriptor

var file_plugin_ids_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x73, 0x22, 0x68, 0x0a, 0x09, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x7a, 0x65, 0x6c, 0x6c, 0x69, 0x6a, 0x5f, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x7a, 0x65, 0x6c, 0x6c, 0x69, 0x6a, 0x50, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x77, 0x64, 0x22, 0x29, 0x0a, 0x0d,
	0x5a, 0x65, 0x6c, 0x6c, 0x69, 0x6a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_ids_proto_rawDescOnce sync.Once
	file_plugin_ids_proto_rawDescData = file_plugin_ids_proto_rawDesc
)

func file_plugin_ids_proto_rawDescGZIP() []byte {
	file_plugin_ids_proto_rawDescOnce.Do(func() {
		file_plugin_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_ids_proto_rawDescData)
	})
	return file_plugin_ids_proto_rawDescData
}

var file_plugin_ids_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_plugin_ids_proto_goTypes = []interface{}{
	(*PluginIds)(nil),     // 0: api.plugin_ids.PluginIds
	(*ZellijVersion)(nil), // 1: api.plugin_ids.ZellijVersion
}
var file_plugin_ids_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugin_ids_proto_init() }
func file_plugin_ids_proto_init() {
	if File_plugin_ids_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_ids_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginIds); i {
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
		file_plugin_ids_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZellijVersion); i {
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
			RawDescriptor: file_plugin_ids_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugin_ids_proto_goTypes,
		DependencyIndexes: file_plugin_ids_proto_depIdxs,
		MessageInfos:      file_plugin_ids_proto_msgTypes,
	}.Build()
	File_plugin_ids_proto = out.File
	file_plugin_ids_proto_rawDesc = nil
	file_plugin_ids_proto_goTypes = nil
	file_plugin_ids_proto_depIdxs = nil
}
