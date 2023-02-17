package dto

import (
	"cloud.google.com/go/translate"
	"strconv"
	"user-service/entities"
	"user-service/models"
)

func UserModelToUserEntity(user models.User, version int) (userEntity entities.User) {
	userEntity.Name = user.Name
	userEntity.PhoneNumber = user.PhoneNumber
	userEntity.State = user.State
	userEntity.District = user.District
	userEntity.Village = user.Village
	userEntity.Version = version
	return
}

func UserModelsToUserEntities(users []models.User, version int) (userEntities []entities.User) {
	for _, user := range users {
		userEntities = append(userEntities, UserModelToUserEntity(user, version))
	}
	return
}

func UserEntityToUserModel(user entities.User) (userModel models.User) {
	userModel.Name = user.Name
	userModel.State = user.State
	userModel.Village = user.Village
	userModel.PhoneNumber = user.PhoneNumber
	userModel.District = user.District
	return
}

func UserModelToStringList(user models.User) (stringList []string) {
	stringList = append(stringList, strconv.FormatInt(user.PhoneNumber, 10))
	stringList = append(stringList, user.Name)
	stringList = append(stringList, user.State)
	stringList = append(stringList, user.Village)
	stringList = append(stringList, user.District)
	return
}

func TranslatedListToUserModel(translatedList []translate.Translation) (user models.User) {
	user.PhoneNumber, _ = strconv.ParseInt(translatedList[0].Text, 10, 64)
	user.Name = translatedList[1].Text
	user.State = translatedList[2].Text
	user.Village = translatedList[3].Text
	user.District = translatedList[4].Text
	return
}
