// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: service.proto

package accessibility_bridge_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

const file_service_proto_rawDesc = "" +
	"\n" +
	"\rservice.proto\x12\x1cservice.accessibility.bridge\x1a\x1bgoogle/protobuf/empty.proto\x1a\x14action_manager.proto2\x86\x04\n" +
	"\rActionManager\x12N\n" +
	"\n" +
	"ScreenDump\x12\x16.google.protobuf.Empty\x1a(.service.accessibility.bridge.ScreenView\x12Q\n" +
	"\fPerformSwipe\x12).service.accessibility.bridge.ActionSwipe\x1a\x16.google.protobuf.Empty\x12[\n" +
	"\x11PerformMultiTouch\x12..service.accessibility.bridge.ActionMultiTouch\x1a\x16.google.protobuf.Empty\x12Q\n" +
	"\fPerformClick\x12).service.accessibility.bridge.ActionClick\x1a\x16.google.protobuf.Empty\x12P\n" +
	"\bTypeText\x12,.service.accessibility.bridge.ActionTypeText\x1a\x16.google.protobuf.Empty\x12P\n" +
	"\rPerformAction\x12'.service.accessibility.bridge.ActionKey\x1a\x16.google.protobuf.EmptyB\x02P\x01b\x06proto3"

var file_service_proto_goTypes = []any{
	(*emptypb.Empty)(nil),    // 0: google.protobuf.Empty
	(*ActionSwipe)(nil),      // 1: service.accessibility.bridge.ActionSwipe
	(*ActionMultiTouch)(nil), // 2: service.accessibility.bridge.ActionMultiTouch
	(*ActionClick)(nil),      // 3: service.accessibility.bridge.ActionClick
	(*ActionTypeText)(nil),   // 4: service.accessibility.bridge.ActionTypeText
	(*ActionKey)(nil),        // 5: service.accessibility.bridge.ActionKey
	(*ScreenView)(nil),       // 6: service.accessibility.bridge.ScreenView
}
var file_service_proto_depIdxs = []int32{
	0, // 0: service.accessibility.bridge.ActionManager.ScreenDump:input_type -> google.protobuf.Empty
	1, // 1: service.accessibility.bridge.ActionManager.PerformSwipe:input_type -> service.accessibility.bridge.ActionSwipe
	2, // 2: service.accessibility.bridge.ActionManager.PerformMultiTouch:input_type -> service.accessibility.bridge.ActionMultiTouch
	3, // 3: service.accessibility.bridge.ActionManager.PerformClick:input_type -> service.accessibility.bridge.ActionClick
	4, // 4: service.accessibility.bridge.ActionManager.TypeText:input_type -> service.accessibility.bridge.ActionTypeText
	5, // 5: service.accessibility.bridge.ActionManager.PerformAction:input_type -> service.accessibility.bridge.ActionKey
	6, // 6: service.accessibility.bridge.ActionManager.ScreenDump:output_type -> service.accessibility.bridge.ScreenView
	0, // 7: service.accessibility.bridge.ActionManager.PerformSwipe:output_type -> google.protobuf.Empty
	0, // 8: service.accessibility.bridge.ActionManager.PerformMultiTouch:output_type -> google.protobuf.Empty
	0, // 9: service.accessibility.bridge.ActionManager.PerformClick:output_type -> google.protobuf.Empty
	0, // 10: service.accessibility.bridge.ActionManager.TypeText:output_type -> google.protobuf.Empty
	0, // 11: service.accessibility.bridge.ActionManager.PerformAction:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_action_manager_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
