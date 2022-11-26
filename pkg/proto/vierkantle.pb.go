// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: vierkantle.proto

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

type CreateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type JoinTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *JoinTeamRequest) Reset() {
	*x = JoinTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinTeamRequest) ProtoMessage() {}

func (x *JoinTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinTeamRequest.ProtoReflect.Descriptor instead.
func (*JoinTeamRequest) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{1}
}

func (x *JoinTeamRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JoinTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChangeNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ChangeNameRequest) Reset() {
	*x = ChangeNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNameRequest) ProtoMessage() {}

func (x *ChangeNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNameRequest.ProtoReflect.Descriptor instead.
func (*ChangeNameRequest) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WordGuessedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *WordGuessedRequest) Reset() {
	*x = WordGuessedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordGuessedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordGuessedRequest) ProtoMessage() {}

func (x *WordGuessedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordGuessedRequest.ProtoReflect.Descriptor instead.
func (*WordGuessedRequest) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{3}
}

func (x *WordGuessedRequest) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type TeamStreamClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*TeamStreamClientMessage_Create
	//	*TeamStreamClientMessage_Join
	//	*TeamStreamClientMessage_Name
	//	*TeamStreamClientMessage_Word
	Request isTeamStreamClientMessage_Request `protobuf_oneof:"request"`
}

func (x *TeamStreamClientMessage) Reset() {
	*x = TeamStreamClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamStreamClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamStreamClientMessage) ProtoMessage() {}

func (x *TeamStreamClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamStreamClientMessage.ProtoReflect.Descriptor instead.
func (*TeamStreamClientMessage) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{4}
}

func (m *TeamStreamClientMessage) GetRequest() isTeamStreamClientMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *TeamStreamClientMessage) GetCreate() *CreateTeamRequest {
	if x, ok := x.GetRequest().(*TeamStreamClientMessage_Create); ok {
		return x.Create
	}
	return nil
}

func (x *TeamStreamClientMessage) GetJoin() *JoinTeamRequest {
	if x, ok := x.GetRequest().(*TeamStreamClientMessage_Join); ok {
		return x.Join
	}
	return nil
}

func (x *TeamStreamClientMessage) GetName() *ChangeNameRequest {
	if x, ok := x.GetRequest().(*TeamStreamClientMessage_Name); ok {
		return x.Name
	}
	return nil
}

func (x *TeamStreamClientMessage) GetWord() *WordGuessedRequest {
	if x, ok := x.GetRequest().(*TeamStreamClientMessage_Word); ok {
		return x.Word
	}
	return nil
}

type isTeamStreamClientMessage_Request interface {
	isTeamStreamClientMessage_Request()
}

type TeamStreamClientMessage_Create struct {
	// The stream must start with one create or join message. After this, the
	// server responds with TeamInfoResponse, after which the create or join
	// must not be sent anymore. If the player name already exists in the team,
	// the server will respond with an ErrorResponse and the stream will be
	// closed.
	Create *CreateTeamRequest `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type TeamStreamClientMessage_Join struct {
	Join *JoinTeamRequest `protobuf:"bytes,2,opt,name=join,proto3,oneof"`
}

type TeamStreamClientMessage_Name struct {
	// When the player changes their name, send a ChangeNameRequest; the
	// server will respond with a new TeamInfoResponse to everyone. If the
	// name is already in use, the server will respond with an ErrorResponse
	// and the name will not be changed.
	Name *ChangeNameRequest `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

type TeamStreamClientMessage_Word struct {
	// When the player guesses a word, send this request; the server will
	// send a WordGuessedResponse to all other players.
	Word *WordGuessedRequest `protobuf:"bytes,4,opt,name=word,proto3,oneof"`
}

func (*TeamStreamClientMessage_Create) isTeamStreamClientMessage_Request() {}

func (*TeamStreamClientMessage_Join) isTeamStreamClientMessage_Request() {}

func (*TeamStreamClientMessage_Name) isTeamStreamClientMessage_Request() {}

func (*TeamStreamClientMessage_Word) isTeamStreamClientMessage_Request() {}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TeamInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	YourName string   `protobuf:"bytes,2,opt,name=your_name,json=yourName,proto3" json:"your_name,omitempty"`
	Players  []string `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *TeamInfoResponse) Reset() {
	*x = TeamInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamInfoResponse) ProtoMessage() {}

func (x *TeamInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamInfoResponse.ProtoReflect.Descriptor instead.
func (*TeamInfoResponse) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{6}
}

func (x *TeamInfoResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TeamInfoResponse) GetYourName() string {
	if x != nil {
		return x.YourName
	}
	return ""
}

func (x *TeamInfoResponse) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

type WordGuessedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player string `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Word   string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *WordGuessedResponse) Reset() {
	*x = WordGuessedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordGuessedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordGuessedResponse) ProtoMessage() {}

