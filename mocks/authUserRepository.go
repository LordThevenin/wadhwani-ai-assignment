package mocks

import (
	"user-service/entities"
)

type AuthUserSQLRepositoryMock struct {
}

func (s *AuthUserSQLRepositoryMock) AddUser(user entities.AuthUser) (err error) {
	return
}

func (s *AuthUserSQLRepositoryMock) GetUser(userName string) (user []entities.AuthUser, err error) {
	return
}
