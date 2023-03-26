// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VierkantleServiceClient is the client API for VierkantleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VierkantleServiceClient interface {
	GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error)
	SubmitScore(ctx context.Context, in *SubmitScoreRequest, opts ...grpc.CallOption) (*SubmitScoreResponse, error)
	GetScores(ctx context.Context, in *GetScoresRequest, opts ...grpc.CallOption) (*GetScoresResponse, error)
	Whoami(ctx context.Context, in *WhoamiRequest, opts ...grpc.CallOption) (*WhoamiResponse, error)
	StartLogin(ctx context.Context, in *StartLoginRequest, opts ...grpc.CallOption) (*StartLoginResponse, error)
	FinishLogin(ctx context.Context, in *FinishLoginRequest, opts ...grpc.CallOption) (*FinishLoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Generator
	WordsForBoard(ctx context.Context, in *WordsForBoardRequest, opts ...grpc.CallOption) (*WordsForBoardResponse, error)
	SeedBoard(ctx context.Context, in *SeedBoardRequest, opts ...grpc.CallOption) (VierkantleService_SeedBoardClient, error)
	FillInBoard(ctx context.Context, in *FillInBoardRequest, opts ...grpc.CallOption) (VierkantleService_FillInBoardClient, error)
	TeamStream(ctx context.Context, opts ...grpc.CallOption) (VierkantleService_TeamStreamClient, error)
}

type vierkantleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVierkantleServiceClient(cc grpc.ClientConnInterface) VierkantleServiceClient {
	return &vierkantleServiceClient{cc}
}

