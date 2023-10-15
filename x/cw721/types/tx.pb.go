// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: uptick/cw721/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

// MsgConvertNFT defines a Msg to convert a native Cosmos nft to a CW721 token
type MsgConvertNFT struct {
	// nft classID to cnvert to CW721
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// nftID to cnvert to CW721
	NftIds []string `protobuf:"bytes,2,rep,name=nft_ids,json=nftIds,proto3" json:"nft_ids,omitempty"`
	// recipient hex address to receive CW721 token
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// cosmos bech32 address from the owner of the given Cosmos coins
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// CW721 token contract address registered in a token pair
	ContractAddress string `protobuf:"bytes,5,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// CW721 token id registered in a token pair
	TokenIds []string `protobuf:"bytes,6,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
}

func (m *MsgConvertNFT) Reset()         { *m = MsgConvertNFT{} }
func (m *MsgConvertNFT) String() string { return proto.CompactTextString(m) }
func (*MsgConvertNFT) ProtoMessage()    {}
func (*MsgConvertNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eed5155f32633f2, []int{0}
}
func (m *MsgConvertNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertNFT.Merge(m, src)
}
func (m *MsgConvertNFT) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertNFT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertNFT proto.InternalMessageInfo

func (m *MsgConvertNFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *MsgConvertNFT) GetNftIds() []string {
	if m != nil {
		return m.NftIds
	}
	return nil
}

func (m *MsgConvertNFT) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertNFT) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgConvertNFT) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgConvertNFT) GetTokenIds() []string {
	if m != nil {
		return m.TokenIds
	}
	return nil
}

// MsgConvertNFTResponse returns no fields
type MsgConvertNFTResponse struct {
}

func (m *MsgConvertNFTResponse) Reset()         { *m = MsgConvertNFTResponse{} }
func (m *MsgConvertNFTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertNFTResponse) ProtoMessage()    {}
func (*MsgConvertNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eed5155f32633f2, []int{1}
}
func (m *MsgConvertNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertNFTResponse.Merge(m, src)
}
func (m *MsgConvertNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertNFTResponse proto.InternalMessageInfo

// MsgConvertCW721 defines a Msg to convert a CW721 token to a native Cosmos
// nft.
type MsgConvertCW721 struct {
	// CW721 token contract address registered in a token pair
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// tokenID to convert
	TokenIds []string `protobuf:"bytes,2,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
	// bech32 address to receive native Cosmos coins
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender hex address from the owner of the given CW721 tokens
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// nft classID to cnvert to CW721
	ClassId string `protobuf:"bytes,5,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// nftID to cnvert to CW721
	NftIds []string `protobuf:"bytes,6,rep,name=nft_ids,json=nftIds,proto3" json:"nft_ids,omitempty"`
}

func (m *MsgConvertCW721) Reset()         { *m = MsgConvertCW721{} }
func (m *MsgConvertCW721) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCW721) ProtoMessage()    {}
func (*MsgConvertCW721) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eed5155f32633f2, []int{2}
}
func (m *MsgConvertCW721) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCW721) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCW721.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCW721) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCW721.Merge(m, src)
}
func (m *MsgConvertCW721) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCW721) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCW721.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCW721 proto.InternalMessageInfo

func (m *MsgConvertCW721) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgConvertCW721) GetTokenIds() []string {
	if m != nil {
		return m.TokenIds
	}
	return nil
}

func (m *MsgConvertCW721) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertCW721) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgConvertCW721) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *MsgConvertCW721) GetNftIds() []string {
	if m != nil {
		return m.NftIds
	}
	return nil
}

// MsgConvertCW721Response returns no fields
type MsgConvertCW721Response struct {
}

func (m *MsgConvertCW721Response) Reset()         { *m = MsgConvertCW721Response{} }
func (m *MsgConvertCW721Response) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCW721Response) ProtoMessage()    {}
func (*MsgConvertCW721Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eed5155f32633f2, []int{3}
}
func (m *MsgConvertCW721Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCW721Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCW721Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCW721Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCW721Response.Merge(m, src)
}
func (m *MsgConvertCW721Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCW721Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCW721Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCW721Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvertNFT)(nil), "uptick.cw721.v1.MsgConvertNFT")
	proto.RegisterType((*MsgConvertNFTResponse)(nil), "uptick.cw721.v1.MsgConvertNFTResponse")
	proto.RegisterType((*MsgConvertCW721)(nil), "uptick.cw721.v1.MsgConvertCW721")
	proto.RegisterType((*MsgConvertCW721Response)(nil), "uptick.cw721.v1.MsgConvertCW721Response")
}

func init() { proto.RegisterFile("uptick/cw721/v1/tx.proto", fileDescriptor_7eed5155f32633f2) }

