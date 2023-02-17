package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/models"
	"user-service/services"
)

type IAuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type AuthController struct {
	authService services.IAuthenticationService
}

func InitAuthController() *AuthController {
	ac := new(AuthController)
	ac.authService = services.InitAuthorizationService()
	return ac
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	err = c.authService.Register(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "user upload failed")
		return
	}
	ctx.JSON(http.StatusOK, "success")
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user models.AuthUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	token, err := c.authService.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "user upload failed")
		return
	}
	ctx.JSON(http.StatusOK, token)
}
