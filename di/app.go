package injections

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MountAppRouters(app *fiber.App, db *gorm.DB) {
	// validate := validator.New()

	//auth
	authModule := InitAuth(db)
	app.Mount("/api", authModule)
}
