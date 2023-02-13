package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/facades"
	"user-service/models"
)

type IUserController interface {
	UploadUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
}

type UserController struct {
	userFacade facades.IUserFacade
}

func InitUserController() *UserController {
	uc := new(UserController)
	uc.userFacade = facades.InitUserFacade()
	return uc
}

func (c *UserController) UploadUsers(ctx *gin.Context) {
	var uploadData models.UserFileUpload
	err := ctx.ShouldBind(&uploadData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	err = c.userFacade.UploadUsers(uploadData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "user upload failed")
		return
	}
	ctx.JSON(http.StatusOK, "success")
}

func (c *UserController) GetUser(ctx *gin.Context) {

}
