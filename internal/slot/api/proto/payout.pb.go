// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: base/payout.proto

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

type Payout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol uint32    `protobuf:"varint,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Payout []float32 `protobuf:"fixed32,2,rep,packed,name=payout,proto3" json:"payout,omitempty"`
}

func (x *Payout) Reset() {
	*x = Payout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_payout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payout) ProtoMessage() {}

func (x *Payout) ProtoReflect() protoreflect.Message {
	mi := &file_base_payout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payout.ProtoReflect.Descriptor instead.
func (*Payout) Descriptor() ([]byte, []int) {
	return file_base_payout_proto_rawDescGZIP(), []int{0}
}

func (x *Payout) GetSymbol() uint32 {
	if x != nil {
		return x.Symbol
	}
	return 0
}

func (x *Payout) GetPayout() []float32 {
	if x != nil {
		return x.Payout
	}
	return nil
}

var File_base_payout_proto protoreflect.FileDescriptor

var file_base_payout_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x38, 0x0a, 0x06, 0x50, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x70, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_payout_proto_rawDescOnce sync.Once
	file_base_payout_proto_rawDescData = file_base_payout_proto_rawDesc
)

func file_base_payout_proto_rawDescGZIP() []byte {
	file_base_payout_proto_rawDescOnce.Do(func() {
		file_base_payout_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_payout_proto_rawDescData)
	})
	return file_base_payout_proto_rawDescData
}

var file_base_payout_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_base_payout_proto_goTypes = []interface{}{
	(*Payout)(nil), // 0: slot.Payout
}
var file_base_payout_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_payout_proto_init() }
func file_base_payout_proto_init() {
	if File_base_payout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_payout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payout); i {
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
			RawDescriptor: file_base_payout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_payout_proto_goTypes,
		DependencyIndexes: file_base_payout_proto_depIdxs,
		MessageInfos:      file_base_payout_proto_msgTypes,
	}.Build()
	File_base_payout_proto = out.File
	file_base_payout_proto_rawDesc = nil
	file_base_payout_proto_goTypes = nil
	file_base_payout_proto_depIdxs = nil
}
