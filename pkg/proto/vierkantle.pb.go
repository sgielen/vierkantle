// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: vierkantle.proto

package proto

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M string `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetM() string {
	if x != nil {
		return x.M
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M string `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetM() string {
	if x != nil {
		return x.M
	}
	return ""
}

type HelloStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M string `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *HelloStreamRequest) Reset() {
	*x = HelloStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloStreamRequest) ProtoMessage() {}

func (x *HelloStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloStreamRequest.ProtoReflect.Descriptor instead.
func (*HelloStreamRequest) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{2}
}

func (x *HelloStreamRequest) GetM() string {
	if x != nil {
		return x.M
	}
	return ""
}

type HelloStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M string `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *HelloStreamResponse) Reset() {
	*x = HelloStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloStreamResponse) ProtoMessage() {}

func (x *HelloStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloStreamResponse.ProtoReflect.Descriptor instead.
func (*HelloStreamResponse) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{3}
}

func (x *HelloStreamResponse) GetM() string {
	if x != nil {
		return x.M
	}
	return ""
}

var File_vierkantle_proto protoreflect.FileDescriptor

var file_vierkantle_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c,
	0x65, 0x22, 0x1c, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6d, 0x22,
	0x1d, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6d, 0x22, 0x22,
	0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x6d, 0x22, 0x23, 0x0a, 0x13, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6d, 0x32, 0xb5, 0x01, 0x0a, 0x11, 0x56, 0x69, 0x65, 0x72,
	0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1b, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72,
	0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e,
	0x74, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x21, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74,
	0x6c, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b,
	0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x67,
	0x69, 0x65, 0x6c, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vierkantle_proto_rawDescOnce sync.Once
	file_vierkantle_proto_rawDescData = file_vierkantle_proto_rawDesc
)

func file_vierkantle_proto_rawDescGZIP() []byte {
	file_vierkantle_proto_rawDescOnce.Do(func() {
		file_vierkantle_proto_rawDescData = protoimpl.X.CompressGZIP(file_vierkantle_proto_rawDescData)
	})
	return file_vierkantle_proto_rawDescData
}

var file_vierkantle_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vierkantle_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),        // 0: nl.vierkantle.HelloRequest
	(*HelloResponse)(nil),       // 1: nl.vierkantle.HelloResponse
	(*HelloStreamRequest)(nil),  // 2: nl.vierkantle.HelloStreamRequest
	(*HelloStreamResponse)(nil), // 3: nl.vierkantle.HelloStreamResponse
}
var file_vierkantle_proto_depIdxs = []int32{
	0, // 0: nl.vierkantle.VierkantleService.Hello:input_type -> nl.vierkantle.HelloRequest
	2, // 1: nl.vierkantle.VierkantleService.HelloStream:input_type -> nl.vierkantle.HelloStreamRequest
	1, // 2: nl.vierkantle.VierkantleService.Hello:output_type -> nl.vierkantle.HelloResponse
	3, // 3: nl.vierkantle.VierkantleService.HelloStream:output_type -> nl.vierkantle.HelloStreamResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vierkantle_proto_init() }
func file_vierkantle_proto_init() {
	if File_vierkantle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vierkantle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_vierkantle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_vierkantle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloStreamRequest); i {
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
		file_vierkantle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloStreamResponse); i {
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
			RawDescriptor: file_vierkantle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vierkantle_proto_goTypes,
		DependencyIndexes: file_vierkantle_proto_depIdxs,
		MessageInfos:      file_vierkantle_proto_msgTypes,
	}.Build()
	File_vierkantle_proto = out.File
	file_vierkantle_proto_rawDesc = nil
	file_vierkantle_proto_goTypes = nil
	file_vierkantle_proto_depIdxs = nil
}