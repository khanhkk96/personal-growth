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

// ProvideIssueRepository creates a new Issue repository.
func ProvideIssueRepository(db *gorm.DB) repositories.IssueRepository {
	return repositories.NewIssueRepository(db)
}

// ProvideIssueService creates a new Issue service.
func ProvideIssueService(repo repositories.IssueRepository, validate *validator.Validate) service_interfaces.IssueService {
	return services.NewIssueServiceImpl(repo, validate)
}

// ProvideIssueController creates a new Issue controller.
func ProvideIssueController(service service_interfaces.IssueService) *controllers.IssueController {
	return controllers.NewIssueController(service)
}

// ProvideIssueRouter creates a new Issue router.
func ProvideIssueRouter(controller *controllers.IssueController, db *gorm.DB) *fiber.App {
	return routers.NewIssueRouter(controller, db)
}

// InitIssue initializes the Issue module using Wire.
func InitIssue(db *gorm.DB) *fiber.App {
	wire.Build(
		ProvideValidator,
		ProvideIssueRepository,
		ProvideIssueService,
		ProvideIssueController,
		ProvideIssueRouter,
	)
	return nil
}
