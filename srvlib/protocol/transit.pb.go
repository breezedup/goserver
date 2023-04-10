// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: transit.proto

package protocol

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

type SSPacketTransit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SArea    int32  `protobuf:"varint,1,opt,name=SArea,proto3" json:"SArea,omitempty"`
	SType    int32  `protobuf:"varint,2,opt,name=SType,proto3" json:"SType,omitempty"`
	SId      int32  `protobuf:"varint,3,opt,name=SId,proto3" json:"SId,omitempty"`
	PacketId int32  `protobuf:"varint,4,opt,name=PacketId,proto3" json:"PacketId,omitempty"`
	Data     []byte `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *SSPacketTransit) Reset() {
	*x = SSPacketTransit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSPacketTransit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSPacketTransit) ProtoMessage() {}

func (x *SSPacketTransit) ProtoReflect() protoreflect.Message {
	mi := &file_transit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSPacketTransit.ProtoReflect.Descriptor instead.
func (*SSPacketTransit) Descriptor() ([]byte, []int) {
	return file_transit_proto_rawDescGZIP(), []int{0}
}

func (x *SSPacketTransit) GetSArea() int32 {
	if x != nil {
		return x.SArea
	}
	return 0
}

func (x *SSPacketTransit) GetSType() int32 {
	if x != nil {
		return x.SType
	}
	return 0
}

func (x *SSPacketTransit) GetSId() int32 {
	if x != nil {
		return x.SId
	}
	return 0
}

func (x *SSPacketTransit) GetPacketId() int32 {
	if x != nil {
		return x.PacketId
	}
	return 0
}

func (x *SSPacketTransit) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_transit_proto protoreflect.FileDescriptor

var file_transit_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x7f, 0x0a, 0x0f, 0x53, 0x53, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x41, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x41, 0x72,
	0x65, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x53, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transit_proto_rawDescOnce sync.Once
	file_transit_proto_rawDescData = file_transit_proto_rawDesc
)

func file_transit_proto_rawDescGZIP() []byte {
	file_transit_proto_rawDescOnce.Do(func() {
		file_transit_proto_rawDescData = protoimpl.X.CompressGZIP(file_transit_proto_rawDescData)
	})
	return file_transit_proto_rawDescData
}

var file_transit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transit_proto_goTypes = []interface{}{
	(*SSPacketTransit)(nil), // 0: protocol.SSPacketTransit
}
var file_transit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transit_proto_init() }
func file_transit_proto_init() {
	if File_transit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSPacketTransit); i {
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
			RawDescriptor: file_transit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transit_proto_goTypes,
		DependencyIndexes: file_transit_proto_depIdxs,
		MessageInfos:      file_transit_proto_msgTypes,
	}.Build()
	File_transit_proto = out.File
	file_transit_proto_rawDesc = nil
	file_transit_proto_goTypes = nil
	file_transit_proto_depIdxs = nil
}