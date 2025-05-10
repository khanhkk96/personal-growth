package injections

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitAppRouters(app *fiber.App, db *gorm.DB) {
	validate := validator.New()
	//auth
	authModule := InitAuth(db, validate)
	app.Mount("/api", authModule)

}
