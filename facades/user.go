package facades

import (
	"user-service/models"
	"user-service/services"
)

type IUserFacade interface {
	UploadUsers(models.UserFileUpload) error
	GetUser()
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
	// Give user service to handle list of user models
	err = f.userService.UploadUsers([]models.User{})
	if err != nil {
		// Log failure to upload users
	}
	return
}

func (f *UserFacade) GetUser() {

}
