package mocks

import (
	"user-service/entities"
)

type UserSQLRepositoryMock struct {
}

// UpsertUsers upserts user values based on phone number
func (s *UserSQLRepositoryMock) UpsertUsers(users []entities.User) (err error) {
	return
}

// GetUser gets valid user value based on id
func (s *UserSQLRepositoryMock) GetUser(id int64) (user []entities.User, err error) {
	return
}
