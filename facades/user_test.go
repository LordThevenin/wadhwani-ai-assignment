package facades

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"reflect"
	"testing"
	"user-service/mocks"
	"user-service/models"
)

func TestUserFacade_GetUser(t *testing.T) {
	type args struct {
		ctx    *gin.Context
		userId int64
		lang   language.Tag
	}
	type testStruct struct {
		name     string
		args     args
		wantUser models.User
		wantErr  bool
	}
	tests := []testStruct{
		// TODO: Add test cases.
		{
			name:     "",
			args:     args{},
			wantUser: models.User{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &UserFacade{
				userService:        &mocks.UserServiceMock{},
				translationService: &mocks.GoogleTranslationServiceMock{},
			}
			gotUser, err := f.GetUser(tt.args.ctx, tt.args.userId, tt.args.lang)
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

func TestUserFacade_UploadUsers(t *testing.T) {
	type args struct {
		uploadData models.UserFileUpload
	}
	type testStruct struct {
		name    string
		args    args
		wantErr bool
	}
	tests := []testStruct{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &UserFacade{
				userService:        &mocks.UserServiceMock{},
				translationService: &mocks.GoogleTranslationServiceMock{},
			}
			if err := f.UploadUsers(tt.args.uploadData); (err != nil) != tt.wantErr {
				t.Errorf("UploadUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserFacade_translateUserModelToTargetLanguage(t *testing.T) {
	type args struct {
		ctx  *gin.Context
		user models.User
		lang language.Tag
	}
	type testStruct struct {
		name               string
		args               args
		wantTranslatedUser models.User
		wantErr            bool
	}
	tests := []testStruct{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &UserFacade{
				userService:        &mocks.UserServiceMock{},
				translationService: &mocks.GoogleTranslationServiceMock{},
			}
			gotTranslatedUser, err := f.translateUserModelToTargetLanguage(tt.args.ctx, tt.args.user, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("translateUserModelToTargetLanguage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTranslatedUser, tt.wantTranslatedUser) {
				t.Errorf("translateUserModelToTargetLanguage() gotTranslatedUser = %v, want %v", gotTranslatedUser, tt.wantTranslatedUser)
			}
		})
	}
}
