package facades

import "user-service/services"

type IUserFacade interface {
	UploadUsers()
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

func (f *UserFacade) UploadUsers() {

}

func (f *UserFacade) GetUser() {

}
