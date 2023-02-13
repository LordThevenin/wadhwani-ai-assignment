package services

import (
	"user-service/dto"
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
	userEntities := dto.UserModelsToUserEntities(users)
	// Upsert users list
	err = s.userRepository.UpsertUsers(userEntities)
	if err != nil {
		// Log repository error for upsert
	}
	return
}

func (s *UserService) GetUser() {

}
