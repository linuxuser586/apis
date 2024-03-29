// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/cassandra/nodetool/v1/nodetool.proto

package v1 // import "github.com/linuxuser586/apis/grpc/cassandra/nodetool/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Args struct {
	Host         string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port         string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password     string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	PasswordFile string `protobuf:"bytes,5,opt,name=password_file,json=passwordFile,proto3" json:"password_file,omitempty"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}
func (*Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodetool_2043173220c88596, []int{0}
}
func (m *Args) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Args.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args.Merge(dst, src)
}
func (m *Args) XXX_Size() int {
	return m.Size()
}
func (m *Args) XXX_DiscardUnknown() {
	xxx_messageInfo_Args.DiscardUnknown(m)
}

var xxx_messageInfo_Args proto.InternalMessageInfo

func (m *Args) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Args) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Args) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Args) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Args) GetPasswordFile() string {
	if m != nil {
		return m.PasswordFile
	}
	return ""
}

func init() {
	proto.RegisterType((*Args)(nil), "linuxuser586.cassandra.nodetool.v1.Args")
}
func (m *Args) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Args) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Host) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNodetool(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	if len(m.Port) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNodetool(dAtA, i, uint64(len(m.Port)))
		i += copy(dAtA[i:], m.Port)
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNodetool(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintNodetool(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.PasswordFile) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintNodetool(dAtA, i, uint64(len(m.PasswordFile)))
		i += copy(dAtA[i:], m.PasswordFile)
	}
	return i, nil
}

func encodeVarintNodetool(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Args) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovNodetool(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovNodetool(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovNodetool(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovNodetool(uint64(l))
	}
	l = len(m.PasswordFile)
	if l > 0 {
		n += 1 + l + sovNodetool(uint64(l))
	}
	return n
}

func sovNodetool(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNodetool(x uint64) (n int) {
	return sovNodetool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Args) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodetool
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
			return fmt.Errorf("proto: Args: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Args: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodetool
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
				return ErrInvalidLengthNodetool
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodetool
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
				return ErrInvalidLengthNodetool
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodetool
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
				return ErrInvalidLengthNodetool
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodetool
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
				return ErrInvalidLengthNodetool
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PasswordFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodetool
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
				return ErrInvalidLengthNodetool
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PasswordFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodetool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodetool
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
func skipNodetool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodetool
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
					return 0, ErrIntOverflowNodetool
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
					return 0, ErrIntOverflowNodetool
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
				return 0, ErrInvalidLengthNodetool
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNodetool
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
				next, err := skipNodetool(dAtA[start:])
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
	ErrInvalidLengthNodetool = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodetool   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("protos/cassandra/nodetool/v1/nodetool.proto", fileDescriptor_nodetool_2043173220c88596)
}

var fileDescriptor_nodetool_2043173220c88596 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x4e, 0x2c, 0x2e, 0x4e, 0xcc, 0x4b, 0x29, 0x4a, 0xd4, 0xcf, 0xcb, 0x4f,
	0x49, 0x2d, 0xc9, 0xcf, 0xcf, 0xd1, 0x2f, 0x33, 0x84, 0xb3, 0xf5, 0xc0, 0xaa, 0x84, 0x94, 0x72,
	0x32, 0xf3, 0x4a, 0x2b, 0x4a, 0x8b, 0x53, 0x8b, 0x4c, 0x2d, 0xcc, 0xf4, 0xe0, 0x5a, 0xf4, 0xe0,
	0xca, 0xca, 0x0c, 0x95, 0xba, 0x19, 0xb9, 0x58, 0x1c, 0x8b, 0xd2, 0x8b, 0x85, 0x84, 0xb8, 0x58,
	0x32, 0xf2, 0x8b, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x41,
	0x7e, 0x51, 0x89, 0x04, 0x13, 0x44, 0x0c, 0xc4, 0x16, 0x92, 0xe2, 0xe2, 0x00, 0x99, 0x98, 0x97,
	0x98, 0x9b, 0x2a, 0xc1, 0x0c, 0x16, 0x87, 0xf3, 0x41, 0x72, 0x05, 0x89, 0xc5, 0xc5, 0xe5, 0xf9,
	0x45, 0x29, 0x12, 0x2c, 0x10, 0x39, 0x18, 0x5f, 0x48, 0x99, 0x8b, 0x17, 0xc6, 0x8e, 0x4f, 0xcb,
	0xcc, 0x49, 0x95, 0x60, 0x05, 0x2b, 0xe0, 0x81, 0x09, 0xba, 0x65, 0xe6, 0xa4, 0x3a, 0x05, 0x9e,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xb2, 0xb7, 0xf4, 0x13, 0x0b, 0x32, 0x8b, 0xf5, 0xd3, 0x8b,
	0x0a, 0x92, 0xb1, 0x87, 0x49, 0x12, 0x1b, 0x38, 0x2c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0xcd, 0xc1, 0xb2, 0x3a, 0x01, 0x00, 0x00,
}
