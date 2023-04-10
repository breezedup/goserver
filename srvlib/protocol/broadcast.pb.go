// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: broadcast.proto

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

type SSPacketBroadcast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessParam *BCSessionUnion `protobuf:"bytes,1,opt,name=SessParam,proto3" json:"SessParam,omitempty"`
	PacketId  int32           `protobuf:"varint,2,opt,name=PacketId,proto3" json:"PacketId,omitempty"`
	Data      []byte          `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *SSPacketBroadcast) Reset() {
	*x = SSPacketBroadcast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSPacketBroadcast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSPacketBroadcast) ProtoMessage() {}

func (x *SSPacketBroadcast) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSPacketBroadcast.ProtoReflect.Descriptor instead.
func (*SSPacketBroadcast) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{0}
}

func (x *SSPacketBroadcast) GetSessParam() *BCSessionUnion {
	if x != nil {
		return x.SessParam
	}
	return nil
}

func (x *SSPacketBroadcast) GetPacketId() int32 {
	if x != nil {
		return x.PacketId
	}
	return 0
}

func (x *SSPacketBroadcast) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BCSessionUnion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bccs *BCClientSession `protobuf:"bytes,1,opt,name=Bccs,proto3" json:"Bccs,omitempty"`
	Bcss *BCServerSession `protobuf:"bytes,2,opt,name=Bcss,proto3" json:"Bcss,omitempty"`
}

func (x *BCSessionUnion) Reset() {
	*x = BCSessionUnion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BCSessionUnion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BCSessionUnion) ProtoMessage() {}

func (x *BCSessionUnion) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BCSessionUnion.ProtoReflect.Descriptor instead.
func (*BCSessionUnion) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{1}
}

func (x *BCSessionUnion) GetBccs() *BCClientSession {
	if x != nil {
		return x.Bccs
	}
	return nil
}

func (x *BCSessionUnion) GetBcss() *BCServerSession {
	if x != nil {
		return x.Bcss
	}
	return nil
}

type BCClientSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
}

func (x *BCClientSession) Reset() {
	*x = BCClientSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BCClientSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BCClientSession) ProtoMessage() {}

func (x *BCClientSession) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BCClientSession.ProtoReflect.Descriptor instead.
func (*BCClientSession) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{2}
}

func (x *BCClientSession) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

type BCServerSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SArea int32 `protobuf:"varint,1,opt,name=SArea,proto3" json:"SArea,omitempty"`
	SType int32 `protobuf:"varint,2,opt,name=SType,proto3" json:"SType,omitempty"`
}

func (x *BCServerSession) Reset() {
	*x = BCServerSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broadcast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BCServerSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BCServerSession) ProtoMessage() {}

func (x *BCServerSession) ProtoReflect() protoreflect.Message {
	mi := &file_broadcast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BCServerSession.ProtoReflect.Descriptor instead.
func (*BCServerSession) Descriptor() ([]byte, []int) {
	return file_broadcast_proto_rawDescGZIP(), []int{3}
}

func (x *BCServerSession) GetSArea() int32 {
	if x != nil {
		return x.SArea
	}
	return 0
}

func (x *BCServerSession) GetSType() int32 {
	if x != nil {
		return x.SType
	}
	return 0
}

var File_broadcast_proto protoreflect.FileDescriptor

var file_broadcast_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x7b, 0x0a, 0x11, 0x53,
	0x53, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x42,
	0x43, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x53,
	0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x6e, 0x0a, 0x0e, 0x42, 0x43, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x04, 0x42, 0x63,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x42, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x04, 0x42, 0x63, 0x63, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x42, 0x63, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x42, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x42, 0x63, 0x73, 0x73, 0x22, 0x27, 0x0a, 0x0f, 0x42, 0x43, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x22, 0x3d, 0x0a, 0x0f, 0x42, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x41, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x41, 0x72, 0x65, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_broadcast_proto_rawDescOnce sync.Once
	file_broadcast_proto_rawDescData = file_broadcast_proto_rawDesc
)

func file_broadcast_proto_rawDescGZIP() []byte {
	file_broadcast_proto_rawDescOnce.Do(func() {
		file_broadcast_proto_rawDescData = protoimpl.X.CompressGZIP(file_broadcast_proto_rawDescData)
	})
	return file_broadcast_proto_rawDescData
}

var file_broadcast_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_broadcast_proto_goTypes = []interface{}{
	(*SSPacketBroadcast)(nil), // 0: protocol.SSPacketBroadcast
	(*BCSessionUnion)(nil),    // 1: protocol.BCSessionUnion
	(*BCClientSession)(nil),   // 2: protocol.BCClientSession
	(*BCServerSession)(nil),   // 3: protocol.BCServerSession
}
var file_broadcast_proto_depIdxs = []int32{
	1, // 0: protocol.SSPacketBroadcast.SessParam:type_name -> protocol.BCSessionUnion
	2, // 1: protocol.BCSessionUnion.Bccs:type_name -> protocol.BCClientSession
	3, // 2: protocol.BCSessionUnion.Bcss:type_name -> protocol.BCServerSession
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_broadcast_proto_init() }
func file_broadcast_proto_init() {
	if File_broadcast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broadcast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSPacketBroadcast); i {
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
		file_broadcast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BCSessionUnion); i {
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
		file_broadcast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BCClientSession); i {
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
		file_broadcast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BCServerSession); i {
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
			RawDescriptor: file_broadcast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_broadcast_proto_goTypes,
		DependencyIndexes: file_broadcast_proto_depIdxs,
		MessageInfos:      file_broadcast_proto_msgTypes,
	}.Build()
	File_broadcast_proto = out.File
	file_broadcast_proto_rawDesc = nil
	file_broadcast_proto_goTypes = nil
	file_broadcast_proto_depIdxs = nil
}
