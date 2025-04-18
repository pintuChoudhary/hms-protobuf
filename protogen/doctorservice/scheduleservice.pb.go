// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: doctorservice/scheduleservice.proto

package doctorservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DoctorScheduleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DoctorId      string                 `protobuf:"bytes,1,opt,name=DoctorId,proto3" json:"DoctorId,omitempty"`
	Date          string                 `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoctorScheduleRequest) Reset() {
	*x = DoctorScheduleRequest{}
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoctorScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorScheduleRequest) ProtoMessage() {}

func (x *DoctorScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorScheduleRequest.ProtoReflect.Descriptor instead.
func (*DoctorScheduleRequest) Descriptor() ([]byte, []int) {
	return file_doctorservice_scheduleservice_proto_rawDescGZIP(), []int{0}
}

func (x *DoctorScheduleRequest) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *DoctorScheduleRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type DoctorScheduleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schedule      *DoctorSchedule        `protobuf:"bytes,1,opt,name=Schedule,proto3" json:"Schedule,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoctorScheduleResponse) Reset() {
	*x = DoctorScheduleResponse{}
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoctorScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorScheduleResponse) ProtoMessage() {}

func (x *DoctorScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorScheduleResponse.ProtoReflect.Descriptor instead.
func (*DoctorScheduleResponse) Descriptor() ([]byte, []int) {
	return file_doctorservice_scheduleservice_proto_rawDescGZIP(), []int{1}
}

func (x *DoctorScheduleResponse) GetSchedule() *DoctorSchedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *DoctorScheduleResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TimeSlot struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Duration      int32                  `protobuf:"varint,1,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Capacity      int32                  `protobuf:"varint,3,opt,name=Capacity,proto3" json:"Capacity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeSlot) Reset() {
	*x = TimeSlot{}
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSlot) ProtoMessage() {}

func (x *TimeSlot) ProtoReflect() protoreflect.Message {
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSlot.ProtoReflect.Descriptor instead.
func (*TimeSlot) Descriptor() ([]byte, []int) {
	return file_doctorservice_scheduleservice_proto_rawDescGZIP(), []int{2}
}

func (x *TimeSlot) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *TimeSlot) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TimeSlot) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

type RegularHour struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DayOfWeek     string                 `protobuf:"bytes,1,opt,name=DayOfWeek,proto3" json:"DayOfWeek,omitempty"`
	StartTime     string                 `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime       string                 `protobuf:"bytes,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegularHour) Reset() {
	*x = RegularHour{}
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegularHour) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegularHour) ProtoMessage() {}

func (x *RegularHour) ProtoReflect() protoreflect.Message {
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegularHour.ProtoReflect.Descriptor instead.
func (*RegularHour) Descriptor() ([]byte, []int) {
	return file_doctorservice_scheduleservice_proto_rawDescGZIP(), []int{3}
}

func (x *RegularHour) GetDayOfWeek() string {
	if x != nil {
		return x.DayOfWeek
	}
	return ""
}

func (x *RegularHour) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *RegularHour) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type TimeOff struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StartTime     string                 `protobuf:"bytes,1,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime       string                 `protobuf:"bytes,2,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Reason        string                 `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeOff) Reset() {
	*x = TimeOff{}
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeOff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeOff) ProtoMessage() {}

