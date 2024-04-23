// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: pipe_message.proto

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

type PipeSource int32

const (
	PipeSource_Cli     PipeSource = 0
	PipeSource_Plugin  PipeSource = 1
	PipeSource_Keybind PipeSource = 2
)

// Enum value maps for PipeSource.
var (
	PipeSource_name = map[int32]string{
		0: "Cli",
		1: "Plugin",
		2: "Keybind",
	}
	PipeSource_value = map[string]int32{
		"Cli":     0,
		"Plugin":  1,
		"Keybind": 2,
	}
)

func (x PipeSource) Enum() *PipeSource {
	p := new(PipeSource)
	*p = x
	return p
}

func (x PipeSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PipeSource) Descriptor() protoreflect.EnumDescriptor {
	return file_pipe_message_proto_enumTypes[0].Descriptor()
}

func (PipeSource) Type() protoreflect.EnumType {
	return &file_pipe_message_proto_enumTypes[0]
}

func (x PipeSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PipeSource.Descriptor instead.
func (PipeSource) EnumDescriptor() ([]byte, []int) {
	return file_pipe_message_proto_rawDescGZIP(), []int{0}
}

type PipeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source         PipeSource `protobuf:"varint,1,opt,name=source,proto3,enum=api.pipe_message.PipeSource" json:"source,omitempty"`
	CliSourceId    *string    `protobuf:"bytes,2,opt,name=cli_source_id,json=cliSourceId,proto3,oneof" json:"cli_source_id,omitempty"`
	PluginSourceId *uint32    `protobuf:"varint,3,opt,name=plugin_source_id,json=pluginSourceId,proto3,oneof" json:"plugin_source_id,omitempty"`
	Name           string     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Payload        *string    `protobuf:"bytes,5,opt,name=payload,proto3,oneof" json:"payload,omitempty"`
	Args           []*Arg     `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	IsPrivate      bool       `protobuf:"varint,7,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
}

func (x *PipeMessage) Reset() {
	*x = PipeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipe_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipeMessage) ProtoMessage() {}

func (x *PipeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pipe_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipeMessage.ProtoReflect.Descriptor instead.
func (*PipeMessage) Descriptor() ([]byte, []int) {
	return file_pipe_message_proto_rawDescGZIP(), []int{0}
}

func (x *PipeMessage) GetSource() PipeSource {
	if x != nil {
		return x.Source
	}
	return PipeSource_Cli
}

func (x *PipeMessage) GetCliSourceId() string {
	if x != nil && x.CliSourceId != nil {
		return *x.CliSourceId
	}
	return ""
}

func (x *PipeMessage) GetPluginSourceId() uint32 {
	if x != nil && x.PluginSourceId != nil {
		return *x.PluginSourceId
	}
	return 0
}

func (x *PipeMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PipeMessage) GetPayload() string {
	if x != nil && x.Payload != nil {
		return *x.Payload
	}
	return ""
}

func (x *PipeMessage) GetArgs() []*Arg {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *PipeMessage) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type Arg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Arg) Reset() {
	*x = Arg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipe_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Arg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Arg) ProtoMessage() {}

func (x *Arg) ProtoReflect() protoreflect.Message {
	mi := &file_pipe_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Arg.ProtoReflect.Descriptor instead.
func (*Arg) Descriptor() ([]byte, []int) {
	return file_pipe_message_proto_rawDescGZIP(), []int{1}
}

func (x *Arg) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Arg) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pipe_message_proto protoreflect.FileDescriptor

var file_pipe_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x69, 0x70, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xcb, 0x02, 0x0a, 0x0b, 0x50, 0x69, 0x70, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x01, 0x52, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x72, 0x67, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6c, 0x69, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2d, 0x0a, 0x03, 0x41, 0x72, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0x2e, 0x0a, 0x0a, 0x50, 0x69, 0x70, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x6c, 0x69, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x62, 0x69, 0x6e,
	0x64, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pipe_message_proto_rawDescOnce sync.Once
	file_pipe_message_proto_rawDescData = file_pipe_message_proto_rawDesc
)

func file_pipe_message_proto_rawDescGZIP() []byte {
	file_pipe_message_proto_rawDescOnce.Do(func() {
		file_pipe_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipe_message_proto_rawDescData)
	})
	return file_pipe_message_proto_rawDescData
}

var file_pipe_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pipe_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pipe_message_proto_goTypes = []interface{}{
	(PipeSource)(0),     // 0: api.pipe_message.PipeSource
	(*PipeMessage)(nil), // 1: api.pipe_message.PipeMessage
	(*Arg)(nil),         // 2: api.pipe_message.Arg
}
var file_pipe_message_proto_depIdxs = []int32{
	0, // 0: api.pipe_message.PipeMessage.source:type_name -> api.pipe_message.PipeSource
	2, // 1: api.pipe_message.PipeMessage.args:type_name -> api.pipe_message.Arg
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pipe_message_proto_init() }
func file_pipe_message_proto_init() {
	if File_pipe_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipe_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipeMessage); i {
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
		file_pipe_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Arg); i {
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
	file_pipe_message_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pipe_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pipe_message_proto_goTypes,
		DependencyIndexes: file_pipe_message_proto_depIdxs,
		EnumInfos:         file_pipe_message_proto_enumTypes,
		MessageInfos:      file_pipe_message_proto_msgTypes,
	}.Build()
	File_pipe_message_proto = out.File
	file_pipe_message_proto_rawDesc = nil
	file_pipe_message_proto_goTypes = nil
	file_pipe_message_proto_depIdxs = nil
}
