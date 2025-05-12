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
// @Success 	200 {object} responses.ProjectResponse
// @Router 		/api/project [post]
func (controller *ProjectController) AddNewProject(ctx *fiber.Ctx) error {
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

// @Summary 	Update project
// @Description Update project
// @Tags 		Project
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param		id path string true "Project ID"
// @Param 		project body requests.CreateOrUpdateProjectRequest true "Project Info"
// @Success 	200 {object} responses.ProjectResponse
// @Router 		/api/project/{id} [put]
func (controller *ProjectController) UpdateProject(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	id := ctx.Params("id")

	request := requests.CreateOrUpdateProjectRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	project, cerr := controller.service.Update(id, request, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Update the project successfully",
		Data:    project,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Delete project
// @Description Delete project
// @Tags 		Project
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param		id path string true "Project ID"
// @Success 	200 {object} responses.ProjectResponse
// @Router 		/api/project/{id} [delete]
func (controller *ProjectController) DeleteProject(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	id := ctx.Params("id")

	project, cerr := controller.service.Delete(id, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Delete the project successfully",
		Data:    project,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Get project details
// @Description Get the details of the project
// @Tags 		Project
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param		id path string true "Project ID"
// @Success 	200 {object} responses.ProjectResponse
// @Router 		/api/project/{id} [get]
func (controller *ProjectController) GetProjectDetail(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	project, cerr := controller.service.Detail(id)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Delete the project successfully",
		Data:    project,
	}

	return ctx.Status(200).JSON(response)
}
