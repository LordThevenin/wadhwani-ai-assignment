package services

import "user-service/repositories"

type IUserService interface {
	UploadUsers()
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

func (s *UserService) UploadUsers() {

}

func (s *UserService) GetUser() {

}
