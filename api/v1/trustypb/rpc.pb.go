// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
	Package trustypb is a generated protocol buffer package.

	It is generated from these files:
		rpc.proto

	It has these top-level messages:
		EmptyRequest
		ServerVersion
		ServerStatus
		ServerStatusResponse
		CallerStatusResponse
		Error
*/
package trustypb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{0} }

type ServerVersion struct {
	// Build is the server build version.
	Build string `protobuf:"bytes,1,opt,name=Build,proto3" json:"Build,omitempty"`
	// Runtime is the runtime version.
	Runtime string `protobuf:"bytes,2,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
}

func (m *ServerVersion) Reset()                    { *m = ServerVersion{} }
func (m *ServerVersion) String() string            { return proto.CompactTextString(m) }
func (*ServerVersion) ProtoMessage()               {}
func (*ServerVersion) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{1} }

func (m *ServerVersion) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *ServerVersion) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

type ServerStatus struct {
	// Name of the server or application.
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Nodename is the human-readable name of the cluster member,
	// or empty for single host.
	Nodename string `protobuf:"bytes,2,opt,name=Nodename,proto3" json:"Nodename,omitempty"`
	// Hostname is operating system's host name.
	Hostname string `protobuf:"bytes,3,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	// ListenURLs is the list of URLs the service is listening on.
	ListenURLs []string `protobuf:"bytes,4,rep,name=ListenURLs" json:"ListenURLs,omitempty"`
	// StartedAt is the Unix time when the server has started.
	StartedAt int64 `protobuf:"varint,5,opt,name=StartedAt,proto3" json:"StartedAt,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{2} }

func (m *ServerStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerStatus) GetNodename() string {
	if m != nil {
		return m.Nodename
	}
	return ""
}

func (m *ServerStatus) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ServerStatus) GetListenURLs() []string {
	if m != nil {
		return m.ListenURLs
	}
	return nil
}

func (m *ServerStatus) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

type ServerStatusResponse struct {
	// Status of the server.
	Status *ServerStatus `protobuf:"bytes,1,opt,name=Status" json:"Status,omitempty"`
	// Version of the server.
	Version *ServerVersion `protobuf:"bytes,2,opt,name=Version" json:"Version,omitempty"`
}

func (m *ServerStatusResponse) Reset()                    { *m = ServerStatusResponse{} }
func (m *ServerStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerStatusResponse) ProtoMessage()               {}
func (*ServerStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{3} }

func (m *ServerStatusResponse) GetStatus() *ServerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ServerStatusResponse) GetVersion() *ServerVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

type CallerStatusResponse struct {
	// ID of the caller.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name of the caller.
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// Role of the caller.
	Role string `protobuf:"bytes,3,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (m *CallerStatusResponse) Reset()                    { *m = CallerStatusResponse{} }
func (m *CallerStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*CallerStatusResponse) ProtoMessage()               {}
func (*CallerStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{4} }

func (m *CallerStatusResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CallerStatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CallerStatusResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

// Error is the generic error returned from unary RPCs.
type Error struct {
	Error   string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{5} }

func (m *Error) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "trustypb.EmptyRequest")
	proto.RegisterType((*ServerVersion)(nil), "trustypb.ServerVersion")
	proto.RegisterType((*ServerStatus)(nil), "trustypb.ServerStatus")
	proto.RegisterType((*ServerStatusResponse)(nil), "trustypb.ServerStatusResponse")
	proto.RegisterType((*CallerStatusResponse)(nil), "trustypb.CallerStatusResponse")
	proto.RegisterType((*Error)(nil), "trustypb.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Status service

type StatusClient interface {
	// Version returns the server version.
	Version(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServerVersion, error)
	// Server returns the server status.
	Server(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServerStatusResponse, error)
	// Caller returns the caller status.
	Caller(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CallerStatusResponse, error)
}

type statusClient struct {
	cc *grpc.ClientConn
}

func NewStatusClient(cc *grpc.ClientConn) StatusClient {
	return &statusClient{cc}
}

func (c *statusClient) Version(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServerVersion, error) {
	out := new(ServerVersion)
	err := grpc.Invoke(ctx, "/trustypb.Status/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusClient) Server(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ServerStatusResponse, error) {
	out := new(ServerStatusResponse)
	err := grpc.Invoke(ctx, "/trustypb.Status/Server", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusClient) Caller(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CallerStatusResponse, error) {
	out := new(CallerStatusResponse)
	err := grpc.Invoke(ctx, "/trustypb.Status/Caller", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Status service

type StatusServer interface {
	// Version returns the server version.
	Version(context.Context, *EmptyRequest) (*ServerVersion, error)
	// Server returns the server status.
	Server(context.Context, *EmptyRequest) (*ServerStatusResponse, error)
	// Caller returns the caller status.
	Caller(context.Context, *EmptyRequest) (*CallerStatusResponse, error)
}

func RegisterStatusServer(s *grpc.Server, srv StatusServer) {
	s.RegisterService(&_Status_serviceDesc, srv)
}

func _Status_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustypb.Status/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).Version(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Status_Server_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).Server(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustypb.Status/Server",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).Server(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Status_Caller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServer).Caller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trustypb.Status/Caller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServer).Caller(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Status_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trustypb.Status",
	HandlerType: (*StatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Status_Version_Handler,
		},
		{
			MethodName: "Server",
			Handler:    _Status_Server_Handler,
		},
		{
			MethodName: "Caller",
			Handler:    _Status_Caller_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func (m *EmptyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ServerVersion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServerVersion) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Build) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Build)))
		i += copy(dAtA[i:], m.Build)
	}
	if len(m.Runtime) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Runtime)))
		i += copy(dAtA[i:], m.Runtime)
	}
	return i, nil
}