func (c *vierkantleServiceClient) GetBoard(ctx context.Context, in *GetBoardRequest, opts ...grpc.CallOption) (*GetBoardResponse, error) {
	out := new(GetBoardResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/GetBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) SubmitScore(ctx context.Context, in *SubmitScoreRequest, opts ...grpc.CallOption) (*SubmitScoreResponse, error) {
	out := new(SubmitScoreResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/SubmitScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) GetScores(ctx context.Context, in *GetScoresRequest, opts ...grpc.CallOption) (*GetScoresResponse, error) {
	out := new(GetScoresResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/GetScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) Whoami(ctx context.Context, in *WhoamiRequest, opts ...grpc.CallOption) (*WhoamiResponse, error) {
	out := new(WhoamiResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/Whoami", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) StartLogin(ctx context.Context, in *StartLoginRequest, opts ...grpc.CallOption) (*StartLoginResponse, error) {
	out := new(StartLoginResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/StartLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) FinishLogin(ctx context.Context, in *FinishLoginRequest, opts ...grpc.CallOption) (*FinishLoginResponse, error) {
	out := new(FinishLoginResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/FinishLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) WordsForBoard(ctx context.Context, in *WordsForBoardRequest, opts ...grpc.CallOption) (*WordsForBoardResponse, error) {
	out := new(WordsForBoardResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/WordsForBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) SeedBoard(ctx context.Context, in *SeedBoardRequest, opts ...grpc.CallOption) (VierkantleService_SeedBoardClient, error) {
	stream, err := c.cc.NewStream(ctx, &VierkantleService_ServiceDesc.Streams[0], "/nl.vierkantle.VierkantleService/SeedBoard", opts...)
	if err != nil {
		return nil, err
	}
	x := &vierkantleServiceSeedBoardClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VierkantleService_SeedBoardClient interface {
	Recv() (*SeedBoardResponse, error)
	grpc.ClientStream
}

type vierkantleServiceSeedBoardClient struct {
	grpc.ClientStream
}

func (x *vierkantleServiceSeedBoardClient) Recv() (*SeedBoardResponse, error) {
	m := new(SeedBoardResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vierkantleServiceClient) FillInBoard(ctx context.Context, in *FillInBoardRequest, opts ...grpc.CallOption) (VierkantleService_FillInBoardClient, error) {
	stream, err := c.cc.NewStream(ctx, &VierkantleService_ServiceDesc.Streams[1], "/nl.vierkantle.VierkantleService/FillInBoard", opts...)
	if err != nil {
		return nil, err
	}
	x := &vierkantleServiceFillInBoardClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VierkantleService_FillInBoardClient interface {
	Recv() (*FillInBoardResponse, error)
	grpc.ClientStream
}

type vierkantleServiceFillInBoardClient struct {
	grpc.ClientStream
}

func (x *vierkantleServiceFillInBoardClient) Recv() (*FillInBoardResponse, error) {
	m := new(FillInBoardResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vierkantleServiceClient) TeamStream(ctx context.Context, opts ...grpc.CallOption) (VierkantleService_TeamStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &VierkantleService_ServiceDesc.Streams[2], "/nl.vierkantle.VierkantleService/TeamStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &vierkantleServiceTeamStreamClient{stream}
	return x, nil
}

type VierkantleService_TeamStreamClient interface {
	Send(*TeamStreamClientMessage) error
	Recv() (*TeamStreamServerMessage, error)
	grpc.ClientStream
}

type vierkantleServiceTeamStreamClient struct {
	grpc.ClientStream
}

func (x *vierkantleServiceTeamStreamClient) Send(m *TeamStreamClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vierkantleServiceTeamStreamClient) Recv() (*TeamStreamServerMessage, error) {
	m := new(TeamStreamServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VierkantleServiceServer is the server API for VierkantleService service.
// All implementations should embed UnimplementedVierkantleServiceServer
// for forward compatibility
type VierkantleServiceServer interface {
	GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error)
	SubmitScore(context.Context, *SubmitScoreRequest) (*SubmitScoreResponse, error)
	GetScores(context.Context, *GetScoresRequest) (*GetScoresResponse, error)
	Whoami(context.Context, *WhoamiRequest) (*WhoamiResponse, error)
	StartLogin(context.Context, *StartLoginRequest) (*StartLoginResponse, error)
	FinishLogin(context.Context, *FinishLoginRequest) (*FinishLoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Generator
	WordsForBoard(context.Context, *WordsForBoardRequest) (*WordsForBoardResponse, error)
	SeedBoard(*SeedBoardRequest, VierkantleService_SeedBoardServer) error
	FillInBoard(*FillInBoardRequest, VierkantleService_FillInBoardServer) error
	TeamStream(VierkantleService_TeamStreamServer) error
}

// UnimplementedVierkantleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVierkantleServiceServer struct {
}

func (UnimplementedVierkantleServiceServer) GetBoard(context.Context, *GetBoardRequest) (*GetBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoard not implemented")
}
func (UnimplementedVierkantleServiceServer) SubmitScore(context.Context, *SubmitScoreRequest) (*SubmitScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitScore not implemented")
}
func (UnimplementedVierkantleServiceServer) GetScores(context.Context, *GetScoresRequest) (*GetScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScores not implemented")
}
func (UnimplementedVierkantleServiceServer) Whoami(context.Context, *WhoamiRequest) (*WhoamiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Whoami not implemented")
}
func (UnimplementedVierkantleServiceServer) StartLogin(context.Context, *StartLoginRequest) (*StartLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLogin not implemented")
}
func (UnimplementedVierkantleServiceServer) FinishLogin(context.Context, *FinishLoginRequest) (*FinishLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishLogin not implemented")
}
func (UnimplementedVierkantleServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedVierkantleServiceServer) WordsForBoard(context.Context, *WordsForBoardRequest) (*WordsForBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WordsForBoard not implemented")
}
func (UnimplementedVierkantleServiceServer) SeedBoard(*SeedBoardRequest, VierkantleService_SeedBoardServer) error {
	return status.Errorf(codes.Unimplemented, "method SeedBoard not implemented")
}
func (UnimplementedVierkantleServiceServer) FillInBoard(*FillInBoardRequest, VierkantleService_FillInBoardServer) error {
	return status.Errorf(codes.Unimplemented, "method FillInBoard not implemented")
}
func (UnimplementedVierkantleServiceServer) TeamStream(VierkantleService_TeamStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TeamStream not implemented")
}

// UnsafeVierkantleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VierkantleServiceServer will
// result in compilation errors.
type UnsafeVierkantleServiceServer interface {
	mustEmbedUnimplementedVierkantleServiceServer()
}

func RegisterVierkantleServiceServer(s grpc.ServiceRegistrar, srv VierkantleServiceServer) {
	s.RegisterService(&VierkantleService_ServiceDesc, srv)
}

func _VierkantleService_GetBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).GetBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/GetBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).GetBoard(ctx, req.(*GetBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_SubmitScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).SubmitScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/SubmitScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).SubmitScore(ctx, req.(*SubmitScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_GetScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).GetScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/GetScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).GetScores(ctx, req.(*GetScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_Whoami_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoamiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).Whoami(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/Whoami",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).Whoami(ctx, req.(*WhoamiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_StartLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).StartLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/StartLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).StartLogin(ctx, req.(*StartLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_FinishLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).FinishLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/FinishLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).FinishLogin(ctx, req.(*FinishLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_WordsForBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WordsForBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).WordsForBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/WordsForBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).WordsForBoard(ctx, req.(*WordsForBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_SeedBoard_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SeedBoardRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VierkantleServiceServer).SeedBoard(m, &vierkantleServiceSeedBoardServer{stream})
}

type VierkantleService_SeedBoardServer interface {
	Send(*SeedBoardResponse) error
	grpc.ServerStream
}

type vierkantleServiceSeedBoardServer struct {
	grpc.ServerStream
}

func (x *vierkantleServiceSeedBoardServer) Send(m *SeedBoardResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VierkantleService_FillInBoard_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FillInBoardRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VierkantleServiceServer).FillInBoard(m, &vierkantleServiceFillInBoardServer{stream})
}

type VierkantleService_FillInBoardServer interface {
	Send(*FillInBoardResponse) error
	grpc.ServerStream
}

type vierkantleServiceFillInBoardServer struct {
	grpc.ServerStream
}

func (x *vierkantleServiceFillInBoardServer) Send(m *FillInBoardResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VierkantleService_TeamStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VierkantleServiceServer).TeamStream(&vierkantleServiceTeamStreamServer{stream})
}

type VierkantleService_TeamStreamServer interface {
	Send(*TeamStreamServerMessage) error
	Recv() (*TeamStreamClientMessage, error)
	grpc.ServerStream
}

type vierkantleServiceTeamStreamServer struct {
	grpc.ServerStream
}

func (x *vierkantleServiceTeamStreamServer) Send(m *TeamStreamServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vierkantleServiceTeamStreamServer) Recv() (*TeamStreamClientMessage, error) {
	m := new(TeamStreamClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VierkantleService_ServiceDesc is the grpc.ServiceDesc for VierkantleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VierkantleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nl.vierkantle.VierkantleService",
	HandlerType: (*VierkantleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBoard",
			Handler:    _VierkantleService_GetBoard_Handler,
		},
		{
			MethodName: "SubmitScore",
			Handler:    _VierkantleService_SubmitScore_Handler,
		},
		{
			MethodName: "GetScores",
			Handler:    _VierkantleService_GetScores_Handler,
		},
		{
			MethodName: "Whoami",
			Handler:    _VierkantleService_Whoami_Handler,
		},
		{
			MethodName: "StartLogin",
			Handler:    _VierkantleService_StartLogin_Handler,
		},
		{
			MethodName: "FinishLogin",
			Handler:    _VierkantleService_FinishLogin_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _VierkantleService_Register_Handler,
		},
		{
			MethodName: "WordsForBoard",
			Handler:    _VierkantleService_WordsForBoard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SeedBoard",
			Handler:       _VierkantleService_SeedBoard_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FillInBoard",
			Handler:       _VierkantleService_FillInBoard_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TeamStream",
			Handler:       _VierkantleService_TeamStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "vierkantle.proto",
}
