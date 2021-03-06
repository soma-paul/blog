// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: practice/blog/gunk/v1/article/all.proto

package article

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Articles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32                  `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=Title,json=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=Description,json=description,proto3" json:"description,omitempty"`
	Author      string                 `protobuf:"bytes,4,opt,name=Author,json=author,proto3" json:"author,omitempty"`
	UserID      int32                  `protobuf:"varint,5,opt,name=UserID,json=user_id,proto3" json:"user_id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,json=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=UpdatedAt,json=updated_at,proto3" json:"updated_at,omitempty"`
}

func (x *Articles) Reset() {
	*x = Articles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Articles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Articles) ProtoMessage() {}

func (x *Articles) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Articles.ProtoReflect.Descriptor instead.
func (*Articles) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{0}
}

func (x *Articles) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Articles) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Articles) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Articles) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Articles) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Articles) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Articles) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"id,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{1}
}

func (x *GetArticleRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Articles `protobuf:"bytes,1,opt,name=Article,json=article,proto3" json:"article,omitempty"`
}

func (x *GetArticleResponse) Reset() {
	*x = GetArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleResponse) ProtoMessage() {}

func (x *GetArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleResponse.ProtoReflect.Descriptor instead.
func (*GetArticleResponse) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticleResponse) GetArticle() *Articles {
	if x != nil {
		return x.Article
	}
	return nil
}

type GetAllArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllArticleRequest) Reset() {
	*x = GetAllArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllArticleRequest) ProtoMessage() {}

func (x *GetAllArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllArticleRequest.ProtoReflect.Descriptor instead.
func (*GetAllArticleRequest) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{3}
}

type GetAllArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Articles `protobuf:"bytes,1,rep,name=Articles,json=articles,proto3" json:"articles,omitempty"`
}

func (x *GetAllArticleResponse) Reset() {
	*x = GetAllArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllArticleResponse) ProtoMessage() {}

func (x *GetAllArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllArticleResponse.ProtoReflect.Descriptor instead.
func (*GetAllArticleResponse) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllArticleResponse) GetArticles() []*Articles {
	if x != nil {
		return x.Articles
	}
	return nil
}

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Articles `protobuf:"bytes,1,opt,name=Article,json=article,proto3" json:"article,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{5}
}

func (x *CreateArticleRequest) GetArticle() *Articles {
	if x != nil {
		return x.Article
	}
	return nil
}

type CreateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"id,omitempty"`
}

func (x *CreateArticleResponse) Reset() {
	*x = CreateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleResponse) ProtoMessage() {}

func (x *CreateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleResponse.ProtoReflect.Descriptor instead.
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{6}
}

func (x *CreateArticleResponse) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type UpdateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Articles `protobuf:"bytes,1,opt,name=Article,json=article,proto3" json:"article,omitempty"`
}

func (x *UpdateArticleRequest) Reset() {
	*x = UpdateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleRequest) ProtoMessage() {}

func (x *UpdateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleRequest.ProtoReflect.Descriptor instead.
func (*UpdateArticleRequest) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateArticleRequest) GetArticle() *Articles {
	if x != nil {
		return x.Article
	}
	return nil
}

type UpdateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateArticleResponse) Reset() {
	*x = UpdateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArticleResponse) ProtoMessage() {}

func (x *UpdateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArticleResponse.ProtoReflect.Descriptor instead.
func (*UpdateArticleResponse) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{8}
}

type DeleteArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"id,omitempty"`
}

func (x *DeleteArticleRequest) Reset() {
	*x = DeleteArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleRequest) ProtoMessage() {}

func (x *DeleteArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleRequest.ProtoReflect.Descriptor instead.
func (*DeleteArticleRequest) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteArticleRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"id,omitempty"`
}

func (x *DeleteArticleResponse) Reset() {
	*x = DeleteArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleResponse) ProtoMessage() {}

func (x *DeleteArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_practice_blog_gunk_v1_article_all_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleResponse.ProtoReflect.Descriptor instead.
func (*DeleteArticleResponse) Descriptor() ([]byte, []int) {
	return file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteArticleResponse) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_practice_blog_gunk_v1_article_all_proto protoreflect.FileDescriptor

var file_practice_blog_gunk_v1_article_all_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f,
	0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f,
	0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2c,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x23, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x45, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x37, 0x0a, 0x11, 0x67,
	0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x06, 0x08, 0x00,
	0x10, 0x00, 0x18, 0x00, 0x22, 0x55, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x1e, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x5a, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x3a,
	0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x57, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x37, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x22, 0x3b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x57, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3a, 0x06,
	0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x1f, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a,
	0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x3a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x06, 0x08, 0x00, 0x10,
	0x00, 0x18, 0x00, 0x22, 0x3b, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x32, 0xd2, 0x03, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x51, 0x0a, 0x0a,
	0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12,
	0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x5a, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02,
	0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x1a, 0x03, 0x88, 0x02, 0x00, 0x42, 0x40, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x25, 0x70, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01,
	0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_practice_blog_gunk_v1_article_all_proto_rawDescOnce sync.Once
	file_practice_blog_gunk_v1_article_all_proto_rawDescData = file_practice_blog_gunk_v1_article_all_proto_rawDesc
)

