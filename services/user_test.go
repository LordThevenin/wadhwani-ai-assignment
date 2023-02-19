package services

import (
	"golang.org/x/text/language"
	"reflect"
	"testing"
	"user-service/mocks"
	"user-service/models"
)

func TestUserService_GetUser(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name     string
		args     args
		wantUser models.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      &mocks.UserSQLRepositoryMock{},
				userCacheRepository: &mocks.RedisUserCacheRepositoryMock{},
			}
			gotUser, err := s.GetUser(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUser() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserService_GetUserFromCache(t *testing.T) {
	type args struct {
		userId int64
		lang   language.Tag
	}
	tests := []struct {
		name     string
		args     args
		wantUser models.User
		wantHit  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      &mocks.UserSQLRepositoryMock{},
				userCacheRepository: &mocks.RedisUserCacheRepositoryMock{},
			}
			gotUser, gotHit := s.GetUserFromCache(tt.args.userId, tt.args.lang)
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserFromCache() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
			if gotHit != tt.wantHit {
				t.Errorf("GetUserFromCache() gotHit = %v, want %v", gotHit, tt.wantHit)
			}
		})
	}
}

func TestUserService_SetUserInCache(t *testing.T) {
	type args struct {
		userId int64
		lang   language.Tag
		user   models.User
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      &mocks.UserSQLRepositoryMock{},
				userCacheRepository: &mocks.RedisUserCacheRepositoryMock{},
			}
			s.SetUserInCache(tt.args.userId, tt.args.lang, tt.args.user)
		})
	}
}

func TestUserService_UploadUsers(t *testing.T) {
	type args struct {
		users []models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				userRepository:      &mocks.UserSQLRepositoryMock{},
				userCacheRepository: &mocks.RedisUserCacheRepositoryMock{},
			}
			if err := s.UploadUsers(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("UploadUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
