package services

import (
	"cloud.google.com/go/translate"
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"user-service/config"
	"user-service/utils"
)

type ITranslationService interface {
	TranslateText(ctx *gin.Context, input []string, lang language.Tag) ([]translate.Translation, error)
}

type GoogleTranslationService struct {
	translateClient *translate.Client
}

func InitGoogleTranslationService() *GoogleTranslationService {
	ctx := context.Background()
	cfg := config.GetConfig()
	gts := new(GoogleTranslationService)
	translateClient, err := translate.NewClient(ctx, option.WithCredentialsJSON(cfg.TranslationConfiguration))
	if err != nil {
		utils.Logger().Errorf("InitGoogleTranslationService: failed to initialize google translate service with error: %s", err.Error())
		panic(err.Error())
	}
	gts.translateClient = translateClient
	return gts
}

func (s *GoogleTranslationService) TranslateText(ctx *gin.Context, input []string, lang language.Tag) ([]translate.Translation, error) {
	resp, err := s.translateClient.Translate(ctx, input, lang, nil)
	if err != nil {
		// Log error in translation
		utils.Logger().Errorf("GoogleTranslationService: failed to translate text to target language with error: %s", err.Error())
		return nil, err
	}
	return resp, nil
}
