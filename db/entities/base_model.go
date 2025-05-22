package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Model       gorm.Model `gorm:"embedded"`
	Id          uuid.UUID  `gorm:"type:uuid; default:gen_random_uuid(); primaryKey"`
	CreatedById uuid.UUID  `gorm:"type:uuid"`
	CreatedBy   User       `gorm:"foreignKey:CreatedById"`
}
