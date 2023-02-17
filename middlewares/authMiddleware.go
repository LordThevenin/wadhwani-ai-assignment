package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.ValidateToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, _ := claims["id"].(int64)
			role, _ := claims["role"].(string)
			c.Set("user_id", userID)
			c.Set("custom_role", role)
		}
		c.Next()
	}
}
