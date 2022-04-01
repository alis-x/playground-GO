// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BooksServiceClient is the client API for BooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksServiceClient interface {
	// Create a book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Get a book.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Delete a book.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List books.
	// Features are listed in Descending order.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
}

type booksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksServiceClient(cc grpc.ClientConnInterface) BooksServiceClient {
	return &booksServiceClient{cc}
}

func (c *booksServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/foo.br.resources.books.v1.BooksService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/foo.br.resources.books.v1.BooksService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/foo.br.resources.books.v1.BooksService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/foo.br.resources.books.v1.BooksService/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServiceServer is the server API for BooksService service.
// All implementations must embed UnimplementedBooksServiceServer
// for forward compatibility
type BooksServiceServer interface {
	// Create a book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	// Get a book.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// Delete a book.
	DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error)
	// List books.
	// Features are listed in Descending order.
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	mustEmbedUnimplementedBooksServiceServer()
}

// UnimplementedBooksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBooksServiceServer struct {
}

func (UnimplementedBooksServiceServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBooksServiceServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBooksServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBooksServiceServer) ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBooksServiceServer) mustEmbedUnimplementedBooksServiceServer() {}

// UnsafeBooksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServiceServer will
// result in compilation errors.
type UnsafeBooksServiceServer interface {
	mustEmbedUnimplementedBooksServiceServer()
}

func RegisterBooksServiceServer(s grpc.ServiceRegistrar, srv BooksServiceServer) {
	s.RegisterService(&BooksService_ServiceDesc, srv)
}

func _BooksService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.br.resources.books.v1.BooksService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.br.resources.books.v1.BooksService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.br.resources.books.v1.BooksService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.br.resources.books.v1.BooksService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BooksService_ServiceDesc is the grpc.ServiceDesc for BooksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BooksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "foo.br.resources.books.v1.BooksService",
	HandlerType: (*BooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BooksService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BooksService_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BooksService_DeleteBook_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _BooksService_ListBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo/br/resources/books/v1/books.proto",
}