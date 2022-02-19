// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: book_proto/crud.proto

package book_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author       string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Pages        int32  `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	CreationDate uint64 `protobuf:"varint,5,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	UploadTime   uint64 `protobuf:"varint,6,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_book_proto_crud_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *Book) GetCreationDate() uint64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *Book) GetUploadTime() uint64 {
	if x != nil {
		return x.UploadTime
	}
	return 0
}

type BookData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId string `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BookData) Reset() {
	*x = BookData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookData) ProtoMessage() {}

func (x *BookData) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookData.ProtoReflect.Descriptor instead.
func (*BookData) Descriptor() ([]byte, []int) {
	return file_book_proto_crud_proto_rawDescGZIP(), []int{1}
}

func (x *BookData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookData) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BookData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId string `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Score  uint32 `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"` // from 1 to 5
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_book_proto_crud_proto_rawDescGZIP(), []int{2}
}

func (x *Score) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Score) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *Score) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Score) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// CurrentPage is made for saving last page user read
type CurrentPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId         string `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId         string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookPageNumber int32  `protobuf:"varint,4,opt,name=book_page_number,json=bookPageNumber,proto3" json:"book_page_number,omitempty"`
}

func (x *CurrentPage) Reset() {
	*x = CurrentPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentPage) ProtoMessage() {}

func (x *CurrentPage) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentPage.ProtoReflect.Descriptor instead.
func (*CurrentPage) Descriptor() ([]byte, []int) {
	return file_book_proto_crud_proto_rawDescGZIP(), []int{3}
}

func (x *CurrentPage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CurrentPage) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *CurrentPage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CurrentPage) GetBookPageNumber() int32 {
	if x != nil {
		return x.BookPageNumber
	}
	return 0
}

// requests
type UploadBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// please note that you should not fill the id
	Book *Book     `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	Data *BookData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Hash string    `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *UploadBookRequest) Reset() {
	*x = UploadBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBookRequest) ProtoMessage() {}

func (x *UploadBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBookRequest.ProtoReflect.Descriptor instead.
func (*UploadBookRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_crud_proto_rawDescGZIP(), []int{4}
}

func (x *UploadBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UploadBookRequest) GetData() *BookData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadBookRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type UploadBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	BookId string `protobuf:"bytes,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *UploadBookResponse) Reset() {
	*x = UploadBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBookResponse) ProtoMessage() {}

func (x *UploadBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBookResponse.ProtoReflect.Descriptor instead.
func (*UploadBookResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_crud_proto_rawDescGZIP(), []int{5}
}

func (x *UploadBookResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *UploadBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

var File_book_proto_crud_proto protoreflect.FileDescriptor

var file_book_proto_crud_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62,
	0x22, 0x9e, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x47, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5f, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x79, 0x0a, 0x0b, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x25,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x3d, 0x0a, 0x12, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12,
	0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x32, 0x57, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_proto_crud_proto_rawDescOnce sync.Once
	file_book_proto_crud_proto_rawDescData = file_book_proto_crud_proto_rawDesc
)

func file_book_proto_crud_proto_rawDescGZIP() []byte {
	file_book_proto_crud_proto_rawDescOnce.Do(func() {
		file_book_proto_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_crud_proto_rawDescData)
	})
	return file_book_proto_crud_proto_rawDescData
}

var file_book_proto_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_book_proto_crud_proto_goTypes = []interface{}{
	(*Book)(nil),               // 0: bookspb.Book
	(*BookData)(nil),           // 1: bookspb.BookData
	(*Score)(nil),              // 2: bookspb.Score
	(*CurrentPage)(nil),        // 3: bookspb.CurrentPage
	(*UploadBookRequest)(nil),  // 4: bookspb.UploadBookRequest
	(*UploadBookResponse)(nil), // 5: bookspb.UploadBookResponse
}
var file_book_proto_crud_proto_depIdxs = []int32{
	0, // 0: bookspb.UploadBookRequest.book:type_name -> bookspb.Book
	1, // 1: bookspb.UploadBookRequest.data:type_name -> bookspb.BookData
	4, // 2: bookspb.BooksService.UploadBook:input_type -> bookspb.UploadBookRequest
	5, // 3: bookspb.BooksService.UploadBook:output_type -> bookspb.UploadBookResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_book_proto_crud_proto_init() }
func file_book_proto_crud_proto_init() {
	if File_book_proto_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_proto_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_book_proto_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookData); i {
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
		file_book_proto_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_book_proto_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentPage); i {
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
		file_book_proto_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBookRequest); i {
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
		file_book_proto_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBookResponse); i {
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
			RawDescriptor: file_book_proto_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_proto_crud_proto_goTypes,
		DependencyIndexes: file_book_proto_crud_proto_depIdxs,
		MessageInfos:      file_book_proto_crud_proto_msgTypes,
	}.Build()
	File_book_proto_crud_proto = out.File
	file_book_proto_crud_proto_rawDesc = nil
	file_book_proto_crud_proto_goTypes = nil
	file_book_proto_crud_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BooksServiceClient is the client API for BooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BooksServiceClient interface {
	UploadBook(ctx context.Context, in *UploadBookRequest, opts ...grpc.CallOption) (*UploadBookResponse, error)
}

type booksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksServiceClient(cc grpc.ClientConnInterface) BooksServiceClient {
	return &booksServiceClient{cc}
}

func (c *booksServiceClient) UploadBook(ctx context.Context, in *UploadBookRequest, opts ...grpc.CallOption) (*UploadBookResponse, error) {
	out := new(UploadBookResponse)
	err := c.cc.Invoke(ctx, "/bookspb.BooksService/UploadBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServiceServer is the server API for BooksService service.
type BooksServiceServer interface {
	UploadBook(context.Context, *UploadBookRequest) (*UploadBookResponse, error)
}

// UnimplementedBooksServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBooksServiceServer struct {
}

func (*UnimplementedBooksServiceServer) UploadBook(context.Context, *UploadBookRequest) (*UploadBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}

func RegisterBooksServiceServer(s *grpc.Server, srv BooksServiceServer) {
	s.RegisterService(&_BooksService_serviceDesc, srv)
}

func _BooksService_UploadBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).UploadBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookspb.BooksService/UploadBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).UploadBook(ctx, req.(*UploadBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BooksService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bookspb.BooksService",
	HandlerType: (*BooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadBook",
			Handler:    _BooksService_UploadBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book_proto/crud.proto",
}