package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	Model         gorm.Model `gorm:"embedded"`
	Id            uuid.UUID  `gorm:"type:uuid; default:gen_random_uuid(); primaryKey"`
	Amount        float64    `gorm:"type:money"`
	TransactionId string     `gorm:"type:character"`
	PayBy         string     `gorm:"type:character; default:'momo'"`
}
