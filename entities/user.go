package entities

type User struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	PhoneNumber int64  `gorm:"column:phone_number;unique"`
	Name        string `gorm:"column:name"`
	Location
	Base
}
