// Code generated by protoc-gen-go. DO NOT EDIT.
// source: products.proto

package products

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SendProductRequest struct {
	IdProducto           int32    `protobuf:"varint,1,opt,name=id_producto,json=idProducto,proto3" json:"id_producto,omitempty"`
	Nombre               string   `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	CantidadDisponible   int32    `protobuf:"varint,3,opt,name=cantidad_disponible,json=cantidadDisponible,proto3" json:"cantidad_disponible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendProductRequest) Reset()         { *m = SendProductRequest{} }
func (m *SendProductRequest) String() string { return proto.CompactTextString(m) }
func (*SendProductRequest) ProtoMessage()    {}
func (*SendProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{0}
}

func (m *SendProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendProductRequest.Unmarshal(m, b)
}
func (m *SendProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendProductRequest.Marshal(b, m, deterministic)
}
func (m *SendProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendProductRequest.Merge(m, src)
}
func (m *SendProductRequest) XXX_Size() int {
	return xxx_messageInfo_SendProductRequest.Size(m)
}
func (m *SendProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendProductRequest proto.InternalMessageInfo

func (m *SendProductRequest) GetIdProducto() int32 {
	if m != nil {
		return m.IdProducto
	}
	return 0
}

func (m *SendProductRequest) GetNombre() string {
	if m != nil {
		return m.Nombre
	}
	return ""
}

func (m *SendProductRequest) GetCantidadDisponible() int32 {
	if m != nil {
		return m.CantidadDisponible
	}
	return 0
}

type SendProductResponse struct {
	IdProducto           int32    `protobuf:"varint,1,opt,name=id_producto,json=idProducto,proto3" json:"id_producto,omitempty"`
	CantidadDisponible   int32    `protobuf:"varint,2,opt,name=cantidad_disponible,json=cantidadDisponible,proto3" json:"cantidad_disponible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendProductResponse) Reset()         { *m = SendProductResponse{} }
func (m *SendProductResponse) String() string { return proto.CompactTextString(m) }
func (*SendProductResponse) ProtoMessage()    {}
func (*SendProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{1}
}

func (m *SendProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendProductResponse.Unmarshal(m, b)
}
func (m *SendProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendProductResponse.Marshal(b, m, deterministic)
}
func (m *SendProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendProductResponse.Merge(m, src)
}
func (m *SendProductResponse) XXX_Size() int {
	return xxx_messageInfo_SendProductResponse.Size(m)
}
func (m *SendProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendProductResponse proto.InternalMessageInfo

func (m *SendProductResponse) GetIdProducto() int32 {
	if m != nil {
		return m.IdProducto
	}
	return 0
}

func (m *SendProductResponse) GetCantidadDisponible() int32 {
	if m != nil {
		return m.CantidadDisponible
	}
	return 0
}

func init() {
	proto.RegisterType((*SendProductRequest)(nil), "SendProductRequest")
	proto.RegisterType((*SendProductResponse)(nil), "SendProductResponse")
}

func init() { proto.RegisterFile("products.proto", fileDescriptor_8c6e54f42122eb82) }

var fileDescriptor_8c6e54f42122eb82 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xaa, 0xe3, 0x12, 0x0a, 0x4e,
	0xcd, 0x4b, 0x09, 0x80, 0x88, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x73, 0x71,
	0x67, 0xa6, 0xc4, 0x43, 0x95, 0xe6, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x71, 0x65, 0xc2,
	0x94, 0xe5, 0x0b, 0x89, 0x71, 0xb1, 0xe5, 0xe5, 0xe7, 0x26, 0x15, 0xa5, 0x4a, 0x30, 0x29, 0x30,
	0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0xfa, 0x5c, 0xc2, 0xc9, 0x89, 0x79, 0x25, 0x99, 0x29, 0x89,
	0x29, 0xf1, 0x29, 0x99, 0xc5, 0x05, 0xf9, 0x79, 0x99, 0x49, 0x39, 0xa9, 0x12, 0xcc, 0x60, 0x03,
	0x84, 0x60, 0x52, 0x2e, 0x70, 0x19, 0xa5, 0x74, 0x2e, 0x61, 0x14, 0xfb, 0x41, 0x12, 0xc5, 0xa9,
	0x84, 0x1d, 0x80, 0xc3, 0x22, 0x26, 0x5c, 0x16, 0x19, 0xb9, 0x70, 0x71, 0x40, 0x35, 0x17, 0x0b,
	0x59, 0x70, 0x71, 0x23, 0x59, 0x2a, 0x24, 0xac, 0x87, 0x19, 0x04, 0x52, 0x22, 0x7a, 0x58, 0xdc,
	0x95, 0xc4, 0x06, 0x0e, 0x35, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0xcf, 0xcc, 0x67,
	0x47, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductsClient interface {
	SendProduct(ctx context.Context, in *SendProductRequest, opts ...grpc.CallOption) (*SendProductResponse, error)
}

type productsClient struct {
	cc *grpc.ClientConn
}

func NewProductsClient(cc *grpc.ClientConn) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) SendProduct(ctx context.Context, in *SendProductRequest, opts ...grpc.CallOption) (*SendProductResponse, error) {
	out := new(SendProductResponse)
	err := c.cc.Invoke(ctx, "/Products/SendProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
type ProductsServer interface {
	SendProduct(context.Context, *SendProductRequest) (*SendProductResponse, error)
}

// UnimplementedProductsServer can be embedded to have forward compatible implementations.
type UnimplementedProductsServer struct {
}

func (*UnimplementedProductsServer) SendProduct(ctx context.Context, req *SendProductRequest) (*SendProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProduct not implemented")
}

func RegisterProductsServer(s *grpc.Server, srv ProductsServer) {
	s.RegisterService(&_Products_serviceDesc, srv)
}

func _Products_SendProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).SendProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Products/SendProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).SendProduct(ctx, req.(*SendProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Products_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendProduct",
			Handler:    _Products_SendProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}
