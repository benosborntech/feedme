// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: updatesservice.proto

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

type GetUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongX  float32 `protobuf:"fixed32,1,opt,name=longX,proto3" json:"longX,omitempty"`
	LatY   float32 `protobuf:"fixed32,2,opt,name=latY,proto3" json:"latY,omitempty"`
	Radius float32 `protobuf:"fixed32,3,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *GetUpdatesRequest) Reset() {
	*x = GetUpdatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatesservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatesRequest) ProtoMessage() {}

func (x *GetUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_updatesservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatesRequest.ProtoReflect.Descriptor instead.
func (*GetUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_updatesservice_proto_rawDescGZIP(), []int{0}
}

func (x *GetUpdatesRequest) GetLongX() float32 {
	if x != nil {
		return x.LongX
	}
	return 0
}

func (x *GetUpdatesRequest) GetLatY() float32 {
	if x != nil {
		return x.LatY
	}
	return 0
}

func (x *GetUpdatesRequest) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type GetUpdatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *ItemData `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetUpdatesResponse) Reset() {
	*x = GetUpdatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updatesservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatesResponse) ProtoMessage() {}

func (x *GetUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_updatesservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatesResponse.ProtoReflect.Descriptor instead.
func (*GetUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_updatesservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetUpdatesResponse) GetItem() *ItemData {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_updatesservice_proto protoreflect.FileDescriptor

var file_updatesservice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x55, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x6e, 0x67, 0x58, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x6f, 0x6e, 0x67, 0x58, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x61, 0x74, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x61, 0x74, 0x59,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x32, 0x5e, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x53, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_updatesservice_proto_rawDescOnce sync.Once
	file_updatesservice_proto_rawDescData = file_updatesservice_proto_rawDesc
)

func file_updatesservice_proto_rawDescGZIP() []byte {
	file_updatesservice_proto_rawDescOnce.Do(func() {
		file_updatesservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_updatesservice_proto_rawDescData)
	})
	return file_updatesservice_proto_rawDescData
}

var file_updatesservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_updatesservice_proto_goTypes = []any{
	(*GetUpdatesRequest)(nil),  // 0: updatesservie.GetUpdatesRequest
	(*GetUpdatesResponse)(nil), // 1: updatesservie.GetUpdatesResponse
	(*ItemData)(nil),           // 2: item.ItemData
}
var file_updatesservice_proto_depIdxs = []int32{
	2, // 0: updatesservie.GetUpdatesResponse.item:type_name -> item.ItemData
	0, // 1: updatesservie.Updates.GetUpdates:input_type -> updatesservie.GetUpdatesRequest
	1, // 2: updatesservie.Updates.GetUpdates:output_type -> updatesservie.GetUpdatesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_updatesservice_proto_init() }
func file_updatesservice_proto_init() {
	if File_updatesservice_proto != nil {
		return
	}
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_updatesservice_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUpdatesRequest); i {
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
		file_updatesservice_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUpdatesResponse); i {
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
			RawDescriptor: file_updatesservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_updatesservice_proto_goTypes,
		DependencyIndexes: file_updatesservice_proto_depIdxs,
		MessageInfos:      file_updatesservice_proto_msgTypes,
	}.Build()
	File_updatesservice_proto = out.File
	file_updatesservice_proto_rawDesc = nil
	file_updatesservice_proto_goTypes = nil
	file_updatesservice_proto_depIdxs = nil
}
