// HomeDashboard Indoor Climate Event Schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: GenerateReportRequest.proto

package core

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

// ReportType defines which kind of report should be generated.
type ReportType int32

const (
	ReportType_NO_TYPE        ReportType = 0
	ReportType_MONTHLY_REPORT ReportType = 1
)

// Enum value maps for ReportType.
var (
	ReportType_name = map[int32]string{
		0: "NO_TYPE",
		1: "MONTHLY_REPORT",
	}
	ReportType_value = map[string]int32{
		"NO_TYPE":        0,
		"MONTHLY_REPORT": 1,
	}
)

func (x ReportType) Enum() *ReportType {
	p := new(ReportType)
	*p = x
	return p
}

func (x ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_GenerateReportRequest_proto_enumTypes[0].Descriptor()
}

func (ReportType) Type() protoreflect.EnumType {
	return &file_GenerateReportRequest_proto_enumTypes[0]
}

func (x ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportType.Descriptor instead.
func (ReportType) EnumDescriptor() ([]byte, []int) {
	return file_GenerateReportRequest_proto_rawDescGZIP(), []int{0}
}

// ReportFormat defines the format a report should be generated.
type ReportFormat int32

const (
	ReportFormat_NO_FORMAT ReportFormat = 0
	ReportFormat_EXCEL     ReportFormat = 1
)

// Enum value maps for ReportFormat.
var (
	ReportFormat_name = map[int32]string{
		0: "NO_FORMAT",
		1: "EXCEL",
	}
	ReportFormat_value = map[string]int32{
		"NO_FORMAT": 0,
		"EXCEL":     1,
	}
)

func (x ReportFormat) Enum() *ReportFormat {
	p := new(ReportFormat)
	*p = x
	return p
}

func (x ReportFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_GenerateReportRequest_proto_enumTypes[1].Descriptor()
}

func (ReportFormat) Type() protoreflect.EnumType {
	return &file_GenerateReportRequest_proto_enumTypes[1]
}

func (x ReportFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportFormat.Descriptor instead.
func (ReportFormat) EnumDescriptor() ([]byte, []int) {
	return file_GenerateReportRequest_proto_rawDescGZIP(), []int{1}
}

// GenerateReportRequest is used to trigger report generation.
type GenerateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format a report should be generated in.
	Format ReportFormat `protobuf:"varint,1,opt,name=Format,proto3,enum=core.ReportFormat" json:"Format,omitempty"`
	// Type defines which kind of report should be generated.
	Type ReportType `protobuf:"varint,2,opt,name=Type,proto3,enum=core.ReportType" json:"Type,omitempty"`
	// Year a report shoul be generated for.
	Year int64 `protobuf:"varint,3,opt,name=Year,proto3" json:"Year,omitempty"`
	// Month a report shoul be generated for.
	Month int64 `protobuf:"varint,4,opt,name=Month,proto3" json:"Month,omitempty"`
	// NamePattern defined a template for report file names.
	NamePattern string `protobuf:"bytes,5,opt,name=NamePattern,proto3" json:"NamePattern,omitempty"`
	// Delivery defines targets a eport should be send to.
	Delivery *ReportDelivery `protobuf:"bytes,6,opt,name=Delivery,proto3" json:"Delivery,omitempty"`
}

func (x *GenerateReportRequest) Reset() {
	*x = GenerateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GenerateReportRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateReportRequest) ProtoMessage() {}

func (x *GenerateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GenerateReportRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateReportRequest.ProtoReflect.Descriptor instead.
func (*GenerateReportRequest) Descriptor() ([]byte, []int) {
	return file_GenerateReportRequest_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateReportRequest) GetFormat() ReportFormat {
	if x != nil {
		return x.Format
	}
	return ReportFormat_NO_FORMAT
}

func (x *GenerateReportRequest) GetType() ReportType {
	if x != nil {
		return x.Type
	}
	return ReportType_NO_TYPE
}

func (x *GenerateReportRequest) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *GenerateReportRequest) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *GenerateReportRequest) GetNamePattern() string {
	if x != nil {
		return x.NamePattern
	}
	return ""
}

func (x *GenerateReportRequest) GetDelivery() *ReportDelivery {
	if x != nil {
		return x.Delivery
	}
	return nil
}

// ReportDelivery defines different targets a report can be disributed to.
type ReportDelivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// S3 defeines target bucket in AWS S4 a report should be uploaded to.
	S3 *S3Target `protobuf:"bytes,1,opt,name=S3,proto3" json:"S3,omitempty"`
	// Mail defines settings to send a report via email.
	Mail *MailTarget `protobuf:"bytes,2,opt,name=Mail,proto3" json:"Mail,omitempty"`
}

func (x *ReportDelivery) Reset() {
	*x = ReportDelivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GenerateReportRequest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDelivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDelivery) ProtoMessage() {}

