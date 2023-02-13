package models

import "mime/multipart"

type User struct {
	PhoneNumber int64  `json:"phoneNumber"`
	Name        string `json:"name"`
	State       string `json:"state"`
	District    string `json:"district"`
	Village     string `json:"village"`
}

type UserFileUpload struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
