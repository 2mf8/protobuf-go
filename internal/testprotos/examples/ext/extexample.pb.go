// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/examples/ext/extexample.proto

package ext

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
)

type Concert struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_HeadlinerName *string                `protobuf:"bytes,1,opt,name=headliner_name,json=headlinerName"`
	XXX_raceDetectHookData   protoimpl.RaceDetectHookData
	XXX_presence             [1]uint32
	extensionFields          protoimpl.ExtensionFields
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *Concert) Reset() {
	*x = Concert{}
	mi := &file_internal_testprotos_examples_ext_extexample_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Concert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Concert) ProtoMessage() {}

func (x *Concert) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_examples_ext_extexample_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Concert) GetHeadlinerName() string {
	if x != nil {
		if x.xxx_hidden_HeadlinerName != nil {
			return *x.xxx_hidden_HeadlinerName
		}
		return ""
	}
	return ""
}

func (x *Concert) SetHeadlinerName(v string) {
	x.xxx_hidden_HeadlinerName = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Concert) HasHeadlinerName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Concert) ClearHeadlinerName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_HeadlinerName = nil
}

type Concert_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	HeadlinerName *string
}

func (b0 Concert_builder) Build() *Concert {
	m0 := &Concert{}
	b, x := &b0, m0
	_, _ = b, x
	if b.HeadlinerName != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_HeadlinerName = b.HeadlinerName
	}
	return m0
}

var file_internal_testprotos_examples_ext_extexample_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*Concert)(nil),
		ExtensionType: (*int32)(nil),
		Field:         123,
		Name:          "goproto.proto.test.promo_id",
		Tag:           "varint,123,opt,name=promo_id",
		Filename:      "internal/testprotos/examples/ext/extexample.proto",
	},
}

// Extension fields to Concert.
var (
	// optional int32 promo_id = 123;
	E_PromoId = &file_internal_testprotos_examples_ext_extexample_proto_extTypes[0]
)

var File_internal_testprotos_examples_ext_extexample_proto protoreflect.FileDescriptor

var file_internal_testprotos_examples_ext_extexample_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65,
	0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x63, 0x65, 0x72, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x05, 0x08, 0x64,
	0x10, 0xc8, 0x01, 0x3a, 0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x69, 0x64, 0x12,
	0x1b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x18, 0x7b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x49, 0x64, 0x42, 0x45, 0x5a, 0x3b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02,
	0x10, 0x03, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_internal_testprotos_examples_ext_extexample_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_examples_ext_extexample_proto_goTypes = []any{
	(*Concert)(nil), // 0: goproto.proto.test.Concert
}
var file_internal_testprotos_examples_ext_extexample_proto_depIdxs = []int32{
	0, // 0: goproto.proto.test.promo_id:extendee -> goproto.proto.test.Concert
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_examples_ext_extexample_proto_init() }
func file_internal_testprotos_examples_ext_extexample_proto_init() {
	if File_internal_testprotos_examples_ext_extexample_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_examples_ext_extexample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_examples_ext_extexample_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_examples_ext_extexample_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_examples_ext_extexample_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_examples_ext_extexample_proto_extTypes,
	}.Build()
	File_internal_testprotos_examples_ext_extexample_proto = out.File
	file_internal_testprotos_examples_ext_extexample_proto_rawDesc = nil
	file_internal_testprotos_examples_ext_extexample_proto_goTypes = nil
	file_internal_testprotos_examples_ext_extexample_proto_depIdxs = nil
}