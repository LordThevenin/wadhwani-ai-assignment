package services

import (
	"cloud.google.com/go/translate"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"reflect"
	"testing"
)

func TestGoogleTranslationService_TranslateText(t *testing.T) {
	type fields struct {
		translateClient *translate.Client
	}
	type args struct {
		ctx   *gin.Context
		input []string
		lang  language.Tag
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []translate.Translation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GoogleTranslationService{
				translateClient: tt.fields.translateClient,
			}
			got, err := s.TranslateText(tt.args.ctx, tt.args.input, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("TranslateText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TranslateText() got = %v, want %v", got, tt.want)
			}
		})
	}
}
