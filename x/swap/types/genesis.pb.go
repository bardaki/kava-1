// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/swap/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the swap module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params       Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PoolRecords  []PoolRecord  `protobuf:"bytes,2,rep,name=pool_records,json=poolRecords,proto3" json:"pool_records" yaml:"pool_records,omitempty"`
	ShareRecords []ShareRecord `protobuf:"bytes,3,rep,name=share_records,json=shareRecords,proto3" json:"share_records" yaml:"share_records,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a1a1687f484a21, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPoolRecords() []PoolRecord {
	if m != nil {
		return m.PoolRecords
	}
	return nil
}

func (m *GenesisState) GetShareRecords() []ShareRecord {
	if m != nil {
		return m.ShareRecords
	}
	return nil
}

// Params defines the parameters for the swap module.
type Params struct {
	AllowedPools []*AllowedPool                         `protobuf:"bytes,1,rep,name=allowed_pools,json=allowedPools,proto3" json:"allowed_pools,omitempty" yaml:"allowed_pools,omitempty"`
	SwapFee      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee,omitempty" yaml:"swap_fee"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a1a1687f484a21, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAllowedPools() []*AllowedPool {
	if m != nil {
		return m.AllowedPools
	}
	return nil
}

// AllowedPool defines a tradable pool.
type AllowedPool struct {
	TokenA string `protobuf:"bytes,1,opt,name=token_a,json=tokenA,proto3" json:"token_a,omitempty" yaml:"token_a,omitempty"`
	TokenB string `protobuf:"bytes,2,opt,name=token_b,json=tokenB,proto3" json:"token_b,omitempty" yaml:"token_b,omitempty"`
}

func (m *AllowedPool) Reset()      { *m = AllowedPool{} }
func (*AllowedPool) ProtoMessage() {}
func (*AllowedPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a1a1687f484a21, []int{2}
}
func (m *AllowedPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedPool.Merge(m, src)
}
func (m *AllowedPool) XXX_Size() int {
	return m.Size()
}
func (m *AllowedPool) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedPool.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedPool proto.InternalMessageInfo

func (m *AllowedPool) GetTokenA() string {
	if m != nil {
		return m.TokenA
	}
	return ""
}

func (m *AllowedPool) GetTokenB() string {
	if m != nil {
		return m.TokenB
	}
	return ""
}

// PoolRecord represents the state of a liquidity pool
// and is used to store the state of a denominated pool
type PoolRecord struct {
	PoolId      string                                  `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id,omitempty"`
	ReservesA   github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,opt,name=reserves_a,json=reservesA,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"reserves_a,omitempty" yaml:"reserves_a"`
	ReservesB   github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=reserves_b,json=reservesB,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"reserves_b,omitempty" yaml:"reserves_b"`
	TotalShares github_com_cosmos_cosmos_sdk_types.Int  `protobuf:"bytes,4,opt,name=total_shares,json=totalShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_shares,omitempty" yaml:"total_shares"`
}

func (m *PoolRecord) Reset()         { *m = PoolRecord{} }
func (m *PoolRecord) String() string { return proto.CompactTextString(m) }
func (*PoolRecord) ProtoMessage()    {}
func (*PoolRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a1a1687f484a21, []int{3}
}
func (m *PoolRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolRecord.Merge(m, src)
}
func (m *PoolRecord) XXX_Size() int {
	return m.Size()
}
func (m *PoolRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolRecord.DiscardUnknown(m)
}

var xxx_messageInfo_PoolRecord proto.InternalMessageInfo

func (m *PoolRecord) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

// ShareRecord stores the shares owned for a depositor and pool
type ShareRecord struct {
	Depositor   string                                 `protobuf:"bytes,1,opt,name=depositor,proto3" json:"depositor,omitempty"`
	PoolId      string                                 `protobuf:"bytes,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id,omitempty"`
	SharesOwned github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=shares_owned,json=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shares_owned,omitempty" yaml:"shares_owned"`
}

func (m *ShareRecord) Reset()         { *m = ShareRecord{} }
func (m *ShareRecord) String() string { return proto.CompactTextString(m) }
func (*ShareRecord) ProtoMessage()    {}
func (*ShareRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1a1a1687f484a21, []int{4}
}
func (m *ShareRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShareRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShareRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShareRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareRecord.Merge(m, src)
}
func (m *ShareRecord) XXX_Size() int {
	return m.Size()
}
func (m *ShareRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ShareRecord proto.InternalMessageInfo

func (m *ShareRecord) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *ShareRecord) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kava.swap.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "kava.swap.v1beta1.Params")
	proto.RegisterType((*AllowedPool)(nil), "kava.swap.v1beta1.AllowedPool")
	proto.RegisterType((*PoolRecord)(nil), "kava.swap.v1beta1.PoolRecord")
	proto.RegisterType((*ShareRecord)(nil), "kava.swap.v1beta1.ShareRecord")
}

func init() { proto.RegisterFile("kava/swap/v1beta1/genesis.proto", fileDescriptor_b1a1a1687f484a21) }

