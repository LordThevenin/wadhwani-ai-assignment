package mocks

import (
	"cloud.google.com/go/translate"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type GoogleTranslationServiceMock struct {
}

func (s *GoogleTranslationServiceMock) TranslateText(ctx *gin.Context, input []string, lang language.Tag) ([]translate.Translation, error) {
	switch lang {
	case language.Hindi:
		{
			return []translate.Translation{
				{
					"73572",
					language.Hindi,
					"",
				},
				{
					"testTranslatedName2",
					language.Hindi,
					"",
				},
				{
					"testTranslatedState2",
					language.Hindi,
					"",
				},
				{
					"testTranslatedVillage2",
					language.Hindi,
					"",
				},
				{
					"testTranslatedDistrict2",
					language.Hindi,
					"",
				},
			}, nil
		}

	}
	return nil, fmt.Errorf("error")
}
