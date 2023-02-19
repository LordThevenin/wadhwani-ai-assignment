package mocks

import (
	"fmt"
	"user-service/models"
)

type RedisUserCacheRepositoryMock struct {
}

func (r *RedisUserCacheRepositoryMock) Get(key string) (user models.User, err error) {
	switch key {
	case "1_en":
		{
			user = models.User{
				PhoneNumber: 7357,
				Name:        "test",
				State:       "testState",
				District:    "testDistrict",
				Village:     "testVillage",
			}
		}
	case "1_hi":
		{
			err = fmt.Errorf("error")
		}

	}
	return
}

func (r *RedisUserCacheRepositoryMock) Set(key string, user models.User) {
}
