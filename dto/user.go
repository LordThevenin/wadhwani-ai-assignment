package dto

import (
	"user-service/entities"
	"user-service/models"
)

func UserModelToUserEntity(user models.User) (userEntity entities.User) {
	userEntity.Name = user.Name
	userEntity.PhoneNumber = user.PhoneNumber
	userEntity.State = user.State
	userEntity.District = user.District
	userEntity.Village = user.Village
	return
}

func UserModelsToUserEntities(users []models.User) (userEntities []entities.User) {
	for _, user := range users {
		userEntities = append(userEntities, UserModelToUserEntity(user))
	}
	return
}
