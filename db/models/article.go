package models

import (
	"time"
)

type Article struct {
	BaseModel
	Title       string    `gorm:"type:string; not null"`
	Summary     string    `gorm:"type:string; not null"`
	Content     string    `gorm:"type:text; not null"`
	Quote       string    `gorm:"type:string"`
	ReadTurns   int       `gorm:"type:int; default:0"`
	IsPublished bool      `gorm:"type:boolean; default:true"`
	IsFeatured  bool      `gorm:"type:boolean; default:false"`
	PublishedAt time.Time `gorm:"type:time"`
}
