package models

import (
	"database/sql"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model        gorm.Model     `gorm:"embedded"`
	Id           uuid.UUID      `gorm:"type:uuid; default:gen_random_uuid(); primarykey"`
	Username     string         `gorm:"type:string; not null; unique"`
	Password     string         `gorm:"type:string; not null;"`
	FullName     string         `gorm:"type:string; not null"`
	Email        string         `gorm:"type:string; not null; unique"`
	Phone        string         `gorm:"type:string; unique"`
	Otp          sql.NullString `gorm:"type:string"`
	OtpExpiredAt sql.NullTime   `gorm:"type:time"`
	OtpCounter   int            `gorm:"type:int; default: 0"`
	IsActive     bool           `gorm:"type:boolean; default:false"`
	Type         string         `gorm:"type:string; default:'user'"` // user, admin
	Avatar       sql.NullString `gorm:"type:string"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)

	return err
}

func (u User) CompareHashAndPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
