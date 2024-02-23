// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: pingpong.proto

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

type Ball struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player     string           `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
	Properties []*Ball_Property `protobuf:"bytes,2,rep,name=Properties,proto3" json:"Properties,omitempty"`
	Raw        []byte           `protobuf:"bytes,3,opt,name=Raw,proto3" json:"Raw,omitempty"`
}

func (x *Ball) Reset() {
	*x = Ball{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ball) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ball) ProtoMessage() {}

func (x *Ball) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ball.ProtoReflect.Descriptor instead.
func (*Ball) Descriptor() ([]byte, []int) {
	return file_pingpong_proto_rawDescGZIP(), []int{0}
}

func (x *Ball) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *Ball) GetProperties() []*Ball_Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Ball) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

type Ball_Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Ball_Property) Reset() {
	*x = Ball_Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ball_Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ball_Property) ProtoMessage() {}

func (x *Ball_Property) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ball_Property.ProtoReflect.Descriptor instead.
func (*Ball_Property) Descriptor() ([]byte, []int) {
	return file_pingpong_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ball_Property) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Ball_Property) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pingpong_proto protoreflect.FileDescriptor

var file_pingpong_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x97, 0x01, 0x0a, 0x04, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x61, 0x6c, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x52, 0x61, 0x77, 0x1a, 0x32, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x28,
	0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x04, 0x50, 0x6c,
	0x61, 0x79, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x1a, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x34, 0x31, 0x37, 0x39, 0x37, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x6f, 0x70, 0x6f, 0x6e, 0x67, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pingpong_proto_rawDescOnce sync.Once
	file_pingpong_proto_rawDescData = file_pingpong_proto_rawDesc
)

func file_pingpong_proto_rawDescGZIP() []byte {
	file_pingpong_proto_rawDescOnce.Do(func() {
		file_pingpong_proto_rawDescData = protoimpl.X.CompressGZIP(file_pingpong_proto_rawDescData)
	})
	return file_pingpong_proto_rawDescData
}

var file_pingpong_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pingpong_proto_goTypes = []interface{}{
	(*Ball)(nil),          // 0: pb.Ball
	(*Ball_Property)(nil), // 1: pb.Ball.Property
}
var file_pingpong_proto_depIdxs = []int32{
	1, // 0: pb.Ball.Properties:type_name -> pb.Ball.Property
	0, // 1: pb.PingPong.Play:input_type -> pb.Ball
	0, // 2: pb.PingPong.Play:output_type -> pb.Ball
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pingpong_proto_init() }
func file_pingpong_proto_init() {
	if File_pingpong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pingpong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ball); i {
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
		file_pingpong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ball_Property); i {
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
			RawDescriptor: file_pingpong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pingpong_proto_goTypes,
		DependencyIndexes: file_pingpong_proto_depIdxs,
		MessageInfos:      file_pingpong_proto_msgTypes,
	}.Build()
	File_pingpong_proto = out.File
	file_pingpong_proto_rawDesc = nil
	file_pingpong_proto_goTypes = nil
	file_pingpong_proto_depIdxs = nil
}