func (x *WordGuessedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordGuessedResponse.ProtoReflect.Descriptor instead.
func (*WordGuessedResponse) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{7}
}

func (x *WordGuessedResponse) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *WordGuessedResponse) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type TeamStreamServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*TeamStreamServerMessage_Error
	//	*TeamStreamServerMessage_Team
	//	*TeamStreamServerMessage_Word
	Response isTeamStreamServerMessage_Response `protobuf_oneof:"response"`
}

func (x *TeamStreamServerMessage) Reset() {
	*x = TeamStreamServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vierkantle_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamStreamServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamStreamServerMessage) ProtoMessage() {}

func (x *TeamStreamServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_vierkantle_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamStreamServerMessage.ProtoReflect.Descriptor instead.
func (*TeamStreamServerMessage) Descriptor() ([]byte, []int) {
	return file_vierkantle_proto_rawDescGZIP(), []int{8}
}

func (m *TeamStreamServerMessage) GetResponse() isTeamStreamServerMessage_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *TeamStreamServerMessage) GetError() *ErrorResponse {
	if x, ok := x.GetResponse().(*TeamStreamServerMessage_Error); ok {
		return x.Error
	}
	return nil
}

func (x *TeamStreamServerMessage) GetTeam() *TeamInfoResponse {
	if x, ok := x.GetResponse().(*TeamStreamServerMessage_Team); ok {
		return x.Team
	}
	return nil
}

func (x *TeamStreamServerMessage) GetWord() *WordGuessedResponse {
	if x, ok := x.GetResponse().(*TeamStreamServerMessage_Word); ok {
		return x.Word
	}
	return nil
}

type isTeamStreamServerMessage_Response interface {
	isTeamStreamServerMessage_Response()
}

type TeamStreamServerMessage_Error struct {
	Error *ErrorResponse `protobuf:"bytes,1,opt,name=error,proto3,oneof"`
}

type TeamStreamServerMessage_Team struct {
	Team *TeamInfoResponse `protobuf:"bytes,2,opt,name=team,proto3,oneof"`
}

type TeamStreamServerMessage_Word struct {
	Word *WordGuessedResponse `protobuf:"bytes,3,opt,name=word,proto3,oneof"`
}

func (*TeamStreamServerMessage_Error) isTeamStreamServerMessage_Response() {}

func (*TeamStreamServerMessage_Team) isTeamStreamServerMessage_Response() {}

func (*TeamStreamServerMessage_Word) isTeamStreamServerMessage_Response() {}

var File_vierkantle_proto protoreflect.FileDescriptor

var file_vierkantle_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c,
	0x65, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0f, 0x4a, 0x6f,
	0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x28, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x87, 0x02, 0x0a, 0x17, 0x54,
	0x65, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72,
	0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72,
	0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x37, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x57,
	0x6f, 0x72, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x10, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x79, 0x6f, 0x75, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x79, 0x6f, 0x75, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x41, 0x0a, 0x13,
	0x57, 0x6f, 0x72, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0xcc, 0x01, 0x0a, 0x17, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x6c, 0x2e,
	0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x38, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72,
	0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x47, 0x75, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x77,
	0x0a, 0x11, 0x56, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0a, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x26, 0x2e, 0x6e, 0x6c, 0x2e, 0x76, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c,
	0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x26, 0x2e, 0x6e, 0x6c, 0x2e, 0x76,
	0x69, 0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x67, 0x69, 0x65, 0x6c, 0x65, 0x6e, 0x2f, 0x76, 0x69,
	0x65, 0x72, 0x6b, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vierkantle_proto_rawDescOnce sync.Once
	file_vierkantle_proto_rawDescData = file_vierkantle_proto_rawDesc
)

