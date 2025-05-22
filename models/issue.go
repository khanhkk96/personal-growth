package models

import (
	"database/sql"
	"personal-growth/common/enums"
)

type Issue struct {
	BaseModel
	Name        string              `gorm:"type:string; not null; unique"`
	Description sql.NullString      `gorm:"type:text"`
	Files       sql.NullString      `gorm:"type:string"`
	ProjectId   sql.NullString      `gorm:"type:uuid"`
	Project     Project             `gorm:"foreignKey:ProjectId"`
	Status      enums.IssueStatus   `gorm:"type:string; not null; default='pending'"`
	Priority    enums.IssuePriority `gorm:"type:string; not null"`
	IssuedAt    sql.NullTime        `gorm:"type:time"`
	NeedToSolve sql.NullInt32       `gorm:"type:float"`
	References  sql.NullString      `gorm:"type:string"`
}
