package facades

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"user-service/dto"
	"user-service/models"
	"user-service/services"
	"user-service/utils"
)

type IUserFacade interface {
	UploadUsers(models.UserFileUpload) error
	GetUser(*gin.Context, int64, language.Tag) (models.User, error)
}

type UserFacade struct {
	userService        services.IUserService
	translationService services.ITranslationService
}

func InitUserFacade() *UserFacade {
	uf := new(UserFacade)
	uf.userService = services.InitUserService()
	uf.translationService = services.InitGoogleTranslationService()
	return uf
}

func (f *UserFacade) UploadUsers(uploadData models.UserFileUpload) (err error) {
	// Transform upload form data to list of user models
	users, err := utils.ParseUserUploadFile(uploadData)
	if err != nil {
		// Log failure to parse upload file
		return
	}
	// Give user service to handle list of user models
	err = f.userService.UploadUsers(users)
	if err != nil {
		// Log failure to upload users
		return
	}
	return
}

func (f *UserFacade) GetUser(ctx *gin.Context, userId int64, lang language.Tag) (user models.User, err error) {
	// Fetch user from service
	user, err = f.userService.GetUser(userId)
	if err != nil {
		// Log failure to get user
		return
	}
	// Translate to requested language
	user, err = f.translateUserModelToTargetLanguage(ctx, user, lang)
	if err != nil {
		// Log failure to translate user
		return
	}
	return
}

func (f *UserFacade) translateUserModelToTargetLanguage(ctx *gin.Context, user models.User, lang language.Tag) (translatedUser models.User, err error) {
	// Convert user model to list of strings
	stringList := dto.UserModelToStringList(user)
	// Translate list of strings
	translatedList, err := f.translationService.TranslateText(ctx, stringList, lang)
	if err != nil {
		// Log error in translating user values
		return
	}
	// Convert list of strings to model
	translatedUser = dto.TranslatedListToUserModel(translatedList)
	return
}
