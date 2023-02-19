package mocks

import (
	"cloud.google.com/go/translate"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type GoogleTranslationServiceMock struct {
}

func (s *GoogleTranslationServiceMock) TranslateText(ctx *gin.Context, input []string, lang language.Tag) ([]translate.Translation, error) {
	return nil, nil
}
