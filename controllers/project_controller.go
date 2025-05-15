package controllers

import (
	"personal-growth/common/enums"
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/helpers"
	"personal-growth/models"
	service_interfaces "personal-growth/services/interfaces"
	"strconv"

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

// @Summary 	Get project list
// @Description Get the list of the project
// @Tags 		Project
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} responses.ProjectPageResponse
// @Param 		page query int false "Page number"
// @Param 		limit query int false "Page size"
// @Param 		q query string false "Search by name/stack"
// @Param 		status query string false "Status" Enums(planning, postpone, ongoing, finished)
// @Param 		type query string false "Type" Enums(web, desktop_app, mobile_app, library)
// @Router 		/api/project [get]
func (controller *ProjectController) GetProjects(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)

	var filter requests.ProjectFilters
	query := ctx.Query("q")
	filter.Query = &query
	filter.Page, _ = strconv.Atoi(ctx.Query("page", "1"))
	filter.Limit, _ = strconv.Atoi(ctx.Query("limit", "10"))
	status := enums.ProjectStatus(ctx.Query("status"))
	filter.Status = &status
	projectType := enums.ProjectType(ctx.Query("type"))
	filter.Type = &projectType

	data, cerr := controller.service.List(filter, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Get list of project successfully",
		Data:    data,
	}

	return ctx.Status(200).JSON(response)
}