func file_vierkantle_proto_rawDescGZIP() []byte {
	file_vierkantle_proto_rawDescOnce.Do(func() {
		file_vierkantle_proto_rawDescData = protoimpl.X.CompressGZIP(file_vierkantle_proto_rawDescData)
	})
	return file_vierkantle_proto_rawDescData
}

var file_vierkantle_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vierkantle_proto_goTypes = []interface{}{
	(*CreateTeamRequest)(nil),       // 0: nl.vierkantle.CreateTeamRequest
	(*JoinTeamRequest)(nil),         // 1: nl.vierkantle.JoinTeamRequest
	(*ChangeNameRequest)(nil),       // 2: nl.vierkantle.ChangeNameRequest
	(*WordGuessedRequest)(nil),      // 3: nl.vierkantle.WordGuessedRequest
	(*TeamStreamClientMessage)(nil), // 4: nl.vierkantle.TeamStreamClientMessage
	(*ErrorResponse)(nil),           // 5: nl.vierkantle.ErrorResponse
	(*TeamInfoResponse)(nil),        // 6: nl.vierkantle.TeamInfoResponse
	(*WordGuessedResponse)(nil),     // 7: nl.vierkantle.WordGuessedResponse
	(*TeamStreamServerMessage)(nil), // 8: nl.vierkantle.TeamStreamServerMessage
}
var file_vierkantle_proto_depIdxs = []int32{
	0, // 0: nl.vierkantle.TeamStreamClientMessage.create:type_name -> nl.vierkantle.CreateTeamRequest
	1, // 1: nl.vierkantle.TeamStreamClientMessage.join:type_name -> nl.vierkantle.JoinTeamRequest
	2, // 2: nl.vierkantle.TeamStreamClientMessage.name:type_name -> nl.vierkantle.ChangeNameRequest
	3, // 3: nl.vierkantle.TeamStreamClientMessage.word:type_name -> nl.vierkantle.WordGuessedRequest
	5, // 4: nl.vierkantle.TeamStreamServerMessage.error:type_name -> nl.vierkantle.ErrorResponse
	6, // 5: nl.vierkantle.TeamStreamServerMessage.team:type_name -> nl.vierkantle.TeamInfoResponse
	7, // 6: nl.vierkantle.TeamStreamServerMessage.word:type_name -> nl.vierkantle.WordGuessedResponse
	4, // 7: nl.vierkantle.VierkantleService.TeamStream:input_type -> nl.vierkantle.TeamStreamClientMessage
	8, // 8: nl.vierkantle.VierkantleService.TeamStream:output_type -> nl.vierkantle.TeamStreamServerMessage
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_vierkantle_proto_init() }
func file_vierkantle_proto_init() {
	if File_vierkantle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vierkantle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamRequest); i {
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
		file_vierkantle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinTeamRequest); i {
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
		file_vierkantle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNameRequest); i {
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
		file_vierkantle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordGuessedRequest); i {
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
		file_vierkantle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamStreamClientMessage); i {
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
		file_vierkantle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
		file_vierkantle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamInfoResponse); i {
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
		file_vierkantle_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordGuessedResponse); i {
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
		file_vierkantle_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamStreamServerMessage); i {
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
	file_vierkantle_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*TeamStreamClientMessage_Create)(nil),
		(*TeamStreamClientMessage_Join)(nil),
		(*TeamStreamClientMessage_Name)(nil),
		(*TeamStreamClientMessage_Word)(nil),
	}
	file_vierkantle_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*TeamStreamServerMessage_Error)(nil),
		(*TeamStreamServerMessage_Team)(nil),
		(*TeamStreamServerMessage_Word)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vierkantle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vierkantle_proto_goTypes,
		DependencyIndexes: file_vierkantle_proto_depIdxs,
		MessageInfos:      file_vierkantle_proto_msgTypes,
	}.Build()
	File_vierkantle_proto = out.File
	file_vierkantle_proto_rawDesc = nil
	file_vierkantle_proto_goTypes = nil
	file_vierkantle_proto_depIdxs = nil
}
