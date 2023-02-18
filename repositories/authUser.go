package repositories

import (
	"gorm.io/gorm"
	"user-service/db"
	"user-service/entities"
	"user-service/utils"
)

type IAuthUserRepository interface {
	AddUser(user entities.AuthUser) error
	GetUser(userName string) ([]entities.AuthUser, error)
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

func (s *AuthUserSQLRepository) GetUser(userName string) (user []entities.AuthUser, err error) {
	result := s.db.Where("user_name = ? AND isDeleted = ?", userName, false).First(&user)
	if result.Error != nil {
		err = result.Error
		utils.Logger().Debugf("AuthUserSQLRepository: error in get user by user name with error: %s", err.Error())
	}
	return
}
