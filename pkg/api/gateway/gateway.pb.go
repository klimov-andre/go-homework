// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: gateway/gateway.proto

package gateway

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
	return file_gateway_gateway_proto_enumTypes[0].Descriptor()
}

func (ListOrder) Type() protoreflect.EnumType {
	return &file_gateway_gateway_proto_enumTypes[0]
}

func (x ListOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListOrder.Descriptor instead.
func (ListOrder) EnumDescriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{0}
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
		mi := &file_gateway_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[0]
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
	return file_gateway_gateway_proto_rawDescGZIP(), []int{0}
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

// GatewayMovieCreateRequest endpoint messages
type GatewayMovieCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year  int32  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *GatewayMovieCreateRequest) Reset() {
	*x = GatewayMovieCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieCreateRequest) ProtoMessage() {}

func (x *GatewayMovieCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieCreateRequest.ProtoReflect.Descriptor instead.
func (*GatewayMovieCreateRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayMovieCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GatewayMovieCreateRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type GatewayMovieCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GatewayMovieCreateResponse) Reset() {
	*x = GatewayMovieCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieCreateResponse) ProtoMessage() {}

func (x *GatewayMovieCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieCreateResponse.ProtoReflect.Descriptor instead.
func (*GatewayMovieCreateResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayMovieCreateResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// MovieList endpoint messages
type GatewayMovieListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64     `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Order ListOrder `protobuf:"varint,2,opt,name=order,proto3,enum=ozon.dev.homework.api.gateway.ListOrder" json:"order,omitempty"`
}

func (x *GatewayMovieListRequest) Reset() {
	*x = GatewayMovieListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieListRequest) ProtoMessage() {}

func (x *GatewayMovieListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[3]
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
	return file_gateway_gateway_proto_rawDescGZIP(), []int{3}
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

type GatewayMovieListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie []*Movie `protobuf:"bytes,1,rep,name=movie,proto3" json:"movie,omitempty"`
}

func (x *GatewayMovieListResponse) Reset() {
	*x = GatewayMovieListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieListResponse) ProtoMessage() {}

func (x *GatewayMovieListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieListResponse.ProtoReflect.Descriptor instead.
func (*GatewayMovieListResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *GatewayMovieListResponse) GetMovie() []*Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

// GatewayMovieUpdate endpoint messages
type GatewayMovieUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *GatewayMovieUpdateRequest) Reset() {
	*x = GatewayMovieUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieUpdateRequest) ProtoMessage() {}

func (x *GatewayMovieUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieUpdateRequest.ProtoReflect.Descriptor instead.
func (*GatewayMovieUpdateRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *GatewayMovieUpdateRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

// GatewayMovieDelete endpoint messages
type GatewayMovieDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GatewayMovieDeleteRequest) Reset() {
	*x = GatewayMovieDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieDeleteRequest) ProtoMessage() {}

func (x *GatewayMovieDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieDeleteRequest.ProtoReflect.Descriptor instead.
func (*GatewayMovieDeleteRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *GatewayMovieDeleteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// GatewayMovieGetOne endpoint messages
type GatewayMovieGetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GatewayMovieGetOneRequest) Reset() {
	*x = GatewayMovieGetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieGetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieGetOneRequest) ProtoMessage() {}

func (x *GatewayMovieGetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieGetOneRequest.ProtoReflect.Descriptor instead.
func (*GatewayMovieGetOneRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *GatewayMovieGetOneRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GatewayMovieGetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *GatewayMovieGetOneResponse) Reset() {
	*x = GatewayMovieGetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayMovieGetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayMovieGetOneResponse) ProtoMessage() {}

func (x *GatewayMovieGetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayMovieGetOneResponse.ProtoReflect.Descriptor instead.
func (*GatewayMovieGetOneResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *GatewayMovieGetOneResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

var File_gateway_gateway_proto protoreflect.FileDescriptor

var file_gateway_gateway_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x41, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x22, 0x45, 0x0a, 0x19, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x2c, 0x0a, 0x1a, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x17, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x18, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x22, 0x57, 0x0a, 0x19, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x19, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x1a, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2a,
	0x50, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x16,
	0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x49, 0x53, 0x54,
	0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10,
	0x02, 0x32, 0xd4, 0x08, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x98, 0x01,
	0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x38, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x12, 0x38, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x98, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x36, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x7c, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x38, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x82,
	0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x64, 0x12, 0x38, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10,
	0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x38, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x11,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x64, 0x12, 0x38, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x32,
	0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9a, 0x01, 0x0a, 0x0b,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x38, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_gateway_proto_rawDescOnce sync.Once
	file_gateway_gateway_proto_rawDescData = file_gateway_gateway_proto_rawDesc
)

func file_gateway_gateway_proto_rawDescGZIP() []byte {
	file_gateway_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_gateway_proto_rawDescData)
	})
	return file_gateway_gateway_proto_rawDescData
}

var file_gateway_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gateway_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gateway_gateway_proto_goTypes = []interface{}{
	(ListOrder)(0),                     // 0: ozon.dev.homework.api.gateway.ListOrder
	(*Movie)(nil),                      // 1: ozon.dev.homework.api.gateway.Movie
	(*GatewayMovieCreateRequest)(nil),  // 2: ozon.dev.homework.api.gateway.GatewayMovieCreateRequest
	(*GatewayMovieCreateResponse)(nil), // 3: ozon.dev.homework.api.gateway.GatewayMovieCreateResponse
	(*GatewayMovieListRequest)(nil),    // 4: ozon.dev.homework.api.gateway.GatewayMovieListRequest
	(*GatewayMovieListResponse)(nil),   // 5: ozon.dev.homework.api.gateway.GatewayMovieListResponse
	(*GatewayMovieUpdateRequest)(nil),  // 6: ozon.dev.homework.api.gateway.GatewayMovieUpdateRequest
	(*GatewayMovieDeleteRequest)(nil),  // 7: ozon.dev.homework.api.gateway.GatewayMovieDeleteRequest
	(*GatewayMovieGetOneRequest)(nil),  // 8: ozon.dev.homework.api.gateway.GatewayMovieGetOneRequest
	(*GatewayMovieGetOneResponse)(nil), // 9: ozon.dev.homework.api.gateway.GatewayMovieGetOneResponse
	(*emptypb.Empty)(nil),              // 10: google.protobuf.Empty
}
var file_gateway_gateway_proto_depIdxs = []int32{
	0,  // 0: ozon.dev.homework.api.gateway.GatewayMovieListRequest.order:type_name -> ozon.dev.homework.api.gateway.ListOrder
	1,  // 1: ozon.dev.homework.api.gateway.GatewayMovieListResponse.movie:type_name -> ozon.dev.homework.api.gateway.Movie
	1,  // 2: ozon.dev.homework.api.gateway.GatewayMovieUpdateRequest.movie:type_name -> ozon.dev.homework.api.gateway.Movie
	1,  // 3: ozon.dev.homework.api.gateway.GatewayMovieGetOneResponse.movie:type_name -> ozon.dev.homework.api.gateway.Movie
	2,  // 4: ozon.dev.homework.api.gateway.Gateway.MovieCreate:input_type -> ozon.dev.homework.api.gateway.GatewayMovieCreateRequest
	2,  // 5: ozon.dev.homework.api.gateway.Gateway.MovieCreateQueued:input_type -> ozon.dev.homework.api.gateway.GatewayMovieCreateRequest
	4,  // 6: ozon.dev.homework.api.gateway.Gateway.MovieList:input_type -> ozon.dev.homework.api.gateway.GatewayMovieListRequest
	6,  // 7: ozon.dev.homework.api.gateway.Gateway.MovieUpdate:input_type -> ozon.dev.homework.api.gateway.GatewayMovieUpdateRequest
	6,  // 8: ozon.dev.homework.api.gateway.Gateway.MovieUpdateQueued:input_type -> ozon.dev.homework.api.gateway.GatewayMovieUpdateRequest
	7,  // 9: ozon.dev.homework.api.gateway.Gateway.MovieDelete:input_type -> ozon.dev.homework.api.gateway.GatewayMovieDeleteRequest
	7,  // 10: ozon.dev.homework.api.gateway.Gateway.MovieDeleteQueued:input_type -> ozon.dev.homework.api.gateway.GatewayMovieDeleteRequest
	8,  // 11: ozon.dev.homework.api.gateway.Gateway.MovieGetOne:input_type -> ozon.dev.homework.api.gateway.GatewayMovieGetOneRequest
	3,  // 12: ozon.dev.homework.api.gateway.Gateway.MovieCreate:output_type -> ozon.dev.homework.api.gateway.GatewayMovieCreateResponse
	10, // 13: ozon.dev.homework.api.gateway.Gateway.MovieCreateQueued:output_type -> google.protobuf.Empty
	5,  // 14: ozon.dev.homework.api.gateway.Gateway.MovieList:output_type -> ozon.dev.homework.api.gateway.GatewayMovieListResponse
	10, // 15: ozon.dev.homework.api.gateway.Gateway.MovieUpdate:output_type -> google.protobuf.Empty
	10, // 16: ozon.dev.homework.api.gateway.Gateway.MovieUpdateQueued:output_type -> google.protobuf.Empty
	10, // 17: ozon.dev.homework.api.gateway.Gateway.MovieDelete:output_type -> google.protobuf.Empty
	10, // 18: ozon.dev.homework.api.gateway.Gateway.MovieDeleteQueued:output_type -> google.protobuf.Empty
	9,  // 19: ozon.dev.homework.api.gateway.Gateway.MovieGetOne:output_type -> ozon.dev.homework.api.gateway.GatewayMovieGetOneResponse
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_gateway_gateway_proto_init() }
func file_gateway_gateway_proto_init() {
	if File_gateway_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_gateway_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieCreateRequest); i {
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
		file_gateway_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieCreateResponse); i {
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
		file_gateway_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_gateway_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieListResponse); i {
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
		file_gateway_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieUpdateRequest); i {
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
		file_gateway_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieDeleteRequest); i {
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
		file_gateway_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieGetOneRequest); i {
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
		file_gateway_gateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayMovieGetOneResponse); i {
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
			RawDescriptor: file_gateway_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_proto_depIdxs,
		EnumInfos:         file_gateway_gateway_proto_enumTypes,
		MessageInfos:      file_gateway_gateway_proto_msgTypes,
	}.Build()
	File_gateway_gateway_proto = out.File
	file_gateway_gateway_proto_rawDesc = nil
	file_gateway_gateway_proto_goTypes = nil
	file_gateway_gateway_proto_depIdxs = nil
}
