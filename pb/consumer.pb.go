// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: consumer.proto

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

type Consumer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Consumer) Reset() {
	*x = Consumer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consumer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consumer) ProtoMessage() {}

func (x *Consumer) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consumer.ProtoReflect.Descriptor instead.
func (*Consumer) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *Consumer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ConsumerOpions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConsumerOpions) Reset() {
	*x = ConsumerOpions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerOpions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerOpions) ProtoMessage() {}

func (x *ConsumerOpions) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerOpions.ProtoReflect.Descriptor instead.
func (*ConsumerOpions) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{1}
}

type ConsumingOpions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConsumingOpions) Reset() {
	*x = ConsumingOpions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumingOpions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumingOpions) ProtoMessage() {}

func (x *ConsumingOpions) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumingOpions.ProtoReflect.Descriptor instead.
func (*ConsumingOpions) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{2}
}

type CreateConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station  *Station        `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
	Consumer *Consumer       `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Options  *ConsumerOpions `protobuf:"bytes,3,opt,name=options,proto3,oneof" json:"options,omitempty"`
}

func (x *CreateConsumerRequest) Reset() {
	*x = CreateConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsumerRequest) ProtoMessage() {}

func (x *CreateConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsumerRequest.ProtoReflect.Descriptor instead.
func (*CreateConsumerRequest) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{3}
}

func (x *CreateConsumerRequest) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

func (x *CreateConsumerRequest) GetConsumer() *Consumer {
	if x != nil {
		return x.Consumer
	}
	return nil
}

func (x *CreateConsumerRequest) GetOptions() *ConsumerOpions {
	if x != nil {
		return x.Options
	}
	return nil
}

type DestroyConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station  *Station  `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
	Consumer *Consumer `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
}

func (x *DestroyConsumerRequest) Reset() {
	*x = DestroyConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyConsumerRequest) ProtoMessage() {}

func (x *DestroyConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyConsumerRequest.ProtoReflect.Descriptor instead.
func (*DestroyConsumerRequest) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{4}
}

func (x *DestroyConsumerRequest) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

func (x *DestroyConsumerRequest) GetConsumer() *Consumer {
	if x != nil {
		return x.Consumer
	}
	return nil
}

var File_consumer_proto protoreflect.FileDescriptor

var file_consumer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4f,
	0x70, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x4f, 0x70, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x69, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x69, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x34, 0x31, 0x37,
	0x39, 0x37, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consumer_proto_rawDescOnce sync.Once
	file_consumer_proto_rawDescData = file_consumer_proto_rawDesc
)

func file_consumer_proto_rawDescGZIP() []byte {
	file_consumer_proto_rawDescOnce.Do(func() {
		file_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_consumer_proto_rawDescData)
	})
	return file_consumer_proto_rawDescData
}

var file_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_consumer_proto_goTypes = []interface{}{
	(*Consumer)(nil),               // 0: pb.Consumer
	(*ConsumerOpions)(nil),         // 1: pb.ConsumerOpions
	(*ConsumingOpions)(nil),        // 2: pb.ConsumingOpions
	(*CreateConsumerRequest)(nil),  // 3: pb.CreateConsumerRequest
	(*DestroyConsumerRequest)(nil), // 4: pb.DestroyConsumerRequest
	(*Station)(nil),                // 5: pb.Station
}
var file_consumer_proto_depIdxs = []int32{
	5, // 0: pb.CreateConsumerRequest.station:type_name -> pb.Station
	0, // 1: pb.CreateConsumerRequest.consumer:type_name -> pb.Consumer
	1, // 2: pb.CreateConsumerRequest.options:type_name -> pb.ConsumerOpions
	5, // 3: pb.DestroyConsumerRequest.station:type_name -> pb.Station
	0, // 4: pb.DestroyConsumerRequest.consumer:type_name -> pb.Consumer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_consumer_proto_init() }
func file_consumer_proto_init() {
	if File_consumer_proto != nil {
		return
	}
	file_station_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consumer); i {
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
		file_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerOpions); i {
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
		file_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumingOpions); i {
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
		file_consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsumerRequest); i {
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
		file_consumer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyConsumerRequest); i {
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
	file_consumer_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consumer_proto_goTypes,
		DependencyIndexes: file_consumer_proto_depIdxs,
		MessageInfos:      file_consumer_proto_msgTypes,
	}.Build()
	File_consumer_proto = out.File
	file_consumer_proto_rawDesc = nil
	file_consumer_proto_goTypes = nil
	file_consumer_proto_depIdxs = nil
}
