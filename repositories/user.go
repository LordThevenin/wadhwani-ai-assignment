package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"user-service/db"
	"user-service/entities"
)

type IUserRepository interface {
	UpsertUsers(users []entities.User) error
	GetUser(id int64) ([]entities.User, error)
}

type UserSQLRepository struct {
	db *gorm.DB
}

func InitUserSQLRepository() *UserSQLRepository {
	ur := new(UserSQLRepository)
	ur.db = db.GetSqlDB()
	return ur
}

// UpsertUsers upserts user values based on phone number
func (s *UserSQLRepository) UpsertUsers(users []entities.User) (err error) {
	err = s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "phone_number"}},
		DoUpdates: clause.AssignmentColumns([]string{"state", "district", "village", "version", "updated_by"}),
	}).Create(&users).Error
	return
}

// GetUser gets valid user value based on id
func (s *UserSQLRepository) GetUser(id int64) (user []entities.User, err error) {
	result := s.db.Where("id = ? AND isDeleted = ?", id, false).First(&user)
	if result.Error != nil {
		err = result.Error
	}
	return
}
