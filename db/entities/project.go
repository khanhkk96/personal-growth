package entities

import (
	"database/sql"
	"personal-growth/common/enums"
)

type Project struct {
	BaseModel
	Name        string              `gorm:"type:string; not null; unique"`
	Type        enums.ProjectType   `gorm:"type:string; not null"`
	Summary     string              `gorm:"type:string; not null"`
	Stack       string              `gorm:"type:string; not null"`
	Description sql.NullString      `gorm:"type:text"`
	StartAt     sql.NullTime        `gorm:"type:time"`
	EndAt       sql.NullTime        `gorm:"type:time"`
	Status      enums.ProjectStatus `gorm:"type:string; not null; default:'ongoing'"`
}
