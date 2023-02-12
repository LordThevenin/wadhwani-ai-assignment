package services

type ITranslationService interface {
}

type GoogleTranslationService struct {
}

func InitGoogleTranslationService() *GoogleTranslationService {
	gts := new(GoogleTranslationService)
	return gts
}
