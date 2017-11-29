// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/proto/ca_service.proto

/*
	Package istio_v1_auth is a generated protocol buffer package.

	It is generated from these files:
		security/proto/ca_service.proto

	It has these top-level messages:
		Request
		Response
*/
package istio_v1_auth

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_rpc "github.com/googleapis/googleapis/google/rpc"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import strings "strings"
import reflect "reflect"

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
	// PEM-encoded certificate signing request
	CsrPem []byte `protobuf:"bytes,1,opt,name=csr_pem,json=csrPem,proto3" json:"csr_pem,omitempty"`
	// opaque credential for node agent
	NodeAgentCredential []byte `protobuf:"bytes,2,opt,name=node_agent_credential,json=nodeAgentCredential,proto3" json:"node_agent_credential,omitempty"`
	// type of the node_agent_credential (aws/gcp/onprem/custom...)
	CredentialType string `protobuf:"bytes,3,opt,name=credential_type,json=credentialType,proto3" json:"credential_type,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorCaService, []int{0} }

type Response struct {
	IsApproved      bool               `protobuf:"varint,1,opt,name=is_approved,json=isApproved,proto3" json:"is_approved,omitempty"`
	Status          *google_rpc.Status `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	SignedCertChain []byte             `protobuf:"bytes,3,opt,name=signed_cert_chain,json=signedCertChain,proto3" json:"signed_cert_chain,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorCaService, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "istio.v1.auth.Request")
	proto.RegisterType((*Response)(nil), "istio.v1.auth.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IstioCAService service

type IstioCAServiceClient interface {
	// A request object includes a PEM-encoded certificate signing request that
	// is generated on the Node Agent. Additionaly credential can be attached
	// within the request object for a server to authenticate the originating
	// node agent.
	HandleCSR(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type istioCAServiceClient struct {
	cc *grpc.ClientConn
}

func NewIstioCAServiceClient(cc *grpc.ClientConn) IstioCAServiceClient {
	return &istioCAServiceClient{cc}
}

func (c *istioCAServiceClient) HandleCSR(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/istio.v1.auth.IstioCAService/HandleCSR", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IstioCAService service

type IstioCAServiceServer interface {
	// A request object includes a PEM-encoded certificate signing request that
	// is generated on the Node Agent. Additionaly credential can be attached
	// within the request object for a server to authenticate the originating
	// node agent.
	HandleCSR(context.Context, *Request) (*Response, error)
}

func RegisterIstioCAServiceServer(s *grpc.Server, srv IstioCAServiceServer) {
	s.RegisterService(&_IstioCAService_serviceDesc, srv)
}

func _IstioCAService_HandleCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioCAServiceServer).HandleCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/istio.v1.auth.IstioCAService/HandleCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioCAServiceServer).HandleCSR(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _IstioCAService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "istio.v1.auth.IstioCAService",
	HandlerType: (*IstioCAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleCSR",
			Handler:    _IstioCAService_HandleCSR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/proto/ca_service.proto",
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
	if len(m.CsrPem) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.CsrPem)))
		i += copy(dAtA[i:], m.CsrPem)
	}
	if len(m.NodeAgentCredential) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.NodeAgentCredential)))
		i += copy(dAtA[i:], m.NodeAgentCredential)
	}
	if len(m.CredentialType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.CredentialType)))
		i += copy(dAtA[i:], m.CredentialType)
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
	if m.IsApproved {
		dAtA[i] = 0x8
		i++
		if m.IsApproved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Status != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaService(dAtA, i, uint64(m.Status.Size()))
		n1, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.SignedCertChain) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCaService(dAtA, i, uint64(len(m.SignedCertChain)))
		i += copy(dAtA[i:], m.SignedCertChain)
	}
	return i, nil
}

func encodeVarintCaService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	l = len(m.CsrPem)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	l = len(m.NodeAgentCredential)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	l = len(m.CredentialType)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.IsApproved {
		n += 2
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovCaService(uint64(l))
	}
	l = len(m.SignedCertChain)
	if l > 0 {
		n += 1 + l + sovCaService(uint64(l))
	}
	return n
}

func sovCaService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCaService(x uint64) (n int) {
	return sovCaService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Request) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Request{`,
		`CsrPem:` + fmt.Sprintf("%v", this.CsrPem) + `,`,
		`NodeAgentCredential:` + fmt.Sprintf("%v", this.NodeAgentCredential) + `,`,
		`CredentialType:` + fmt.Sprintf("%v", this.CredentialType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Response{`,
		`IsApproved:` + fmt.Sprintf("%v", this.IsApproved) + `,`,
		`Status:` + strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "google_rpc.Status", 1) + `,`,
		`SignedCertChain:` + fmt.Sprintf("%v", this.SignedCertChain) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCaService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaService
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
				return fmt.Errorf("proto: wrong wireType = %d for field CsrPem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CsrPem = append(m.CsrPem[:0], dAtA[iNdEx:postIndex]...)
			if m.CsrPem == nil {
				m.CsrPem = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAgentCredential", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAgentCredential = append(m.NodeAgentCredential[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeAgentCredential == nil {
				m.NodeAgentCredential = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
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
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredentialType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaService
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
				return ErrIntOverflowCaService
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsApproved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsApproved = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
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
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &google_rpc.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedCertChain", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCaService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignedCertChain = append(m.SignedCertChain[:0], dAtA[iNdEx:postIndex]...)
			if m.SignedCertChain == nil {
				m.SignedCertChain = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaService
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
func skipCaService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCaService
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
					return 0, ErrIntOverflowCaService
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
					return 0, ErrIntOverflowCaService
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
				return 0, ErrInvalidLengthCaService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCaService
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
				next, err := skipCaService(dAtA[start:])
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
	ErrInvalidLengthCaService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCaService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("security/proto/ca_service.proto", fileDescriptorCaService) }

var fileDescriptorCaService = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbd, 0x8e, 0x13, 0x31,
	0x14, 0x85, 0xc7, 0x20, 0x65, 0x77, 0xbd, 0xcb, 0xae, 0x30, 0x3f, 0x89, 0x52, 0x78, 0x57, 0xdb,
	0xb0, 0x4a, 0xe1, 0x51, 0x42, 0x8b, 0x84, 0xc2, 0x34, 0xd0, 0x20, 0xe4, 0x50, 0xd1, 0x58, 0xc6,
	0x73, 0x35, 0xb1, 0x94, 0x8c, 0x8d, 0xed, 0x19, 0x29, 0x15, 0x48, 0xbc, 0x00, 0x8f, 0xc1, 0xa3,
	0xa4, 0x4c, 0x49, 0xc9, 0x0c, 0x0d, 0x65, 0x1e, 0x01, 0xcd, 0x38, 0x22, 0x42, 0x74, 0xd6, 0x77,
	0x8e, 0xae, 0xef, 0x39, 0x17, 0x5f, 0x7b, 0x50, 0x95, 0xd3, 0x61, 0x93, 0x5a, 0x67, 0x82, 0x49,
	0x95, 0x14, 0x1e, 0x5c, 0xad, 0x15, 0xb0, 0x1e, 0x90, 0x07, 0xda, 0x07, 0x6d, 0x58, 0x3d, 0x65,
	0xb2, 0x0a, 0xcb, 0xf1, 0xb0, 0x30, 0xa6, 0x58, 0x41, 0xea, 0xac, 0x4a, 0x7d, 0x90, 0xa1, 0xf2,
	0xd1, 0x37, 0x7e, 0x5c, 0x98, 0xc2, 0xc4, 0x19, 0xdd, 0x2b, 0xd2, 0xdb, 0xcf, 0xf8, 0x84, 0xc3,
	0xa7, 0x0a, 0x7c, 0x20, 0x43, 0x7c, 0xa2, 0xbc, 0x13, 0x16, 0xd6, 0x23, 0x74, 0x83, 0xee, 0x2e,
	0xf8, 0x40, 0x79, 0xf7, 0x0e, 0xd6, 0x64, 0x86, 0x9f, 0x94, 0x26, 0x07, 0x21, 0x0b, 0x28, 0x83,
	0x50, 0x0e, 0x72, 0x28, 0x83, 0x96, 0xab, 0xd1, 0xbd, 0xde, 0xf6, 0xa8, 0x13, 0xe7, 0x9d, 0x96,
	0xfd, 0x95, 0xc8, 0x33, 0x7c, 0x75, 0x34, 0x8a, 0xb0, 0xb1, 0x30, 0xba, 0x7f, 0x83, 0xee, 0xce,
	0xf8, 0xe5, 0x11, 0xbf, 0xdf, 0x58, 0xb8, 0xfd, 0x8a, 0xf0, 0x29, 0x07, 0x6f, 0x4d, 0xe9, 0x81,
	0x5c, 0xe3, 0x73, 0xed, 0x85, 0xb4, 0xd6, 0x99, 0x1a, 0xf2, 0x7e, 0x8d, 0x53, 0x8e, 0xb5, 0x9f,
	0x1f, 0x08, 0x99, 0xe0, 0x41, 0x0c, 0xd5, 0xff, 0x7d, 0x3e, 0x23, 0x2c, 0xc6, 0x65, 0xce, 0x2a,
	0xb6, 0xe8, 0x15, 0x7e, 0x70, 0x90, 0x09, 0x7e, 0xe8, 0x75, 0x51, 0x42, 0x2e, 0x14, 0xb8, 0x20,
	0xd4, 0x52, 0xea, 0xb2, 0x5f, 0xe2, 0x82, 0x5f, 0x45, 0x21, 0x03, 0x17, 0xb2, 0x0e, 0xcf, 0xde,
	0xe2, 0xcb, 0x37, 0x5d, 0x8d, 0xd9, 0x7c, 0x11, 0xcb, 0x25, 0x2f, 0xf0, 0xd9, 0x6b, 0x59, 0xe6,
	0x2b, 0xc8, 0x16, 0x9c, 0x3c, 0x65, 0xff, 0x94, 0xcc, 0x0e, 0x95, 0x8d, 0x87, 0xff, 0xf1, 0x18,
	0xe4, 0xd5, 0xcb, 0x6d, 0x43, 0x93, 0x5d, 0x43, 0x93, 0x1f, 0x0d, 0x4d, 0xf6, 0x0d, 0x4d, 0xbe,
	0xb4, 0x14, 0x7d, 0x6f, 0x69, 0xb2, 0x6d, 0x29, 0xda, 0xb5, 0x14, 0xfd, 0x6c, 0x29, 0xfa, 0xdd,
	0xd2, 0x64, 0xdf, 0x52, 0xf4, 0xed, 0x17, 0x4d, 0x3e, 0xc4, 0x33, 0x8a, 0x7a, 0x2a, 0xba, 0x49,
	0x1f, 0x07, 0xfd, 0x79, 0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x02, 0x68, 0x59, 0x98, 0xff,
	0x01, 0x00, 0x00,
}
