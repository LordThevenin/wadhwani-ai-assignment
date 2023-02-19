package mocks

import (
	"user-service/models"
)

type RedisUserCacheRepositoryMock struct {
}

func (r *RedisUserCacheRepositoryMock) Get(key string) (user models.User, err error) {
	return
}

func (r *RedisUserCacheRepositoryMock) Set(key string, user models.User) {
}
