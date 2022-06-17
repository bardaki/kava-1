// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/liquidstaking/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest defines the request type for querying x/liquidstaking
// parameters.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d94773bc8bc995f, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse defines the response type for querying x/liquidstaking
// parameters.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d94773bc8bc995f, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

// QueryDelegationHoldersRequest defines the request type for querying x/liquidstaking
// delegation holders.
type QueryDelegationHoldersRequest struct {
	Validator  string             `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDelegationHoldersRequest) Reset()         { *m = QueryDelegationHoldersRequest{} }
func (m *QueryDelegationHoldersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegationHoldersRequest) ProtoMessage()    {}
func (*QueryDelegationHoldersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d94773bc8bc995f, []int{2}
}
func (m *QueryDelegationHoldersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegationHoldersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegationHoldersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegationHoldersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegationHoldersRequest.Merge(m, src)
}
func (m *QueryDelegationHoldersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegationHoldersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegationHoldersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegationHoldersRequest proto.InternalMessageInfo

func (m *QueryDelegationHoldersRequest) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *QueryDelegationHoldersRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryDelegationHoldersResponse defines the response type for querying x/liquidstaking
// delegation holders.
type QueryDelegationHoldersResponse struct {
	DelegationHolders DelegationHolders   `protobuf:"bytes,1,rep,name=delegationHolders,proto3,castrepeated=DelegationHolders" json:"delegationHolders"`
	Pagination        *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDelegationHoldersResponse) Reset()         { *m = QueryDelegationHoldersResponse{} }
func (m *QueryDelegationHoldersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegationHoldersResponse) ProtoMessage()    {}
func (*QueryDelegationHoldersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d94773bc8bc995f, []int{3}
}
func (m *QueryDelegationHoldersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegationHoldersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegationHoldersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegationHoldersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegationHoldersResponse.Merge(m, src)
}
func (m *QueryDelegationHoldersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegationHoldersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegationHoldersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegationHoldersResponse proto.InternalMessageInfo

func (m *QueryDelegationHoldersResponse) GetDelegationHolders() DelegationHolders {
	if m != nil {
		return m.DelegationHolders
	}
	return nil
}

func (m *QueryDelegationHoldersResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "kava.liquidstaking.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "kava.liquidstaking.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryDelegationHoldersRequest)(nil), "kava.liquidstaking.v1beta1.QueryDelegationHoldersRequest")
	proto.RegisterType((*QueryDelegationHoldersResponse)(nil), "kava.liquidstaking.v1beta1.QueryDelegationHoldersResponse")
}

func init() {
	proto.RegisterFile("kava/liquidstaking/v1beta1/query.proto", fileDescriptor_5d94773bc8bc995f)
}

var fileDescriptor_5d94773bc8bc995f = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x32, 0x2a, 0xcd, 0x3b, 0xcd, 0xec, 0xd0, 0x45, 0x23, 0x9d, 0x22, 0x34, 0xa6,
	0x89, 0xd9, 0x5a, 0x11, 0x07, 0x76, 0x42, 0x15, 0xe2, 0xcf, 0x0d, 0x72, 0x44, 0x42, 0xc8, 0x59,
	0x2c, 0x63, 0x2d, 0x8d, 0xd3, 0xd8, 0xa9, 0xd8, 0x15, 0x09, 0x89, 0x23, 0x12, 0xe2, 0xce, 0x99,
	0x6f, 0xc1, 0x6d, 0xc7, 0x49, 0x5c, 0x38, 0xc1, 0xd4, 0xf0, 0x41, 0x50, 0x6c, 0x43, 0xb7, 0x86,
	0xa6, 0x68, 0xb7, 0xc6, 0x7e, 0xde, 0xe7, 0xfd, 0x3d, 0xaf, 0xdf, 0xc2, 0x9d, 0x63, 0x3a, 0xa1,
	0x24, 0x15, 0xe3, 0x52, 0x24, 0x4a, 0xd3, 0x63, 0x91, 0x71, 0x32, 0x39, 0x88, 0x99, 0xa6, 0x07,
	0x64, 0x5c, 0xb2, 0xe2, 0x04, 0xe7, 0x85, 0xd4, 0x12, 0xf9, 0xb5, 0x0e, 0x5f, 0xd2, 0x61, 0xa7,
	0xf3, 0x37, 0xb8, 0xe4, 0xd2, 0xc8, 0x48, 0xfd, 0xcb, 0x56, 0xf8, 0x5b, 0x5c, 0x4a, 0x9e, 0x32,
	0x42, 0x73, 0x41, 0x68, 0x96, 0x49, 0x4d, 0xb5, 0x90, 0x99, 0x72, 0xb7, 0x9b, 0x47, 0x52, 0x8d,
	0xa4, 0x7a, 0x65, 0xcb, 0xec, 0x87, 0xbb, 0xda, 0xb3, 0x5f, 0x24, 0xa6, 0x8a, 0x59, 0x86, 0xbf,
	0x44, 0x39, 0xe5, 0x22, 0x33, 0x3e, 0x4e, 0xdb, 0x86, 0xaf, 0xb4, 0x2c, 0x98, 0xd5, 0x85, 0x1b,
	0x10, 0x3d, 0xaf, 0x9d, 0x9e, 0xd1, 0x82, 0x8e, 0x54, 0xc4, 0xc6, 0x25, 0x53, 0x3a, 0x7c, 0x09,
	0x6f, 0x5c, 0x3a, 0x55, 0xb9, 0xcc, 0x14, 0x43, 0x0f, 0x60, 0x37, 0x37, 0x27, 0x3d, 0xb0, 0x0d,
	0x76, 0xd7, 0x06, 0x21, 0x5e, 0x1c, 0x1e, 0xdb, 0xda, 0xe1, 0xca, 0xe9, 0x8f, 0xbe, 0x17, 0xb9,
	0xba, 0xc3, 0x95, 0xf7, 0x9f, 0xfb, 0x5e, 0xf8, 0x0e, 0xc0, 0x9b, 0xc6, 0xff, 0x21, 0x4b, 0x19,
	0x37, 0xd8, 0x4f, 0x64, 0x9a, 0xb0, 0xe2, 0x0f, 0x00, 0xda, 0x82, 0xab, 0x13, 0x9a, 0x8a, 0x84,
	0x6a, 0x59, 0x98, 0x66, 0xab, 0xd1, 0xec, 0x00, 0x3d, 0x82, 0x70, 0x16, 0xb8, 0xd7, 0x31, 0x2c,
	0x3b, 0xd8, 0xcd, 0xaa, 0x9e, 0x0e, 0xb6, 0x2f, 0x34, 0x43, 0xe1, 0xcc, 0x39, 0x47, 0x17, 0x2a,
	0xc3, 0x73, 0x00, 0x83, 0x45, 0x1c, 0x2e, 0x72, 0x09, 0xd7, 0x93, 0xf9, 0xcb, 0x1e, 0xd8, 0xbe,
	0xb6, 0xbb, 0x36, 0xb8, 0xd3, 0x96, 0x7e, 0xde, 0x71, 0xb8, 0x59, 0xcf, 0xe1, 0xcb, 0xcf, 0xfe,
	0x7a, 0xb3, 0x57, 0xb3, 0x03, 0x7a, 0xfc, 0x8f, 0x84, 0xb7, 0x97, 0x26, 0xb4, 0xcc, 0x17, 0x23,
	0x0e, 0xaa, 0x0e, 0xbc, 0x6e, 0x22, 0xa2, 0x4f, 0x00, 0x76, 0xed, 0x9b, 0x20, 0xdc, 0x46, 0xde,
	0x5c, 0x07, 0x9f, 0xfc, 0xb7, 0xde, 0x12, 0x84, 0x7b, 0x6f, 0xbf, 0xfd, 0xfa, 0xd8, 0xb9, 0x85,
	0x42, 0xd2, 0xb2, 0x86, 0x76, 0x25, 0xd0, 0x57, 0x00, 0x9b, 0x33, 0x41, 0xf7, 0x97, 0xb6, 0x5c,
	0xb4, 0x3b, 0xfe, 0xe1, 0x55, 0x4a, 0x1d, 0xf8, 0x3d, 0x03, 0x4e, 0xd0, 0x7e, 0x1b, 0x78, 0xe3,
	0xb9, 0x86, 0x4f, 0x4f, 0xa7, 0x01, 0x38, 0x9b, 0x06, 0xe0, 0x7c, 0x1a, 0x80, 0x0f, 0x55, 0xe0,
	0x9d, 0x55, 0x81, 0xf7, 0xbd, 0x0a, 0xbc, 0x17, 0x84, 0x0b, 0xfd, 0xba, 0x8c, 0xf1, 0x91, 0x1c,
	0x19, 0xcb, 0xfd, 0x94, 0xc6, 0xca, 0x9a, 0xbf, 0x99, 0xb3, 0xd7, 0x27, 0x39, 0x53, 0x71, 0xd7,
	0xfc, 0x2f, 0xef, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x29, 0xfe, 0x05, 0x65, 0x80, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries all parameters of the liquidstaking module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// DelegationHolders queries liquidsavings DelegationHolders.
	DelegationHolders(ctx context.Context, in *QueryDelegationHoldersRequest, opts ...grpc.CallOption) (*QueryDelegationHoldersResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/kava.liquidstaking.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DelegationHolders(ctx context.Context, in *QueryDelegationHoldersRequest, opts ...grpc.CallOption) (*QueryDelegationHoldersResponse, error) {
	out := new(QueryDelegationHoldersResponse)
	err := c.cc.Invoke(ctx, "/kava.liquidstaking.v1beta1.Query/DelegationHolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries all parameters of the liquidstaking module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// DelegationHolders queries liquidsavings DelegationHolders.
	DelegationHolders(context.Context, *QueryDelegationHoldersRequest) (*QueryDelegationHoldersResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DelegationHolders(ctx context.Context, req *QueryDelegationHoldersRequest) (*QueryDelegationHoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegationHolders not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.liquidstaking.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DelegationHolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationHoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationHolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.liquidstaking.v1beta1.Query/DelegationHolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationHolders(ctx, req.(*QueryDelegationHoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kava.liquidstaking.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DelegationHolders",
			Handler:    _Query_DelegationHolders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kava/liquidstaking/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDelegationHoldersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegationHoldersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegationHoldersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegationHoldersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegationHoldersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegationHoldersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegationHolders) > 0 {
		for iNdEx := len(m.DelegationHolders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationHolders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDelegationHoldersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegationHoldersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DelegationHolders) > 0 {
		for _, e := range m.DelegationHolders {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDelegationHoldersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDelegationHoldersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegationHoldersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDelegationHoldersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDelegationHoldersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegationHoldersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationHolders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationHolders = append(m.DelegationHolders, DelegationHolder{})
			if err := m.DelegationHolders[len(m.DelegationHolders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)