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
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	// Generator
	WordsForBoard(ctx context.Context, in *WordsForBoardRequest, opts ...grpc.CallOption) (*WordsForBoardResponse, error)
	SeedBoard(ctx context.Context, in *SeedBoardRequest, opts ...grpc.CallOption) (VierkantleService_SeedBoardClient, error)
	FillInBoard(ctx context.Context, in *FillInBoardRequest, opts ...grpc.CallOption) (VierkantleService_FillInBoardClient, error)
	MarkWordType(ctx context.Context, in *MarkWordTypeRequest, opts ...grpc.CallOption) (*MarkWordTypeResponse, error)
	GetNewestBoard(ctx context.Context, in *GetNewestBoardRequest, opts ...grpc.CallOption) (*GetNewestBoardResponse, error)
	AddBoardToQueue(ctx context.Context, in *AddBoardToQueueRequest, opts ...grpc.CallOption) (*AddBoardToQueueResponse, error)
	ListBoardQueue(ctx context.Context, in *ListBoardQueueRequest, opts ...grpc.CallOption) (*ListBoardQueueResponse, error)
	GetBoardFromQueue(ctx context.Context, in *GetBoardFromQueueRequest, opts ...grpc.CallOption) (*GetBoardFromQueueResponse, error)
	UpdateBoardInQueue(ctx context.Context, in *UpdateBoardInQueueRequest, opts ...grpc.CallOption) (*UpdateBoardInQueueResponse, error)
	RemoveBoardsFromQueue(ctx context.Context, in *RemoveBoardsFromQueueRequest, opts ...grpc.CallOption) (*RemoveBoardsFromQueueResponse, error)
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

