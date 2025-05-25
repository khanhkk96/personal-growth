package injections

import (
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

// ProvidePaymentRepository creates a new Payment repository.
func ProvidePaymentRepository(db *gorm.DB) repositories.PaymentRepository {
	return repositories.NewPaymentRepository(db)
}

// ProvidePaymentService creates a new Payment service.
func ProvidePaymentService(repo repositories.PaymentRepository, validate *validator.Validate) service_interfaces.PaymentService {
	return services.NewPaymentServiceImpl(repo, validate)
}

// ProvidePaymentController creates a new Payment controller.
func ProvidePaymentController(service service_interfaces.PaymentService) *controllers.PaymentController {
	return controllers.NewPaymentController(service)
}

// ProvidePaymentRouter creates a new Payment router.
func ProvidePaymentRouter(controller *controllers.PaymentController, db *gorm.DB) *fiber.App {
	return routers.NewPaymentRouter(controller, db)
}

// InitPayment initializes the Payment module using Wire.
func InitPayment(db *gorm.DB) *fiber.App {
	wire.Build(
		ProvideValidator,
		ProvidePaymentRepository,
		ProvidePaymentService,
		ProvidePaymentController,
		ProvidePaymentRouter,
	)
	return nil
}
