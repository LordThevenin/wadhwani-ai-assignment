package controllers

import "user-service/facades"

type IUserController interface {
	UploadUsers()
	GetUser()
}

type UserController struct {
	userFacade facades.IUserFacade
}

func InitUserController() *UserController {
	uc := new(UserController)
	uc.userFacade = facades.InitUserFacade()
	return uc
}

func (c *UserController) UploadUsers() {

}

func (c *UserController) GetUser() {

}
