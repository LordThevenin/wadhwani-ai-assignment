package server

import (
	"github.com/gin-gonic/gin"
	"user-service/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	var uc controllers.IUserController
	uc = controllers.InitUserController()

	root := router.Group("/user-service")
	{
		ur := root.Group("/user")
		ur.POST("/upload", uc.UploadUsers)
		ur.GET("/:id", uc.GetUser)
	}

	return router
}
