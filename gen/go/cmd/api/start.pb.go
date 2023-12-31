// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cmd/api/start.proto

package start_proto

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

type GetAuthorsByBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAuthorsByBooksRequest) Reset() {
	*x = GetAuthorsByBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_api_start_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsByBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsByBooksRequest) ProtoMessage() {}

func (x *GetAuthorsByBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_api_start_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsByBooksRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorsByBooksRequest) Descriptor() ([]byte, []int) {
	return file_cmd_api_start_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthorsByBooksRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAuthorsByBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author []*Author `protobuf:"bytes,1,rep,name=author,proto3" json:"author,omitempty"`
}

func (x *GetAuthorsByBooksResponse) Reset() {
	*x = GetAuthorsByBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_api_start_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsByBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsByBooksResponse) ProtoMessage() {}

func (x *GetAuthorsByBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_api_start_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsByBooksResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorsByBooksResponse) Descriptor() ([]byte, []int) {
	return file_cmd_api_start_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthorsByBooksResponse) GetAuthor() []*Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type GetBooksByAuthorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBooksByAuthorsRequest) Reset() {
	*x = GetBooksByAuthorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_api_start_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksByAuthorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksByAuthorsRequest) ProtoMessage() {}

func (x *GetBooksByAuthorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_api_start_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksByAuthorsRequest.ProtoReflect.Descriptor instead.
func (*GetBooksByAuthorsRequest) Descriptor() ([]byte, []int) {
	return file_cmd_api_start_proto_rawDescGZIP(), []int{2}
}

func (x *GetBooksByAuthorsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBooksByAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBooksByAuthorsResponse) Reset() {
	*x = GetBooksByAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_api_start_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksByAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksByAuthorsResponse) ProtoMessage() {}

func (x *GetBooksByAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_api_start_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksByAuthorsResponse.ProtoReflect.Descriptor instead.
func (*GetBooksByAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_cmd_api_start_proto_rawDescGZIP(), []int{3}
}

func (x *GetBooksByAuthorsResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author []*Author `protobuf:"bytes,3,rep,name=author,proto3" json:"author,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_api_start_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_api_start_proto_msgTypes[4]
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
	return file_cmd_api_start_proto_rawDescGZIP(), []int{4}
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

func (x *Book) GetAuthor() []*Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_api_start_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_api_start_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_cmd_api_start_proto_rawDescGZIP(), []int{5}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_cmd_api_start_proto protoreflect.FileDescriptor

var file_cmd_api_start_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6d, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x4a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x22, 0x59, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x2c, 0x0a, 0x06, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xdc, 0x01, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x68, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x79, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x63, 0x6d, 0x64, 0x2f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cmd_api_start_proto_rawDescOnce sync.Once
	file_cmd_api_start_proto_rawDescData = file_cmd_api_start_proto_rawDesc
)

func file_cmd_api_start_proto_rawDescGZIP() []byte {
	file_cmd_api_start_proto_rawDescOnce.Do(func() {
		file_cmd_api_start_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_api_start_proto_rawDescData)
	})
	return file_cmd_api_start_proto_rawDescData
}

var file_cmd_api_start_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cmd_api_start_proto_goTypes = []interface{}{
	(*GetAuthorsByBooksRequest)(nil),  // 0: api.search.v1.GetAuthorsByBooksRequest
	(*GetAuthorsByBooksResponse)(nil), // 1: api.search.v1.GetAuthorsByBooksResponse
	(*GetBooksByAuthorsRequest)(nil),  // 2: api.search.v1.GetBooksByAuthorsRequest
	(*GetBooksByAuthorsResponse)(nil), // 3: api.search.v1.GetBooksByAuthorsResponse
	(*Book)(nil),                      // 4: api.search.v1.Book
	(*Author)(nil),                    // 5: api.search.v1.Author
}
var file_cmd_api_start_proto_depIdxs = []int32{
	5, // 0: api.search.v1.GetAuthorsByBooksResponse.author:type_name -> api.search.v1.Author
	4, // 1: api.search.v1.GetBooksByAuthorsResponse.books:type_name -> api.search.v1.Book
	5, // 2: api.search.v1.Book.author:type_name -> api.search.v1.Author
	0, // 3: api.search.v1.search.GetAuthorsByBooks:input_type -> api.search.v1.GetAuthorsByBooksRequest
	2, // 4: api.search.v1.search.GetBooksByAuthors:input_type -> api.search.v1.GetBooksByAuthorsRequest
	1, // 5: api.search.v1.search.GetAuthorsByBooks:output_type -> api.search.v1.GetAuthorsByBooksResponse
	3, // 6: api.search.v1.search.GetBooksByAuthors:output_type -> api.search.v1.GetBooksByAuthorsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cmd_api_start_proto_init() }
func file_cmd_api_start_proto_init() {
	if File_cmd_api_start_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_api_start_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsByBooksRequest); i {
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
		file_cmd_api_start_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsByBooksResponse); i {
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
		file_cmd_api_start_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksByAuthorsRequest); i {
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
		file_cmd_api_start_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksByAuthorsResponse); i {
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
		file_cmd_api_start_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_cmd_api_start_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
			RawDescriptor: file_cmd_api_start_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_api_start_proto_goTypes,
		DependencyIndexes: file_cmd_api_start_proto_depIdxs,
		MessageInfos:      file_cmd_api_start_proto_msgTypes,
	}.Build()
	File_cmd_api_start_proto = out.File
	file_cmd_api_start_proto_rawDesc = nil
	file_cmd_api_start_proto_goTypes = nil
	file_cmd_api_start_proto_depIdxs = nil
}
