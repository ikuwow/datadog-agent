// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc
// source: pkg/security/proto/api/api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SecurityModule_GetEvents_FullMethodName             = "/api.SecurityModule/GetEvents"
	SecurityModule_DumpProcessCache_FullMethodName      = "/api.SecurityModule/DumpProcessCache"
	SecurityModule_GetConfig_FullMethodName             = "/api.SecurityModule/GetConfig"
	SecurityModule_GetStatus_FullMethodName             = "/api.SecurityModule/GetStatus"
	SecurityModule_RunSelfTest_FullMethodName           = "/api.SecurityModule/RunSelfTest"
	SecurityModule_GetRuleSetReport_FullMethodName      = "/api.SecurityModule/GetRuleSetReport"
	SecurityModule_ReloadPolicies_FullMethodName        = "/api.SecurityModule/ReloadPolicies"
	SecurityModule_DumpNetworkNamespace_FullMethodName  = "/api.SecurityModule/DumpNetworkNamespace"
	SecurityModule_DumpDiscarders_FullMethodName        = "/api.SecurityModule/DumpDiscarders"
	SecurityModule_DumpActivity_FullMethodName          = "/api.SecurityModule/DumpActivity"
	SecurityModule_ListActivityDumps_FullMethodName     = "/api.SecurityModule/ListActivityDumps"
	SecurityModule_StopActivityDump_FullMethodName      = "/api.SecurityModule/StopActivityDump"
	SecurityModule_TranscodingRequest_FullMethodName    = "/api.SecurityModule/TranscodingRequest"
	SecurityModule_GetActivityDumpStream_FullMethodName = "/api.SecurityModule/GetActivityDumpStream"
	SecurityModule_ListSecurityProfiles_FullMethodName  = "/api.SecurityModule/ListSecurityProfiles"
	SecurityModule_SaveSecurityProfile_FullMethodName   = "/api.SecurityModule/SaveSecurityProfile"
)

// SecurityModuleClient is the client API for SecurityModule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityModuleClient interface {
	GetEvents(ctx context.Context, in *GetEventParams, opts ...grpc.CallOption) (SecurityModule_GetEventsClient, error)
	DumpProcessCache(ctx context.Context, in *DumpProcessCacheParams, opts ...grpc.CallOption) (*SecurityDumpProcessCacheMessage, error)
	GetConfig(ctx context.Context, in *GetConfigParams, opts ...grpc.CallOption) (*SecurityConfigMessage, error)
	GetStatus(ctx context.Context, in *GetStatusParams, opts ...grpc.CallOption) (*Status, error)
	RunSelfTest(ctx context.Context, in *RunSelfTestParams, opts ...grpc.CallOption) (*SecuritySelfTestResultMessage, error)
	GetRuleSetReport(ctx context.Context, in *GetRuleSetReportParams, opts ...grpc.CallOption) (*GetRuleSetReportResultMessage, error)
	ReloadPolicies(ctx context.Context, in *ReloadPoliciesParams, opts ...grpc.CallOption) (*ReloadPoliciesResultMessage, error)
	DumpNetworkNamespace(ctx context.Context, in *DumpNetworkNamespaceParams, opts ...grpc.CallOption) (*DumpNetworkNamespaceMessage, error)
	DumpDiscarders(ctx context.Context, in *DumpDiscardersParams, opts ...grpc.CallOption) (*DumpDiscardersMessage, error)
	// Activity dumps
	DumpActivity(ctx context.Context, in *ActivityDumpParams, opts ...grpc.CallOption) (*ActivityDumpMessage, error)
	ListActivityDumps(ctx context.Context, in *ActivityDumpListParams, opts ...grpc.CallOption) (*ActivityDumpListMessage, error)
	StopActivityDump(ctx context.Context, in *ActivityDumpStopParams, opts ...grpc.CallOption) (*ActivityDumpStopMessage, error)
	TranscodingRequest(ctx context.Context, in *TranscodingRequestParams, opts ...grpc.CallOption) (*TranscodingRequestMessage, error)
	GetActivityDumpStream(ctx context.Context, in *ActivityDumpStreamParams, opts ...grpc.CallOption) (SecurityModule_GetActivityDumpStreamClient, error)
	// Security Profiles
	ListSecurityProfiles(ctx context.Context, in *SecurityProfileListParams, opts ...grpc.CallOption) (*SecurityProfileListMessage, error)
	SaveSecurityProfile(ctx context.Context, in *SecurityProfileSaveParams, opts ...grpc.CallOption) (*SecurityProfileSaveMessage, error)
}

type securityModuleClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityModuleClient(cc grpc.ClientConnInterface) SecurityModuleClient {
	return &securityModuleClient{cc}
}

func (c *securityModuleClient) GetEvents(ctx context.Context, in *GetEventParams, opts ...grpc.CallOption) (SecurityModule_GetEventsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SecurityModule_ServiceDesc.Streams[0], SecurityModule_GetEvents_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &securityModuleGetEventsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SecurityModule_GetEventsClient interface {
	Recv() (*SecurityEventMessage, error)
	grpc.ClientStream
}

type securityModuleGetEventsClient struct {
	grpc.ClientStream
}

func (x *securityModuleGetEventsClient) Recv() (*SecurityEventMessage, error) {
	m := new(SecurityEventMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *securityModuleClient) DumpProcessCache(ctx context.Context, in *DumpProcessCacheParams, opts ...grpc.CallOption) (*SecurityDumpProcessCacheMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecurityDumpProcessCacheMessage)
	err := c.cc.Invoke(ctx, SecurityModule_DumpProcessCache_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) GetConfig(ctx context.Context, in *GetConfigParams, opts ...grpc.CallOption) (*SecurityConfigMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecurityConfigMessage)
	err := c.cc.Invoke(ctx, SecurityModule_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) GetStatus(ctx context.Context, in *GetStatusParams, opts ...grpc.CallOption) (*Status, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Status)
	err := c.cc.Invoke(ctx, SecurityModule_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) RunSelfTest(ctx context.Context, in *RunSelfTestParams, opts ...grpc.CallOption) (*SecuritySelfTestResultMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecuritySelfTestResultMessage)
	err := c.cc.Invoke(ctx, SecurityModule_RunSelfTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) GetRuleSetReport(ctx context.Context, in *GetRuleSetReportParams, opts ...grpc.CallOption) (*GetRuleSetReportResultMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRuleSetReportResultMessage)
	err := c.cc.Invoke(ctx, SecurityModule_GetRuleSetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) ReloadPolicies(ctx context.Context, in *ReloadPoliciesParams, opts ...grpc.CallOption) (*ReloadPoliciesResultMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReloadPoliciesResultMessage)
	err := c.cc.Invoke(ctx, SecurityModule_ReloadPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) DumpNetworkNamespace(ctx context.Context, in *DumpNetworkNamespaceParams, opts ...grpc.CallOption) (*DumpNetworkNamespaceMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DumpNetworkNamespaceMessage)
	err := c.cc.Invoke(ctx, SecurityModule_DumpNetworkNamespace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) DumpDiscarders(ctx context.Context, in *DumpDiscardersParams, opts ...grpc.CallOption) (*DumpDiscardersMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DumpDiscardersMessage)
	err := c.cc.Invoke(ctx, SecurityModule_DumpDiscarders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) DumpActivity(ctx context.Context, in *ActivityDumpParams, opts ...grpc.CallOption) (*ActivityDumpMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityDumpMessage)
	err := c.cc.Invoke(ctx, SecurityModule_DumpActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) ListActivityDumps(ctx context.Context, in *ActivityDumpListParams, opts ...grpc.CallOption) (*ActivityDumpListMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityDumpListMessage)
	err := c.cc.Invoke(ctx, SecurityModule_ListActivityDumps_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) StopActivityDump(ctx context.Context, in *ActivityDumpStopParams, opts ...grpc.CallOption) (*ActivityDumpStopMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivityDumpStopMessage)
	err := c.cc.Invoke(ctx, SecurityModule_StopActivityDump_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) TranscodingRequest(ctx context.Context, in *TranscodingRequestParams, opts ...grpc.CallOption) (*TranscodingRequestMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranscodingRequestMessage)
	err := c.cc.Invoke(ctx, SecurityModule_TranscodingRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) GetActivityDumpStream(ctx context.Context, in *ActivityDumpStreamParams, opts ...grpc.CallOption) (SecurityModule_GetActivityDumpStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SecurityModule_ServiceDesc.Streams[1], SecurityModule_GetActivityDumpStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &securityModuleGetActivityDumpStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SecurityModule_GetActivityDumpStreamClient interface {
	Recv() (*ActivityDumpStreamMessage, error)
	grpc.ClientStream
}

type securityModuleGetActivityDumpStreamClient struct {
	grpc.ClientStream
}

func (x *securityModuleGetActivityDumpStreamClient) Recv() (*ActivityDumpStreamMessage, error) {
	m := new(ActivityDumpStreamMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *securityModuleClient) ListSecurityProfiles(ctx context.Context, in *SecurityProfileListParams, opts ...grpc.CallOption) (*SecurityProfileListMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecurityProfileListMessage)
	err := c.cc.Invoke(ctx, SecurityModule_ListSecurityProfiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityModuleClient) SaveSecurityProfile(ctx context.Context, in *SecurityProfileSaveParams, opts ...grpc.CallOption) (*SecurityProfileSaveMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecurityProfileSaveMessage)
	err := c.cc.Invoke(ctx, SecurityModule_SaveSecurityProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityModuleServer is the server API for SecurityModule service.
// All implementations must embed UnimplementedSecurityModuleServer
// for forward compatibility
type SecurityModuleServer interface {
	GetEvents(*GetEventParams, SecurityModule_GetEventsServer) error
	DumpProcessCache(context.Context, *DumpProcessCacheParams) (*SecurityDumpProcessCacheMessage, error)
	GetConfig(context.Context, *GetConfigParams) (*SecurityConfigMessage, error)
	GetStatus(context.Context, *GetStatusParams) (*Status, error)
	RunSelfTest(context.Context, *RunSelfTestParams) (*SecuritySelfTestResultMessage, error)
	GetRuleSetReport(context.Context, *GetRuleSetReportParams) (*GetRuleSetReportResultMessage, error)
	ReloadPolicies(context.Context, *ReloadPoliciesParams) (*ReloadPoliciesResultMessage, error)
	DumpNetworkNamespace(context.Context, *DumpNetworkNamespaceParams) (*DumpNetworkNamespaceMessage, error)
	DumpDiscarders(context.Context, *DumpDiscardersParams) (*DumpDiscardersMessage, error)
	// Activity dumps
	DumpActivity(context.Context, *ActivityDumpParams) (*ActivityDumpMessage, error)
	ListActivityDumps(context.Context, *ActivityDumpListParams) (*ActivityDumpListMessage, error)
	StopActivityDump(context.Context, *ActivityDumpStopParams) (*ActivityDumpStopMessage, error)
	TranscodingRequest(context.Context, *TranscodingRequestParams) (*TranscodingRequestMessage, error)
	GetActivityDumpStream(*ActivityDumpStreamParams, SecurityModule_GetActivityDumpStreamServer) error
	// Security Profiles
	ListSecurityProfiles(context.Context, *SecurityProfileListParams) (*SecurityProfileListMessage, error)
	SaveSecurityProfile(context.Context, *SecurityProfileSaveParams) (*SecurityProfileSaveMessage, error)
	mustEmbedUnimplementedSecurityModuleServer()
}

// UnimplementedSecurityModuleServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityModuleServer struct {
}

func (UnimplementedSecurityModuleServer) GetEvents(*GetEventParams, SecurityModule_GetEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedSecurityModuleServer) DumpProcessCache(context.Context, *DumpProcessCacheParams) (*SecurityDumpProcessCacheMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpProcessCache not implemented")
}
func (UnimplementedSecurityModuleServer) GetConfig(context.Context, *GetConfigParams) (*SecurityConfigMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedSecurityModuleServer) GetStatus(context.Context, *GetStatusParams) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedSecurityModuleServer) RunSelfTest(context.Context, *RunSelfTestParams) (*SecuritySelfTestResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunSelfTest not implemented")
}
func (UnimplementedSecurityModuleServer) GetRuleSetReport(context.Context, *GetRuleSetReportParams) (*GetRuleSetReportResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRuleSetReport not implemented")
}
func (UnimplementedSecurityModuleServer) ReloadPolicies(context.Context, *ReloadPoliciesParams) (*ReloadPoliciesResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadPolicies not implemented")
}
func (UnimplementedSecurityModuleServer) DumpNetworkNamespace(context.Context, *DumpNetworkNamespaceParams) (*DumpNetworkNamespaceMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpNetworkNamespace not implemented")
}
func (UnimplementedSecurityModuleServer) DumpDiscarders(context.Context, *DumpDiscardersParams) (*DumpDiscardersMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpDiscarders not implemented")
}
func (UnimplementedSecurityModuleServer) DumpActivity(context.Context, *ActivityDumpParams) (*ActivityDumpMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpActivity not implemented")
}
func (UnimplementedSecurityModuleServer) ListActivityDumps(context.Context, *ActivityDumpListParams) (*ActivityDumpListMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivityDumps not implemented")
}
func (UnimplementedSecurityModuleServer) StopActivityDump(context.Context, *ActivityDumpStopParams) (*ActivityDumpStopMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopActivityDump not implemented")
}
func (UnimplementedSecurityModuleServer) TranscodingRequest(context.Context, *TranscodingRequestParams) (*TranscodingRequestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranscodingRequest not implemented")
}
func (UnimplementedSecurityModuleServer) GetActivityDumpStream(*ActivityDumpStreamParams, SecurityModule_GetActivityDumpStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActivityDumpStream not implemented")
}
func (UnimplementedSecurityModuleServer) ListSecurityProfiles(context.Context, *SecurityProfileListParams) (*SecurityProfileListMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecurityProfiles not implemented")
}
func (UnimplementedSecurityModuleServer) SaveSecurityProfile(context.Context, *SecurityProfileSaveParams) (*SecurityProfileSaveMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSecurityProfile not implemented")
}
func (UnimplementedSecurityModuleServer) mustEmbedUnimplementedSecurityModuleServer() {}

// UnsafeSecurityModuleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityModuleServer will
// result in compilation errors.
type UnsafeSecurityModuleServer interface {
	mustEmbedUnimplementedSecurityModuleServer()
}

func RegisterSecurityModuleServer(s grpc.ServiceRegistrar, srv SecurityModuleServer) {
	s.RegisterService(&SecurityModule_ServiceDesc, srv)
}

func _SecurityModule_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecurityModuleServer).GetEvents(m, &securityModuleGetEventsServer{ServerStream: stream})
}

type SecurityModule_GetEventsServer interface {
	Send(*SecurityEventMessage) error
	grpc.ServerStream
}

type securityModuleGetEventsServer struct {
	grpc.ServerStream
}

func (x *securityModuleGetEventsServer) Send(m *SecurityEventMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _SecurityModule_DumpProcessCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpProcessCacheParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).DumpProcessCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_DumpProcessCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).DumpProcessCache(ctx, req.(*DumpProcessCacheParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).GetConfig(ctx, req.(*GetConfigParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).GetStatus(ctx, req.(*GetStatusParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_RunSelfTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunSelfTestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).RunSelfTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_RunSelfTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).RunSelfTest(ctx, req.(*RunSelfTestParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_GetRuleSetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuleSetReportParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).GetRuleSetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_GetRuleSetReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).GetRuleSetReport(ctx, req.(*GetRuleSetReportParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_ReloadPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadPoliciesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).ReloadPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_ReloadPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).ReloadPolicies(ctx, req.(*ReloadPoliciesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_DumpNetworkNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpNetworkNamespaceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).DumpNetworkNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_DumpNetworkNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).DumpNetworkNamespace(ctx, req.(*DumpNetworkNamespaceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_DumpDiscarders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpDiscardersParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).DumpDiscarders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_DumpDiscarders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).DumpDiscarders(ctx, req.(*DumpDiscardersParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_DumpActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityDumpParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).DumpActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_DumpActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).DumpActivity(ctx, req.(*ActivityDumpParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_ListActivityDumps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityDumpListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).ListActivityDumps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_ListActivityDumps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).ListActivityDumps(ctx, req.(*ActivityDumpListParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_StopActivityDump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityDumpStopParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).StopActivityDump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_StopActivityDump_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).StopActivityDump(ctx, req.(*ActivityDumpStopParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_TranscodingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscodingRequestParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).TranscodingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_TranscodingRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).TranscodingRequest(ctx, req.(*TranscodingRequestParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_GetActivityDumpStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ActivityDumpStreamParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecurityModuleServer).GetActivityDumpStream(m, &securityModuleGetActivityDumpStreamServer{ServerStream: stream})
}

type SecurityModule_GetActivityDumpStreamServer interface {
	Send(*ActivityDumpStreamMessage) error
	grpc.ServerStream
}

type securityModuleGetActivityDumpStreamServer struct {
	grpc.ServerStream
}

func (x *securityModuleGetActivityDumpStreamServer) Send(m *ActivityDumpStreamMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _SecurityModule_ListSecurityProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecurityProfileListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).ListSecurityProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_ListSecurityProfiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).ListSecurityProfiles(ctx, req.(*SecurityProfileListParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityModule_SaveSecurityProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecurityProfileSaveParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityModuleServer).SaveSecurityProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityModule_SaveSecurityProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityModuleServer).SaveSecurityProfile(ctx, req.(*SecurityProfileSaveParams))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityModule_ServiceDesc is the grpc.ServiceDesc for SecurityModule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityModule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SecurityModule",
	HandlerType: (*SecurityModuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DumpProcessCache",
			Handler:    _SecurityModule_DumpProcessCache_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _SecurityModule_GetConfig_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _SecurityModule_GetStatus_Handler,
		},
		{
			MethodName: "RunSelfTest",
			Handler:    _SecurityModule_RunSelfTest_Handler,
		},
		{
			MethodName: "GetRuleSetReport",
			Handler:    _SecurityModule_GetRuleSetReport_Handler,
		},
		{
			MethodName: "ReloadPolicies",
			Handler:    _SecurityModule_ReloadPolicies_Handler,
		},
		{
			MethodName: "DumpNetworkNamespace",
			Handler:    _SecurityModule_DumpNetworkNamespace_Handler,
		},
		{
			MethodName: "DumpDiscarders",
			Handler:    _SecurityModule_DumpDiscarders_Handler,
		},
		{
			MethodName: "DumpActivity",
			Handler:    _SecurityModule_DumpActivity_Handler,
		},
		{
			MethodName: "ListActivityDumps",
			Handler:    _SecurityModule_ListActivityDumps_Handler,
		},
		{
			MethodName: "StopActivityDump",
			Handler:    _SecurityModule_StopActivityDump_Handler,
		},
		{
			MethodName: "TranscodingRequest",
			Handler:    _SecurityModule_TranscodingRequest_Handler,
		},
		{
			MethodName: "ListSecurityProfiles",
			Handler:    _SecurityModule_ListSecurityProfiles_Handler,
		},
		{
			MethodName: "SaveSecurityProfile",
			Handler:    _SecurityModule_SaveSecurityProfile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _SecurityModule_GetEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetActivityDumpStream",
			Handler:       _SecurityModule_GetActivityDumpStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/security/proto/api/api.proto",
}