var fileDescriptor_b1a1a1687f484a21 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x8b, 0xd3, 0x4e,
	0x14, 0xc7, 0x9b, 0x76, 0xe9, 0xfe, 0x3a, 0xe9, 0xf2, 0x63, 0xe3, 0x22, 0x51, 0x76, 0x93, 0x12,
	0x70, 0xed, 0xc1, 0x4d, 0xd8, 0x15, 0x11, 0xf6, 0x64, 0xa3, 0x28, 0x7b, 0x72, 0x49, 0x0f, 0x82,
	0x08, 0x61, 0xd2, 0x8c, 0xdd, 0xd0, 0xa4, 0x13, 0x32, 0x63, 0x6b, 0xf1, 0xa4, 0x7f, 0x81, 0x47,
	0x2f, 0x82, 0x7f, 0x8b, 0xa7, 0x3d, 0xee, 0x51, 0x3c, 0x04, 0x69, 0x0f, 0x42, 0xbd, 0xf5, 0x2f,
	0x90, 0x99, 0x4c, 0x37, 0x09, 0xad, 0xe2, 0xea, 0xa9, 0x79, 0xf3, 0xde, 0xfb, 0x7e, 0xde, 0x7b,
	0xf3, 0x3a, 0x40, 0x1f, 0xc0, 0x11, 0xb4, 0xc8, 0x18, 0xc6, 0xd6, 0xe8, 0xd0, 0x43, 0x14, 0x1e,
	0x5a, 0x7d, 0x34, 0x44, 0x24, 0x20, 0x66, 0x9c, 0x60, 0x8a, 0x95, 0x6d, 0x16, 0x60, 0xb2, 0x00,
	0x53, 0x04, 0xdc, 0xdc, 0xe9, 0xe3, 0x3e, 0xe6, 0x5e, 0x8b, 0x7d, 0x65, 0x81, 0xc6, 0xc7, 0x2a,
	0x68, 0x3e, 0xc9, 0x52, 0xbb, 0x14, 0x52, 0xa4, 0xdc, 0x07, 0xf5, 0x18, 0x26, 0x30, 0x22, 0xaa,
	0xd4, 0x92, 0xda, 0xf2, 0xd1, 0x0d, 0x73, 0x45, 0xca, 0x3c, 0xe5, 0x01, 0xf6, 0xc6, 0x79, 0xaa,
	0x57, 0x1c, 0x11, 0xae, 0xf4, 0x41, 0x33, 0xc6, 0x38, 0x74, 0x13, 0xd4, 0xc3, 0x89, 0x4f, 0xd4,
	0x6a, 0xab, 0xd6, 0x96, 0x8f, 0xf6, 0xd6, 0xa5, 0x63, 0x1c, 0x3a, 0x3c, 0xca, 0xbe, 0xc5, 0x24,
	0x16, 0xa9, 0xbe, 0x37, 0x81, 0x51, 0x78, 0x6c, 0x14, 0x05, 0xee, 0xe0, 0x28, 0xa0, 0x28, 0x8a,
	0xe9, 0xc4, 0x70, 0xe4, 0xf8, 0x32, 0x85, 0x28, 0x03, 0xb0, 0x45, 0xce, 0x60, 0x82, 0x2e, 0x49,
	0x35, 0x4e, 0xd2, 0xd6, 0x90, 0xba, 0x2c, 0x4e, 0xa0, 0xf6, 0x05, 0x4a, 0xcb, 0x50, 0x25, 0x89,
	0x22, 0xab, 0x49, 0xf2, 0x24, 0x62, 0xfc, 0x90, 0x40, 0xfd, 0x74, 0xd9, 0xe0, 0x16, 0x0c, 0x43,
	0x3c, 0x46, 0xbe, 0xcb, 0xca, 0x61, 0x03, 0xfa, 0x15, 0xb7, 0x93, 0xc5, 0xb1, 0x46, 0x6d, 0x23,
	0x67, 0x96, 0xd2, 0x4b, 0x4c, 0x98, 0x27, 0x10, 0x25, 0x01, 0xff, 0x31, 0x35, 0xf7, 0x25, 0x42,
	0x6a, 0xb5, 0x25, 0xb5, 0x9b, 0xf6, 0x33, 0x56, 0xfb, 0xd7, 0x54, 0xdf, 0xef, 0x07, 0xf4, 0xec,
	0x95, 0x67, 0xf6, 0x70, 0x64, 0xf5, 0x30, 0x89, 0x30, 0x11, 0x3f, 0x07, 0xc4, 0x1f, 0x58, 0x74,
	0x12, 0x23, 0x62, 0x3e, 0x42, 0xbd, 0x79, 0xaa, 0x2b, 0x4b, 0x85, 0x1c, 0xb3, 0x48, 0xf5, 0xff,
	0x45, 0xef, 0xc2, 0x67, 0x38, 0x9b, 0xec, 0xf3, 0x31, 0x42, 0xc7, 0x1b, 0x1f, 0x3e, 0xe9, 0x15,
	0xe3, 0xad, 0x04, 0xe4, 0x42, 0xed, 0xca, 0x3d, 0xb0, 0x49, 0xf1, 0x00, 0x0d, 0x5d, 0xc8, 0xb7,
	0xa1, 0x61, 0xef, 0x2e, 0x52, 0x5d, 0xcd, 0x44, 0x84, 0xa3, 0xd8, 0x46, 0x9d, 0x9f, 0x75, 0xf2,
	0x34, 0x8f, 0xd7, 0xbf, 0x26, 0xcd, 0x5b, 0x4d, 0xb3, 0x45, 0x0d, 0x9f, 0x6b, 0x00, 0xe4, 0x1b,
	0xc2, 0xb4, 0xf8, 0x56, 0x04, 0xfe, 0x6a, 0x09, 0xc2, 0x51, 0xd2, 0x62, 0x67, 0x27, 0xbe, 0xf2,
	0x06, 0x80, 0x04, 0x11, 0x94, 0x8c, 0x10, 0x71, 0xa1, 0x98, 0xe2, 0x0b, 0x31, 0xc5, 0xdb, 0x7f,
	0x30, 0xc5, 0x87, 0x38, 0x18, 0xce, 0x53, 0x7d, 0x27, 0x17, 0x29, 0x0d, 0x72, 0x3b, 0x2b, 0x20,
	0xf7, 0x1a, 0x4e, 0x63, 0x69, 0x74, 0x4a, 0x70, 0x4f, 0xad, 0xfd, 0x3b, 0xdc, 0xfb, 0x2d, 0xdc,
	0x2b, 0xc0, 0x6d, 0xe5, 0x9d, 0x04, 0x9a, 0x14, 0x53, 0x18, 0xba, 0x7c, 0x91, 0x89, 0xba, 0xc1,
	0xf9, 0xee, 0x15, 0x56, 0xe8, 0x64, 0x48, 0xe7, 0xa9, 0x7e, 0xbd, 0xa8, 0x52, 0x2a, 0xe0, 0xda,
	0xf2, 0x2a, 0x73, 0xbf, 0xe1, 0xc8, 0xdc, 0xec, 0x66, 0xd6, 0x77, 0x09, 0xc8, 0x85, 0x3f, 0x9f,
	0xb2, 0x0b, 0x1a, 0x3e, 0x8a, 0x31, 0x09, 0x28, 0x4e, 0xb2, 0x7b, 0x74, 0xf2, 0x83, 0xe2, 0x1d,
	0x57, 0xaf, 0x70, 0xc7, 0xac, 0xd3, 0x8c, 0xee, 0xe2, 0xf1, 0x10, 0xf9, 0x62, 0xd2, 0x7f, 0xd1,
	0x69, 0x51, 0x65, 0x5d, 0xa7, 0x45, 0xbf, 0xe1, 0xc8, 0x99, 0xf9, 0x94, 0x59, 0xf6, 0x83, 0xf3,
	0xa9, 0x26, 0x5d, 0x4c, 0x35, 0xe9, 0xdb, 0x54, 0x93, 0xde, 0xcf, 0xb4, 0xca, 0xc5, 0x4c, 0xab,
	0x7c, 0x99, 0x69, 0x95, 0xe7, 0x45, 0x3e, 0x7b, 0x22, 0x0e, 0x42, 0xe8, 0x11, 0xfe, 0x65, 0xbd,
	0xce, 0xde, 0x6e, 0x5e, 0x83, 0x57, 0xe7, 0x2f, 0xf1, 0xdd, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xa5, 0x3e, 0x8a, 0x53, 0xd5, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShareRecords) > 0 {
		for iNdEx := len(m.ShareRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShareRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PoolRecords) > 0 {
		for iNdEx := len(m.PoolRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.AllowedPools) > 0 {
		for iNdEx := len(m.AllowedPools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedPools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AllowedPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenB) > 0 {
		i -= len(m.TokenB)
		copy(dAtA[i:], m.TokenB)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokenB)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenA) > 0 {
		i -= len(m.TokenA)
		copy(dAtA[i:], m.TokenA)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokenA)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PoolRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalShares.Size()
		i -= size
		if _, err := m.TotalShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ReservesB.Size()
		i -= size
		if _, err := m.ReservesB.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.ReservesA.Size()
		i -= size
		if _, err := m.ReservesA.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ShareRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShareRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShareRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PoolRecords) > 0 {
		for _, e := range m.PoolRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ShareRecords) > 0 {
		for _, e := range m.ShareRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AllowedPools) > 0 {
		for _, e := range m.AllowedPools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.SwapFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *AllowedPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenA)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.TokenB)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *PoolRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ReservesA.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ReservesB.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TotalShares.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ShareRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolRecords = append(m.PoolRecords, PoolRecord{})
			if err := m.PoolRecords[len(m.PoolRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareRecords = append(m.ShareRecords, ShareRecord{})
			if err := m.ShareRecords[len(m.ShareRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedPools = append(m.AllowedPools, &AllowedPool{})
			if err := m.AllowedPools[len(m.AllowedPools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *AllowedPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: AllowedPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *PoolRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PoolRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesA", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesB", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesB.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ShareRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ShareRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShareRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesOwned", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
