// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: encounter.proto

package encounter

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

type SocialEnocunter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncounterId                   string  `protobuf:"bytes,1,opt,name=encounterId,proto3" json:"encounterId,omitempty"`
	Name                          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description                   string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XpPoints                      int32   `protobuf:"varint,4,opt,name=xpPoints,proto3" json:"xpPoints,omitempty"`
	Status                        string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Type                          string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Latitude                      float64 `protobuf:"fixed64,7,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude                     float64 `protobuf:"fixed64,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Id                            string  `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
	TouristsRequiredForCompletion int32   `protobuf:"varint,10,opt,name=touristsRequiredForCompletion,proto3" json:"touristsRequiredForCompletion,omitempty"`
	DistanceTreshold              float64 `protobuf:"fixed64,11,opt,name=distanceTreshold,proto3" json:"distanceTreshold,omitempty"`
	TouristIDs                    []int64 `protobuf:"varint,12,rep,packed,name=touristIDs,proto3" json:"touristIDs,omitempty"`
	ShouldBeApproved              bool    `protobuf:"varint,13,opt,name=shouldBeApproved,proto3" json:"shouldBeApproved,omitempty"`
}

func (x *SocialEnocunter) Reset() {
	*x = SocialEnocunter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialEnocunter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialEnocunter) ProtoMessage() {}

func (x *SocialEnocunter) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialEnocunter.ProtoReflect.Descriptor instead.
func (*SocialEnocunter) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{0}
}

func (x *SocialEnocunter) GetEncounterId() string {
	if x != nil {
		return x.EncounterId
	}
	return ""
}

func (x *SocialEnocunter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SocialEnocunter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SocialEnocunter) GetXpPoints() int32 {
	if x != nil {
		return x.XpPoints
	}
	return 0
}

func (x *SocialEnocunter) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SocialEnocunter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SocialEnocunter) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *SocialEnocunter) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *SocialEnocunter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SocialEnocunter) GetTouristsRequiredForCompletion() int32 {
	if x != nil {
		return x.TouristsRequiredForCompletion
	}
	return 0
}

func (x *SocialEnocunter) GetDistanceTreshold() float64 {
	if x != nil {
		return x.DistanceTreshold
	}
	return 0
}

func (x *SocialEnocunter) GetTouristIDs() []int64 {
	if x != nil {
		return x.TouristIDs
	}
	return nil
}

func (x *SocialEnocunter) GetShouldBeApproved() bool {
	if x != nil {
		return x.ShouldBeApproved
	}
	return false
}

var File_encounter_proto protoreflect.FileDescriptor

var file_encounter_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb9, 0x03, 0x0a, 0x0f, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x6f, 0x63,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x78, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x78, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x44, 0x0a, 0x1d, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1d, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x10, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x49, 0x44, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x49, 0x44,
	0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x42, 0x65, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x68, 0x6f,
	0x75, 0x6c, 0x64, 0x42, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x32, 0x4a, 0x0a,
	0x09, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x6f, 0x63,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x45, 0x6e,
	0x6f, 0x63, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encounter_proto_rawDescOnce sync.Once
	file_encounter_proto_rawDescData = file_encounter_proto_rawDesc
)

func file_encounter_proto_rawDescGZIP() []byte {
	file_encounter_proto_rawDescOnce.Do(func() {
		file_encounter_proto_rawDescData = protoimpl.X.CompressGZIP(file_encounter_proto_rawDescData)
	})
	return file_encounter_proto_rawDescData
}

var file_encounter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_encounter_proto_goTypes = []interface{}{
	(*SocialEnocunter)(nil), // 0: SocialEnocunter
}
var file_encounter_proto_depIdxs = []int32{
	0, // 0: Encounter.CreateSocialEncounter:input_type -> SocialEnocunter
	0, // 1: Encounter.CreateSocialEncounter:output_type -> SocialEnocunter
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_encounter_proto_init() }
func file_encounter_proto_init() {
	if File_encounter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encounter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialEnocunter); i {
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
			RawDescriptor: file_encounter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_encounter_proto_goTypes,
		DependencyIndexes: file_encounter_proto_depIdxs,
		MessageInfos:      file_encounter_proto_msgTypes,
	}.Build()
	File_encounter_proto = out.File
	file_encounter_proto_rawDesc = nil
	file_encounter_proto_goTypes = nil
	file_encounter_proto_depIdxs = nil
}
