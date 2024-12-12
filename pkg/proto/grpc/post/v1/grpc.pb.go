// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: post/v1/grpc.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	mi := &file_post_v1_grpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	mi := &file_post_v1_grpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	mi := &file_post_v1_grpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeletePostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BatchGetPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *BatchGetPostsRequest) Reset() {
	*x = BatchGetPostsRequest{}
	mi := &file_post_v1_grpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchGetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetPostsRequest) ProtoMessage() {}

func (x *BatchGetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetPostsRequest.ProtoReflect.Descriptor instead.
func (*BatchGetPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetPostsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BatchGetPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *BatchGetPostsResponse) Reset() {
	*x = BatchGetPostsResponse{}
	mi := &file_post_v1_grpc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchGetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetPostsResponse) ProtoMessage() {}

func (x *BatchGetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetPostsResponse.ProtoReflect.Descriptor instead.
func (*BatchGetPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *BatchGetPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type ListPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// A page token, received from a previous `ListPosts` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListPosts` must match
	// the call that provided the page token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The maximum number of posts to return. The service may return fewer than
	// this value. If unspecified, at most 100 items will be returned.
	// The maximum value is 100; values above 100 will be coerced to 100.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListPostsRequest) Reset() {
	*x = ListPostsRequest{}
	mi := &file_post_v1_grpc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsRequest) ProtoMessage() {}

func (x *ListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsRequest.ProtoReflect.Descriptor instead.
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *ListPostsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListPostsRequest) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListPostsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	mi := &file_post_v1_grpc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResponse.ProtoReflect.Descriptor instead.
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *ListPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *ListPostsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ListPostIdProjectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parents []string `protobuf:"bytes,1,rep,name=parents,proto3" json:"parents,omitempty"`
	// A page token, received from a previous `ListPostIdProjections` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListPostIdProjections` must match
	// the call that provided the page token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The maximum number of post projections to return. The service may return fewer than
	// this value. If unspecified, at most 100 items will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListPostIdProjectionsRequest) Reset() {
	*x = ListPostIdProjectionsRequest{}
	mi := &file_post_v1_grpc_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostIdProjectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostIdProjectionsRequest) ProtoMessage() {}

func (x *ListPostIdProjectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostIdProjectionsRequest.ProtoReflect.Descriptor instead.
func (*ListPostIdProjectionsRequest) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *ListPostIdProjectionsRequest) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

func (x *ListPostIdProjectionsRequest) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListPostIdProjectionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListPostIdProjectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostIdProjections []*PostIdProjection `protobuf:"bytes,1,rep,name=post_id_projections,json=postIdProjections,proto3" json:"post_id_projections,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPostIdProjectionsResponse) Reset() {
	*x = ListPostIdProjectionsResponse{}
	mi := &file_post_v1_grpc_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostIdProjectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostIdProjectionsResponse) ProtoMessage() {}

func (x *ListPostIdProjectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostIdProjectionsResponse.ProtoReflect.Descriptor instead.
func (*ListPostIdProjectionsResponse) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *ListPostIdProjectionsResponse) GetPostIdProjections() []*PostIdProjection {
	if x != nil {
		return x.PostIdProjections
	}
	return nil
}

func (x *ListPostIdProjectionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is unique and sortable post identifier.
	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text        string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	AuthorId    string                 `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	Deleted     bool                   `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_post_v1_grpc_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Post) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Post) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Post) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type PostIdProjection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *PostIdProjection) Reset() {
	*x = PostIdProjection{}
	mi := &file_post_v1_grpc_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostIdProjection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostIdProjection) ProtoMessage() {}

func (x *PostIdProjection) ProtoReflect() protoreflect.Message {
	mi := &file_post_v1_grpc_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostIdProjection.ProtoReflect.Descriptor instead.
func (*PostIdProjection) Descriptor() ([]byte, []int) {
	return file_post_v1_grpc_proto_rawDescGZIP(), []int{10}
}

func (x *PostIdProjection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostIdProjection) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

var File_post_v1_grpc_proto protoreflect.FileDescriptor

var file_post_v1_grpc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x14, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x3c, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x22, 0x6f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x60, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x04, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x10,
	0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x32, 0xb7, 0x03,
	0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x0d, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x66, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_v1_grpc_proto_rawDescOnce sync.Once
	file_post_v1_grpc_proto_rawDescData = file_post_v1_grpc_proto_rawDesc
)

func file_post_v1_grpc_proto_rawDescGZIP() []byte {
	file_post_v1_grpc_proto_rawDescOnce.Do(func() {
		file_post_v1_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_v1_grpc_proto_rawDescData)
	})
	return file_post_v1_grpc_proto_rawDescData
}

var file_post_v1_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_post_v1_grpc_proto_goTypes = []any{
	(*CreatePostRequest)(nil),             // 0: post.v1.CreatePostRequest
	(*GetPostRequest)(nil),                // 1: post.v1.GetPostRequest
	(*DeletePostRequest)(nil),             // 2: post.v1.DeletePostRequest
	(*BatchGetPostsRequest)(nil),          // 3: post.v1.BatchGetPostsRequest
	(*BatchGetPostsResponse)(nil),         // 4: post.v1.BatchGetPostsResponse
	(*ListPostsRequest)(nil),              // 5: post.v1.ListPostsRequest
	(*ListPostsResponse)(nil),             // 6: post.v1.ListPostsResponse
	(*ListPostIdProjectionsRequest)(nil),  // 7: post.v1.ListPostIdProjectionsRequest
	(*ListPostIdProjectionsResponse)(nil), // 8: post.v1.ListPostIdProjectionsResponse
	(*Post)(nil),                          // 9: post.v1.Post
	(*PostIdProjection)(nil),              // 10: post.v1.PostIdProjection
	(*timestamppb.Timestamp)(nil),         // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                 // 12: google.protobuf.Empty
}
var file_post_v1_grpc_proto_depIdxs = []int32{
	9,  // 0: post.v1.CreatePostRequest.post:type_name -> post.v1.Post
	9,  // 1: post.v1.BatchGetPostsResponse.posts:type_name -> post.v1.Post
	9,  // 2: post.v1.ListPostsResponse.posts:type_name -> post.v1.Post
	10, // 3: post.v1.ListPostIdProjectionsResponse.post_id_projections:type_name -> post.v1.PostIdProjection
	11, // 4: post.v1.Post.updated_time:type_name -> google.protobuf.Timestamp
	0,  // 5: post.v1.PostService.CreatePost:input_type -> post.v1.CreatePostRequest
	1,  // 6: post.v1.PostService.GetPost:input_type -> post.v1.GetPostRequest
	2,  // 7: post.v1.PostService.DeletePost:input_type -> post.v1.DeletePostRequest
	3,  // 8: post.v1.PostService.BatchGetPosts:input_type -> post.v1.BatchGetPostsRequest
	5,  // 9: post.v1.PostService.ListPosts:input_type -> post.v1.ListPostsRequest
	7,  // 10: post.v1.PostService.ListPostIdProjections:input_type -> post.v1.ListPostIdProjectionsRequest
	9,  // 11: post.v1.PostService.CreatePost:output_type -> post.v1.Post
	9,  // 12: post.v1.PostService.GetPost:output_type -> post.v1.Post
	12, // 13: post.v1.PostService.DeletePost:output_type -> google.protobuf.Empty
	4,  // 14: post.v1.PostService.BatchGetPosts:output_type -> post.v1.BatchGetPostsResponse
	6,  // 15: post.v1.PostService.ListPosts:output_type -> post.v1.ListPostsResponse
	8,  // 16: post.v1.PostService.ListPostIdProjections:output_type -> post.v1.ListPostIdProjectionsResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_post_v1_grpc_proto_init() }
func file_post_v1_grpc_proto_init() {
	if File_post_v1_grpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_v1_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_v1_grpc_proto_goTypes,
		DependencyIndexes: file_post_v1_grpc_proto_depIdxs,
		MessageInfos:      file_post_v1_grpc_proto_msgTypes,
	}.Build()
	File_post_v1_grpc_proto = out.File
	file_post_v1_grpc_proto_rawDesc = nil
	file_post_v1_grpc_proto_goTypes = nil
	file_post_v1_grpc_proto_depIdxs = nil
}
