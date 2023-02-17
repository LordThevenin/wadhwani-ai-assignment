package models

type AuthUser struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
