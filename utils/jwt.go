package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"user-service/config"
	"user-service/entities"
)

func GenerateToken(user entities.AuthUser) (string, error) {
	cfg := config.GetConfig()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.ApiSecret))

}
