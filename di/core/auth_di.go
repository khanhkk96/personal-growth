package injections

import (
	"personal-growth/common/enums"
	"personal-growth/controllers"
	"personal-growth/repositories"
	"personal-growth/routers"
	"personal-growth/services"
	service_interfaces "personal-growth/services/interfaces"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// func InitAuth(db *gorm.DB, validate *validator.Validate) *fiber.App {
// 	// Init repository
// 	userRepository := repositories.NewUserRepository(db)

// 	// Init service
// 	authService := services.NewAuthServiceImpl(userRepository, validate)

// 	// Init controller
// 	authController := controllers.NewAuthController(authService)

// 	// Init routes
// 	return routers.NewAuthRouter(authController, db)
// }

// ProvideValidator creates a new validator instance.
func ProvideValidator() *validator.Validate {
	v := validator.New()
	enums.RegisterCustomValidations(v)
	return v
}

// ProvideUserRepository creates a new User repository.
func ProvideUserRepository(db *gorm.DB) repositories.UserRepository {
	return repositories.NewUserRepository(db)
}

// ProvideAuthService creates a new Auth service.
func ProvideAuthService(repo repositories.UserRepository, validate *validator.Validate) service_interfaces.AuthService {
	return services.NewAuthServiceImpl(repo, validate)
}

// ProvideAuthController creates a new Auth controller.
func ProvideAuthController(service service_interfaces.AuthService) *controllers.AuthController {
	return controllers.NewAuthController(service)
}

// ProvideAuthRouter creates a new Auth router.
func ProvideAuthRouter(controller *controllers.AuthController, db *gorm.DB) *fiber.App {
	return routers.NewAuthRouter(controller, db)
}

// InitAuth initializes the Auth module using Wire.
func InitAuth(db *gorm.DB) *fiber.App {
	wire.Build(
		ProvideValidator,
		ProvideUserRepository,
		ProvideAuthService,
		ProvideAuthController,
		ProvideAuthRouter,
	)
	return nil
}
