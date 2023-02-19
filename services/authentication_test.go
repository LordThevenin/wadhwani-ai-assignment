package services

import (
	"testing"
	"user-service/config"
	"user-service/mocks"
	"user-service/models"
)

func TestAuthenticationService_Login(t *testing.T) {
	config.Init("../test.env")
	type args struct {
		user models.AuthUser
	}
	tests := []struct {
		name    string
		args    args
		wantJwt bool
		wantErr bool
	}{
		{
			"Success",
			args{
				models.AuthUser{
					UserName: "test",
					Password: "password",
				},
			},
			true,
			false,
		},
		{
			"Fail Password",
			args{
				models.AuthUser{
					UserName: "test2",
					Password: "password",
				},
			},
			false,
			true,
		},
		{
			"Fail GetUser",
			args{
				models.AuthUser{
					UserName: "test3",
					Password: "password",
				},
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthenticationService{
				authRepository: &mocks.AuthUserSQLRepositoryMock{},
			}
			gotJwt, err := s.Login(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotJwt != "") != tt.wantJwt {
				t.Errorf("Login() gotJwt = %v, want %v", gotJwt, tt.wantJwt)
			}
		})
	}
}

func TestAuthenticationService_Register(t *testing.T) {
	type args struct {
		user models.AuthUser
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"Success",
			args{models.AuthUser{
				UserName: "test",
				Password: "password",
			}},
			false,
		},
		{
			"Fail",
			args{models.AuthUser{
				UserName: "test2",
				Password: "password2",
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthenticationService{
				authRepository: &mocks.AuthUserSQLRepositoryMock{},
			}
			if err := s.Register(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
