// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: adapterservice.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_adapterservice_proto protoreflect.FileDescriptor

var file_adapterservice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xef, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x39,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x48, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x34, 0x31, 0x37, 0x39, 0x37, 0x2f, 0x6d,
	0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x62, 0xca, 0x02, 0x0d,
	0x4d, 0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x70, 0x68, 0x70, 0x5c, 0x50, 0x62, 0xe2, 0x02, 0x13,
	0x4d, 0x65, 0x6d, 0x70, 0x68, 0x69, 0x73, 0x70, 0x68, 0x70, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_adapterservice_proto_goTypes = []interface{}{
	(*CreateStationRequest)(nil),  // 0: pb.CreateStationRequest
	(*DestroyStationRequest)(nil), // 1: pb.DestroyStationRequest
	(*ProduceMessages)(nil),       // 2: pb.ProduceMessages
	(*ConsumeMessages)(nil),       // 3: pb.ConsumeMessages
	(*Status)(nil),                // 4: pb.Status
	(*ConsumeResponse)(nil),       // 5: pb.ConsumeResponse
}
var file_adapterservice_proto_depIdxs = []int32{
	0, // 0: pb.AdapterService.CreateStation:input_type -> pb.CreateStationRequest
	1, // 1: pb.AdapterService.DestroyStation:input_type -> pb.DestroyStationRequest
	2, // 2: pb.AdapterService.Produce:input_type -> pb.ProduceMessages
	3, // 3: pb.AdapterService.Consume:input_type -> pb.ConsumeMessages
	4, // 4: pb.AdapterService.CreateStation:output_type -> pb.Status
	4, // 5: pb.AdapterService.DestroyStation:output_type -> pb.Status
	4, // 6: pb.AdapterService.Produce:output_type -> pb.Status
	5, // 7: pb.AdapterService.Consume:output_type -> pb.ConsumeResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adapterservice_proto_init() }
func file_adapterservice_proto_init() {
	if File_adapterservice_proto != nil {
		return
	}
	file_status_proto_init()
	file_station_proto_init()
	file_producer_proto_init()
	file_consumer_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adapterservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adapterservice_proto_goTypes,
		DependencyIndexes: file_adapterservice_proto_depIdxs,
	}.Build()
	File_adapterservice_proto = out.File
	file_adapterservice_proto_rawDesc = nil
	file_adapterservice_proto_goTypes = nil
	file_adapterservice_proto_depIdxs = nil
}
