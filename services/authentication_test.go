package services

import (
	"reflect"
	"testing"
	"user-service/entities"
	"user-service/models"
	"user-service/repositories"
)

func TestAuthenticationService_Login(t *testing.T) {
	type fields struct {
		authRepository repositories.IAuthUserRepository
	}
	type args struct {
		user models.AuthUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantJwt string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthenticationService{
				authRepository: tt.fields.authRepository,
			}
			gotJwt, err := s.Login(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotJwt != tt.wantJwt {
				t.Errorf("Login() gotJwt = %v, want %v", gotJwt, tt.wantJwt)
			}
		})
	}
}

func TestAuthenticationService_Register(t *testing.T) {
	type fields struct {
		authRepository repositories.IAuthUserRepository
	}
	type args struct {
		user models.AuthUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthenticationService{
				authRepository: tt.fields.authRepository,
			}
			if err := s.Register(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthenticationService_getAuthUser(t *testing.T) {
	type fields struct {
		authRepository repositories.IAuthUserRepository
	}
	type args struct {
		user models.AuthUser
		err  error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.AuthUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthenticationService{
				authRepository: tt.fields.authRepository,
			}
			got, err := s.getAuthUser(tt.args.user, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAuthUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAuthUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthenticationService_validatePassword(t *testing.T) {
	type fields struct {
		authRepository repositories.IAuthUserRepository
	}
	type args struct {
		user       models.AuthUser
		err        error
		userEntity []entities.AuthUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthenticationService{
				authRepository: tt.fields.authRepository,
			}
			if err := s.validatePassword(tt.args.user, tt.args.err, tt.args.userEntity); (err != nil) != tt.wantErr {
				t.Errorf("validatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
