package dto

import (
	"user-service/entities"
	"user-service/models"
)

func AuthUserModelToEntity(user models.AuthUser, version int) (userEntity entities.AuthUser) {
	userEntity.Username = user.UserName
	userEntity.Password = user.Password
	userEntity.Role = "user"
	userEntity.Version = version
	return
}
