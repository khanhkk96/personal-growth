package model

import (
	"time"
)

type Project struct {
	BaseModel
	Name        string    `gorm:"type:string; not null; unique"`
	Type        string    `gorm:"type:string; not null"`
	Summary     string    `gorm:"type:string; not null"`
	Description string    `gorm:"type:text"`
	Stack       string    `gorm:"type:string"`
	StartAt     time.Time `gorm:"type:time"`
	EndAt       time.Time `gorm:"type:time"`
}