func file_practice_blog_gunk_v1_article_all_proto_rawDescGZIP() []byte {
	file_practice_blog_gunk_v1_article_all_proto_rawDescOnce.Do(func() {
		file_practice_blog_gunk_v1_article_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_practice_blog_gunk_v1_article_all_proto_rawDescData)
	})
	return file_practice_blog_gunk_v1_article_all_proto_rawDescData
}

var file_practice_blog_gunk_v1_article_all_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_practice_blog_gunk_v1_article_all_proto_goTypes = []interface{}{
	(*Articles)(nil),              // 0: article.Articles
	(*GetArticleRequest)(nil),     // 1: article.getArticleRequest
	(*GetArticleResponse)(nil),    // 2: article.getArticleResponse
	(*GetAllArticleRequest)(nil),  // 3: article.GetAllArticleRequest
	(*GetAllArticleResponse)(nil), // 4: article.GetAllArticleResponse
	(*CreateArticleRequest)(nil),  // 5: article.CreateArticleRequest
	(*CreateArticleResponse)(nil), // 6: article.CreateArticleResponse
	(*UpdateArticleRequest)(nil),  // 7: article.UpdateArticleRequest
	(*UpdateArticleResponse)(nil), // 8: article.UpdateArticleResponse
	(*DeleteArticleRequest)(nil),  // 9: article.DeleteArticleRequest
	(*DeleteArticleResponse)(nil), // 10: article.DeleteArticleResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_practice_blog_gunk_v1_article_all_proto_depIdxs = []int32{
	11, // 0: article.Articles.CreatedAt:type_name -> google.protobuf.Timestamp
	11, // 1: article.Articles.UpdatedAt:type_name -> google.protobuf.Timestamp
	0,  // 2: article.getArticleResponse.Article:type_name -> article.Articles
	0,  // 3: article.GetAllArticleResponse.Articles:type_name -> article.Articles
	0,  // 4: article.CreateArticleRequest.Article:type_name -> article.Articles
	0,  // 5: article.UpdateArticleRequest.Article:type_name -> article.Articles
	1,  // 6: article.Article.getArticle:input_type -> article.getArticleRequest
	3,  // 7: article.Article.GetAllArticles:input_type -> article.GetAllArticleRequest
	5,  // 8: article.Article.CreateArticle:input_type -> article.CreateArticleRequest
	7,  // 9: article.Article.UpdateArticle:input_type -> article.UpdateArticleRequest
	9,  // 10: article.Article.DeleteArticle:input_type -> article.DeleteArticleRequest
	2,  // 11: article.Article.getArticle:output_type -> article.getArticleResponse
	4,  // 12: article.Article.GetAllArticles:output_type -> article.GetAllArticleResponse
	6,  // 13: article.Article.CreateArticle:output_type -> article.CreateArticleResponse
	8,  // 14: article.Article.UpdateArticle:output_type -> article.UpdateArticleResponse
	10, // 15: article.Article.DeleteArticle:output_type -> article.DeleteArticleResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_practice_blog_gunk_v1_article_all_proto_init() }
func file_practice_blog_gunk_v1_article_all_proto_init() {
	if File_practice_blog_gunk_v1_article_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Articles); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleRequest); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleResponse); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllArticleRequest); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllArticleResponse); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleResponse); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleRequest); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArticleResponse); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleRequest); i {
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
		file_practice_blog_gunk_v1_article_all_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleResponse); i {
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
			RawDescriptor: file_practice_blog_gunk_v1_article_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_practice_blog_gunk_v1_article_all_proto_goTypes,
		DependencyIndexes: file_practice_blog_gunk_v1_article_all_proto_depIdxs,
		MessageInfos:      file_practice_blog_gunk_v1_article_all_proto_msgTypes,
	}.Build()
	File_practice_blog_gunk_v1_article_all_proto = out.File
	file_practice_blog_gunk_v1_article_all_proto_rawDesc = nil
	file_practice_blog_gunk_v1_article_all_proto_goTypes = nil
	file_practice_blog_gunk_v1_article_all_proto_depIdxs = nil
}
