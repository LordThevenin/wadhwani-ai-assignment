package models

type User struct {
	PhoneNumber int64  `json:"phoneNumber"`
	Name        string `json:"name"`
	State       string `json:"state"`
	District    string `json:"district"`
	Village     string `json:"village"`
}
