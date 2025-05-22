package controllers

import (
	"fmt"
	"personal-growth/common/enums"
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/helpers"
	"personal-growth/models"
	service_interfaces "personal-growth/services/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type IssueController struct {
	service service_interfaces.IssueService
}

func NewIssueController(service service_interfaces.IssueService) *IssueController {
	return &IssueController{
		service: service,
	}
}

// @Summary 	Add new issue
// @Description Add new issue
// @Tags 		Issue
// @Security  	BearerAuth
// @Produce 	json
// @Accept 		multipart/form-data
// @Param 		issue formData requests.CreateOrUpdateIssueRequest false "Issue Info"
// @Param 		files formData []file false "File to upload"
// @Success 	200 {object} responses.IssueResponse
// @Router 		/api/issue [post]
func (controller *IssueController) AddNewIssue(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	files := ctx.Locals("files").([]string)

	request := requests.CreateOrUpdateIssueRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)
	fmt.Printf("issue::%v", request)

	issue, cerr := controller.service.Add(request, files, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Add new issue successfully",
		Data:    issue,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Update issue
// @Description Update issue
// @Tags 		Issue
// @Security  	BearerAuth
// @Produce 	json
// @Accept 		multipart/form-data
// @Param 		issue formData requests.CreateOrUpdateIssueRequest true "Issue Info"
// @Param 		files formData []file true "File to upload"
// @Param		id path string true "Issue ID"
// @Success 	200 {object} responses.IssueResponse
// @Router 		/api/issue/{id} [put]
func (controller *IssueController) UpdateIssue(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	files := ctx.Locals("files").([]string)
	id := ctx.Params("id")

	request := requests.CreateOrUpdateIssueRequest{}
	err := ctx.BodyParser(&request)
	helpers.ErrorPanic(err)

	issue, cerr := controller.service.Update(id, request, files, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Update the issue successfully",
		Data:    issue,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Delete issue
// @Description Delete issue
// @Tags 		Issue
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param		id path string true "Issue ID"
// @Success 	200 {object} responses.IssueResponse
// @Router 		/api/issue/{id} [delete]
func (controller *IssueController) DeleteIssue(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)
	id := ctx.Params("id")

	issue, cerr := controller.service.Delete(id, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Delete the issue successfully",
		Data:    issue,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Get issue details
// @Description Get the details of the issue
// @Tags 		Issue
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param		id path string true "Issue ID"
// @Success 	200 {object} responses.IssueResponse
// @Router 		/api/issue/{id} [get]
func (controller *IssueController) GetIssueDetail(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	issue, cerr := controller.service.Detail(id)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Delete the issue successfully",
		Data:    issue,
	}

	return ctx.Status(200).JSON(response)
}

// @Summary 	Get issue list
// @Description Get the list of the issue
// @Tags 		Issue
// @Security  	BearerAuth
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} responses.IssuePageResponse
// @Param 		page query int false "Page number"
// @Param 		limit query int false "Page size"
// @Param 		q query string false "Search by name/stack"
// @Param 		status query string false "Status" Enums(pending, processing, resolved)
// @Param 		priority query string false "Priority" Enums(low, medium, high)
// @Router 		/api/issue [get]
func (controller *IssueController) GetIssues(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.User)

	var filter requests.IssueFilters
	query := ctx.Query("q")
	filter.Query = &query
	filter.Page, _ = strconv.Atoi(ctx.Query("page", "1"))
	filter.Limit, _ = strconv.Atoi(ctx.Query("limit", "10"))
	status := enums.IssueStatus(ctx.Query("status"))
	filter.Status = &status
	issuePriority := enums.IssuePriority(ctx.Query("priority"))
	filter.Priority = &issuePriority

	data, cerr := controller.service.List(filter, user)
	if cerr != nil {
		return ctx.Status(cerr.Code).JSON(cerr.Message)
	}

	response := responses.Response{
		Code:    200,
		Status:  "ok",
		Message: "Get list of issue successfully",
		Data:    data,
	}

	return ctx.Status(200).JSON(response)
}
