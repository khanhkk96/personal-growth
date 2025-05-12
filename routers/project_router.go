package routers

import (
	"personal-growth/controllers"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewProjectRouter(controller *controllers.ProjectController, db *gorm.DB) *fiber.App {
	projectRouter := fiber.New()

	projectRouter.Group("/", middlewares.Authenticate(), middlewares.GetProfileHandler(db)).Route("/project",
		func(router fiber.Router) {
			router.Post("/", controller.AddNewProject)
			router.Put("/:id", controller.UpdateProject)
			router.Delete("/:id", controller.DeleteProject)
			router.Get("/:id", controller.GetProjectDetail)
			router.Get("/", controller.GetProjects)
		})

	return projectRouter
}