var fileDescriptor_7eed5155f32633f2 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0xad, 0x53, 0x2e, 0xd7, 0x5a, 0xa0, 0x22, 0x0b, 0x68, 0x2e, 0xa0, 0xd0, 0x0b, 0x12, 0xf4,
	0x06, 0x62, 0xa5, 0x0c, 0x37, 0xc3, 0x49, 0xa0, 0x0e, 0xd7, 0xa1, 0x02, 0x21, 0xb1, 0x44, 0xa9,
	0xed, 0x86, 0xa8, 0xc5, 0x8e, 0x6c, 0x37, 0x2d, 0x03, 0x0b, 0x12, 0x1b, 0x03, 0x12, 0x3f, 0x08,
	0xb1, 0x31, 0x56, 0x62, 0x61, 0x44, 0x2d, 0x3f, 0x04, 0xe1, 0xb4, 0xa5, 0x29, 0x50, 0x10, 0x9b,
	0xbf, 0xf7, 0x5e, 0xbe, 0xf7, 0x7d, 0x2f, 0x36, 0x74, 0x26, 0x99, 0x4e, 0xc9, 0x08, 0x93, 0xe9,
	0x69, 0x27, 0xc4, 0x79, 0x88, 0xf5, 0x2c, 0xc8, 0xa4, 0xd0, 0x02, 0x35, 0x0a, 0x26, 0x30, 0x4c,
	0x90, 0x87, 0xee, 0x8d, 0x44, 0x88, 0x64, 0xcc, 0x70, 0x9c, 0xa5, 0x38, 0xe6, 0x5c, 0xe8, 0x58,
	0xa7, 0x82, 0xab, 0x42, 0xee, 0x5e, 0x49, 0x44, 0x22, 0xcc, 0x11, 0xff, 0x38, 0x15, 0xa8, 0xff,
	0x01, 0xc0, 0x4b, 0xe7, 0x2a, 0x39, 0x13, 0x3c, 0x67, 0x52, 0xf7, 0x1e, 0x3e, 0x46, 0x47, 0xb0,
	0x46, 0xc6, 0xb1, 0x52, 0x51, 0x4a, 0x1d, 0xd0, 0x02, 0xed, 0x7a, 0xff, 0xd0, 0xd4, 0x5d, 0x8a,
	0x9a, 0xf0, 0x90, 0x0f, 0x75, 0x94, 0x52, 0xe5, 0x58, 0xad, 0x6a, 0xbb, 0xde, 0xb7, 0xf9, 0x50,
	0x77, 0xa9, 0x42, 0x2e, 0xac, 0x49, 0x46, 0x58, 0x9a, 0x33, 0xe9, 0x54, 0xcd, 0x37, 0x9b, 0x1a,
	0x5d, 0x83, 0xb6, 0x62, 0x9c, 0x32, 0xe9, 0x5c, 0x30, 0xcc, 0xaa, 0x42, 0x27, 0xf0, 0x32, 0x11,
	0x5c, 0xcb, 0x98, 0xe8, 0x28, 0xa6, 0x54, 0x32, 0xa5, 0x9c, 0x03, 0xa3, 0x68, 0xac, 0xf1, 0xfb,
	0x05, 0x8c, 0xae, 0xc3, 0xba, 0x16, 0x23, 0xc6, 0x8d, 0xb3, 0x6d, 0x9c, 0x6b, 0x06, 0xe8, 0x52,
	0xe5, 0x37, 0xe1, 0xd5, 0xd2, 0x02, 0x7d, 0xa6, 0x32, 0xc1, 0x15, 0xf3, 0x3f, 0x02, 0xd8, 0xf8,
	0xc9, 0x9c, 0x3d, 0x3d, 0xed, 0x84, 0xbf, 0x35, 0x05, 0xff, 0x60, 0x6a, 0x95, 0x4d, 0xff, 0x6b,
	0xe1, 0xed, 0x60, 0x0f, 0xfe, 0x18, 0xac, 0xbd, 0x1d, 0xac, 0x7f, 0x04, 0x9b, 0x3b, 0x2b, 0xac,
	0xd7, 0xeb, 0xbc, 0xb5, 0x60, 0xf5, 0x5c, 0x25, 0xe8, 0x15, 0x84, 0x5b, 0x7f, 0xcf, 0x0b, 0x76,
	0x6e, 0x45, 0x50, 0x0a, 0xc7, 0xbd, 0xbd, 0x9f, 0xdf, 0x84, 0x77, 0xe7, 0xf5, 0xe7, 0x6f, 0xef,
	0xad, 0x63, 0x74, 0x13, 0xff, 0x7a, 0xff, 0x30, 0x29, 0xf4, 0x11, 0x1f, 0x6a, 0xf4, 0x06, 0xc0,
	0x8b, 0xa5, 0x88, 0x5b, 0x7b, 0x1c, 0x8c, 0xc2, 0x6d, 0xff, 0x4d, 0xb1, 0x99, 0xe2, 0xc4, 0x4c,
	0x71, 0x0b, 0x1d, 0xef, 0x9b, 0xc2, 0x60, 0x0f, 0x1e, 0x7d, 0x5a, 0x78, 0x60, 0xbe, 0xf0, 0xc0,
	0xd7, 0x85, 0x07, 0xde, 0x2d, 0xbd, 0xca, 0x7c, 0xe9, 0x55, 0xbe, 0x2c, 0xbd, 0xca, 0xb3, 0xbb,
	0x49, 0xaa, 0x9f, 0x4f, 0x06, 0x01, 0x11, 0x2f, 0xf0, 0x13, 0xd3, 0xa6, 0xc7, 0xf4, 0x54, 0xc8,
	0xd1, 0xba, 0xe9, 0x6c, 0xd5, 0x56, 0xbf, 0xcc, 0x98, 0x1a, 0xd8, 0xe6, 0x61, 0xdc, 0xfb, 0x1e,
	0x00, 0x00, 0xff, 0xff, 0x3e, 0x15, 0x86, 0x44, 0x79, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertNFT mints a CW721 representation of the native Cosmos nft
	// that is registered on the token mapping.
	ConvertNFT(ctx context.Context, in *MsgConvertNFT, opts ...grpc.CallOption) (*MsgConvertNFTResponse, error)
	// ConvertCW721 mints a native Cosmos coin representation of the CW721 token
	// contract that is registered on the token mapping.
	ConvertCW721(ctx context.Context, in *MsgConvertCW721, opts ...grpc.CallOption) (*MsgConvertCW721Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertNFT(ctx context.Context, in *MsgConvertNFT, opts ...grpc.CallOption) (*MsgConvertNFTResponse, error) {
	out := new(MsgConvertNFTResponse)
	err := c.cc.Invoke(ctx, "/uptick.cw721.v1.Msg/ConvertNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConvertCW721(ctx context.Context, in *MsgConvertCW721, opts ...grpc.CallOption) (*MsgConvertCW721Response, error) {
	out := new(MsgConvertCW721Response)
	err := c.cc.Invoke(ctx, "/uptick.cw721.v1.Msg/ConvertCW721", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConvertNFT mints a CW721 representation of the native Cosmos nft
	// that is registered on the token mapping.
	ConvertNFT(context.Context, *MsgConvertNFT) (*MsgConvertNFTResponse, error)
	// ConvertCW721 mints a native Cosmos coin representation of the CW721 token
	// contract that is registered on the token mapping.
	ConvertCW721(context.Context, *MsgConvertCW721) (*MsgConvertCW721Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ConvertNFT(ctx context.Context, req *MsgConvertNFT) (*MsgConvertNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertNFT not implemented")
}
func (*UnimplementedMsgServer) ConvertCW721(ctx context.Context, req *MsgConvertCW721) (*MsgConvertCW721Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCW721 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ConvertNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uptick.cw721.v1.Msg/ConvertNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertNFT(ctx, req.(*MsgConvertNFT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConvertCW721_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertCW721)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertCW721(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uptick.cw721.v1.Msg/ConvertCW721",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertCW721(ctx, req.(*MsgConvertCW721))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uptick.cw721.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertNFT",
			Handler:    _Msg_ConvertNFT_Handler,
		},
		{
			MethodName: "ConvertCW721",
			Handler:    _Msg_ConvertCW721_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uptick/cw721/v1/tx.proto",
}

func (m *MsgConvertNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenIds) > 0 {
		for iNdEx := len(m.TokenIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TokenIds[iNdEx])
			copy(dAtA[i:], m.TokenIds[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.TokenIds[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftIds) > 0 {
		for iNdEx := len(m.NftIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.NftIds[iNdEx])
			copy(dAtA[i:], m.NftIds[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.NftIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgConvertCW721) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCW721) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCW721) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftIds) > 0 {
		for iNdEx := len(m.NftIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.NftIds[iNdEx])
			copy(dAtA[i:], m.NftIds[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.NftIds[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenIds) > 0 {
		for iNdEx := len(m.TokenIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TokenIds[iNdEx])
			copy(dAtA[i:], m.TokenIds[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.TokenIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertCW721Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCW721Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCW721Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgConvertNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.NftIds) > 0 {
		for _, s := range m.NftIds {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.TokenIds) > 0 {
		for _, s := range m.TokenIds {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgConvertNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgConvertCW721) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.TokenIds) > 0 {
		for _, s := range m.TokenIds {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.NftIds) > 0 {
		for _, s := range m.NftIds {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgConvertCW721Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgConvertNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftIds = append(m.NftIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIds = append(m.TokenIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertCW721) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertCW721: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCW721: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIds = append(m.TokenIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftIds = append(m.NftIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgConvertCW721Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgConvertCW721Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCW721Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
