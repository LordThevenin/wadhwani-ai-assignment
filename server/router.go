package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"user-service/controllers"
	"user-service/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	var uc controllers.IUserController
	var ac controllers.IAuthController
	uc = controllers.InitUserController()
	ac = controllers.InitAuthController()

	root := router.Group("/user-service")
	{
		root.Static("/swaggerui/", "cmd/api/swaggerui")
		ur := root.Group("/user")
		ur.POST("/upload", middlewares.CORSMiddleware(), middlewares.JwtAuthMiddleware(), uc.UploadUsers)
		ur.GET("/:id", middlewares.CORSMiddleware(), middlewares.JwtAuthMiddleware(), uc.GetUser)

		ar := root.Group("/auth")
		ar.POST("/register", ac.Register)
		ar.POST("/login", ac.Login)
	}

	return router
}
