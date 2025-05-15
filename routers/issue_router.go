package routers

import (
	"personal-growth/controllers"
	"personal-growth/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewIssueRouter(controller *controllers.IssueController, db *gorm.DB) *fiber.App {
	issueRouter := fiber.New()

	issueRouter.Group("/", middlewares.Authenticate(), middlewares.GetProfileHandler(db)).Route("/issue",
		func(router fiber.Router) {
			router.Post("/", controller.AddNewIssue)
			router.Put("/:id", controller.UpdateIssue)
			router.Delete("/:id", controller.DeleteIssue)
			router.Get("/:id", controller.GetIssueDetail)
			router.Get("/", controller.GetIssues)
		})

	return issueRouter
}
