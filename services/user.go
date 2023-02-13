package services

import (
	"user-service/entities"
	"user-service/models"
	"user-service/repositories"
)

type IUserService interface {
	UploadUsers([]models.User) error
	GetUser()
}

type UserService struct {
	userRepository      repositories.IUserRepository
	userCacheRepository repositories.IUserCacheRepository
}

func InitUserService() *UserService {
	us := new(UserService)
	us.userRepository = repositories.InitUserSQLRepository()
	us.userCacheRepository = repositories.InitRedisUserCacheRepository()
	return us
}

func (s *UserService) UploadUsers(users []models.User) (err error) {
	// Transform users list to entities
	// Upsert users list
	err = s.userRepository.UpsertUsers([]entities.User{})
	if err != nil {
		// Log repository error for upsert
	}
	return
}

func (s *UserService) GetUser() {

}