func (x *TimeOff) ProtoReflect() protoreflect.Message {
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeOff.ProtoReflect.Descriptor instead.
func (*TimeOff) Descriptor() ([]byte, []int) {
	return file_doctorservice_scheduleservice_proto_rawDescGZIP(), []int{4}
}

func (x *TimeOff) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *TimeOff) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *TimeOff) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type DoctorSchedule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DoctorId      string                 `protobuf:"bytes,1,opt,name=DoctorId,proto3" json:"DoctorId,omitempty"`
	Date          string                 `protobuf:"bytes,2,opt,name=Date,proto3" json:"Date,omitempty"`
	Availability  bool                   `protobuf:"varint,3,opt,name=Availability,proto3" json:"Availability,omitempty"` // No use
	TimeSlots     []*TimeSlot            `protobuf:"bytes,4,rep,name=TimeSlots,proto3" json:"TimeSlots,omitempty"`
	RegularHours  []*RegularHour         `protobuf:"bytes,5,rep,name=RegularHours,proto3" json:"RegularHours,omitempty"`
	TimeOffs      []*TimeOff             `protobuf:"bytes,6,rep,name=TimeOffs,proto3" json:"TimeOffs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoctorSchedule) Reset() {
	*x = DoctorSchedule{}
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoctorSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorSchedule) ProtoMessage() {}

func (x *DoctorSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_doctorservice_scheduleservice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorSchedule.ProtoReflect.Descriptor instead.
func (*DoctorSchedule) Descriptor() ([]byte, []int) {
	return file_doctorservice_scheduleservice_proto_rawDescGZIP(), []int{5}
}

func (x *DoctorSchedule) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *DoctorSchedule) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *DoctorSchedule) GetAvailability() bool {
	if x != nil {
		return x.Availability
	}
	return false
}

func (x *DoctorSchedule) GetTimeSlots() []*TimeSlot {
	if x != nil {
		return x.TimeSlots
	}
	return nil
}

func (x *DoctorSchedule) GetRegularHours() []*RegularHour {
	if x != nil {
		return x.RegularHours
	}
	return nil
}

func (x *DoctorSchedule) GetTimeOffs() []*TimeOff {
	if x != nil {
		return x.TimeOffs
	}
	return nil
}

var File_doctorservice_scheduleservice_proto protoreflect.FileDescriptor

const file_doctorservice_scheduleservice_proto_rawDesc = "" +
	"\n" +
	"#doctorservice/scheduleservice.proto\x12\rdoctorservice\"G\n" +
	"\x15DoctorScheduleRequest\x12\x1a\n" +
	"\bDoctorId\x18\x01 \x01(\tR\bDoctorId\x12\x12\n" +
	"\x04date\x18\x02 \x01(\tR\x04date\"i\n" +
	"\x16DoctorScheduleResponse\x129\n" +
	"\bSchedule\x18\x01 \x01(\v2\x1d.doctorservice.DoctorScheduleR\bSchedule\x12\x14\n" +
	"\x05Error\x18\x02 \x01(\tR\x05Error\"V\n" +
	"\bTimeSlot\x12\x1a\n" +
	"\bDuration\x18\x01 \x01(\x05R\bDuration\x12\x12\n" +
	"\x04Type\x18\x02 \x01(\tR\x04Type\x12\x1a\n" +
	"\bCapacity\x18\x03 \x01(\x05R\bCapacity\"c\n" +
	"\vRegularHour\x12\x1c\n" +
	"\tDayOfWeek\x18\x01 \x01(\tR\tDayOfWeek\x12\x1c\n" +
	"\tStartTime\x18\x02 \x01(\tR\tStartTime\x12\x18\n" +
	"\aEndTime\x18\x03 \x01(\tR\aEndTime\"Y\n" +
	"\aTimeOff\x12\x1c\n" +
	"\tStartTime\x18\x01 \x01(\tR\tStartTime\x12\x18\n" +
	"\aEndTime\x18\x02 \x01(\tR\aEndTime\x12\x16\n" +
	"\x06Reason\x18\x03 \x01(\tR\x06Reason\"\x8f\x02\n" +
	"\x0eDoctorSchedule\x12\x1a\n" +
	"\bDoctorId\x18\x01 \x01(\tR\bDoctorId\x12\x12\n" +
	"\x04Date\x18\x02 \x01(\tR\x04Date\x12\"\n" +
	"\fAvailability\x18\x03 \x01(\bR\fAvailability\x125\n" +
	"\tTimeSlots\x18\x04 \x03(\v2\x17.doctorservice.TimeSlotR\tTimeSlots\x12>\n" +
	"\fRegularHours\x18\x05 \x03(\v2\x1a.doctorservice.RegularHourR\fRegularHours\x122\n" +
	"\bTimeOffs\x18\x06 \x03(\v2\x16.doctorservice.TimeOffR\bTimeOffs2y\n" +
	"\x15DoctorScheduleService\x12`\n" +
	"\x11GetDoctorSchedule\x12$.doctorservice.DoctorScheduleRequest\x1a%.doctorservice.DoctorScheduleResponseB\x1fZ\x1d./doctorservice;doctorserviceb\x06proto3"

var (
	file_doctorservice_scheduleservice_proto_rawDescOnce sync.Once
	file_doctorservice_scheduleservice_proto_rawDescData []byte
)

func file_doctorservice_scheduleservice_proto_rawDescGZIP() []byte {
	file_doctorservice_scheduleservice_proto_rawDescOnce.Do(func() {
		file_doctorservice_scheduleservice_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_doctorservice_scheduleservice_proto_rawDesc), len(file_doctorservice_scheduleservice_proto_rawDesc)))
	})
	return file_doctorservice_scheduleservice_proto_rawDescData
}

var file_doctorservice_scheduleservice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_doctorservice_scheduleservice_proto_goTypes = []any{
	(*DoctorScheduleRequest)(nil),  // 0: doctorservice.DoctorScheduleRequest
	(*DoctorScheduleResponse)(nil), // 1: doctorservice.DoctorScheduleResponse
	(*TimeSlot)(nil),               // 2: doctorservice.TimeSlot
	(*RegularHour)(nil),            // 3: doctorservice.RegularHour
	(*TimeOff)(nil),                // 4: doctorservice.TimeOff
	(*DoctorSchedule)(nil),         // 5: doctorservice.DoctorSchedule
}
var file_doctorservice_scheduleservice_proto_depIdxs = []int32{
	5, // 0: doctorservice.DoctorScheduleResponse.Schedule:type_name -> doctorservice.DoctorSchedule
	2, // 1: doctorservice.DoctorSchedule.TimeSlots:type_name -> doctorservice.TimeSlot
	3, // 2: doctorservice.DoctorSchedule.RegularHours:type_name -> doctorservice.RegularHour
	4, // 3: doctorservice.DoctorSchedule.TimeOffs:type_name -> doctorservice.TimeOff
	0, // 4: doctorservice.DoctorScheduleService.GetDoctorSchedule:input_type -> doctorservice.DoctorScheduleRequest
	1, // 5: doctorservice.DoctorScheduleService.GetDoctorSchedule:output_type -> doctorservice.DoctorScheduleResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_doctorservice_scheduleservice_proto_init() }
func file_doctorservice_scheduleservice_proto_init() {
	if File_doctorservice_scheduleservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_doctorservice_scheduleservice_proto_rawDesc), len(file_doctorservice_scheduleservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doctorservice_scheduleservice_proto_goTypes,
		DependencyIndexes: file_doctorservice_scheduleservice_proto_depIdxs,
		MessageInfos:      file_doctorservice_scheduleservice_proto_msgTypes,
	}.Build()
	File_doctorservice_scheduleservice_proto = out.File
	file_doctorservice_scheduleservice_proto_goTypes = nil
	file_doctorservice_scheduleservice_proto_depIdxs = nil
}
