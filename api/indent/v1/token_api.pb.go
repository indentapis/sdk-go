// Code generated by protoc-gen-go. DO NOT EDIT.
// source: indent/v1/token_api.proto

package indentv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ExchangeTokenRequest struct {
	GrantType            string   `protobuf:"bytes,1,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeTokenRequest) Reset()         { *m = ExchangeTokenRequest{} }
func (m *ExchangeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ExchangeTokenRequest) ProtoMessage()    {}
func (*ExchangeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_071f5b8ba7b06907, []int{0}
}

func (m *ExchangeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeTokenRequest.Unmarshal(m, b)
}
func (m *ExchangeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeTokenRequest.Marshal(b, m, deterministic)
}
func (m *ExchangeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeTokenRequest.Merge(m, src)
}
func (m *ExchangeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ExchangeTokenRequest.Size(m)
}
func (m *ExchangeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeTokenRequest proto.InternalMessageInfo

func (m *ExchangeTokenRequest) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

func (m *ExchangeTokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type JWKS struct {
	Keys                 []*JWK   `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWKS) Reset()         { *m = JWKS{} }
func (m *JWKS) String() string { return proto.CompactTextString(m) }
func (*JWKS) ProtoMessage()    {}
func (*JWKS) Descriptor() ([]byte, []int) {
	return fileDescriptor_071f5b8ba7b06907, []int{1}
}

func (m *JWKS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWKS.Unmarshal(m, b)
}
func (m *JWKS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWKS.Marshal(b, m, deterministic)
}
func (m *JWKS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWKS.Merge(m, src)
}
func (m *JWKS) XXX_Size() int {
	return xxx_messageInfo_JWKS.Size(m)
}
func (m *JWKS) XXX_DiscardUnknown() {
	xxx_messageInfo_JWKS.DiscardUnknown(m)
}

var xxx_messageInfo_JWKS proto.InternalMessageInfo

func (m *JWKS) GetKeys() []*JWK {
	if m != nil {
		return m.Keys
	}
	return nil
}

