// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: protos/class.proto

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

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	SchoolYear        string `protobuf:"bytes,3,opt,name=SchoolYear,proto3" json:"SchoolYear,omitempty"`
	HomeroomTeacherId string `protobuf:"bytes,4,opt,name=HomeroomTeacherId,proto3" json:"HomeroomTeacherId,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	mi := &file_protos_class_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_protos_class_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_protos_class_proto_rawDescGZIP(), []int{0}
}

func (x *Class) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Class) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Class) GetSchoolYear() string {
	if x != nil {
		return x.SchoolYear
	}
	return ""
}

func (x *Class) GetHomeroomTeacherId() string {
	if x != nil {
		return x.HomeroomTeacherId
	}
	return ""
}

// Create Class
type CreateClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	SchoolYear        string   `protobuf:"bytes,2,opt,name=SchoolYear,proto3" json:"SchoolYear,omitempty"`
	HomeroomTeacherId string   `protobuf:"bytes,3,opt,name=HomeroomTeacherId,proto3" json:"HomeroomTeacherId,omitempty"`
	StudentsId        []string `protobuf:"bytes,4,rep,name=StudentsId,proto3" json:"StudentsId,omitempty"`
}

func (x *CreateClassRequest) Reset() {
	*x = CreateClassRequest{}
	mi := &file_protos_class_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassRequest) ProtoMessage() {}

func (x *CreateClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_class_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassRequest.ProtoReflect.Descriptor instead.
func (*CreateClassRequest) Descriptor() ([]byte, []int) {
	return file_protos_class_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClassRequest) GetSchoolYear() string {
	if x != nil {
		return x.SchoolYear
	}
	return ""
}

func (x *CreateClassRequest) GetHomeroomTeacherId() string {
	if x != nil {
		return x.HomeroomTeacherId
	}
	return ""
}

func (x *CreateClassRequest) GetStudentsId() []string {
	if x != nil {
		return x.StudentsId
	}
	return nil
}

type CreateClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ClassCreated *Class `protobuf:"bytes,2,opt,name=classCreated,proto3" json:"classCreated,omitempty"`
}

func (x *CreateClassResponse) Reset() {
	*x = CreateClassResponse{}
	mi := &file_protos_class_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClassResponse) ProtoMessage() {}

func (x *CreateClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_class_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClassResponse.ProtoReflect.Descriptor instead.
func (*CreateClassResponse) Descriptor() ([]byte, []int) {
	return file_protos_class_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClassResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateClassResponse) GetClassCreated() *Class {
	if x != nil {
		return x.ClassCreated
	}
	return nil
}

// Get all classes
type GetAllClassesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ClassList []*Class `protobuf:"bytes,2,rep,name=classList,proto3" json:"classList,omitempty"`
}

func (x *GetAllClassesResponse) Reset() {
	*x = GetAllClassesResponse{}
	mi := &file_protos_class_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllClassesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllClassesResponse) ProtoMessage() {}

func (x *GetAllClassesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_class_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllClassesResponse.ProtoReflect.Descriptor instead.
func (*GetAllClassesResponse) Descriptor() ([]byte, []int) {
	return file_protos_class_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllClassesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllClassesResponse) GetClassList() []*Class {
	if x != nil {
		return x.ClassList
	}
	return nil
}

type GetClassByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClassByIdRequest) Reset() {
	*x = GetClassByIdRequest{}
	mi := &file_protos_class_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClassByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByIdRequest) ProtoMessage() {}

func (x *GetClassByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_class_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByIdRequest.ProtoReflect.Descriptor instead.
func (*GetClassByIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_class_proto_rawDescGZIP(), []int{4}
}

func (x *GetClassByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetClassByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Class    *Class     `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	Students []*Student `protobuf:"bytes,3,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *GetClassByIdResponse) Reset() {
	*x = GetClassByIdResponse{}
	mi := &file_protos_class_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClassByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClassByIdResponse) ProtoMessage() {}

func (x *GetClassByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_class_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClassByIdResponse.ProtoReflect.Descriptor instead.
func (*GetClassByIdResponse) Descriptor() ([]byte, []int) {
	return file_protos_class_proto_rawDescGZIP(), []int{5}
}

func (x *GetClassByIdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetClassByIdResponse) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

func (x *GetClassByIdResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

var File_protos_class_proto protoreflect.FileDescriptor

var file_protos_class_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x59, 0x65, 0x61, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x48, 0x6f, 0x6d, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x12, 0x2c, 0x0a,
	0x11, 0x48, 0x6f, 0x6d, 0x65, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x48, 0x6f, 0x6d, 0x65, 0x72, 0x6f,
	0x6f, 0x6d, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x0c,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x0c, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x57, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x74, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xc9,
	0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_class_proto_rawDescOnce sync.Once
	file_protos_class_proto_rawDescData = file_protos_class_proto_rawDesc
)

func file_protos_class_proto_rawDescGZIP() []byte {
	file_protos_class_proto_rawDescOnce.Do(func() {
		file_protos_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_class_proto_rawDescData)
	})
	return file_protos_class_proto_rawDescData
}

var file_protos_class_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_class_proto_goTypes = []any{
	(*Class)(nil),                 // 0: Class
	(*CreateClassRequest)(nil),    // 1: CreateClassRequest
	(*CreateClassResponse)(nil),   // 2: CreateClassResponse
	(*GetAllClassesResponse)(nil), // 3: GetAllClassesResponse
	(*GetClassByIdRequest)(nil),   // 4: GetClassByIdRequest
	(*GetClassByIdResponse)(nil),  // 5: GetClassByIdResponse
	(*Student)(nil),               // 6: Student
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_protos_class_proto_depIdxs = []int32{
	0, // 0: CreateClassResponse.classCreated:type_name -> Class
	0, // 1: GetAllClassesResponse.classList:type_name -> Class
	0, // 2: GetClassByIdResponse.class:type_name -> Class
	6, // 3: GetClassByIdResponse.students:type_name -> Student
	1, // 4: ClassService.CreateNewClass:input_type -> CreateClassRequest
	7, // 5: ClassService.GetAllClasses:input_type -> google.protobuf.Empty
	4, // 6: ClassService.GetClassById:input_type -> GetClassByIdRequest
	2, // 7: ClassService.CreateNewClass:output_type -> CreateClassResponse
	3, // 8: ClassService.GetAllClasses:output_type -> GetAllClassesResponse
	5, // 9: ClassService.GetClassById:output_type -> GetClassByIdResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_class_proto_init() }
func file_protos_class_proto_init() {
	if File_protos_class_proto != nil {
		return
	}
	file_protos_student_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_class_proto_goTypes,
		DependencyIndexes: file_protos_class_proto_depIdxs,
		MessageInfos:      file_protos_class_proto_msgTypes,
	}.Build()
	File_protos_class_proto = out.File
	file_protos_class_proto_rawDesc = nil
	file_protos_class_proto_goTypes = nil
	file_protos_class_proto_depIdxs = nil
}
