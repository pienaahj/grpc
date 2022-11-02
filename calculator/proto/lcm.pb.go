// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: lcm.proto

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

type LcmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *LcmRequest) Reset() {
	*x = LcmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LcmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LcmRequest) ProtoMessage() {}

func (x *LcmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lcm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LcmRequest.ProtoReflect.Descriptor instead.
func (*LcmRequest) Descriptor() ([]byte, []int) {
	return file_lcm_proto_rawDescGZIP(), []int{0}
}

func (x *LcmRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type LcmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LcmResponse) Reset() {
	*x = LcmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lcm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LcmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LcmResponse) ProtoMessage() {}

func (x *LcmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lcm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LcmResponse.ProtoReflect.Descriptor instead.
func (*LcmResponse) Descriptor() ([]byte, []int) {
	return file_lcm_proto_rawDescGZIP(), []int{1}
}

func (x *LcmResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_lcm_proto protoreflect.FileDescriptor

var file_lcm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x0a, 0x4c, 0x63, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x25, 0x0a,
	0x0b, 0x4c, 0x63, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x65, 0x6e, 0x61, 0x61, 0x68, 0x6a, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lcm_proto_rawDescOnce sync.Once
	file_lcm_proto_rawDescData = file_lcm_proto_rawDesc
)

func file_lcm_proto_rawDescGZIP() []byte {
	file_lcm_proto_rawDescOnce.Do(func() {
		file_lcm_proto_rawDescData = protoimpl.X.CompressGZIP(file_lcm_proto_rawDescData)
	})
	return file_lcm_proto_rawDescData
}

var file_lcm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lcm_proto_goTypes = []interface{}{
	(*LcmRequest)(nil),  // 0: calculator.LcmRequest
	(*LcmResponse)(nil), // 1: calculator.LcmResponse
}
var file_lcm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lcm_proto_init() }
func file_lcm_proto_init() {
	if File_lcm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lcm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LcmRequest); i {
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
		file_lcm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LcmResponse); i {
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
			RawDescriptor: file_lcm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lcm_proto_goTypes,
		DependencyIndexes: file_lcm_proto_depIdxs,
		MessageInfos:      file_lcm_proto_msgTypes,
	}.Build()
	File_lcm_proto = out.File
	file_lcm_proto_rawDesc = nil
	file_lcm_proto_goTypes = nil
	file_lcm_proto_depIdxs = nil
}