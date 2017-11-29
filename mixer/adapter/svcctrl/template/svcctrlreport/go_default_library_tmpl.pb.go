// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bazel-out/local-fastbuild/genfiles/mixer/adapter/svcctrl/template/svcctrlreport/go_default_library_tmpl.proto

/*
	Package svcctrlreport is a generated protocol buffer package.

	It is generated from these files:
		bazel-out/local-fastbuild/genfiles/mixer/adapter/svcctrl/template/svcctrlreport/go_default_library_tmpl.proto

	It has these top-level messages:
		Type
		InstanceParam
*/
package svcctrlreport

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/v1/template"

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

// A template used by Google Service Control (svcctrl) adapter. The adapter
// generates metrics and logentry for each request based on the data point
// defined by this template.
//
// Config example:
// ```
// apiVersion: "config.istio.io/v1alpha2"
// kind: svcctrlreport
// metadata:
//   name: report
//   namespace: istio-system
// spec:
//   api_version : api.version | ""
//   api_operation : api.operation | ""
//   api_protocol : api.protocol | ""
//   api_service : api.service | ""
//   api_key : api.key | ""
//   request_time : request.time
//   request_method : request.method
//   request_path : request.path
//   request_bytes: request.size
//   response_time : response.time
//   response_code : response.code | 520
//   response_bytes : response.size | 0
//   response_latency : response.duration | "0ms"
// ```
type Type struct {
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorGoDefaultLibraryTmpl, []int{0} }

