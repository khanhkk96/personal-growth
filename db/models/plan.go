package models

import (
	"time"
)

type Plan struct {
	BaseModel
	Name              string    `gorm:"type:string; not null; unique"`
	Description       string    `gorm:"type:text"`
	Objective         string    `gorm:"type:string; not null"`
	ExpectedStartTime time.Time `gorm:"type:time; not null"`
	ExpectedEndTime   time.Time `gorm:"type:time"`
	ActualStartTime   time.Time `gorm:"type:time"`
	ActualEndTime     time.Time `gorm:"type:time"`
	Progress          int       `gorm:"type:int; default:0"`
	Status            string    `gorm:"type:string; not null"`
}
