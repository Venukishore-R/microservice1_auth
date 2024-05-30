package transports

import (
	"context"

	"github.com/Venukishore-R/microservice1_auth/endpoints"
	"github.com/Venukishore-R/microservice1_auth/protos"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type MyServer struct {
	register          grpc.Handler
	login             grpc.Handler
	authenticate      grpc.Handler
	generateNewAccTok grpc.Handler
	protos.UnimplementedUserServiceServer
}

func NewMyServer(endpoints endpoints.Endpoints, logger log.Logger) MyServer {
	options := []grpc.ServerOption{grpc.ServerBefore(jwt.GRPCToContext())}
	//clientOption := grpc.ClientBefore(jwt.ContextToGRPC())

	return MyServer{
		register: grpc.NewServer(
			endpoints.Register,
			decodeRegisterRequest,
			encodeRegisterResponse,
		),
		login: grpc.NewServer(
			endpoints.Login,
			decodeLoginRequest,
			encodeLoginResponse,
		),
		authenticate: grpc.NewServer(
			endpoints.Authenticate,
			decodeAuthReq,
			encodeAuthResp,
			//clientOption,
			options...,
		),
		generateNewAccTok: grpc.NewServer(
			endpoints.GenerateNewToken,
			decodeGenerateNewTokReq,
			encodeGenerateNewTokResp,
			options...,
		),
	}
}

func (s *MyServer) Register(ctx context.Context, request *protos.User) (*protos.UserRegResp, error) {
	_, resp, err := s.register.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*protos.UserRegResp), err
}

func decodeRegisterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.User)
	return endpoints.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}, nil
}

func encodeRegisterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.RegisterResponse)
	return &protos.UserRegResp{
		Status:      resp.Status,
		Description: resp.Description,
	}, nil
}

func (s *MyServer) Login(ctx context.Context, request *protos.UserLoginReq) (*protos.UserLoginResp, error) {
	_, resp, err := s.login.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*protos.UserLoginResp), err
}

func decodeLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.UserLoginReq)
	return endpoints.LoginRequest{Email: req.Email, Password: req.Password}, nil
}

func encodeLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.LoginResponse)
	return &protos.UserLoginResp{Status: resp.Status, AccessToken: resp.AccessToken, RefreshToken: resp.RefreshToken}, nil
}

func (s *MyServer) Authenticate(ctx context.Context, request *protos.Empty) (*protos.AuthenticateUserResp, error) {
	_, resp, err := s.authenticate.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.AuthenticateUserResp), nil
}

func decodeAuthReq(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*protos.Empty)
	return endpoints.Empty{}, nil
}

func encodeAuthResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.AuthUserResp)
	return &protos.AuthenticateUserResp{Status: resp.Status, Name: resp.Name, Email: resp.Email, Phone: resp.Phone}, nil
}

func (s *MyServer) GenerateNewToken(ctx context.Context, request *protos.Empty) (*protos.GenerateNewAccTokResp, error) {
	_, resp, err := s.generateNewAccTok.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.GenerateNewAccTokResp), nil
}

func decodeGenerateNewTokReq(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*protos.Empty)
	return endpoints.Empty{}, nil
}

func encodeGenerateNewTokResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GenerateNewTokResp)
	return &protos.GenerateNewAccTokResp{AccessToken: resp.AccessToken}, nil
}
