package entities

import "time"

// Base entity which is part of every table model as manadatory.
type Base struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`
	CreatedBy int64     `gorm:"column:created_by;default:0"`
	UpdatedBy int64     `gorm:"column:updated_by;default:0"`
	Version   int       `gorm:"column:version"`
	IsDeleted bool      `gorm:"column:isDeleted"`
}

type Location struct {
	State    string `gorm:"column:state"`
	District string `gorm:"column:district"`
	Village  string `gorm:"column:village"`
}
