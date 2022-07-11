// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: enumerations.proto

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

type EyeColor int32

const (
	EyeColor_EYE_COLOR_UNSPECIFIED EyeColor = 0
	EyeColor_EYE_COLOR_GREEN       EyeColor = 1
	EyeColor_EYE_COLOR_BLUE        EyeColor = 2
	EyeColor_EYE_COLOR_BROWN       EyeColor = 3
)

// Enum value maps for EyeColor.
var (
	EyeColor_name = map[int32]string{
		0: "EYE_COLOR_UNSPECIFIED",
		1: "EYE_COLOR_GREEN",
		2: "EYE_COLOR_BLUE",
		3: "EYE_COLOR_BROWN",
	}
	EyeColor_value = map[string]int32{
		"EYE_COLOR_UNSPECIFIED": 0,
		"EYE_COLOR_GREEN":       1,
		"EYE_COLOR_BLUE":        2,
		"EYE_COLOR_BROWN":       3,
	}
)

func (x EyeColor) Enum() *EyeColor {
	p := new(EyeColor)
	*p = x
	return p
}

func (x EyeColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EyeColor) Descriptor() protoreflect.EnumDescriptor {
	return file_enumerations_proto_enumTypes[0].Descriptor()
}

func (EyeColor) Type() protoreflect.EnumType {
	return &file_enumerations_proto_enumTypes[0]
}

func (x EyeColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EyeColor.Descriptor instead.
func (EyeColor) EnumDescriptor() ([]byte, []int) {
	return file_enumerations_proto_rawDescGZIP(), []int{0}
}

type Enumeration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EyeColor EyeColor `protobuf:"varint,1,opt,name=eye_color,json=eyeColor,proto3,enum=example.enumerations.EyeColor" json:"eye_color,omitempty"`
}

func (x *Enumeration) Reset() {
	*x = Enumeration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enumerations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enumeration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enumeration) ProtoMessage() {}

func (x *Enumeration) ProtoReflect() protoreflect.Message {
	mi := &file_enumerations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enumeration.ProtoReflect.Descriptor instead.
func (*Enumeration) Descriptor() ([]byte, []int) {
	return file_enumerations_proto_rawDescGZIP(), []int{0}
}

func (x *Enumeration) GetEyeColor() EyeColor {
	if x != nil {
		return x.EyeColor
	}
	return EyeColor_EYE_COLOR_UNSPECIFIED
}

var File_enumerations_proto protoreflect.FileDescriptor

var file_enumerations_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4a, 0x0a, 0x0b, 0x45, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x09, 0x65, 0x79, 0x65,
	0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x65, 0x79,
	0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x2a, 0x63, 0x0a, 0x08, 0x45, 0x79, 0x65, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x59, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x59, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x47, 0x52, 0x45, 0x45, 0x4e,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x59, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f,
	0x42, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x59, 0x45, 0x5f, 0x43, 0x4f,
	0x4c, 0x4f, 0x52, 0x5f, 0x42, 0x52, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x78, 0x70, 0x6f, 0x73,
	0x61, 0x64, 0x61, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enumerations_proto_rawDescOnce sync.Once
	file_enumerations_proto_rawDescData = file_enumerations_proto_rawDesc
)

func file_enumerations_proto_rawDescGZIP() []byte {
	file_enumerations_proto_rawDescOnce.Do(func() {
		file_enumerations_proto_rawDescData = protoimpl.X.CompressGZIP(file_enumerations_proto_rawDescData)
	})
	return file_enumerations_proto_rawDescData
}

var file_enumerations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enumerations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enumerations_proto_goTypes = []interface{}{
	(EyeColor)(0),       // 0: example.enumerations.EyeColor
	(*Enumeration)(nil), // 1: example.enumerations.Enumeration
}
var file_enumerations_proto_depIdxs = []int32{
	0, // 0: example.enumerations.Enumeration.eye_color:type_name -> example.enumerations.EyeColor
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_enumerations_proto_init() }
func file_enumerations_proto_init() {
	if File_enumerations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enumerations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enumeration); i {
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
			RawDescriptor: file_enumerations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enumerations_proto_goTypes,
		DependencyIndexes: file_enumerations_proto_depIdxs,
		EnumInfos:         file_enumerations_proto_enumTypes,
		MessageInfos:      file_enumerations_proto_msgTypes,
	}.Build()
	File_enumerations_proto = out.File
	file_enumerations_proto_rawDesc = nil
	file_enumerations_proto_goTypes = nil
	file_enumerations_proto_depIdxs = nil
}
