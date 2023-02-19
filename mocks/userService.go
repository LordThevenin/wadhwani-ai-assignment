package mocks

import (
	"fmt"
	"golang.org/x/text/language"
	"user-service/models"
)

type UserServiceMock struct {
}

func (s *UserServiceMock) UploadUsers(users []models.User) (err error) {
	return
}

func (s *UserServiceMock) GetUser(userId int64) (user models.User, err error) {
	switch userId {
	case 2:
		{
			user = models.User{
				PhoneNumber: 73572,
				Name:        "testName2",
				State:       "testState2",
				District:    "testDistrict2",
				Village:     "testVillage2",
			}
		}

	default:
		{
			err = fmt.Errorf("error")
		}
	}
	return
}

func (s *UserServiceMock) GetUserFromCache(userId int64, lang language.Tag) (user models.User, hit bool) {
	switch userId {
	case 1:
		{
			user = models.User{
				PhoneNumber: 7357,
				Name:        "testName",
				State:       "testState",
				District:    "testDistrict",
				Village:     "testVillage",
			}
			hit = true
		}

	default:
		{
			hit = false
		}
	}
	return
}

func (s *UserServiceMock) SetUserInCache(userId int64, lang language.Tag, user models.User) {
}
