package controllers

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/db/entities"
	"personal-growth/helpers"
	service_interfaces "personal-growth/services/interfaces"

	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	service service_interfaces.ProjectService
}

func NewProjectController(service service_interfaces.ProjectService) *ProjectController {
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
	user := ctx.Locals("user").(*entities.User)

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
	user := ctx.Locals("user").(*entities.User)
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
	user := ctx.Locals("user").(*entities.User)
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

// @Summary 	Get project list
// @Description Get the list of the project
// @Tags 		Project
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} responses.ProjectPageResponse
// @Param 		filters query requests.ProjectFilters false "Project Filter"
// @Router 		/api/project [get]
func (controller *ProjectController) GetProjects(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*entities.User)

	var filters requests.ProjectFilters
	if err := ctx.QueryParser(&filters); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// fmt.Printf("query::::::: %v\n", filters)
	data := controller.service.List(filters, user)

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Get list of project successfully",
		Data:    data,
	}

	return ctx.Status(200).JSON(response)
}
