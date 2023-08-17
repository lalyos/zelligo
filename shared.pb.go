// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0
// source: shared.proto

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

type SwitchToModePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputMode *InputModeMessage `protobuf:"bytes,1,opt,name=input_mode,json=inputMode,proto3" json:"input_mode,omitempty"`
}

func (x *SwitchToModePayload) Reset() {
	*x = SwitchToModePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchToModePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchToModePayload) ProtoMessage() {}

func (x *SwitchToModePayload) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchToModePayload.ProtoReflect.Descriptor instead.
func (*SwitchToModePayload) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{0}
}

func (x *SwitchToModePayload) GetInputMode() *InputModeMessage {
	if x != nil {
		return x.InputMode
	}
	return nil
}

type PaneIdAndShouldFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaneId              uint32 `protobuf:"varint,1,opt,name=pane_id,json=paneId,proto3" json:"pane_id,omitempty"`
	ShouldFloatIfHidden bool   `protobuf:"varint,2,opt,name=should_float_if_hidden,json=shouldFloatIfHidden,proto3" json:"should_float_if_hidden,omitempty"`
}

func (x *PaneIdAndShouldFloat) Reset() {
	*x = PaneIdAndShouldFloat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaneIdAndShouldFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaneIdAndShouldFloat) ProtoMessage() {}

func (x *PaneIdAndShouldFloat) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaneIdAndShouldFloat.ProtoReflect.Descriptor instead.
func (*PaneIdAndShouldFloat) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{1}
}

func (x *PaneIdAndShouldFloat) GetPaneId() uint32 {
	if x != nil {
		return x.PaneId
	}
	return 0
}

func (x *PaneIdAndShouldFloat) GetShouldFloatIfHidden() bool {
	if x != nil {
		return x.ShouldFloatIfHidden
	}
	return false
}

var File_shared_proto protoreflect.FileDescriptor

var file_shared_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x10, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x13,
	0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x22, 0x64, 0x0a, 0x14, 0x50, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x41, 0x6e,
	0x64, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70,
	0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x69, 0x66, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x49, 0x66, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_shared_proto_rawDescOnce sync.Once
	file_shared_proto_rawDescData = file_shared_proto_rawDesc
)

func file_shared_proto_rawDescGZIP() []byte {
	file_shared_proto_rawDescOnce.Do(func() {
		file_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_proto_rawDescData)
	})
	return file_shared_proto_rawDescData
}

var file_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shared_proto_goTypes = []interface{}{
	(*SwitchToModePayload)(nil),  // 0: api.shared.SwitchToModePayload
	(*PaneIdAndShouldFloat)(nil), // 1: api.shared.PaneIdAndShouldFloat
	(*InputModeMessage)(nil),     // 2: api.input_mode.InputModeMessage
}
var file_shared_proto_depIdxs = []int32{
	2, // 0: api.shared.SwitchToModePayload.input_mode:type_name -> api.input_mode.InputModeMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shared_proto_init() }
func file_shared_proto_init() {
	if File_shared_proto != nil {
		return
	}
	file_input_mode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchToModePayload); i {
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
		file_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaneIdAndShouldFloat); i {
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
			RawDescriptor: file_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_proto_goTypes,
		DependencyIndexes: file_shared_proto_depIdxs,
		MessageInfos:      file_shared_proto_msgTypes,
	}.Build()
	File_shared_proto = out.File
	file_shared_proto_rawDesc = nil
	file_shared_proto_goTypes = nil
	file_shared_proto_depIdxs = nil
}