// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: protos/scoreNew.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExamType int32

const (
	ExamType_fifteen         ExamType = 0
	ExamType_fourty_five     ExamType = 1
	ExamType_first_semester  ExamType = 2
	ExamType_second_semester ExamType = 3
	ExamType_mouth           ExamType = 4 //Kiem tra mom
)

// Enum value maps for ExamType.
var (
	ExamType_name = map[int32]string{
		0: "fifteen",
		1: "fourty_five",
		2: "first_semester",
		3: "second_semester",
		4: "mouth",
	}
	ExamType_value = map[string]int32{
		"fifteen":         0,
		"fourty_five":     1,
		"first_semester":  2,
		"second_semester": 3,
		"mouth":           4,
	}
)

func (x ExamType) Enum() *ExamType {
	p := new(ExamType)
	*p = x
	return p
}

func (x ExamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExamType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_scoreNew_proto_enumTypes[0].Descriptor()
}

func (ExamType) Type() protoreflect.EnumType {
	return &file_protos_scoreNew_proto_enumTypes[0]
}

func (x ExamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExamType.Descriptor instead.
func (ExamType) EnumDescriptor() ([]byte, []int) {
	return file_protos_scoreNew_proto_rawDescGZIP(), []int{0}
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	StudentId   string   `protobuf:"bytes,2,opt,name=StudentId,proto3" json:"StudentId,omitempty"`
	SubjectId   string   `protobuf:"bytes,3,opt,name=SubjectId,proto3" json:"SubjectId,omitempty"`
	ExamType    ExamType `protobuf:"varint,4,opt,name=examType,proto3,enum=ExamType" json:"examType,omitempty"`
	Result      float64  `protobuf:"fixed64,5,opt,name=result,proto3" json:"result,omitempty"`
	Semester    int32    `protobuf:"varint,6,opt,name=semester,proto3" json:"semester,omitempty"`
	Weight      int32    `protobuf:"varint,7,opt,name=weight,proto3" json:"weight,omitempty"`
	CreatedById string   `protobuf:"bytes,8,opt,name=CreatedById,proto3" json:"CreatedById,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	mi := &file_protos_scoreNew_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_protos_scoreNew_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_protos_scoreNew_proto_rawDescGZIP(), []int{0}
}

func (x *Score) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Score) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *Score) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Score) GetExamType() ExamType {
	if x != nil {
		return x.ExamType
	}
	return ExamType_fifteen
}

func (x *Score) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *Score) GetSemester() int32 {
	if x != nil {
		return x.Semester
	}
	return 0
}

func (x *Score) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Score) GetCreatedById() string {
	if x != nil {
		return x.CreatedById
	}
	return ""
}

type CreateScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string   `protobuf:"bytes,1,opt,name=StudentId,proto3" json:"StudentId,omitempty"`
	SubjectId string   `protobuf:"bytes,2,opt,name=SubjectId,proto3" json:"SubjectId,omitempty"`
	ExamType  ExamType `protobuf:"varint,3,opt,name=examType,proto3,enum=ExamType" json:"examType,omitempty"`
	Semester  int32    `protobuf:"varint,4,opt,name=semester,proto3" json:"semester,omitempty"`
	Result    float64  `protobuf:"fixed64,5,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateScoreRequest) Reset() {
	*x = CreateScoreRequest{}
	mi := &file_protos_scoreNew_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScoreRequest) ProtoMessage() {}

func (x *CreateScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_scoreNew_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScoreRequest.ProtoReflect.Descriptor instead.
func (*CreateScoreRequest) Descriptor() ([]byte, []int) {
	return file_protos_scoreNew_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScoreRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateScoreRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *CreateScoreRequest) GetExamType() ExamType {
	if x != nil {
		return x.ExamType
	}
	return ExamType_fifteen
}

func (x *CreateScoreRequest) GetSemester() int32 {
	if x != nil {
		return x.Semester
	}
	return 0
}

func (x *CreateScoreRequest) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type CreateScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ScoreCreated *Score `protobuf:"bytes,2,opt,name=scoreCreated,proto3" json:"scoreCreated,omitempty"`
}

func (x *CreateScoreResponse) Reset() {
	*x = CreateScoreResponse{}
	mi := &file_protos_scoreNew_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScoreResponse) ProtoMessage() {}

func (x *CreateScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_scoreNew_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScoreResponse.ProtoReflect.Descriptor instead.
func (*CreateScoreResponse) Descriptor() ([]byte, []int) {
	return file_protos_scoreNew_proto_rawDescGZIP(), []int{2}
}

func (x *CreateScoreResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateScoreResponse) GetScoreCreated() *Score {
	if x != nil {
		return x.ScoreCreated
	}
	return nil
}

type GetAllScoreStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Scores  []*Score `protobuf:"bytes,2,rep,name=scores,proto3" json:"scores,omitempty"`
}

func (x *GetAllScoreStudentResponse) Reset() {
	*x = GetAllScoreStudentResponse{}
	mi := &file_protos_scoreNew_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllScoreStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllScoreStudentResponse) ProtoMessage() {}

func (x *GetAllScoreStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_scoreNew_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllScoreStudentResponse.ProtoReflect.Descriptor instead.
func (*GetAllScoreStudentResponse) Descriptor() ([]byte, []int) {
	return file_protos_scoreNew_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllScoreStudentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllScoreStudentResponse) GetScores() []*Score {
	if x != nil {
		return x.Scores
	}
	return nil
}

var File_protos_scoreNew_proto protoreflect.FileDescriptor

var file_protos_scoreNew_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x4e, 0x65,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x78,
	0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x6d,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x64, 0x22,
	0xab, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x65, 0x78, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x6d,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5b, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0c, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x56, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x2a, 0x5c, 0x0a, 0x08, 0x45, 0x78, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x66, 0x69, 0x66, 0x74, 0x65, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x66,
	0x6f, 0x75, 0x72, 0x74, 0x79, 0x5f, 0x66, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x6d, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x6f, 0x75, 0x74, 0x68, 0x10, 0x04,
	0x32, 0x8f, 0x01, 0x0a, 0x0c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_scoreNew_proto_rawDescOnce sync.Once
	file_protos_scoreNew_proto_rawDescData = file_protos_scoreNew_proto_rawDesc
)

func file_protos_scoreNew_proto_rawDescGZIP() []byte {
	file_protos_scoreNew_proto_rawDescOnce.Do(func() {
		file_protos_scoreNew_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_scoreNew_proto_rawDescData)
	})
	return file_protos_scoreNew_proto_rawDescData
}

