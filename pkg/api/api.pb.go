// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListOrder int32

const (
	ListOrder_LIST_ORDER_UNSPECIFIED ListOrder = 0
	ListOrder_LIST_ORDER_ASC         ListOrder = 1
	ListOrder_LIST_ORDER_DESC        ListOrder = 2
)

// Enum value maps for ListOrder.
var (
	ListOrder_name = map[int32]string{
		0: "LIST_ORDER_UNSPECIFIED",
		1: "LIST_ORDER_ASC",
		2: "LIST_ORDER_DESC",
	}
	ListOrder_value = map[string]int32{
		"LIST_ORDER_UNSPECIFIED": 0,
		"LIST_ORDER_ASC":         1,
		"LIST_ORDER_DESC":        2,
	}
)

func (x ListOrder) Enum() *ListOrder {
	p := new(ListOrder)
	*p = x
	return p
}

func (x ListOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (ListOrder) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x ListOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListOrder.Descriptor instead.
func (ListOrder) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Year  int32  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

// MovieCreate endpoint messages
type MovieCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year  int32  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *MovieCreateRequest) Reset() {
	*x = MovieCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieCreateRequest) ProtoMessage() {}

func (x *MovieCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieCreateRequest.ProtoReflect.Descriptor instead.
func (*MovieCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *MovieCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieCreateRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

// MovieList endpoint messages
type GatewayMovieListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64     `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Order ListOrder `protobuf:"varint,2,opt,name=order,proto3,enum=ozon.dev.homework.api.ListOrder" json:"order,omitempty"`
}

func (x *GatewayMovieListRequest) Reset() {
	*x = GatewayMovieListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieListRequest) ProtoMessage() {}

func (x *GatewayMovieListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieListRequest.ProtoReflect.Descriptor instead.
func (*GatewayMovieListRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayMovieListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GatewayMovieListRequest) GetOrder() ListOrder {
	if x != nil {
		return x.Order
	}
	return ListOrder_LIST_ORDER_UNSPECIFIED
}

type StorageMovieListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Order  string `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *StorageMovieListRequest) Reset() {
	*x = StorageMovieListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageMovieListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageMovieListRequest) ProtoMessage() {}

func (x *StorageMovieListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageMovieListRequest.ProtoReflect.Descriptor instead.
func (*StorageMovieListRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *StorageMovieListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *StorageMovieListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *StorageMovieListRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type MovieListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie []*Movie `protobuf:"bytes,1,rep,name=movie,proto3" json:"movie,omitempty"`
}

func (x *MovieListResponse) Reset() {
	*x = MovieListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieListResponse) ProtoMessage() {}

func (x *MovieListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieListResponse.ProtoReflect.Descriptor instead.
func (*MovieListResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *MovieListResponse) GetMovie() []*Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

// MovieUpdate endpoint messages
type MovieUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *MovieUpdateRequest) Reset() {
	*x = MovieUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieUpdateRequest) ProtoMessage() {}

func (x *MovieUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieUpdateRequest.ProtoReflect.Descriptor instead.
func (*MovieUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *MovieUpdateRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

// MovieDelete endpoint messages
type MovieDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MovieDeleteRequest) Reset() {
	*x = MovieDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDeleteRequest) ProtoMessage() {}

func (x *MovieDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDeleteRequest.ProtoReflect.Descriptor instead.
func (*MovieDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *MovieDeleteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// MovieGetOne endpoint messages
type MovieGetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MovieGetOneRequest) Reset() {
	*x = MovieGetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieGetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieGetOneRequest) ProtoMessage() {}

func (x *MovieGetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieGetOneRequest.ProtoReflect.Descriptor instead.
func (*MovieGetOneRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *MovieGetOneRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MovieGetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *MovieGetOneResponse) Reset() {
	*x = MovieGetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieGetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieGetOneResponse) ProtoMessage() {}

func (x *MovieGetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieGetOneResponse.ProtoReflect.Descriptor instead.
func (*MovieGetOneResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *MovieGetOneResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x22, 0x3e, 0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x22, 0x67, 0x0a, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x36, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x17, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f,
	0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x22, 0x48, 0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x24, 0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2a, 0x50, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x16, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c,
	0x49, 0x53, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x02, 0x32, 0xcc, 0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x12, 0x66, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x29, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x0b,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x0b, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x12, 0x29, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x32, 0xd8, 0x03, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x52, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x29,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2e, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x52,
	0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x29, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47,
	0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x29, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16,
	0x5a, 0x14, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_goTypes = []interface{}{
	(ListOrder)(0),                  // 0: ozon.dev.homework.api.ListOrder
	(*Movie)(nil),                   // 1: ozon.dev.homework.api.Movie
	(*MovieCreateRequest)(nil),      // 2: ozon.dev.homework.api.MovieCreateRequest
	(*GatewayMovieListRequest)(nil), // 3: ozon.dev.homework.api.GatewayMovieListRequest
	(*StorageMovieListRequest)(nil), // 4: ozon.dev.homework.api.StorageMovieListRequest
	(*MovieListResponse)(nil),       // 5: ozon.dev.homework.api.MovieListResponse
	(*MovieUpdateRequest)(nil),      // 6: ozon.dev.homework.api.MovieUpdateRequest
	(*MovieDeleteRequest)(nil),      // 7: ozon.dev.homework.api.MovieDeleteRequest
	(*MovieGetOneRequest)(nil),      // 8: ozon.dev.homework.api.MovieGetOneRequest
	(*MovieGetOneResponse)(nil),     // 9: ozon.dev.homework.api.MovieGetOneResponse
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: ozon.dev.homework.api.GatewayMovieListRequest.order:type_name -> ozon.dev.homework.api.ListOrder
	1,  // 1: ozon.dev.homework.api.MovieListResponse.movie:type_name -> ozon.dev.homework.api.Movie
	1,  // 2: ozon.dev.homework.api.MovieUpdateRequest.movie:type_name -> ozon.dev.homework.api.Movie
	1,  // 3: ozon.dev.homework.api.MovieGetOneResponse.movie:type_name -> ozon.dev.homework.api.Movie
	2,  // 4: ozon.dev.homework.api.Gateway.MovieCreate:input_type -> ozon.dev.homework.api.MovieCreateRequest
	3,  // 5: ozon.dev.homework.api.Gateway.MovieList:input_type -> ozon.dev.homework.api.GatewayMovieListRequest
	6,  // 6: ozon.dev.homework.api.Gateway.MovieUpdate:input_type -> ozon.dev.homework.api.MovieUpdateRequest
	7,  // 7: ozon.dev.homework.api.Gateway.MovieDelete:input_type -> ozon.dev.homework.api.MovieDeleteRequest
	8,  // 8: ozon.dev.homework.api.Gateway.MovieGetOne:input_type -> ozon.dev.homework.api.MovieGetOneRequest
	2,  // 9: ozon.dev.homework.api.Storage.MovieCreate:input_type -> ozon.dev.homework.api.MovieCreateRequest
	4,  // 10: ozon.dev.homework.api.Storage.MovieList:input_type -> ozon.dev.homework.api.StorageMovieListRequest
	6,  // 11: ozon.dev.homework.api.Storage.MovieUpdate:input_type -> ozon.dev.homework.api.MovieUpdateRequest
	7,  // 12: ozon.dev.homework.api.Storage.MovieDelete:input_type -> ozon.dev.homework.api.MovieDeleteRequest
	8,  // 13: ozon.dev.homework.api.Storage.MovieGetOne:input_type -> ozon.dev.homework.api.MovieGetOneRequest
	10, // 14: ozon.dev.homework.api.Gateway.MovieCreate:output_type -> google.protobuf.Empty
	5,  // 15: ozon.dev.homework.api.Gateway.MovieList:output_type -> ozon.dev.homework.api.MovieListResponse
	10, // 16: ozon.dev.homework.api.Gateway.MovieUpdate:output_type -> google.protobuf.Empty
	10, // 17: ozon.dev.homework.api.Gateway.MovieDelete:output_type -> google.protobuf.Empty
	9,  // 18: ozon.dev.homework.api.Gateway.MovieGetOne:output_type -> ozon.dev.homework.api.MovieGetOneResponse
	10, // 19: ozon.dev.homework.api.Storage.MovieCreate:output_type -> google.protobuf.Empty
	5,  // 20: ozon.dev.homework.api.Storage.MovieList:output_type -> ozon.dev.homework.api.MovieListResponse
	10, // 21: ozon.dev.homework.api.Storage.MovieUpdate:output_type -> google.protobuf.Empty
	10, // 22: ozon.dev.homework.api.Storage.MovieDelete:output_type -> google.protobuf.Empty
	9,  // 23: ozon.dev.homework.api.Storage.MovieGetOne:output_type -> ozon.dev.homework.api.MovieGetOneResponse
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieCreateRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieListRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageMovieListRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieListResponse); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieUpdateRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDeleteRequest); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieGetOneRequest); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieGetOneResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
