// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: login.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginWebReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails   string `protobuf:"bytes,1,opt,name=emails,proto3" json:"emails,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Time     string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *LoginWebReq) Reset() {
	*x = LoginWebReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginWebReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginWebReq) ProtoMessage() {}

func (x *LoginWebReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginWebReq.ProtoReflect.Descriptor instead.
func (*LoginWebReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginWebReq) GetEmails() string {
	if x != nil {
		return x.Emails
	}
	return ""
}

func (x *LoginWebReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginWebReq) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type Loginresp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32      `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Data    *LoginInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Loginresp) Reset() {
	*x = Loginresp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loginresp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loginresp) ProtoMessage() {}

func (x *Loginresp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loginresp.ProtoReflect.Descriptor instead.
func (*Loginresp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *Loginresp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Loginresp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Loginresp) GetData() *LoginInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type LoginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nicename string `protobuf:"bytes,1,opt,name=nicename,proto3" json:"nicename,omitempty"`
	Emails   string `protobuf:"bytes,2,opt,name=emails,proto3" json:"emails,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginInfo) Reset() {
	*x = LoginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginInfo) ProtoMessage() {}

func (x *LoginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginInfo.ProtoReflect.Descriptor instead.
func (*LoginInfo) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginInfo) GetNicename() string {
	if x != nil {
		return x.Nicename
	}
	return ""
}

func (x *LoginInfo) GetEmails() string {
	if x != nil {
		return x.Emails
	}
	return ""
}

func (x *LoginInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginPwdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails   string `protobuf:"bytes,1,opt,name=emails,proto3" json:"emails,omitempty"`
	Code     string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginPwdReq) Reset() {
	*x = LoginPwdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginPwdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPwdReq) ProtoMessage() {}

func (x *LoginPwdReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPwdReq.ProtoReflect.Descriptor instead.
func (*LoginPwdReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginPwdReq) GetEmails() string {
	if x != nil {
		return x.Emails
	}
	return ""
}

func (x *LoginPwdReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LoginPwdReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginPwdLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password    string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	Time        string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *LoginPwdLoginReq) Reset() {
	*x = LoginPwdLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginPwdLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPwdLoginReq) ProtoMessage() {}

func (x *LoginPwdLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPwdLoginReq.ProtoReflect.Descriptor instead.
func (*LoginPwdLoginReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{4}
}

func (x *LoginPwdLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginPwdLoginReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *LoginPwdLoginReq) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type LoginInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
}

func (x *LoginInfoReq) Reset() {
	*x = LoginInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginInfoReq) ProtoMessage() {}

func (x *LoginInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginInfoReq.ProtoReflect.Descriptor instead.
func (*LoginInfoReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{5}
}

func (x *LoginInfoReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type LoginUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails   string `protobuf:"bytes,1,opt,name=emails,proto3" json:"emails,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	NickName string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
}

func (x *LoginUserReq) Reset() {
	*x = LoginUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserReq) ProtoMessage() {}

func (x *LoginUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserReq.ProtoReflect.Descriptor instead.
func (*LoginUserReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{6}
}

func (x *LoginUserReq) GetEmails() string {
	if x != nil {
		return x.Emails
	}
	return ""
}

func (x *LoginUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginUserReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LoginUserReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type LoginUserCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emails string `protobuf:"bytes,1,opt,name=emails,proto3" json:"emails,omitempty"`
	Types  string `protobuf:"bytes,2,opt,name=types,proto3" json:"types,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LoginUserCodeReq) Reset() {
	*x = LoginUserCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserCodeReq) ProtoMessage() {}

func (x *LoginUserCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserCodeReq.ProtoReflect.Descriptor instead.
func (*LoginUserCodeReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{7}
}

func (x *LoginUserCodeReq) GetEmails() string {
	if x != nil {
		return x.Emails
	}
	return ""
}

func (x *LoginUserCodeReq) GetTypes() string {
	if x != nil {
		return x.Types
	}
	return ""
}

func (x *LoginUserCodeReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x65, 0x62, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x09, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x09, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x55, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x64, 0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x50, 0x77, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2a,
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x0c, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x54,
	0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x81, 0x04, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x4c,
	0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x65, 0x62, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x65, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x73, 0x70,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x4f, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x77, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x5f, 0x0a,
	0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x77, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x77, 0x64, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x3a, 0x01, 0x2a, 0x1a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x4d,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65,
	0x73, 0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x4d, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x73,
	0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x0d,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x72, 0x65, 0x73, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x79, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_login_proto_goTypes = []interface{}{
	(*LoginWebReq)(nil),      // 0: proto.loginWebReq
	(*Loginresp)(nil),        // 1: proto.loginresp
	(*LoginInfo)(nil),        // 2: proto.loginInfo
	(*LoginPwdReq)(nil),      // 3: proto.loginPwdReq
	(*LoginPwdLoginReq)(nil), // 4: proto.loginPwdLoginReq
	(*LoginInfoReq)(nil),     // 5: proto.loginInfoReq
	(*LoginUserReq)(nil),     // 6: proto.loginUserReq
	(*LoginUserCodeReq)(nil), // 7: proto.loginUserCodeReq
}
var file_login_proto_depIdxs = []int32{
	2, // 0: proto.loginresp.data:type_name -> proto.loginInfo
	0, // 1: proto.Login.loginWeb:input_type -> proto.loginWebReq
	3, // 2: proto.Login.loginPwd:input_type -> proto.loginPwdReq
	4, // 3: proto.Login.loginPwdLogin:input_type -> proto.loginPwdLoginReq
	5, // 4: proto.Login.loginInfo:input_type -> proto.loginInfoReq
	6, // 5: proto.Login.loginUser:input_type -> proto.loginUserReq
	7, // 6: proto.Login.loginUserCode:input_type -> proto.loginUserCodeReq
	1, // 7: proto.Login.loginWeb:output_type -> proto.loginresp
	1, // 8: proto.Login.loginPwd:output_type -> proto.loginresp
	1, // 9: proto.Login.loginPwdLogin:output_type -> proto.loginresp
	1, // 10: proto.Login.loginInfo:output_type -> proto.loginresp
	1, // 11: proto.Login.loginUser:output_type -> proto.loginresp
	1, // 12: proto.Login.loginUserCode:output_type -> proto.loginresp
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginWebReq); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loginresp); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginInfo); i {
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
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginPwdReq); i {
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
		file_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginPwdLoginReq); i {
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
		file_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginInfoReq); i {
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
		file_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserReq); i {
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
		file_login_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserCodeReq); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
