// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/cassandra/clearsnapshot/v1/clearsnapshot.proto

package v1 // import "github.com/linuxuser586/apis/grpc/cassandra/clearsnapshot/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "github.com/linuxuser586/apis/grpc/cassandra/nodetool/v1"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Args         *v1.Args `protobuf:"bytes,1,opt,name=args" json:"args,omitempty"`
	SnapshotName string   `protobuf:"bytes,2,opt,name=snapshot_name,json=snapshotName,proto3" json:"snapshot_name,omitempty"`
	Keyspaces    string   `protobuf:"bytes,3,opt,name=keyspaces,proto3" json:"keyspaces,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_clearsnapshot_dc7dc0b2fd278d54, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetArgs() *v1.Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Request) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *Request) GetKeyspaces() string {
	if m != nil {
		return m.Keyspaces
	}
	return ""
}

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_clearsnapshot_dc7dc0b2fd278d54, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "linuxuser586.cassandra.clearnsnapshot.v1.Request")
	proto.RegisterType((*Response)(nil), "linuxuser586.cassandra.clearnsnapshot.v1.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClearSnapshotServiceClient is the client API for ClearSnapshotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClearSnapshotServiceClient interface {
	Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type clearSnapshotServiceClient struct {
	cc *grpc.ClientConn
}

func NewClearSnapshotServiceClient(cc *grpc.ClientConn) ClearSnapshotServiceClient {
	return &clearSnapshotServiceClient{cc}
}

func (c *clearSnapshotServiceClient) Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/linuxuser586.cassandra.clearnsnapshot.v1.ClearSnapshotService/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClearSnapshotServiceServer is the server API for ClearSnapshotService service.
type ClearSnapshotServiceServer interface {
	Run(context.Context, *Request) (*Response, error)
}

func RegisterClearSnapshotServiceServer(s *grpc.Server, srv ClearSnapshotServiceServer) {
	s.RegisterService(&_ClearSnapshotService_serviceDesc, srv)
}

func _ClearSnapshotService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClearSnapshotServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxuser586.cassandra.clearnsnapshot.v1.ClearSnapshotService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClearSnapshotServiceServer).Run(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClearSnapshotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "linuxuser586.cassandra.clearnsnapshot.v1.ClearSnapshotService",
	HandlerType: (*ClearSnapshotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _ClearSnapshotService_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/cassandra/clearsnapshot/v1/clearsnapshot.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Args != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClearsnapshot(dAtA, i, uint64(m.Args.Size()))
		n1, err := m.Args.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.SnapshotName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClearsnapshot(dAtA, i, uint64(len(m.SnapshotName)))
		i += copy(dAtA[i:], m.SnapshotName)
	}
	if len(m.Keyspaces) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClearsnapshot(dAtA, i, uint64(len(m.Keyspaces)))
		i += copy(dAtA[i:], m.Keyspaces)
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClearsnapshot(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintClearsnapshot(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Args != nil {
		l = m.Args.Size()
		n += 1 + l + sovClearsnapshot(uint64(l))
	}
	l = len(m.SnapshotName)
	if l > 0 {
		n += 1 + l + sovClearsnapshot(uint64(l))
	}
	l = len(m.Keyspaces)
	if l > 0 {
		n += 1 + l + sovClearsnapshot(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovClearsnapshot(uint64(l))
	}
	return n
}

func sovClearsnapshot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClearsnapshot(x uint64) (n int) {
	return sovClearsnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClearsnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClearsnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClearsnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Args == nil {
				m.Args = &v1.Args{}
			}
			if err := m.Args.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClearsnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClearsnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keyspaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClearsnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClearsnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keyspaces = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClearsnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClearsnapshot
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClearsnapshot
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClearsnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClearsnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClearsnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClearsnapshot
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
func skipClearsnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClearsnapshot
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
					return 0, ErrIntOverflowClearsnapshot
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClearsnapshot
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthClearsnapshot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClearsnapshot
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipClearsnapshot(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthClearsnapshot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClearsnapshot   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("protos/cassandra/clearsnapshot/v1/clearsnapshot.proto", fileDescriptor_clearsnapshot_dc7dc0b2fd278d54)
}

var fileDescriptor_clearsnapshot_dc7dc0b2fd278d54 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x73, 0xff, 0xfe, 0xb1, 0xf6, 0xd4, 0xe5, 0x70, 0x08, 0x45, 0x42, 0xa9, 0x0e, 0x01,
	0xe1, 0x42, 0x2a, 0x15, 0x87, 0x2e, 0xea, 0xee, 0x90, 0x82, 0x83, 0x8b, 0x5c, 0xd3, 0x97, 0x34,
	0xd8, 0xdc, 0x9d, 0xf7, 0x26, 0x41, 0xbf, 0x80, 0x83, 0x93, 0x1f, 0xcb, 0xb1, 0xa3, 0xa3, 0xb4,
	0x5f, 0x44, 0x7a, 0x6d, 0xa2, 0x22, 0x8a, 0x6e, 0xc9, 0xc3, 0xf3, 0xfc, 0xde, 0xe7, 0xbd, 0x97,
	0xf6, 0xb5, 0x51, 0xb9, 0xc2, 0x20, 0x16, 0x88, 0x42, 0x8e, 0x8d, 0x08, 0xe2, 0x29, 0x08, 0x83,
	0x52, 0x68, 0x9c, 0xa8, 0x3c, 0x28, 0xc3, 0xcf, 0x02, 0xb7, 0x7e, 0xe6, 0x4f, 0x53, 0x59, 0xdc,
	0x15, 0x08, 0xa6, 0x7f, 0x72, 0xcc, 0xeb, 0x30, 0xb7, 0x5e, 0x59, 0x9b, 0xcb, 0xb0, 0x7d, 0xf8,
	0x65, 0x80, 0x54, 0x63, 0xc8, 0x95, 0x9a, 0x2e, 0xd9, 0xd5, 0xf7, 0x0a, 0xdb, 0x7d, 0x24, 0xb4,
	0x19, 0xc1, 0x6d, 0x01, 0x98, 0xb3, 0x01, 0xfd, 0x2f, 0x4c, 0x82, 0x2e, 0xe9, 0x10, 0x7f, 0xab,
	0xe7, 0xf3, 0x6f, 0x26, 0xd6, 0x84, 0x32, 0xe4, 0xa7, 0x26, 0xc1, 0xc8, 0xa6, 0xd8, 0x3e, 0xdd,
	0xa9, 0x5a, 0x5c, 0x4b, 0x91, 0x81, 0xfb, 0xaf, 0x43, 0xfc, 0x56, 0xb4, 0x5d, 0x89, 0x17, 0x22,
	0x03, 0xb6, 0x47, 0x5b, 0x37, 0x70, 0x8f, 0x5a, 0xc4, 0x80, 0x6e, 0xc3, 0x1a, 0xde, 0x85, 0xee,
	0x01, 0xdd, 0x8c, 0x00, 0xb5, 0x92, 0x08, 0xcc, 0xa5, 0xcd, 0x0c, 0x10, 0x45, 0x02, 0xb6, 0x4f,
	0x2b, 0xaa, 0x7e, 0x7b, 0x0f, 0x84, 0xee, 0x9e, 0x2f, 0xb7, 0x1e, 0xae, 0xc9, 0x43, 0x30, 0x65,
	0x1a, 0x03, 0x93, 0xb4, 0x11, 0x15, 0x92, 0x85, 0xfc, 0xb7, 0x4f, 0xc5, 0xd7, 0x9b, 0xb7, 0x7b,
	0x7f, 0x89, 0xac, 0x0a, 0x76, 0x9d, 0xb3, 0xcb, 0xe7, 0xb9, 0x47, 0x66, 0x73, 0x8f, 0xbc, 0xce,
	0x3d, 0xf2, 0xb4, 0xf0, 0x9c, 0xd9, 0xc2, 0x73, 0x5e, 0x16, 0x9e, 0x73, 0x35, 0x48, 0xd2, 0x7c,
	0x52, 0x8c, 0x78, 0xac, 0xb2, 0xe0, 0x23, 0x39, 0x10, 0x3a, 0xc5, 0x20, 0x31, 0x3a, 0xfe, 0xe1,
	0xfc, 0xa3, 0x0d, 0x7b, 0x9a, 0xa3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xd6, 0x21, 0xf0,
	0x2a, 0x02, 0x00, 0x00,
}