var file_protos_scoreNew_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_scoreNew_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_scoreNew_proto_goTypes = []any{
	(ExamType)(0),                      // 0: ExamType
	(*Score)(nil),                      // 1: Score
	(*CreateScoreRequest)(nil),         // 2: CreateScoreRequest
	(*CreateScoreResponse)(nil),        // 3: CreateScoreResponse
	(*GetAllScoreStudentResponse)(nil), // 4: GetAllScoreStudentResponse
	(*emptypb.Empty)(nil),              // 5: google.protobuf.Empty
}
var file_protos_scoreNew_proto_depIdxs = []int32{
	0, // 0: Score.examType:type_name -> ExamType
	0, // 1: CreateScoreRequest.examType:type_name -> ExamType
	1, // 2: CreateScoreResponse.scoreCreated:type_name -> Score
	1, // 3: GetAllScoreStudentResponse.scores:type_name -> Score
	2, // 4: ScoreService.CreateNewScore:input_type -> CreateScoreRequest
	5, // 5: ScoreService.GetAllScore:input_type -> google.protobuf.Empty
	3, // 6: ScoreService.CreateNewScore:output_type -> CreateScoreResponse
	4, // 7: ScoreService.GetAllScore:output_type -> GetAllScoreStudentResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_scoreNew_proto_init() }
func file_protos_scoreNew_proto_init() {
	if File_protos_scoreNew_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_scoreNew_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_scoreNew_proto_goTypes,
		DependencyIndexes: file_protos_scoreNew_proto_depIdxs,
		EnumInfos:         file_protos_scoreNew_proto_enumTypes,
		MessageInfos:      file_protos_scoreNew_proto_msgTypes,
	}.Build()
	File_protos_scoreNew_proto = out.File
	file_protos_scoreNew_proto_rawDesc = nil
	file_protos_scoreNew_proto_goTypes = nil
	file_protos_scoreNew_proto_depIdxs = nil
}
