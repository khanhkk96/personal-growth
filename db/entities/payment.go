package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	Model             gorm.Model `gorm:"embedded"`
	Id                uuid.UUID  `gorm:"type:uuid; default:gen_random_uuid(); primaryKey"`
	PayBy             string     `gorm:"type:string; default:'momo'"`
	Amount            int64      `gorm:"type:money; not null"`
	BankCode          string     `gorm:"type:string"`
	TransactionNo     string     `gorm:"type:string; not null"`
	BankTransactionNo string     `gorm:"type:string"`
	PayDate           time.Time  `gorm:"type:timestamp"`
	TransactionStatus string     `gorm:"type:string; default:'success'"`
	OrderInfo         string     `gorm:"type:string"`
	TxnRef            string     `gorm:"type:string; not null"`
}
