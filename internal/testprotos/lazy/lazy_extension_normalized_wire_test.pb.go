// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/lazy/lazy_extension_normalized_wire_test.proto

package lazy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Sub struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	C          *uint32 `protobuf:"varint,3,opt,name=c" json:"c,omitempty"`
	Grandchild *Sub    `protobuf:"bytes,4,opt,name=grandchild" json:"grandchild,omitempty"`
}

func (x *Sub) Reset() {
	*x = Sub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub) ProtoMessage() {}

func (x *Sub) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub.ProtoReflect.Descriptor instead.
func (*Sub) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescGZIP(), []int{0}
}

func (x *Sub) GetC() uint32 {
	if x != nil && x.C != nil {
		return *x.C
	}
	return 0
}

func (x *Sub) GetGrandchild() *Sub {
	if x != nil {
		return x.Grandchild
	}
	return nil
}

type Top struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A     *uint32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	Child *Sub    `protobuf:"bytes,2,opt,name=child" json:"child,omitempty"`
}

func (x *Top) Reset() {
	*x = Top{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Top) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Top) ProtoMessage() {}

func (x *Top) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Top.ProtoReflect.Descriptor instead.
func (*Top) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescGZIP(), []int{1}
}

func (x *Top) GetA() uint32 {
	if x != nil && x.A != nil {
		return *x.A
	}
	return 0
}

func (x *Top) GetChild() *Sub {
	if x != nil {
		return x.Child
	}
	return nil
}

type Ext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeFlag *bool `protobuf:"varint,1,opt,name=some_flag,json=someFlag" json:"some_flag,omitempty"`
}

func (x *Ext) Reset() {
	*x = Ext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ext) ProtoMessage() {}

func (x *Ext) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ext.ProtoReflect.Descriptor instead.
func (*Ext) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescGZIP(), []int{2}
}

func (x *Ext) GetSomeFlag() bool {
	if x != nil && x.SomeFlag != nil {
		return *x.SomeFlag
	}
	return false
}

var file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*Sub)(nil),
		ExtensionType: (*Ext)(nil),
		Field:         2,
		Name:          "lazy_extension_normalized_wire_test.Ext.b",
		Tag:           "bytes,2,opt,name=b",
		Filename:      "internal/testprotos/lazy/lazy_extension_normalized_wire_test.proto",
	},
}

// Extension fields to Sub.
var (
	// optional lazy_extension_normalized_wire_test.Ext b = 2;
	E_Ext_B = &file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_extTypes[0]
)

var File_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto protoreflect.FileDescriptor

var file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDesc = []byte{
	0x0a, 0x42, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f,
	0x77, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x03, 0x53, 0x75, 0x62,
	0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x63, 0x12, 0x48,
	0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77,
	0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x0a, 0x67, 0x72,
	0x61, 0x6e, 0x64, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x2a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x53,
	0x0a, 0x03, 0x54, 0x6f, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x01, 0x61, 0x12, 0x3e, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77,
	0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x05, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x03, 0x45, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x73, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x32, 0x60, 0x0a, 0x01, 0x62, 0x12, 0x28, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x72, 0x65, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x72, 0x65, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x52, 0x01, 0x62, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a,
	0x79,
}

var (
	file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescData = file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDesc
)

func file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescData)
	})
	return file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDescData
}

var file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_goTypes = []interface{}{
	(*Sub)(nil), // 0: lazy_extension_normalized_wire_test.Sub
	(*Top)(nil), // 1: lazy_extension_normalized_wire_test.Top
	(*Ext)(nil), // 2: lazy_extension_normalized_wire_test.Ext
}
var file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_depIdxs = []int32{
	0, // 0: lazy_extension_normalized_wire_test.Sub.grandchild:type_name -> lazy_extension_normalized_wire_test.Sub
	0, // 1: lazy_extension_normalized_wire_test.Top.child:type_name -> lazy_extension_normalized_wire_test.Sub
	0, // 2: lazy_extension_normalized_wire_test.Ext.b:extendee -> lazy_extension_normalized_wire_test.Sub
	2, // 3: lazy_extension_normalized_wire_test.Ext.b:type_name -> lazy_extension_normalized_wire_test.Ext
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_init() }
func file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_init() {
	if File_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Top); i {
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
		file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ext); i {
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
			RawDescriptor: file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_extTypes,
	}.Build()
	File_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto = out.File
	file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_rawDesc = nil
	file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_goTypes = nil
	file_internal_testprotos_lazy_lazy_extension_normalized_wire_test_proto_depIdxs = nil
}