func (x *ReportDelivery) ProtoReflect() protoreflect.Message {
	mi := &file_GenerateReportRequest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportDelivery.ProtoReflect.Descriptor instead.
func (*ReportDelivery) Descriptor() ([]byte, []int) {
	return file_GenerateReportRequest_proto_rawDescGZIP(), []int{1}
}

func (x *ReportDelivery) GetS3() *S3Target {
	if x != nil {
		return x.S3
	}
	return nil
}

func (x *ReportDelivery) GetMail() *MailTarget {
	if x != nil {
		return x.Mail
	}
	return nil
}

// S3Target defines settings to upload a report to an AWS S3 bucket.
type S3Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Region is a AWS region of a S3 bucket.
	Region string `protobuf:"bytes,1,opt,name=Region,proto3" json:"Region,omitempty"`
	// Bucket defines the name of a S3 bucket.
	Bucket string `protobuf:"bytes,2,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
	// Path defines a optional prefix path.
	Path string `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
}

func (x *S3Target) Reset() {
	*x = S3Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GenerateReportRequest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3Target) ProtoMessage() {}

func (x *S3Target) ProtoReflect() protoreflect.Message {
	mi := &file_GenerateReportRequest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3Target.ProtoReflect.Descriptor instead.
func (*S3Target) Descriptor() ([]byte, []int) {
	return file_GenerateReportRequest_proto_rawDescGZIP(), []int{2}
}

func (x *S3Target) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *S3Target) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *S3Target) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// MailTarget contains settings for mail delivery.
type MailTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ToAddresses provides a list of email receiver.
	ToAddresses []string `protobuf:"bytes,1,rep,name=ToAddresses,proto3" json:"ToAddresses,omitempty"`
}

func (x *MailTarget) Reset() {
	*x = MailTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GenerateReportRequest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailTarget) ProtoMessage() {}

func (x *MailTarget) ProtoReflect() protoreflect.Message {
	mi := &file_GenerateReportRequest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailTarget.ProtoReflect.Descriptor instead.
func (*MailTarget) Descriptor() ([]byte, []int) {
	return file_GenerateReportRequest_proto_rawDescGZIP(), []int{3}
}

func (x *MailTarget) GetToAddresses() []string {
	if x != nil {
		return x.ToAddresses
	}
	return nil
}

var File_GenerateReportRequest_proto protoreflect.FileDescriptor

var file_GenerateReportRequest_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0xe7, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x61, 0x6d,
	0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x22, 0x56, 0x0a,
	0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12,
	0x1e, 0x0a, 0x02, 0x53, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x02, 0x53, 0x33, 0x12,
	0x24, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x04, 0x4d, 0x61, 0x69, 0x6c, 0x22, 0x4e, 0x0a, 0x08, 0x53, 0x33, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x2a, 0x2d, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x4f,
	0x52, 0x54, 0x10, 0x01, 0x2a, 0x28, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x58, 0x43, 0x45, 0x4c, 0x10, 0x01, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GenerateReportRequest_proto_rawDescOnce sync.Once
	file_GenerateReportRequest_proto_rawDescData = file_GenerateReportRequest_proto_rawDesc
)

func file_GenerateReportRequest_proto_rawDescGZIP() []byte {
	file_GenerateReportRequest_proto_rawDescOnce.Do(func() {
		file_GenerateReportRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_GenerateReportRequest_proto_rawDescData)
	})
	return file_GenerateReportRequest_proto_rawDescData
}

var file_GenerateReportRequest_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_GenerateReportRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_GenerateReportRequest_proto_goTypes = []interface{}{
	(ReportType)(0),               // 0: core.ReportType
	(ReportFormat)(0),             // 1: core.ReportFormat
	(*GenerateReportRequest)(nil), // 2: core.GenerateReportRequest
	(*ReportDelivery)(nil),        // 3: core.ReportDelivery
	(*S3Target)(nil),              // 4: core.S3Target
	(*MailTarget)(nil),            // 5: core.MailTarget
}
var file_GenerateReportRequest_proto_depIdxs = []int32{
	1, // 0: core.GenerateReportRequest.Format:type_name -> core.ReportFormat
	0, // 1: core.GenerateReportRequest.Type:type_name -> core.ReportType
	3, // 2: core.GenerateReportRequest.Delivery:type_name -> core.ReportDelivery
	4, // 3: core.ReportDelivery.S3:type_name -> core.S3Target
	5, // 4: core.ReportDelivery.Mail:type_name -> core.MailTarget
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_GenerateReportRequest_proto_init() }
func file_GenerateReportRequest_proto_init() {
	if File_GenerateReportRequest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GenerateReportRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateReportRequest); i {
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
		file_GenerateReportRequest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportDelivery); i {
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
		file_GenerateReportRequest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3Target); i {
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
		file_GenerateReportRequest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailTarget); i {
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
			RawDescriptor: file_GenerateReportRequest_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GenerateReportRequest_proto_goTypes,
		DependencyIndexes: file_GenerateReportRequest_proto_depIdxs,
		EnumInfos:         file_GenerateReportRequest_proto_enumTypes,
		MessageInfos:      file_GenerateReportRequest_proto_msgTypes,
	}.Build()
	File_GenerateReportRequest_proto = out.File
	file_GenerateReportRequest_proto_rawDesc = nil
	file_GenerateReportRequest_proto_goTypes = nil
	file_GenerateReportRequest_proto_depIdxs = nil
}
