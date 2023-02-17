package server

import (
	"github.com/gin-gonic/gin"
	"user-service/controllers"
	"user-service/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	var uc controllers.IUserController
	var ac controllers.IAuthController
	uc = controllers.InitUserController()
	ac = controllers.InitAuthController()

	root := router.Group("/user-service")
	{
		ur := root.Group("/user")
		ur.POST("/upload", middlewares.JwtAuthMiddleware(), uc.UploadUsers)
		ur.GET("/:id", middlewares.JwtAuthMiddleware(), uc.GetUser)

		ar := root.Group("/auth")
		ar.POST("/register", ac.Register)
		ar.POST("/login", ac.Login)
	}

	return router
}
