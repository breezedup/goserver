// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: sessionauth.proto

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

type SSPacketAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthKey   string `protobuf:"bytes,1,opt,name=AuthKey,proto3" json:"AuthKey,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *SSPacketAuth) Reset() {
	*x = SSPacketAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSPacketAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSPacketAuth) ProtoMessage() {}

func (x *SSPacketAuth) ProtoReflect() protoreflect.Message {
	mi := &file_sessionauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSPacketAuth.ProtoReflect.Descriptor instead.
func (*SSPacketAuth) Descriptor() ([]byte, []int) {
	return file_sessionauth_proto_rawDescGZIP(), []int{0}
}

func (x *SSPacketAuth) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *SSPacketAuth) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SSPacketAuthAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *SSPacketAuthAck) Reset() {
	*x = SSPacketAuthAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sessionauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSPacketAuthAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSPacketAuthAck) ProtoMessage() {}

func (x *SSPacketAuthAck) ProtoReflect() protoreflect.Message {
	mi := &file_sessionauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSPacketAuthAck.ProtoReflect.Descriptor instead.
func (*SSPacketAuthAck) Descriptor() ([]byte, []int) {
	return file_sessionauth_proto_rawDescGZIP(), []int{1}
}

func (x *SSPacketAuthAck) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sessionauth_proto protoreflect.FileDescriptor

var file_sessionauth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x46, 0x0a,
	0x0c, 0x53, 0x53, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x18, 0x0a,
	0x07, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x23, 0x0a, 0x0f, 0x53, 0x53, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sessionauth_proto_rawDescOnce sync.Once
	file_sessionauth_proto_rawDescData = file_sessionauth_proto_rawDesc
)

func file_sessionauth_proto_rawDescGZIP() []byte {
	file_sessionauth_proto_rawDescOnce.Do(func() {
		file_sessionauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_sessionauth_proto_rawDescData)
	})
	return file_sessionauth_proto_rawDescData
}

var file_sessionauth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sessionauth_proto_goTypes = []interface{}{
	(*SSPacketAuth)(nil),    // 0: protocol.SSPacketAuth
	(*SSPacketAuthAck)(nil), // 1: protocol.SSPacketAuthAck
}
var file_sessionauth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sessionauth_proto_init() }
func file_sessionauth_proto_init() {
	if File_sessionauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sessionauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSPacketAuth); i {
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
		file_sessionauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSPacketAuthAck); i {
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
			RawDescriptor: file_sessionauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sessionauth_proto_goTypes,
		DependencyIndexes: file_sessionauth_proto_depIdxs,
		MessageInfos:      file_sessionauth_proto_msgTypes,
	}.Build()
	File_sessionauth_proto = out.File
	file_sessionauth_proto_rawDesc = nil
	file_sessionauth_proto_goTypes = nil
	file_sessionauth_proto_depIdxs = nil
}