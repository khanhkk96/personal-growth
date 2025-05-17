package injections

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MountAppRouters(app *fiber.App, db *gorm.DB) {
	//auth
	authModule := InitAuth(db)
	projectModule := InitProject(db)
	issueModule := InitIssue(db)
	paymentModule := InitPayment(db)

	app.Mount("/api", authModule).Mount("/api", projectModule).Mount("/api", issueModule).Mount("/api", paymentModule)
}
