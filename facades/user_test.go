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
		{
			name: "Cache Success",
			args: args{
				&gin.Context{},
				1,
				language.English,
			},
			wantUser: models.User{
				PhoneNumber: 7357,
				Name:        "testName",
				State:       "testState",
				District:    "testDistrict",
				Village:     "testVillage",
			},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				&gin.Context{},
				3,
				language.English,
			},
			wantUser: models.User{},
			wantErr:  true,
		},
		{
			name: "DB Success",
			args: args{
				&gin.Context{},
				2,
				language.English,
			},
			wantUser: models.User{
				PhoneNumber: 73572,
				Name:        "testName2",
				State:       "testState2",
				District:    "testDistrict2",
				Village:     "testVillage2",
			},
			wantErr: false,
		},
		{
			name: "DB Success Translated",
			args: args{
				&gin.Context{},
				2,
				language.Hindi,
			},
			wantUser: models.User{
				PhoneNumber: 73572,
				Name:        "testTranslatedName2",
				State:       "testTranslatedState2",
				District:    "testTranslatedDistrict2",
				Village:     "testTranslatedVillage2",
			},
			wantErr: false,
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
		{
			name: "Translate Non English",
			args: args{
				&gin.Context{},
				models.User{
					PhoneNumber: 73572,
					Name:        "testName2",
					State:       "testState2",
					District:    "testDistrict2",
					Village:     "testVillage2",
				},
				language.Hindi,
			},
			wantTranslatedUser: models.User{
				PhoneNumber: 73572,
				Name:        "testTranslatedName2",
				State:       "testTranslatedState2",
				District:    "testTranslatedDistrict2",
				Village:     "testTranslatedVillage2",
			},
			wantErr: false,
		},
		{
			name: "Not Translate English",
			args: args{
				&gin.Context{},
				models.User{
					PhoneNumber: 73572,
					Name:        "testName2",
					State:       "testState2",
					District:    "testDistrict2",
					Village:     "testVillage2",
				},
				language.English,
			},
			wantTranslatedUser: models.User{
				PhoneNumber: 73572,
				Name:        "testName2",
				State:       "testState2",
				District:    "testDistrict2",
				Village:     "testVillage2",
			},
			wantErr: false,
		},
		{
			name: "Translate Failed",
			args: args{
				&gin.Context{},
				models.User{
					PhoneNumber: 73572,
					Name:        "testName2",
					State:       "testState2",
					District:    "testDistrict2",
					Village:     "testVillage2",
				},
				language.French,
			},
			wantTranslatedUser: models.User{},
			wantErr:            true,
		},
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
