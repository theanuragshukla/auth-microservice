// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc2
// source: protos/auth.proto

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data   *Tokens `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Errors *Errors `protobuf:"bytes,4,opt,name=errors,proto3,oneof" json:"errors,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *LoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoginResponse) GetData() *Tokens {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoginResponse) GetErrors() *Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string  `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  *string `protobuf:"bytes,2,opt,name=lastName,proto3,oneof" json:"lastName,omitempty"`
	Email     string  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string  `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignupRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SignupRequest) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data   *Tokens `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Errors *Errors `protobuf:"bytes,4,opt,name=errors,proto3,oneof" json:"errors,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignupResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SignupResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SignupResponse) GetData() *Tokens {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SignupResponse) GetErrors() *Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	Uid          string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{4}
}

func (x *Tokens) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Tokens) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Tokens) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type SingleValidationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field *string `protobuf:"bytes,1,opt,name=field,proto3,oneof" json:"field,omitempty"`
	Tag   *string `protobuf:"bytes,2,opt,name=tag,proto3,oneof" json:"tag,omitempty"`
	Value *string `protobuf:"bytes,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *SingleValidationError) Reset() {
	*x = SingleValidationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleValidationError) ProtoMessage() {}

func (x *SingleValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleValidationError.ProtoReflect.Descriptor instead.
func (*SingleValidationError) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{5}
}

func (x *SingleValidationError) GetField() string {
	if x != nil && x.Field != nil {
		return *x.Field
	}
	return ""
}

func (x *SingleValidationError) GetTag() string {
	if x != nil && x.Tag != nil {
		return *x.Tag
	}
	return ""
}

func (x *SingleValidationError) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type Errors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []*SingleValidationError `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Errors) Reset() {
	*x = Errors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errors) ProtoMessage() {}

func (x *Errors) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errors.ProtoReflect.Descriptor instead.
func (*Errors) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{6}
}

func (x *Errors) GetErrors() []*SingleValidationError {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_protos_auth_proto protoreflect.FileDescriptor

var file_protos_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x48, 0x01, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x8d, 0x01,
	0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x01,
	0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x48, 0x01, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x60, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x15, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x19, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x61,
	0x67, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x06, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0x59, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x26, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12,
	0x0e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_auth_proto_rawDescOnce sync.Once
	file_protos_auth_proto_rawDescData = file_protos_auth_proto_rawDesc
)

func file_protos_auth_proto_rawDescGZIP() []byte {
	file_protos_auth_proto_rawDescOnce.Do(func() {
		file_protos_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_auth_proto_rawDescData)
	})
	return file_protos_auth_proto_rawDescData
}

var file_protos_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_auth_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),          // 0: LoginRequest
	(*LoginResponse)(nil),         // 1: LoginResponse
	(*SignupRequest)(nil),         // 2: SignupRequest
	(*SignupResponse)(nil),        // 3: SignupResponse
	(*Tokens)(nil),                // 4: Tokens
	(*SingleValidationError)(nil), // 5: SingleValidationError
	(*Errors)(nil),                // 6: Errors
}
var file_protos_auth_proto_depIdxs = []int32{
	4, // 0: LoginResponse.data:type_name -> Tokens
	6, // 1: LoginResponse.errors:type_name -> Errors
	4, // 2: SignupResponse.data:type_name -> Tokens
	6, // 3: SignupResponse.errors:type_name -> Errors
	5, // 4: Errors.errors:type_name -> SingleValidationError
	0, // 5: Auth.Login:input_type -> LoginRequest
	2, // 6: Auth.Signup:input_type -> SignupRequest
	1, // 7: Auth.Login:output_type -> LoginResponse
	3, // 8: Auth.Signup:output_type -> SignupResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_auth_proto_init() }
func file_protos_auth_proto_init() {
	if File_protos_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_protos_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_protos_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
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
		file_protos_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
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
		file_protos_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
		file_protos_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleValidationError); i {
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
		file_protos_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errors); i {
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
	file_protos_auth_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_protos_auth_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protos_auth_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_protos_auth_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_auth_proto_goTypes,
		DependencyIndexes: file_protos_auth_proto_depIdxs,
		MessageInfos:      file_protos_auth_proto_msgTypes,
	}.Build()
	File_protos_auth_proto = out.File
	file_protos_auth_proto_rawDesc = nil
	file_protos_auth_proto_goTypes = nil
	file_protos_auth_proto_depIdxs = nil
}
