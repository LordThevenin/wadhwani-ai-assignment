package utils

import (
	"encoding/csv"
	"mime/multipart"
	"strconv"
	"user-service/models"
)

func ParseUserUploadFile(uploadData models.UserFileUpload) ([]models.User, error) {
	csvData, err := fetchCSVData(uploadData.File)
	if err != nil {
		// Log error in fetching CSV data
		return nil, err
	}
	users := getUsersFromCSV(csvData)
	return users, nil
}

func getUsersFromCSV(csvData [][]string) (users []models.User) {
	for i, csvLine := range csvData {
		if i > 0 {
			phoneNumber, _ := strconv.Atoi(csvLine[0])
			user := models.User{
				PhoneNumber: int64(phoneNumber),
				Name:        csvLine[1],
				State:       csvLine[2],
				District:    csvLine[3],
				Village:     csvLine[4],
			}
			users = append(users, user)
		}
	}
	return
}

func fetchCSVData(fileHeader *multipart.FileHeader) ([][]string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		// Log error in opening file
		return nil, err
	}
	csvData, err := csv.NewReader(file).ReadAll()
	if err != nil {
		// Log error in reading csv data
		return nil, err
	}
	return csvData, nil
}
