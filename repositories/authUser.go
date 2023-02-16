package repositories

import (
	"gorm.io/gorm"
	"user-service/db"
	"user-service/entities"
)

type IAuthUserRepository interface {
	AddUser(user entities.AuthUser) error
	GetUser(id int64) ([]entities.AuthUser, error)
}

type AuthUserSQLRepository struct {
	db *gorm.DB
}

func InitAuthUserSQLRepository() *AuthUserSQLRepository {
	ur := new(AuthUserSQLRepository)
	ur.db = db.GetSqlDB()
	return ur
}

func (s *AuthUserSQLRepository) AddUser(user entities.AuthUser) (err error) {
	err = s.db.Create(&user).Error
	return
}

func (s *AuthUserSQLRepository) GetUser(id int64) (user entities.AuthUser, err error) {
	result := s.db.Where("id = ? AND isDeleted = ?", id, false).First(&user)
	if result.Error != nil {
		err = result.Error
	}
	return
}
