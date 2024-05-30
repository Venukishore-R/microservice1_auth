package endpoints

import (
	"context"
	"github.com/Venukishore-R/microservice1_auth/services"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Register         endpoint.Endpoint
	Login            endpoint.Endpoint
	Authenticate     endpoint.Endpoint
	GenerateNewToken endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		Register:         makeRegisterEndpoint(s),
		Login:            makeLoginEndpoint(s),
		Authenticate:     makeAuthenticateEndpoint(s),
		GenerateNewToken: makeGenerateNewToken(s),
	}
}

func makeRegisterEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		status, desc, err := s.Register(ctx, req.Name, req.Email, req.Phone, req.Password)

		return RegisterResponse{
			Status:      status,
			Description: desc,
		}, err
	}
}

func makeLoginEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		status, accessToken, refreshToken, err := s.Login(ctx, req.Email, req.Password)
		return LoginResponse{
			Status:       status,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, err
	}
}

func makeAuthenticateEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(Empty)
		status, name, email, phone, err := s.Authenticate(ctx)
		return AuthUserResp{Status: status, Name: name, Email: email, Phone: phone}, err
	}
}

func makeGenerateNewToken(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(Empty)
		accessToken, err := s.GenerateNewToken(ctx)
		return GenerateNewTokResp{
			AccessToken: accessToken,
		}, err
	}
}
