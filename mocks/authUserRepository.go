package mocks

import (
	"fmt"
	"user-service/entities"
)

type AuthUserSQLRepositoryMock struct {
}

func (s *AuthUserSQLRepositoryMock) AddUser(user entities.AuthUser) (err error) {
	switch user.Username {
	case "test2":
		{
			err = fmt.Errorf("Error")
		}

	}
	return
}

func (s *AuthUserSQLRepositoryMock) GetUser(userName string) (user []entities.AuthUser, err error) {
	switch userName {
	case "test":
		{
			user = append(user, entities.AuthUser{
				ID:       1,
				Username: "test",
				Password: "$2a$10$BykiymlGi1W/x6NJHUfUzOspwtYPG3pkweSEf/jzx.N9W.hOcpSE6",
				Role:     "user",
				Base:     entities.Base{},
			})
		}
	case "test2":
		{
			user = append(user, entities.AuthUser{
				ID:       2,
				Username: "test2",
				Password: "password2",
				Role:     "user",
				Base:     entities.Base{},
			})
		}
	}
	return
}
