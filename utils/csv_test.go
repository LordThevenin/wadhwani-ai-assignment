package utils

import (
	"mime/multipart"
	"reflect"
	"testing"
	"user-service/models"
)

func TestParseUserUploadFile(t *testing.T) {
	type args struct {
		uploadData models.UserFileUpload
	}
	tests := []struct {
		name    string
		args    args
		want    []models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUserUploadFile(tt.args.uploadData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUserUploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUserUploadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchCSVData(t *testing.T) {
	type args struct {
		fileHeader *multipart.FileHeader
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchCSVData(tt.args.fileHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchCSVData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchCSVData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUsersFromCSV(t *testing.T) {
	type args struct {
		csvData [][]string
	}
	tests := []struct {
		name      string
		args      args
		wantUsers []models.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUsers := getUsersFromCSV(tt.args.csvData); !reflect.DeepEqual(gotUsers, tt.wantUsers) {
				t.Errorf("getUsersFromCSV() = %v, want %v", gotUsers, tt.wantUsers)
			}
		})
	}
}