func (c *vierkantleServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/Logout", in, out, opts...)
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

func (c *vierkantleServiceClient) MarkWordType(ctx context.Context, in *MarkWordTypeRequest, opts ...grpc.CallOption) (*MarkWordTypeResponse, error) {
	out := new(MarkWordTypeResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/MarkWordType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) GetNewestBoard(ctx context.Context, in *GetNewestBoardRequest, opts ...grpc.CallOption) (*GetNewestBoardResponse, error) {
	out := new(GetNewestBoardResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/GetNewestBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) AddBoardToQueue(ctx context.Context, in *AddBoardToQueueRequest, opts ...grpc.CallOption) (*AddBoardToQueueResponse, error) {
	out := new(AddBoardToQueueResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/AddBoardToQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) ListBoardQueue(ctx context.Context, in *ListBoardQueueRequest, opts ...grpc.CallOption) (*ListBoardQueueResponse, error) {
	out := new(ListBoardQueueResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/ListBoardQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) GetBoardFromQueue(ctx context.Context, in *GetBoardFromQueueRequest, opts ...grpc.CallOption) (*GetBoardFromQueueResponse, error) {
	out := new(GetBoardFromQueueResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/GetBoardFromQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) UpdateBoardInQueue(ctx context.Context, in *UpdateBoardInQueueRequest, opts ...grpc.CallOption) (*UpdateBoardInQueueResponse, error) {
	out := new(UpdateBoardInQueueResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/UpdateBoardInQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vierkantleServiceClient) RemoveBoardsFromQueue(ctx context.Context, in *RemoveBoardsFromQueueRequest, opts ...grpc.CallOption) (*RemoveBoardsFromQueueResponse, error) {
	out := new(RemoveBoardsFromQueueResponse)
	err := c.cc.Invoke(ctx, "/nl.vierkantle.VierkantleService/RemoveBoardsFromQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	// Generator
	WordsForBoard(context.Context, *WordsForBoardRequest) (*WordsForBoardResponse, error)
	SeedBoard(*SeedBoardRequest, VierkantleService_SeedBoardServer) error
	FillInBoard(*FillInBoardRequest, VierkantleService_FillInBoardServer) error
	MarkWordType(context.Context, *MarkWordTypeRequest) (*MarkWordTypeResponse, error)
	GetNewestBoard(context.Context, *GetNewestBoardRequest) (*GetNewestBoardResponse, error)
	AddBoardToQueue(context.Context, *AddBoardToQueueRequest) (*AddBoardToQueueResponse, error)
	ListBoardQueue(context.Context, *ListBoardQueueRequest) (*ListBoardQueueResponse, error)
	GetBoardFromQueue(context.Context, *GetBoardFromQueueRequest) (*GetBoardFromQueueResponse, error)
	UpdateBoardInQueue(context.Context, *UpdateBoardInQueueRequest) (*UpdateBoardInQueueResponse, error)
	RemoveBoardsFromQueue(context.Context, *RemoveBoardsFromQueueRequest) (*RemoveBoardsFromQueueResponse, error)
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
func (UnimplementedVierkantleServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
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
func (UnimplementedVierkantleServiceServer) MarkWordType(context.Context, *MarkWordTypeRequest) (*MarkWordTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkWordType not implemented")
}
func (UnimplementedVierkantleServiceServer) GetNewestBoard(context.Context, *GetNewestBoardRequest) (*GetNewestBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewestBoard not implemented")
}
func (UnimplementedVierkantleServiceServer) AddBoardToQueue(context.Context, *AddBoardToQueueRequest) (*AddBoardToQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBoardToQueue not implemented")
}
func (UnimplementedVierkantleServiceServer) ListBoardQueue(context.Context, *ListBoardQueueRequest) (*ListBoardQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBoardQueue not implemented")
}
func (UnimplementedVierkantleServiceServer) GetBoardFromQueue(context.Context, *GetBoardFromQueueRequest) (*GetBoardFromQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoardFromQueue not implemented")
}
func (UnimplementedVierkantleServiceServer) UpdateBoardInQueue(context.Context, *UpdateBoardInQueueRequest) (*UpdateBoardInQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBoardInQueue not implemented")
}
func (UnimplementedVierkantleServiceServer) RemoveBoardsFromQueue(context.Context, *RemoveBoardsFromQueueRequest) (*RemoveBoardsFromQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBoardsFromQueue not implemented")
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

func _VierkantleService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).Logout(ctx, req.(*LogoutRequest))
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

func _VierkantleService_MarkWordType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkWordTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).MarkWordType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/MarkWordType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).MarkWordType(ctx, req.(*MarkWordTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_GetNewestBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewestBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).GetNewestBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/GetNewestBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).GetNewestBoard(ctx, req.(*GetNewestBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_AddBoardToQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBoardToQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).AddBoardToQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/AddBoardToQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).AddBoardToQueue(ctx, req.(*AddBoardToQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_ListBoardQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBoardQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).ListBoardQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/ListBoardQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).ListBoardQueue(ctx, req.(*ListBoardQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_GetBoardFromQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoardFromQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).GetBoardFromQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/GetBoardFromQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).GetBoardFromQueue(ctx, req.(*GetBoardFromQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_UpdateBoardInQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoardInQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).UpdateBoardInQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/UpdateBoardInQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).UpdateBoardInQueue(ctx, req.(*UpdateBoardInQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VierkantleService_RemoveBoardsFromQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBoardsFromQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VierkantleServiceServer).RemoveBoardsFromQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nl.vierkantle.VierkantleService/RemoveBoardsFromQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VierkantleServiceServer).RemoveBoardsFromQueue(ctx, req.(*RemoveBoardsFromQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "Logout",
			Handler:    _VierkantleService_Logout_Handler,
		},
		{
			MethodName: "WordsForBoard",
			Handler:    _VierkantleService_WordsForBoard_Handler,
		},
		{
			MethodName: "MarkWordType",
			Handler:    _VierkantleService_MarkWordType_Handler,
		},
		{
			MethodName: "GetNewestBoard",
			Handler:    _VierkantleService_GetNewestBoard_Handler,
		},
		{
			MethodName: "AddBoardToQueue",
			Handler:    _VierkantleService_AddBoardToQueue_Handler,
		},
		{
			MethodName: "ListBoardQueue",
			Handler:    _VierkantleService_ListBoardQueue_Handler,
		},
		{
			MethodName: "GetBoardFromQueue",
			Handler:    _VierkantleService_GetBoardFromQueue_Handler,
		},
		{
			MethodName: "UpdateBoardInQueue",
			Handler:    _VierkantleService_UpdateBoardInQueue_Handler,
		},
		{
			MethodName: "RemoveBoardsFromQueue",
			Handler:    _VierkantleService_RemoveBoardsFromQueue_Handler,
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
