package transports

import (
	"context"
	"github.com/Venukishore-R/microservice1_auth/endpoints"
	"github.com/Venukishore-R/microservice1_auth/protos"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/grpc"
)

type MyServer struct {
	register grpc.Handler
	login    grpc.Handler
	protos.UnimplementedUserServiceServer
}

func NewMyServer(endpoints endpoints.Endpoints, logger log.Logger) MyServer {
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
