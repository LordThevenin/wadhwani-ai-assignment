package entities

type AuthUser struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Username string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role"`
	Base
}
