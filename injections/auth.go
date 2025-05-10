package injections

import (
	"personal-growth/controllers"
	"personal-growth/models"
	"personal-growth/repository"
	"personal-growth/routers"
	"personal-growth/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitAuth(db *gorm.DB, validate *validator.Validate) *fiber.App {
	// Init repository
	userRepository := repository.NewBaseRepository[models.User](db)

	// Init service
	authService := services.NewAuthServiceImpl(userRepository, validate)

	// Init controller
	authController := controllers.NewAuthController(authService)

	// Init routes
	return routers.NewAuthRouter(authController, db)
}
