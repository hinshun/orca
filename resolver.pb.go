// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/hinshun/ipcs/resolver.proto

package ipcs

import (
	context "context"
	fmt "fmt"
	types "github.com/containerd/containerd/api/types"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type Resolved struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Target               types.Descriptor `protobuf:"bytes,2,opt,name=target,proto3" json:"target"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Resolved) Reset()      { *m = Resolved{} }
func (*Resolved) ProtoMessage() {}
func (*Resolved) Descriptor() ([]byte, []int) {
	return fileDescriptor_e322347346474075, []int{0}
}
func (m *Resolved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resolved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resolved.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resolved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resolved.Merge(m, src)
}
func (m *Resolved) XXX_Size() int {
	return m.Size()
}
func (m *Resolved) XXX_DiscardUnknown() {
	xxx_messageInfo_Resolved.DiscardUnknown(m)
}

var xxx_messageInfo_Resolved proto.InternalMessageInfo

type ResolveRequest struct {
	Ref                  string   `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResolveRequest) Reset()      { *m = ResolveRequest{} }
func (*ResolveRequest) ProtoMessage() {}
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e322347346474075, []int{1}
}
func (m *ResolveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResolveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResolveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResolveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveRequest.Merge(m, src)
}
func (m *ResolveRequest) XXX_Size() int {
	return m.Size()
}
func (m *ResolveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveRequest proto.InternalMessageInfo

type ResolveResponse struct {
	Resolved             Resolved `protobuf:"bytes,1,opt,name=resolved,proto3" json:"resolved"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResolveResponse) Reset()      { *m = ResolveResponse{} }
func (*ResolveResponse) ProtoMessage() {}
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e322347346474075, []int{2}
}
func (m *ResolveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResolveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResolveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResolveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveResponse.Merge(m, src)
}
func (m *ResolveResponse) XXX_Size() int {
	return m.Size()
}
func (m *ResolveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Resolved)(nil), "contentd.services.resolver.v1.Resolved")
	proto.RegisterType((*ResolveRequest)(nil), "contentd.services.resolver.v1.ResolveRequest")
	proto.RegisterType((*ResolveResponse)(nil), "contentd.services.resolver.v1.ResolveResponse")
}

func init() {
	proto.RegisterFile("github.com/hinshun/ipcs/resolver.proto", fileDescriptor_e322347346474075)
}

var fileDescriptor_e322347346474075 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x63, 0xa8, 0x4a, 0x71, 0x25, 0x40, 0x16, 0x12, 0x55, 0x05, 0xa6, 0xca, 0x00, 0x5d,
	0xb0, 0xd5, 0xb2, 0xc1, 0x56, 0xb1, 0xb0, 0x66, 0x42, 0x15, 0x4b, 0x9a, 0xfc, 0x24, 0x96, 0xa8,
	0x6d, 0x6c, 0xb7, 0x12, 0x1b, 0x8f, 0xd7, 0x91, 0x91, 0x09, 0xd1, 0x3c, 0x09, 0x4a, 0x62, 0x42,
	0x17, 0x50, 0xb7, 0x1b, 0xee, 0xbe, 0xf3, 0xd9, 0xc6, 0x17, 0x99, 0x70, 0xf9, 0x62, 0xc6, 0x12,
	0x35, 0xe7, 0xb9, 0x90, 0x36, 0x5f, 0x48, 0x2e, 0x74, 0x62, 0xb9, 0x01, 0xab, 0x9e, 0x97, 0x60,
	0x98, 0x36, 0xca, 0x29, 0x72, 0x96, 0x28, 0xe9, 0x40, 0xba, 0x94, 0x59, 0x30, 0x4b, 0x91, 0x80,
	0x65, 0x8d, 0x63, 0x39, 0xea, 0x1f, 0x67, 0x2a, 0x53, 0x95, 0x93, 0x97, 0xaa, 0x0e, 0xf5, 0x6f,
	0x37, 0xe0, 0x65, 0x3e, 0x16, 0x12, 0x4c, 0xba, 0x29, 0x63, 0x2d, 0xb8, 0x7b, 0xd5, 0x60, 0x79,
	0x0a, 0x36, 0x31, 0x42, 0x3b, 0xe5, 0x1b, 0xc3, 0x29, 0xee, 0x44, 0x75, 0x43, 0x4a, 0x08, 0x6e,
	0xc9, 0x78, 0x0e, 0x3d, 0x34, 0x40, 0xc3, 0xfd, 0xa8, 0xd2, 0xe4, 0x06, 0xb7, 0x5d, 0x6c, 0x32,
	0x70, 0xbd, 0x9d, 0x01, 0x1a, 0x76, 0xc7, 0xa7, 0xec, 0x97, 0xcb, 0x2a, 0x26, 0xbb, 0x6b, 0x98,
	0x93, 0xd6, 0xea, 0xf3, 0x3c, 0x88, 0x7c, 0x22, 0x0c, 0xf1, 0x81, 0x67, 0x47, 0xf0, 0xb2, 0x00,
	0xeb, 0xc8, 0x11, 0xde, 0x35, 0xf0, 0xe4, 0x0b, 0x4a, 0x19, 0x3e, 0xe2, 0xc3, 0xc6, 0x63, 0xb5,
	0x92, 0x16, 0xc8, 0x3d, 0xee, 0xf8, 0xd1, 0x69, 0xe5, 0xec, 0x8e, 0x2f, 0xd9, 0xbf, 0xf7, 0xc2,
	0x7e, 0x16, 0xf8, 0xfe, 0x26, 0x3e, 0x76, 0xcd, 0x3a, 0x43, 0x72, 0xbc, 0xe7, 0x35, 0xb9, 0xda,
	0x8e, 0xe7, 0x4f, 0xdd, 0x67, 0xdb, 0xda, 0xeb, 0x01, 0x93, 0xd1, 0x6a, 0x4d, 0x83, 0x8f, 0x35,
	0x0d, 0xde, 0x0a, 0x8a, 0x56, 0x05, 0x45, 0xef, 0x05, 0x45, 0x5f, 0x05, 0x45, 0xd3, 0x93, 0x3f,
	0xfe, 0xc1, 0x43, 0x30, 0x6b, 0x57, 0xef, 0x71, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x1e,
	0x38, 0xf2, 0x2b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResolverClient is the client API for Resolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResolverClient interface {
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
}

type resolverClient struct {
	cc *grpc.ClientConn
}

func NewResolverClient(cc *grpc.ClientConn) ResolverClient {
	return &resolverClient{cc}
}

func (c *resolverClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, "/contentd.services.resolver.v1.Resolver/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolverServer is the server API for Resolver service.
type ResolverServer interface {
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
}

// UnimplementedResolverServer can be embedded to have forward compatible implementations.
type UnimplementedResolverServer struct {
}

func (*UnimplementedResolverServer) Resolve(ctx context.Context, req *ResolveRequest) (*ResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}

func RegisterResolverServer(s *grpc.Server, srv ResolverServer) {
	s.RegisterService(&_Resolver_serviceDesc, srv)
}

func _Resolver_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentd.services.resolver.v1.Resolver/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Resolver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contentd.services.resolver.v1.Resolver",
	HandlerType: (*ResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resolve",
			Handler:    _Resolver_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/hinshun/ipcs/resolver.proto",
}

func (m *Resolved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resolved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resolved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.Target.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResolver(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintResolver(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResolveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResolveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResolveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ref) > 0 {
		i -= len(m.Ref)
		copy(dAtA[i:], m.Ref)
		i = encodeVarintResolver(dAtA, i, uint64(len(m.Ref)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResolveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResolveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResolveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.Resolved.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResolver(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintResolver(dAtA []byte, offset int, v uint64) int {
	offset -= sovResolver(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resolved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovResolver(uint64(l))
	}
	l = m.Target.Size()
	n += 1 + l + sovResolver(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResolveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovResolver(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResolveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Resolved.Size()
	n += 1 + l + sovResolver(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovResolver(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResolver(x uint64) (n int) {
	return sovResolver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Resolved) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Resolved{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Target:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Target), "Descriptor", "types.Descriptor", 1), `&`, ``, 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResolveRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResolveRequest{`,
		`Ref:` + fmt.Sprintf("%v", this.Ref) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResolveResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResolveResponse{`,
		`Resolved:` + strings.Replace(strings.Replace(this.Resolved.String(), "Resolved", "Resolved", 1), `&`, ``, 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringResolver(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Resolved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResolver
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
			return fmt.Errorf("proto: Resolved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resolved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResolver
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
				return ErrInvalidLengthResolver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResolver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResolver
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
				return ErrInvalidLengthResolver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResolver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Target.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResolver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResolver
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResolver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResolveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResolver
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
			return fmt.Errorf("proto: ResolveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResolveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResolver
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
				return ErrInvalidLengthResolver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResolver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResolver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResolver
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResolver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResolveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResolver
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
			return fmt.Errorf("proto: ResolveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResolveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resolved", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResolver
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
				return ErrInvalidLengthResolver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResolver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resolved.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResolver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResolver
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthResolver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipResolver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResolver
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
					return 0, ErrIntOverflowResolver
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
					return 0, ErrIntOverflowResolver
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
				return 0, ErrInvalidLengthResolver
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResolver
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResolver
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResolver        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResolver          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResolver = fmt.Errorf("proto: unexpected end of group")
)
