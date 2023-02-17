package server

import (
	"github.com/gin-gonic/gin"
	"user-service/controllers"
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
		ur.POST("/upload", uc.UploadUsers)
		ur.GET("/:id", uc.GetUser)

		ar := root.Group("/auth")
		ar.POST("/register", ac.Register)
		ar.POST("/login", ac.Login)
	}

	return router
}
