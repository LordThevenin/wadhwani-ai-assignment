package entities

import "time"

// Base entity which is part of every table model as manadatory.
type Base struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
	CreatedBy string    `gorm:"column:created_by;default:SYSTEM"`
	UpdatedBy string    `gorm:"column:updated_by;default:SYSTEM"`
	Version   int       `gorm:"column:version"`
	IsDeleted bool      `gorm:"column:isDeleted"`
}

type Location struct {
	State    string `gorm:"column:state"`
	District string `gorm:"column:district"`
	Village  string `gorm:"column:village"`
}
