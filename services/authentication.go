package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"user-service/dto"
	"user-service/entities"
	"user-service/models"
	"user-service/repositories"
	"user-service/utils"
)

type IAuthenticationService interface {
	Register(user models.AuthUser) error
	Login(user models.AuthUser) (string, error)
}

type AuthenticationService struct {
	authRepository repositories.IAuthUserRepository
}

func InitAuthorizationService() *AuthenticationService {
	as := new(AuthenticationService)
	as.authRepository = repositories.InitAuthUserSQLRepository()
	return as
}

func (s *AuthenticationService) Register(user models.AuthUser) (err error) {
	userEntity := dto.AuthUserModelToEntity(user, 1)
	err = s.authRepository.AddUser(userEntity)
	if err != nil {
		// Log error in adding auth user
	}
	return
}

func (s *AuthenticationService) Login(user models.AuthUser) (jwt string, err error) {
	userEntity, err := s.getAuthUser(user, err)
	// Validate user
	err = s.validatePassword(user, err, userEntity)
	if err != nil {
		// Log password match failed
		err = fmt.Errorf("incorrect password")
	}
	jwt, err = utils.GenerateToken(userEntity[0])
	if err != nil {
		// Log error in generating jwt token
		err = fmt.Errorf("incorrect password")
	}
	return
}

func (s *AuthenticationService) validatePassword(user models.AuthUser, err error, userEntity []entities.AuthUser) error {
	err = bcrypt.CompareHashAndPassword([]byte(userEntity[0].Password), []byte(user.Password))
	return err
}

func (s *AuthenticationService) getAuthUser(user models.AuthUser, err error) ([]entities.AuthUser, error) {
	userEntity, err := s.authRepository.GetUser(user.UserName)
	if err != nil {
		// Log error finding user
	}
	if userEntity == nil || len(userEntity) == 0 {
		// Log no user found
		err = fmt.Errorf("user not found")
	}
	return userEntity, err
}
