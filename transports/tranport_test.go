package transports

import (
	"context"
	"reflect"
	"testing"

	"github.com/Venukishore-R/microservice1_auth/protos"
	"github.com/go-kit/kit/transport/grpc"
)

func TestMyServer_Login(t *testing.T) {
	type fields struct {
		register                       grpc.Handler
		login                          grpc.Handler
		authenticate                   grpc.Handler
		generateNewAccTok              grpc.Handler
		UnimplementedUserServiceServer protos.UnimplementedUserServiceServer
	}
	type args struct {
		ctx     context.Context
		request *protos.UserLoginReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *protos.UserLoginResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyServer{
				register:                       tt.fields.register,
				login:                          tt.fields.login,
				authenticate:                   tt.fields.authenticate,
				generateNewAccTok:              tt.fields.generateNewAccTok,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := s.Login(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyServer.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
