package services

import (
	"fmt"
	"golang.org/x/text/language"
	"user-service/dto"
	"user-service/models"
	"user-service/repositories"
	"user-service/utils"
)

type IUserService interface {
	UploadUsers([]models.User) error
	GetUser(int64) (models.User, error)
	GetUserFromCache(int64, language.Tag) (models.User, bool)
	SetUserInCache(int64, language.Tag, models.User)
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
	userEntities := dto.UserModelsToUserEntities(users, 1)
	// Upsert users list
	err = s.userRepository.UpsertUsers(userEntities)
	if err != nil {
		// Log repository error for upsert
		utils.Logger().Errorf("UserService: user upsert failure")
	}
	return
}

func (s *UserService) GetUser(userId int64) (user models.User, err error) {
	// Fetch from repository
	userEntity, err := s.userRepository.GetUser(userId)
	if err != nil {
		// Log repository err for fetch
		utils.Logger().Errorf("UserService: failed to fetch user with given id")
	}
	if len(userEntity) > 0 {
		user = dto.UserEntityToUserModel(userEntity[0])
	} else {
		// Log error for no user found
		err = fmt.Errorf("no user found")
	}
	return
}

func (s *UserService) GetUserFromCache(userId int64, lang language.Tag) (user models.User, hit bool) {
	hit = false
	key := fmt.Sprintf("%d_%s", userId, lang.String())
	user, err := s.userCacheRepository.Get(key)
	if err != nil {
		// Log cache miss for the key
		utils.Logger().Infof("UserService: failed to fetch user from cache")
		return
	}
	hit = true
	return
}

func (s *UserService) SetUserInCache(userId int64, lang language.Tag, user models.User) {
	key := fmt.Sprintf("%d_%s", userId, lang.String())
	s.userCacheRepository.Set(key, user)
}
