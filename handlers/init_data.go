package handlers

import (
	"log"
	"personal-growth/common/enums"
	"personal-growth/db/models"

	"gorm.io/gorm"
)

func InitAdmin(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		users := []models.User{
			{
				Username: "admin",
				Password: "123456",
				FullName: "Admin",
				Email:    "admin@gmail.com",
				IsActive: true,
				Role:     enums.ADMIN,
			},
		}

		if err := db.Create(&users).Error; err != nil {
			log.Fatal("Seed error:", err)
		}
		log.Println("Seeded admin account successfully.")
	} else {
		log.Println("Admin already exist, skipping seed.")
	}
}
