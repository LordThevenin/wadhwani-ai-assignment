package mocks

import (
	"fmt"
	"user-service/entities"
)

type UserSQLRepositoryMock struct {
}

// UpsertUsers upserts user values based on phone number
func (s *UserSQLRepositoryMock) UpsertUsers(users []entities.User) (err error) {
	switch users[0].Name {
	case "test2":
		{
			err = fmt.Errorf("error")
		}

	}
	return
}

// GetUser gets valid user value based on id
func (s *UserSQLRepositoryMock) GetUser(id int64) (user []entities.User, err error) {
	switch id {
	case 0:
		{
			err = fmt.Errorf("error")
		}
	case 1:
		{
			user = append(user, entities.User{
				ID:          1,
				PhoneNumber: 7357,
				Name:        "test",
				Location: entities.Location{
					State:    "testState",
					District: "testDistrict",
					Village:  "testVillage",
				},
				Base: entities.Base{},
			})
		}

	}
	return
}