type InstanceParam struct {
	ApiVersion      string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	ApiOperation    string `protobuf:"bytes,2,opt,name=api_operation,json=apiOperation,proto3" json:"api_operation,omitempty"`
	ApiProtocol     string `protobuf:"bytes,3,opt,name=api_protocol,json=apiProtocol,proto3" json:"api_protocol,omitempty"`
	ApiService      string `protobuf:"bytes,4,opt,name=api_service,json=apiService,proto3" json:"api_service,omitempty"`
	ApiKey          string `protobuf:"bytes,5,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	RequestTime     string `protobuf:"bytes,6,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	RequestMethod   string `protobuf:"bytes,7,opt,name=request_method,json=requestMethod,proto3" json:"request_method,omitempty"`
	RequestPath     string `protobuf:"bytes,8,opt,name=request_path,json=requestPath,proto3" json:"request_path,omitempty"`
	RequestBytes    string `protobuf:"bytes,9,opt,name=request_bytes,json=requestBytes,proto3" json:"request_bytes,omitempty"`
	ResponseTime    string `protobuf:"bytes,10,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	ResponseCode    string `protobuf:"bytes,11,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	ResponseBytes   string `protobuf:"bytes,12,opt,name=response_bytes,json=responseBytes,proto3" json:"response_bytes,omitempty"`
	ResponseLatency string `protobuf:"bytes,13,opt,name=response_latency,json=responseLatency,proto3" json:"response_latency,omitempty"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorGoDefaultLibraryTmpl, []int{1}
}

func (m *InstanceParam) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *InstanceParam) GetApiOperation() string {
	if m != nil {
		return m.ApiOperation
	}
	return ""
}

func (m *InstanceParam) GetApiProtocol() string {
	if m != nil {
		return m.ApiProtocol
	}
	return ""
}

func (m *InstanceParam) GetApiService() string {
	if m != nil {
		return m.ApiService
	}
	return ""
}

func (m *InstanceParam) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *InstanceParam) GetRequestTime() string {
	if m != nil {
		return m.RequestTime
	}
	return ""
}

func (m *InstanceParam) GetRequestMethod() string {
	if m != nil {
		return m.RequestMethod
	}
	return ""
}

func (m *InstanceParam) GetRequestPath() string {
	if m != nil {
		return m.RequestPath
	}
	return ""
}

func (m *InstanceParam) GetRequestBytes() string {
	if m != nil {
		return m.RequestBytes
	}
	return ""
}

func (m *InstanceParam) GetResponseTime() string {
	if m != nil {
		return m.ResponseTime
	}
	return ""
}

func (m *InstanceParam) GetResponseCode() string {
	if m != nil {
		return m.ResponseCode
	}
	return ""
}

func (m *InstanceParam) GetResponseBytes() string {
	if m != nil {
		return m.ResponseBytes
	}
	return ""
}

func (m *InstanceParam) GetResponseLatency() string {
	if m != nil {
		return m.ResponseLatency
	}
	return ""
}

func init() {
	proto.RegisterType((*Type)(nil), "svcctrlreport.Type")
	proto.RegisterType((*InstanceParam)(nil), "svcctrlreport.InstanceParam")
}
func (this *Type) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Type)
	if !ok {
		that2, ok := that.(Type)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *InstanceParam) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InstanceParam)
	if !ok {
		that2, ok := that.(InstanceParam)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ApiVersion != that1.ApiVersion {
		return false
	}
	if this.ApiOperation != that1.ApiOperation {
		return false
	}
	if this.ApiProtocol != that1.ApiProtocol {
		return false
	}
	if this.ApiService != that1.ApiService {
		return false
	}
	if this.ApiKey != that1.ApiKey {
		return false
	}
	if this.RequestTime != that1.RequestTime {
		return false
	}
	if this.RequestMethod != that1.RequestMethod {
		return false
	}
	if this.RequestPath != that1.RequestPath {
		return false
	}
	if this.RequestBytes != that1.RequestBytes {
		return false
	}
	if this.ResponseTime != that1.ResponseTime {
		return false
	}
	if this.ResponseCode != that1.ResponseCode {
		return false
	}
	if this.ResponseBytes != that1.ResponseBytes {
		return false
	}
	if this.ResponseLatency != that1.ResponseLatency {
		return false
	}
	return true
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&svcctrlreport.Type{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceParam) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 17)
	s = append(s, "&svcctrlreport.InstanceParam{")
	s = append(s, "ApiVersion: "+fmt.Sprintf("%#v", this.ApiVersion)+",\n")
	s = append(s, "ApiOperation: "+fmt.Sprintf("%#v", this.ApiOperation)+",\n")
	s = append(s, "ApiProtocol: "+fmt.Sprintf("%#v", this.ApiProtocol)+",\n")
	s = append(s, "ApiService: "+fmt.Sprintf("%#v", this.ApiService)+",\n")
	s = append(s, "ApiKey: "+fmt.Sprintf("%#v", this.ApiKey)+",\n")
	s = append(s, "RequestTime: "+fmt.Sprintf("%#v", this.RequestTime)+",\n")
	s = append(s, "RequestMethod: "+fmt.Sprintf("%#v", this.RequestMethod)+",\n")
	s = append(s, "RequestPath: "+fmt.Sprintf("%#v", this.RequestPath)+",\n")
	s = append(s, "RequestBytes: "+fmt.Sprintf("%#v", this.RequestBytes)+",\n")
	s = append(s, "ResponseTime: "+fmt.Sprintf("%#v", this.ResponseTime)+",\n")
	s = append(s, "ResponseCode: "+fmt.Sprintf("%#v", this.ResponseCode)+",\n")
	s = append(s, "ResponseBytes: "+fmt.Sprintf("%#v", this.ResponseBytes)+",\n")
	s = append(s, "ResponseLatency: "+fmt.Sprintf("%#v", this.ResponseLatency)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGoDefaultLibraryTmpl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ApiVersion) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ApiVersion)))
		i += copy(dAtA[i:], m.ApiVersion)
	}
	if len(m.ApiOperation) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ApiOperation)))
		i += copy(dAtA[i:], m.ApiOperation)
	}
	if len(m.ApiProtocol) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ApiProtocol)))
		i += copy(dAtA[i:], m.ApiProtocol)
	}
	if len(m.ApiService) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ApiService)))
		i += copy(dAtA[i:], m.ApiService)
	}
	if len(m.ApiKey) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ApiKey)))
		i += copy(dAtA[i:], m.ApiKey)
	}
	if len(m.RequestTime) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.RequestTime)))
		i += copy(dAtA[i:], m.RequestTime)
	}
	if len(m.RequestMethod) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.RequestMethod)))
		i += copy(dAtA[i:], m.RequestMethod)
	}
	if len(m.RequestPath) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.RequestPath)))
		i += copy(dAtA[i:], m.RequestPath)
	}
	if len(m.RequestBytes) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.RequestBytes)))
		i += copy(dAtA[i:], m.RequestBytes)
	}
	if len(m.ResponseTime) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ResponseTime)))
		i += copy(dAtA[i:], m.ResponseTime)
	}
	if len(m.ResponseCode) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ResponseCode)))
		i += copy(dAtA[i:], m.ResponseCode)
	}
	if len(m.ResponseBytes) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ResponseBytes)))
		i += copy(dAtA[i:], m.ResponseBytes)
	}
	if len(m.ResponseLatency) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.ResponseLatency)))
		i += copy(dAtA[i:], m.ResponseLatency)
	}
	return i, nil
}

func encodeVarintGoDefaultLibraryTmpl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Type) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	l = len(m.ApiVersion)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ApiOperation)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ApiProtocol)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ApiService)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ApiKey)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.RequestTime)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.RequestMethod)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.RequestPath)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.RequestBytes)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ResponseTime)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ResponseCode)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ResponseBytes)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	l = len(m.ResponseLatency)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	return n
}

func sovGoDefaultLibraryTmpl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGoDefaultLibraryTmpl(x uint64) (n int) {
	return sovGoDefaultLibraryTmpl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceParam{`,
		`ApiVersion:` + fmt.Sprintf("%v", this.ApiVersion) + `,`,
		`ApiOperation:` + fmt.Sprintf("%v", this.ApiOperation) + `,`,
		`ApiProtocol:` + fmt.Sprintf("%v", this.ApiProtocol) + `,`,
		`ApiService:` + fmt.Sprintf("%v", this.ApiService) + `,`,
		`ApiKey:` + fmt.Sprintf("%v", this.ApiKey) + `,`,
		`RequestTime:` + fmt.Sprintf("%v", this.RequestTime) + `,`,
		`RequestMethod:` + fmt.Sprintf("%v", this.RequestMethod) + `,`,
		`RequestPath:` + fmt.Sprintf("%v", this.RequestPath) + `,`,
		`RequestBytes:` + fmt.Sprintf("%v", this.RequestBytes) + `,`,
		`ResponseTime:` + fmt.Sprintf("%v", this.ResponseTime) + `,`,
		`ResponseCode:` + fmt.Sprintf("%v", this.ResponseCode) + `,`,
		`ResponseBytes:` + fmt.Sprintf("%v", this.ResponseBytes) + `,`,
		`ResponseLatency:` + fmt.Sprintf("%v", this.ResponseLatency) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGoDefaultLibraryTmpl(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoDefaultLibraryTmpl
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
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGoDefaultLibraryTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
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
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoDefaultLibraryTmpl
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
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiOperation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiOperation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiProtocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiProtocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiService", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestBytes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseBytes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseLatency", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseLatency = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGoDefaultLibraryTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
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
func skipGoDefaultLibraryTmpl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
					return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
					return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
				return 0, ErrInvalidLengthGoDefaultLibraryTmpl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
				next, err := skipGoDefaultLibraryTmpl(dAtA[start:])
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
	ErrInvalidLengthGoDefaultLibraryTmpl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGoDefaultLibraryTmpl   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("bazel-out/local-fastbuild/genfiles/mixer/adapter/svcctrl/template/svcctrlreport/go_default_library_tmpl.proto", fileDescriptorGoDefaultLibraryTmpl)
}

var fileDescriptorGoDefaultLibraryTmpl = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xb1, 0x6e, 0xd4, 0x30,
	0x1c, 0xc6, 0xcf, 0x6d, 0xb9, 0x52, 0xb7, 0x01, 0x94, 0x05, 0x8b, 0xc1, 0xc0, 0x21, 0x24, 0x18,
	0x4a, 0x04, 0xbc, 0x41, 0x99, 0x10, 0x20, 0x4e, 0x50, 0xb1, 0x46, 0x4e, 0xf2, 0xbf, 0x9e, 0x85,
	0x13, 0x1b, 0xfb, 0x7f, 0xa7, 0x86, 0x89, 0x47, 0x40, 0xe2, 0x25, 0x78, 0x14, 0xc6, 0x8a, 0x89,
	0xb1, 0x17, 0x18, 0x18, 0x3b, 0x32, 0x22, 0xdb, 0x49, 0x7b, 0x37, 0xe6, 0x97, 0x9f, 0xbf, 0xef,
	0xb3, 0x4c, 0xeb, 0x42, 0x7c, 0x06, 0x75, 0xa8, 0x17, 0x98, 0x29, 0x5d, 0x0a, 0x75, 0x38, 0x13,
	0x0e, 0x8b, 0x85, 0x54, 0x55, 0x76, 0x02, 0xcd, 0x4c, 0x2a, 0x70, 0x59, 0x2d, 0x4f, 0xc1, 0x66,
	0xa2, 0x12, 0x06, 0xc1, 0x66, 0x6e, 0x59, 0x96, 0x68, 0x55, 0x86, 0x50, 0x1b, 0x25, 0x10, 0x06,
	0x60, 0xc1, 0x68, 0x8b, 0xd9, 0x89, 0xce, 0x2b, 0x98, 0x89, 0x85, 0xc2, 0x5c, 0xc9, 0xc2, 0x0a,
	0xdb, 0xe6, 0x58, 0x1b, 0xf5, 0xc4, 0x58, 0x8d, 0x3a, 0x4d, 0x36, 0xe4, 0x3b, 0x93, 0x18, 0xbd,
	0x7c, 0x7a, 0x95, 0x06, 0xa7, 0x08, 0x8d, 0x93, 0xba, 0x71, 0xf1, 0xc8, 0x64, 0x4c, 0x77, 0x8e,
	0x5b, 0x03, 0x93, 0xf3, 0x6d, 0x9a, 0xbc, 0x6c, 0x1c, 0x8a, 0xa6, 0x84, 0xa9, 0xb0, 0xa2, 0x4e,
	0xef, 0xd2, 0x7d, 0x61, 0x64, 0xbe, 0x04, 0xeb, 0x7d, 0x46, 0xee, 0x91, 0x47, 0x7b, 0xef, 0xa8,
	0x30, 0xf2, 0x43, 0x24, 0xe9, 0x03, 0x9a, 0x78, 0x41, 0x1b, 0xb0, 0x02, 0xbd, 0xb2, 0x15, 0x94,
	0x03, 0x61, 0xe4, 0xdb, 0x81, 0xa5, 0xf7, 0xa9, 0xff, 0xce, 0x43, 0x59, 0xa9, 0x15, 0xdb, 0x0e,
	0x8e, 0x4f, 0x9e, 0xf6, 0x68, 0x28, 0x72, 0x60, 0x97, 0xb2, 0x04, 0xb6, 0x73, 0x59, 0xf4, 0x3e,
	0x92, 0xf4, 0x36, 0xdd, 0xf5, 0xc2, 0x47, 0x68, 0xd9, 0xb5, 0xf0, 0x73, 0x2c, 0x8c, 0x7c, 0x05,
	0xad, 0x0f, 0xb7, 0xf0, 0x69, 0x01, 0x0e, 0x73, 0x94, 0x35, 0xb0, 0x71, 0x0c, 0xef, 0xd9, 0xb1,
	0xac, 0x21, 0x7d, 0x48, 0x6f, 0x0c, 0x4a, 0x0d, 0x38, 0xd7, 0x15, 0xdb, 0x0d, 0x52, 0xd2, 0xd3,
	0x37, 0x01, 0xae, 0x27, 0x19, 0x81, 0x73, 0x76, 0x7d, 0x23, 0x69, 0x2a, 0x70, 0xee, 0xaf, 0x3b,
	0x28, 0x45, 0x8b, 0xe0, 0xd8, 0x5e, 0xbc, 0x6e, 0x0f, 0x8f, 0x3c, 0x8b, 0x92, 0x33, 0xba, 0x71,
	0x10, 0x27, 0xd1, 0x41, 0x8a, 0x30, 0x6c, 0x5a, 0x97, 0x4a, 0x5d, 0x01, 0xdb, 0xdf, 0x94, 0x5e,
	0xe8, 0xaa, 0x1f, 0xde, 0x4b, 0xb1, 0xef, 0x60, 0x18, 0x1e, 0x69, 0x2c, 0x7c, 0x4c, 0x6f, 0x5d,
	0x6a, 0xfe, 0x85, 0x9b, 0xb2, 0x65, 0x49, 0x10, 0x6f, 0x0e, 0xfc, 0x75, 0xc4, 0x47, 0xcf, 0xce,
	0x56, 0x7c, 0xf4, 0x6b, 0xc5, 0x47, 0x17, 0x2b, 0x4e, 0xbe, 0x74, 0x9c, 0x7c, 0xef, 0x38, 0xf9,
	0xd1, 0x71, 0x72, 0xd6, 0x71, 0x72, 0xde, 0x71, 0xf2, 0xb7, 0xe3, 0xa3, 0x8b, 0x8e, 0x93, 0xaf,
	0xbf, 0xf9, 0xe8, 0xdf, 0xcf, 0x3f, 0xdf, 0xb6, 0x48, 0x31, 0x0e, 0x0f, 0xf7, 0xfc, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe5, 0x6b, 0x7d, 0xba, 0xc9, 0x02, 0x00, 0x00,
}
