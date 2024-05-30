package endpoints

import (
	"reflect"
	"testing"

	"github.com/Venukishore-R/microservice1_auth/services"
	"github.com/go-kit/kit/endpoint"
)

func Test_makeRegisterEndpoint(t *testing.T) {
	type args struct {
		s services.Service
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRegisterEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRegisterEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
