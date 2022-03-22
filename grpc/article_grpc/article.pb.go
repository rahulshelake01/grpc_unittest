// Code generated by protoc-gen-go. DO NOT EDIT.
// source: article.proto

package article_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Article struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Articlecreationdate  string   `protobuf:"bytes,4,opt,name=articlecreationdate,proto3" json:"articlecreationdate,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{0}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Article) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetArticlecreationdate() string {
	if m != nil {
		return m.Articlecreationdate
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ArticleList struct {
	Articles             []*Article `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ArticleList) Reset()         { *m = ArticleList{} }
func (m *ArticleList) String() string { return proto.CompactTextString(m) }
func (*ArticleList) ProtoMessage()    {}
func (*ArticleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{1}
}

func (m *ArticleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleList.Unmarshal(m, b)
}
func (m *ArticleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleList.Marshal(b, m, deterministic)
}
func (m *ArticleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleList.Merge(m, src)
}
func (m *ArticleList) XXX_Size() int {
	return xxx_messageInfo_ArticleList.Size(m)
}
func (m *ArticleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleList.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleList proto.InternalMessageInfo

func (m *ArticleList) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type AddArticle struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Articlecreationdate  string   `protobuf:"bytes,4,opt,name=articlecreationdate,proto3" json:"articlecreationdate,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddArticle) Reset()         { *m = AddArticle{} }
func (m *AddArticle) String() string { return proto.CompactTextString(m) }
func (*AddArticle) ProtoMessage()    {}
func (*AddArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{2}
}

func (m *AddArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddArticle.Unmarshal(m, b)
}
func (m *AddArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddArticle.Marshal(b, m, deterministic)
}
func (m *AddArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddArticle.Merge(m, src)
}
func (m *AddArticle) XXX_Size() int {
	return xxx_messageInfo_AddArticle.Size(m)
}
func (m *AddArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_AddArticle.DiscardUnknown(m)
}

var xxx_messageInfo_AddArticle proto.InternalMessageInfo

func (m *AddArticle) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AddArticle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddArticle) GetArticlecreationdate() string {
	if m != nil {
		return m.Articlecreationdate
	}
	return ""
}

func (m *AddArticle) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ArticleResponse struct {
	Status               int32      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Success              bool       `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Articles             []*Article `protobuf:"bytes,4,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ArticleResponse) Reset()         { *m = ArticleResponse{} }
func (m *ArticleResponse) String() string { return proto.CompactTextString(m) }
func (*ArticleResponse) ProtoMessage()    {}
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{3}
}

func (m *ArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleResponse.Unmarshal(m, b)
}
func (m *ArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleResponse.Marshal(b, m, deterministic)
}
func (m *ArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleResponse.Merge(m, src)
}
func (m *ArticleResponse) XXX_Size() int {
	return xxx_messageInfo_ArticleResponse.Size(m)
}
func (m *ArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleResponse proto.InternalMessageInfo

func (m *ArticleResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ArticleResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ArticleResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ArticleResponse) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type SearchArticle struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchArticle) Reset()         { *m = SearchArticle{} }
func (m *SearchArticle) String() string { return proto.CompactTextString(m) }
func (*SearchArticle) ProtoMessage()    {}
func (*SearchArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{4}
}

func (m *SearchArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchArticle.Unmarshal(m, b)
}
func (m *SearchArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchArticle.Marshal(b, m, deterministic)
}
func (m *SearchArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchArticle.Merge(m, src)
}
func (m *SearchArticle) XXX_Size() int {
	return xxx_messageInfo_SearchArticle.Size(m)
}
func (m *SearchArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchArticle.DiscardUnknown(m)
}

var xxx_messageInfo_SearchArticle proto.InternalMessageInfo

func (m *SearchArticle) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c593d380f9840a2, []int{5}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Article)(nil), "article_grpc.Article")
	proto.RegisterType((*ArticleList)(nil), "article_grpc.ArticleList")
	proto.RegisterType((*AddArticle)(nil), "article_grpc.AddArticle")
	proto.RegisterType((*ArticleResponse)(nil), "article_grpc.ArticleResponse")
	proto.RegisterType((*SearchArticle)(nil), "article_grpc.SearchArticle")
	proto.RegisterType((*Void)(nil), "article_grpc.void")
}

func init() { proto.RegisterFile("article.proto", fileDescriptor_5c593d380f9840a2) }

var fileDescriptor_5c593d380f9840a2 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x9b, 0xb6, 0x61, 0x4a, 0x41, 0x1a, 0x7e, 0x64, 0x15, 0x21, 0xaa, 0xac, 0xba,
	0xaa, 0xa0, 0x5c, 0x80, 0xb2, 0x61, 0x03, 0x9b, 0x70, 0x00, 0x64, 0x1c, 0xab, 0x58, 0xb4, 0x49,
	0x95, 0x71, 0x39, 0x04, 0x2b, 0x76, 0xdc, 0x90, 0x73, 0x20, 0x3b, 0x76, 0x44, 0x50, 0xa5, 0x76,
	0xc3, 0xf2, 0xc5, 0x9f, 0xc7, 0xef, 0xcd, 0x53, 0x60, 0x20, 0x4a, 0xa3, 0xe5, 0x42, 0x4d, 0x56,
	0x65, 0x61, 0x0a, 0x3c, 0xf0, 0xf2, 0x79, 0x5e, 0xae, 0x64, 0xf2, 0xc5, 0xa0, 0x37, 0xab, 0x3e,
	0xe0, 0x10, 0xe2, 0x5c, 0xcb, 0xb7, 0x5c, 0x2c, 0x15, 0x67, 0x23, 0x36, 0xde, 0x4f, 0x6b, 0x8d,
	0x87, 0xd0, 0xd2, 0x19, 0x6f, 0x8d, 0xd8, 0x38, 0x4a, 0x5b, 0x3a, 0xc3, 0x13, 0xe8, 0x18, 0x6d,
	0x16, 0x8a, 0xb7, 0x1d, 0x58, 0x09, 0xbc, 0x82, 0x63, 0x3f, 0x5d, 0x96, 0x4a, 0x18, 0x5d, 0xe4,
	0x99, 0x30, 0x8a, 0x47, 0x8e, 0xd9, 0x74, 0x84, 0x1c, 0x7a, 0xb2, 0xc8, 0x8d, 0xca, 0x0d, 0xef,
	0x38, 0x2a, 0xc8, 0xe4, 0x16, 0xfa, 0xde, 0xd8, 0x83, 0x26, 0x83, 0xd7, 0x10, 0xfb, 0xfb, 0xc4,
	0xd9, 0xa8, 0x3d, 0xee, 0x4f, 0x4f, 0x27, 0xbf, 0x93, 0x4c, 0x3c, 0x9c, 0xd6, 0x58, 0xf2, 0xc1,
	0x00, 0x66, 0x59, 0xb6, 0x4b, 0xbc, 0xff, 0x8f, 0xf3, 0xc9, 0xe0, 0x28, 0x58, 0x54, 0xb4, 0x2a,
	0x72, 0x52, 0x78, 0x06, 0x5d, 0x32, 0xc2, 0xac, 0xc9, 0xf9, 0xe9, 0xa4, 0x5e, 0xd9, 0x29, 0xb4,
	0x96, 0x52, 0x11, 0xb9, 0x8d, 0xc7, 0x69, 0x90, 0xf6, 0x64, 0xa9, 0x88, 0xc4, 0x3c, 0x38, 0x0d,
	0xb2, 0xb1, 0x9f, 0x68, 0xb7, 0xfd, 0x5c, 0xc2, 0xe0, 0x49, 0x89, 0x52, 0xbe, 0x86, 0x0d, 0x55,
	0x25, 0x5b, 0x2f, 0x6d, 0x5b, 0x72, 0xd2, 0x85, 0xe8, 0xbd, 0xd0, 0xd9, 0xf4, 0x9b, 0x41, 0xec,
	0x19, 0xc2, 0xbb, 0xba, 0x97, 0x85, 0xed, 0x05, 0x9b, 0xaf, 0x58, 0x7e, 0x78, 0xb1, 0xf9, 0x65,
	0x1f, 0x3b, 0xd9, 0xc3, 0x7b, 0x57, 0x8c, 0x87, 0x90, 0xff, 0xc1, 0xeb, 0xca, 0xb6, 0x0f, 0x7a,
	0x0c, 0x11, 0xc2, 0xac, 0xf3, 0xe6, 0x8d, 0x46, 0xbe, 0xad, 0xe3, 0x5e, 0xba, 0xee, 0x17, 0xb9,
	0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x72, 0xc1, 0x50, 0x33, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticlesClient interface {
	Articlelist(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ArticleResponse, error)
	Addarticle(ctx context.Context, in *AddArticle, opts ...grpc.CallOption) (*ArticleResponse, error)
	Searcharticle(ctx context.Context, in *SearchArticle, opts ...grpc.CallOption) (*ArticleResponse, error)
}

type articlesClient struct {
	cc *grpc.ClientConn
}

func NewArticlesClient(cc *grpc.ClientConn) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) Articlelist(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, "/article_grpc.Articles/Articlelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Addarticle(ctx context.Context, in *AddArticle, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, "/article_grpc.Articles/Addarticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Searcharticle(ctx context.Context, in *SearchArticle, opts ...grpc.CallOption) (*ArticleResponse, error) {
	out := new(ArticleResponse)
	err := c.cc.Invoke(ctx, "/article_grpc.Articles/Searcharticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
type ArticlesServer interface {
	Articlelist(context.Context, *Void) (*ArticleResponse, error)
	Addarticle(context.Context, *AddArticle) (*ArticleResponse, error)
	Searcharticle(context.Context, *SearchArticle) (*ArticleResponse, error)
}

func RegisterArticlesServer(s *grpc.Server, srv ArticlesServer) {
	s.RegisterService(&_Articles_serviceDesc, srv)
}

func _Articles_Articlelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Articlelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_grpc.Articles/Articlelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Articlelist(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Addarticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Addarticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_grpc.Articles/Addarticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Addarticle(ctx, req.(*AddArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Searcharticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Searcharticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_grpc.Articles/Searcharticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Searcharticle(ctx, req.(*SearchArticle))
	}
	return interceptor(ctx, in, info, handler)
}

var _Articles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "article_grpc.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Articlelist",
			Handler:    _Articles_Articlelist_Handler,
		},
		{
			MethodName: "Addarticle",
			Handler:    _Articles_Addarticle_Handler,
		},
		{
			MethodName: "Searcharticle",
			Handler:    _Articles_Searcharticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
