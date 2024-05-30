package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-kit/log"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func TestConnectDb(t *testing.T) {
	tests := []struct {
		name    string
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectDb()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerService_Register(t *testing.T) {
	type fields struct {
		logger log.Logger
	}
	type args struct {
		ctx      context.Context
		name     string
		email    string
		phone    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LoggerService{
				logger: tt.fields.logger,
			}
			got, got1, err := s.Register(tt.args.ctx, tt.args.name, tt.args.email, tt.args.phone, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoggerService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LoggerService.Register() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LoggerService.Register() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
