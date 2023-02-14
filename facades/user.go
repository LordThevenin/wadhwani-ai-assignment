package facades

import (
	"user-service/models"
	"user-service/services"
	"user-service/utils"
)

type IUserFacade interface {
	UploadUsers(models.UserFileUpload) error
	GetUser(int64) (models.User, error)
}

type UserFacade struct {
	userService        services.IUserService
	translationService services.ITranslationService
}

func InitUserFacade() *UserFacade {
	uf := new(UserFacade)
	uf.userService = services.InitUserService()
	uf.translationService = services.InitGoogleTranslationService()
	return uf
}

func (f *UserFacade) UploadUsers(uploadData models.UserFileUpload) (err error) {
	// Transform upload form data to list of user models
	users, err := utils.ParseUserUploadFile(uploadData)
	if err != nil {
		// Log failure to parse upload file
		return
	}
	// Give user service to handle list of user models
	err = f.userService.UploadUsers(users)
	if err != nil {
		// Log failure to upload users
		return
	}
	return
}

func (f *UserFacade) GetUser(userId int64) (user models.User, err error) {
	// Fetch user from service
	user, err = f.userService.GetUser(userId)
	if err != nil {
		// Log failure to upload users
		return
	}
	// Translate to requested language
	return
}
