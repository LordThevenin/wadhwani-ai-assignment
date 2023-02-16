package entities

type AuthUser struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement"`
	username string `gorm:"column:user_name"`
	password string `gorm:"column:password"`
	Base
}