func (m *ServerStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServerStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Nodename) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Nodename)))
		i += copy(dAtA[i:], m.Nodename)
	}
	if len(m.Hostname) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Hostname)))
		i += copy(dAtA[i:], m.Hostname)
	}
	if len(m.ListenURLs) > 0 {
		for _, s := range m.ListenURLs {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.StartedAt != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StartedAt))
	}
	return i, nil
}

func (m *ServerStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServerStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Status.Size()))
		n1, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Version != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Version.Size()))
		n2, err := m.Version.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *CallerStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallerStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Role) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	return i, nil
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if m.Code != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EmptyRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ServerVersion) Size() (n int) {
	var l int
	_ = l
	l = len(m.Build)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Runtime)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *ServerStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Nodename)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.ListenURLs) > 0 {
		for _, s := range m.ListenURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.StartedAt != 0 {
		n += 1 + sovRpc(uint64(m.StartedAt))
	}
	return n
}

func (m *ServerStatusResponse) Size() (n int) {
	var l int
	_ = l
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *CallerStatusResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *Error) Size() (n int) {
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovRpc(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func sovRpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EmptyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: EmptyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *ServerVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ServerVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Build", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Build = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Runtime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Runtime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *ServerStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ServerStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodename = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenURLs = append(m.ListenURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedAt", wireType)
			}
			m.StartedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *ServerStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ServerStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &ServerStatus{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &ServerVersion{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *CallerStatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: CallerStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallerStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
				return 0, ErrInvalidLengthRpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpc
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
				next, err := skipRpc(dAtA[start:])
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
	ErrInvalidLengthRpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0xf3, 0xa7, 0xcd, 0xb4, 0x54, 0x30, 0x58, 0x60, 0xac, 0xca, 0x8a, 0x7c, 0xca,
	0x29, 0x56, 0xc3, 0x03, 0x20, 0x4a, 0x2b, 0x51, 0x51, 0xe5, 0xb0, 0x11, 0xbd, 0xf4, 0xb4, 0x4d,
	0x56, 0x96, 0x25, 0x67, 0xd7, 0xec, 0xae, 0x23, 0xe5, 0xca, 0x2b, 0x70, 0xa9, 0x78, 0x22, 0x8e,
	0x48, 0xbc, 0x00, 0x0a, 0x3c, 0x08, 0xf2, 0xae, 0x37, 0x18, 0x48, 0xa4, 0xde, 0x66, 0xe6, 0x9b,
	0xf9, 0x32, 0xfb, 0x9b, 0x18, 0x06, 0xb2, 0x9c, 0x8f, 0x4b, 0x29, 0xb4, 0xc0, 0x43, 0x2d, 0x2b,
	0xa5, 0xd7, 0xe5, 0x5d, 0x14, 0x64, 0x22, 0x13, 0xa6, 0x98, 0xd6, 0x91, 0xd5, 0xa3, 0xd3, 0x4c,
	0x88, 0xac, 0x60, 0x29, 0x2d, 0xf3, 0x94, 0x72, 0x2e, 0x34, 0xd5, 0xb9, 0xe0, 0xca, 0xaa, 0xc9,
	0x09, 0x1c, 0x5f, 0x2e, 0x4b, 0xbd, 0x26, 0xec, 0x63, 0xc5, 0x94, 0x4e, 0x5e, 0xc3, 0xe3, 0x19,
	0x93, 0x2b, 0x26, 0x6f, 0x98, 0x54, 0xb9, 0xe0, 0x18, 0x40, 0xef, 0xbc, 0xca, 0x8b, 0x45, 0xe8,
	0x0d, 0xbd, 0xd1, 0x80, 0xd8, 0x04, 0x43, 0x38, 0x20, 0x15, 0xd7, 0xf9, 0x92, 0x85, 0xbe, 0xa9,
	0xbb, 0x34, 0xb9, 0xf7, 0xe0, 0xd8, 0x3a, 0xcc, 0x34, 0xd5, 0x95, 0x42, 0x84, 0xee, 0x94, 0x2e,
	0x59, 0x33, 0x6f, 0x62, 0x8c, 0xe0, 0x70, 0x2a, 0x16, 0x8c, 0xd3, 0xed, 0xfc, 0x36, 0xaf, 0xb5,
	0x77, 0x42, 0x69, 0xa3, 0x75, 0xac, 0xe6, 0x72, 0x8c, 0x01, 0xae, 0x73, 0xa5, 0x19, 0xff, 0x40,
	0xae, 0x55, 0xd8, 0x1d, 0x76, 0x46, 0x03, 0xd2, 0xaa, 0xe0, 0x29, 0x0c, 0x66, 0x9a, 0x4a, 0xcd,
	0x16, 0x6f, 0x74, 0xd8, 0x1b, 0x7a, 0xa3, 0x0e, 0xf9, 0x53, 0x48, 0xd6, 0x10, 0xb4, 0x37, 0x23,
	0x4c, 0x95, 0x82, 0x2b, 0x86, 0x63, 0xe8, 0xdb, 0x8a, 0xd9, 0xf1, 0x68, 0xf2, 0x7c, 0xec, 0x90,
	0x8e, 0xff, 0xea, 0x6f, 0xba, 0xf0, 0x0c, 0x0e, 0x1a, 0x3a, 0x66, 0xf9, 0xa3, 0xc9, 0x8b, 0x7f,
	0x07, 0x1a, 0x99, 0xb8, 0xbe, 0x64, 0x0a, 0xc1, 0x5b, 0x5a, 0x14, 0xff, 0xfd, 0xf4, 0x09, 0xf8,
	0x57, 0x17, 0x0d, 0x1a, 0xff, 0xea, 0x62, 0x0b, 0xcb, 0x6f, 0xc1, 0x42, 0xe8, 0x12, 0x51, 0x38,
	0x18, 0x26, 0x4e, 0xde, 0x43, 0xef, 0x52, 0x4a, 0x21, 0xeb, 0xf3, 0xb0, 0x3a, 0x70, 0xe7, 0x31,
	0x49, 0x3d, 0x32, 0x17, 0x0b, 0x6b, 0xd3, 0x23, 0x26, 0xae, 0x4f, 0xb6, 0x64, 0x4a, 0xd1, 0xcc,
	0x39, 0xb9, 0x74, 0xf2, 0xc5, 0x77, 0x00, 0xf0, 0x66, 0xfb, 0x34, 0x6c, 0x51, 0x68, 0xff, 0x43,
	0xa2, 0x7d, 0x8f, 0x4d, 0xa2, 0x4f, 0xdf, 0x7f, 0x7d, 0xf6, 0x03, 0xc4, 0x74, 0x75, 0x96, 0x2a,
	0xe3, 0x97, 0xae, 0x1a, 0xb3, 0x5b, 0xe8, 0xdb, 0xe6, 0xbd, 0xb6, 0xf1, 0x1e, 0xe8, 0x0d, 0xa9,
	0xe4, 0xa5, 0x71, 0x7f, 0x86, 0x4f, 0x5b, 0xee, 0xca, 0x5a, 0xde, 0x42, 0xdf, 0xc2, 0x7d, 0x88,
	0xf9, 0xae, 0x33, 0xec, 0x34, 0x9f, 0x9b, 0xc6, 0xf3, 0x27, 0x5f, 0x37, 0xb1, 0xf7, 0x6d, 0x13,
	0x7b, 0x3f, 0x36, 0xb1, 0x77, 0xff, 0x33, 0x7e, 0x74, 0xd7, 0x37, 0x5f, 0xce, 0xab, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x73, 0xb6, 0x8a, 0xe6, 0x84, 0x03, 0x00, 0x00,
}
