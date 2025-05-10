package controllers

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/helpers"
	"personal-growth/models"
	"personal-growth/services"

	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	service services.ProjectService
}

func NewProjectController(service services.ProjectService) *ProjectController {
	return &ProjectController{
		service: service,
	}
}

// @Summary 	Add new project
// @Description Add new project
// @Tags 		Project
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		project body requests.CreateOrUpdateProjectRequest true "Project Info"
// @Router 		/api/project [post]
func (controller *ProjectController) AddNewProject(ctx *fiber.Ctx) error {
	println("Add new project")
	user := ctx.Locals("user").(*models.User)

	request := requests.CreateOrUpdateProjectRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	project, cerr := controller.service.Add(request, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Add new project successfully",
		Data:    project,
	}

	return ctx.Status(200).JSON(response)
}
