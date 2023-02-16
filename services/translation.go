package services

import (
	"cloud.google.com/go/translate"
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"user-service/config"
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
		panic(err.Error())
	}
	gts.translateClient = translateClient
	return gts
}

func (s *GoogleTranslationService) TranslateText(ctx *gin.Context, input []string, lang language.Tag) ([]translate.Translation, error) {
	resp, err := s.translateClient.Translate(ctx, input, lang, nil)
	if err != nil {
		// Log error in translation
		return nil, err
	}
	return resp, nil
}
