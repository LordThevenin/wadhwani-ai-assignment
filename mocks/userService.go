package mocks

import (
	"golang.org/x/text/language"
	"user-service/models"
)

type UserServiceMock struct {
}

func (s *UserServiceMock) UploadUsers(users []models.User) (err error) {
	return
}

func (s *UserServiceMock) GetUser(userId int64) (user models.User, err error) {
	return
}

func (s *UserServiceMock) GetUserFromCache(userId int64, lang language.Tag) (user models.User, hit bool) {
	return
}

func (s *UserServiceMock) SetUserInCache(userId int64, lang language.Tag, user models.User) {
}
