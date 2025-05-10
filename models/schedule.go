package models

import "time"

type Schedule struct {
	BaseModel
	Name        string    `gorm:"type:string; not null"`
	Description string    `gorm:"type:text"`
	Note        string    `gorm:"type:string"`
	StartTime   time.Time `gorm:"type:time; not null"`
	EndTime     time.Time `gorm:"type:time"`
	IsRequired  bool      `gorm:"type:boolean; default:true"`
	PlanId      string    `gorm:"type:uuid"`
	Plan        Plan      `gorm:"foreingKey:PlanId"`
	Status      string    `gorm:"type:string; default:pending"`
}
