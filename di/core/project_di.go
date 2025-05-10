package injections

import (
	"personal-growth/controllers"
	"personal-growth/repositories"
	"personal-growth/routers"
	"personal-growth/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProvideProjectRepository creates a new Project repository.
func ProvideProjectRepository(db *gorm.DB) repositories.ProjectRepository {
	return repositories.NewProjectRepository(db)
}

// ProvideProjectService creates a new Project service.
func ProvideProjectService(repo repositories.ProjectRepository, validate *validator.Validate) services.ProjectService {
	return services.NewProjectServiceImpl(repo, validate)
}

// ProvideProjectController creates a new Project controller.
func ProvideProjectController(service services.ProjectService) *controllers.ProjectController {
	return controllers.NewProjectController(service)
}

// ProvideProjectRouter creates a new Project router.
func ProvideProjectRouter(controller *controllers.ProjectController, db *gorm.DB) *fiber.App {
	return routers.NewProjectRouter(controller, db)
}

// InitProject initializes the Project module using Wire.
func InitProject(db *gorm.DB) *fiber.App {
	wire.Build(
		ProvideValidator,
		ProvideProjectRepository,
		ProvideProjectService,
		ProvideProjectController,
		ProvideProjectRouter,
	)
	return nil
}
