package models

import (
	"time"
)

type Issue struct {
	BaseModel
	Name        string    `gorm:"type:string; not null"`
	Description string    `gorm:"type:text"`
	Image       string    `gorm:"type:string"`
	ProjectId   string    `gorm:"type:uuid"`
	Project     Project   `gorm:"foreignKey:ProjectId"`
	Status      string    `gorm:"type:string; not null"`
	Priority    string    `gorm:"type:string; not null"`
	IssuedAt    time.Time `gorm:"type:time"`
	SolvedTime  float64   `gorm:"type:float"`
	References  string    `gorm:"type:string"`
}