type JWK struct {
	// The "kty" (key type) parameter identifies the cryptographic algorithm family used with the key,
	// such as "RSA" or "EC".
	Kty string `protobuf:"bytes,1,opt,name=kty,proto3" json:"kty,omitempty"`
	// The "use" (public key use) parameter identifies the intended use of the public key.
	Use string `protobuf:"bytes,2,opt,name=use,proto3" json:"use,omitempty"`
	//The "alg" (algorithm) parameter identifies the algorithm intended for use with the key.
	Alg string `protobuf:"bytes,3,opt,name=alg,proto3" json:"alg,omitempty"`
	// The "kid" (key ID) parameter is used to match a specific key.
	Kid string `protobuf:"bytes,4,opt,name=kid,proto3" json:"kid,omitempty"`
	// The "crv" (curve) parameter identifies the cryptographic curve used with the key.
	Crv string `protobuf:"bytes,5,opt,name=crv,proto3" json:"crv,omitempty"`
	// The "x" (x coordinate) parameter contains the x coordinate for the elliptic curve point.
	X string `protobuf:"bytes,6,opt,name=x,proto3" json:"x,omitempty"`
	// The "y" (y coordinate) parameter contains the y coordinate for the elliptic curve point.
	Y                    string   `protobuf:"bytes,7,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWK) Reset()         { *m = JWK{} }
func (m *JWK) String() string { return proto.CompactTextString(m) }
func (*JWK) ProtoMessage()    {}
func (*JWK) Descriptor() ([]byte, []int) {
	return fileDescriptor_071f5b8ba7b06907, []int{2}
}

func (m *JWK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWK.Unmarshal(m, b)
}
func (m *JWK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWK.Marshal(b, m, deterministic)
}
func (m *JWK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWK.Merge(m, src)
}
func (m *JWK) XXX_Size() int {
	return xxx_messageInfo_JWK.Size(m)
}
func (m *JWK) XXX_DiscardUnknown() {
	xxx_messageInfo_JWK.DiscardUnknown(m)
}

var xxx_messageInfo_JWK proto.InternalMessageInfo

func (m *JWK) GetKty() string {
	if m != nil {
		return m.Kty
	}
	return ""
}

func (m *JWK) GetUse() string {
	if m != nil {
		return m.Use
	}
	return ""
}

func (m *JWK) GetAlg() string {
	if m != nil {
		return m.Alg
	}
	return ""
}

func (m *JWK) GetKid() string {
	if m != nil {
		return m.Kid
	}
	return ""
}

func (m *JWK) GetCrv() string {
	if m != nil {
		return m.Crv
	}
	return ""
}

func (m *JWK) GetX() string {
	if m != nil {
		return m.X
	}
	return ""
}

func (m *JWK) GetY() string {
	if m != nil {
		return m.Y
	}
	return ""
}

func init() {
	proto.RegisterType((*ExchangeTokenRequest)(nil), "indent.v1.ExchangeTokenRequest")
	proto.RegisterType((*JWKS)(nil), "indent.v1.JWKS")
	proto.RegisterType((*JWK)(nil), "indent.v1.JWK")
}

func init() {
	proto.RegisterFile("indent/v1/token_api.proto", fileDescriptor_071f5b8ba7b06907)
}

var fileDescriptor_071f5b8ba7b06907 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xc6, 0xa1, 0xc5, 0x4b, 0xdd, 0x56, 0xcb, 0x8f, 0xb6, 0x06, 0x44, 0x65, 0x2e, 0x55,
	0x0f, 0xb6, 0x5c, 0x6e, 0x70, 0xa2, 0x52, 0x85, 0x9c, 0x5c, 0x22, 0xb7, 0x6a, 0xaa, 0x0a, 0x14,
	0x6d, 0x9d, 0xa9, 0x63, 0x25, 0xdd, 0x35, 0xeb, 0xb5, 0x15, 0x8b, 0x1b, 0xaf, 0xc0, 0x1b, 0x70,
	0xe4, 0xcc, 0x53, 0x70, 0xe5, 0x15, 0x78, 0x10, 0xb4, 0xbb, 0x6e, 0x48, 0x24, 0x6e, 0x33, 0xdf,
	0xcc, 0x7c, 0x33, 0xf3, 0x7d, 0xf8, 0xa0, 0xe0, 0x53, 0xe0, 0x2a, 0x6a, 0xe2, 0x48, 0x89, 0x39,
	0xf0, 0x09, 0x2b, 0x8b, 0xb0, 0x94, 0x42, 0x09, 0xe2, 0xda, 0x52, 0xd8, 0xc4, 0xfe, 0x8b, 0x5c,
	0x88, 0x7c, 0x01, 0x11, 0x2b, 0x8b, 0x88, 0x71, 0x2e, 0x14, 0x53, 0x85, 0xe0, 0x95, 0x6d, 0xf4,
	0x9f, 0x77, 0x55, 0x93, 0xdd, 0xd4, 0xb7, 0x11, 0xdc, 0x95, 0xaa, 0xed, 0x8a, 0x6b, 0x0b, 0x24,
	0x54, 0xa2, 0x96, 0x19, 0x74, 0x73, 0xc1, 0x35, 0x7e, 0x72, 0xb6, 0xcc, 0x66, 0x8c, 0xe7, 0x70,
	0xa1, 0x77, 0xa7, 0xf0, 0xb9, 0x86, 0x4a, 0x91, 0x97, 0x18, 0xe7, 0x92, 0x71, 0x35, 0x51, 0x6d,
	0x09, 0x14, 0x1d, 0xa2, 0x23, 0x37, 0x75, 0x0d, 0x72, 0xd1, 0x96, 0x40, 0x5e, 0x63, 0x4f, 0xc2,
	0xad, 0x84, 0x6a, 0x36, 0x31, 0x27, 0xd3, 0x9e, 0xe9, 0xd8, 0xe9, 0x40, 0x43, 0x15, 0x1c, 0xe3,
	0xfe, 0x60, 0x3c, 0x3c, 0x27, 0x01, 0xee, 0xcf, 0xa1, 0xad, 0x28, 0x3a, 0x74, 0x8e, 0x1e, 0x9d,
	0xec, 0x86, 0xab, 0x9f, 0xc2, 0xc1, 0x78, 0x98, 0x9a, 0x5a, 0xf0, 0x05, 0x3b, 0x83, 0xf1, 0x90,
	0xec, 0x63, 0x67, 0xae, 0xda, 0x6e, 0x9f, 0x0e, 0x35, 0x52, 0x57, 0xd0, 0xf1, 0xeb, 0x50, 0x23,
	0x6c, 0x91, 0x53, 0xc7, 0x22, 0x6c, 0x91, 0x9b, 0xa9, 0x62, 0x4a, 0xfb, 0xdd, 0x54, 0x31, 0xd5,
	0x48, 0x26, 0x1b, 0xfa, 0xc0, 0x22, 0x99, 0x6c, 0xc8, 0x0e, 0x46, 0x4b, 0xba, 0x65, 0x72, 0xb4,
	0xd4, 0x59, 0x4b, 0xb7, 0x6d, 0xd6, 0x9e, 0xfc, 0x44, 0xf8, 0xa1, 0x39, 0xf9, 0xfd, 0x28, 0x21,
	0x9f, 0xb0, 0xb7, 0xa1, 0x08, 0x79, 0xb5, 0x76, 0xf0, 0xff, 0xb4, 0xf2, 0xf7, 0xd7, 0x1a, 0xec,
	0xe7, 0x07, 0x5f, 0x7f, 0xff, 0xf9, 0xd6, 0x7b, 0x1c, 0xec, 0x6a, 0xc9, 0x59, 0xad, 0x66, 0xd6,
	0xd8, 0xb7, 0xe8, 0x98, 0x24, 0x78, 0xfb, 0x03, 0x28, 0xa3, 0xcb, 0xb3, 0xd0, 0x9a, 0x16, 0xde,
	0x9b, 0x16, 0x9e, 0x69, 0xd3, 0xfc, 0xbd, 0x4d, 0x85, 0xce, 0x83, 0xa7, 0x86, 0x6e, 0x8f, 0x78,
	0x2b, 0x3a, 0xad, 0xd9, 0x69, 0x8a, 0xbd, 0x4c, 0xdc, 0xfd, 0x6b, 0x3e, 0xf5, 0xec, 0x13, 0x65,
	0x31, 0xd2, 0x84, 0x23, 0x74, 0x4d, 0x56, 0xc6, 0xbf, 0xb3, 0x51, 0x13, 0x7f, 0xef, 0x39, 0xc9,
	0xd5, 0xd5, 0x8f, 0x9e, 0x9b, 0xd8, 0xb1, 0xcb, 0xf8, 0xd7, 0x7d, 0xfc, 0xf1, 0x32, 0xbe, 0xd9,
	0x32, 0xb7, 0xbc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xc7, 0x63, 0x69, 0x94, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenAPIClient is the client API for TokenAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenAPIClient interface {
	// ExchangeToken produces access tokens associated with a refresh token.
	ExchangeToken(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*Token, error)
	// GetJWKS produces the JSON Web Key Set for the service.
	GetJWKS(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*JWKS, error)
}

type tokenAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenAPIClient(cc grpc.ClientConnInterface) TokenAPIClient {
	return &tokenAPIClient{cc}
}

func (c *tokenAPIClient) ExchangeToken(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/indent.v1.TokenAPI/ExchangeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenAPIClient) GetJWKS(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*JWKS, error) {
	out := new(JWKS)
	err := c.cc.Invoke(ctx, "/indent.v1.TokenAPI/GetJWKS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenAPIServer is the server API for TokenAPI service.
type TokenAPIServer interface {
	// ExchangeToken produces access tokens associated with a refresh token.
	ExchangeToken(context.Context, *ExchangeTokenRequest) (*Token, error)
	// GetJWKS produces the JSON Web Key Set for the service.
	GetJWKS(context.Context, *emptypb.Empty) (*JWKS, error)
}

// UnimplementedTokenAPIServer can be embedded to have forward compatible implementations.
type UnimplementedTokenAPIServer struct {
}

func (*UnimplementedTokenAPIServer) ExchangeToken(ctx context.Context, req *ExchangeTokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeToken not implemented")
}
func (*UnimplementedTokenAPIServer) GetJWKS(ctx context.Context, req *emptypb.Empty) (*JWKS, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJWKS not implemented")
}

func RegisterTokenAPIServer(s *grpc.Server, srv TokenAPIServer) {
	s.RegisterService(&_TokenAPI_serviceDesc, srv)
}

func _TokenAPI_ExchangeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenAPIServer).ExchangeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.TokenAPI/ExchangeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenAPIServer).ExchangeToken(ctx, req.(*ExchangeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenAPI_GetJWKS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenAPIServer).GetJWKS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/indent.v1.TokenAPI/GetJWKS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenAPIServer).GetJWKS(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "indent.v1.TokenAPI",
	HandlerType: (*TokenAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExchangeToken",
			Handler:    _TokenAPI_ExchangeToken_Handler,
		},
		{
			MethodName: "GetJWKS",
			Handler:    _TokenAPI_GetJWKS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indent/v1/token_api.proto",
}
