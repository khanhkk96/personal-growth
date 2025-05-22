package services

import (
	"database/sql"
	"fmt"
	"log"
	"personal-growth/common/enums"
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/db/entities"
	"personal-growth/repositories"
	service_interfaces "personal-growth/services/interfaces"
	"personal-growth/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type IssueServiceImpl struct {
	repository repositories.IssueRepository
	validate   *validator.Validate
}

func NewIssueServiceImpl(repository repositories.IssueRepository, validate *validator.Validate) service_interfaces.IssueService {
	return &IssueServiceImpl{
		validate:   validate,
		repository: repository,
	}
}

// Add implements service_interfaces.IssueService.
func (i *IssueServiceImpl) Add(data requests.CreateOrUpdateIssueRequest, files []string, user *entities.User) (*responses.IssueResponse, *fiber.Error) {
	//validate input data
	if err := i.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadGateway, "Invalid data")
	}

	// check if issue name already exists
	issue, _ := i.repository.FindOneBy("name = ? AND created_by_id = ?", data.Name, user.Id)
	if issue != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Issue already exists")
	}

	issue = &entities.Issue{}
	// copy data from request to issue
	copier.Copy(issue, data)

	if !utils.IsEmpty(&data.ProjectId) {
		var project entities.Project
		err := i.repository.GetDataSource().First(&project, "id = ?", data.ProjectId).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Project not found")
		}
		issue.Project = project
	} else {
		issue.ProjectId = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	issue.CreatedById = user.Id
	if len(files) > 0 {
		issue.Files = sql.NullString{
			String: strings.Join(files, ","),
			Valid:  true,
		}
	}

	cerr := i.repository.Create(issue)
	if cerr != nil {
		log.Println("Error: ", cerr)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create new issue")
	}

	newIssue := &responses.IssueResponse{}
	copier.Copy(newIssue, issue)
	return newIssue, nil
}

// Delete implements service_interfaces.IssueService.
func (i *IssueServiceImpl) Delete(id string, user *entities.User) (*responses.IssueResponse, *fiber.Error) {
	issue, _ := i.repository.FindByID(id)
	if issue == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Issue not found")
	}

	if user.Role != enums.ADMIN && user.Id != issue.CreatedById {
		return nil, fiber.NewError(fiber.StatusForbidden, "You do not have permission to delete this issue")
	}

	derr := i.repository.Remove(id)
	if derr != nil {
		log.Println("Error: ", derr)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete this issue")
	}

	issueDetail := &responses.IssueResponse{}
	copier.Copy(issueDetail, issue)
	return issueDetail, nil
}

// Detail implements service_interfaces.IssueService.
func (i *IssueServiceImpl) Detail(id string) (*responses.IssueResponse, *fiber.Error) {
	issue, _ := i.repository.FindByID(id)
	if issue == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Issue not found")
	}

	issueDetail := &responses.IssueResponse{}
	copier.Copy(issueDetail, issue)
	return issueDetail, nil
}

// List implements service_interfaces.IssueService.
func (i *IssueServiceImpl) List(options requests.IssueFilters, user *entities.User) (*responses.IssuePageResponse, *fiber.Error) {
	var issues []entities.Issue

	builder := i.repository.GetDataSource().Model(&entities.Issue{})

	if !utils.IsEmpty(options.Query) {
		queryByName := fmt.Sprintf(`%%%s%%`, *options.Query)
		builder = builder.Where("name LIKE ? OR stack LIKE ?", queryByName, queryByName)
	}

	if !utils.IsEmpty((*string)(options.Status)) {
		builder = builder.Where("status = ?", *options.Status)
	}

	if !utils.IsEmpty((*string)(options.Priority)) {
		builder = builder.Where("priority = ?", *options.Priority)
	}

	if user.Role != enums.ADMIN {
		builder = builder.Where("created_by_id = ?", user.Id)
	}

	var totalItem int64
	builder.Count(&totalItem)
	builder.Offset((options.Page - 1) * options.Limit).Limit(options.Limit).Preload("CreatedBy").Preload("Project").Find(&issues)

	// Convert issues to []interface{}
	// jsonS, _ := json.Marshal(issues)
	// fmt.Println("Issues: ", string(jsonS))
	issueResponses := make([]responses.IssueResponse, len(issues))
	for i, issue := range issues {
		copier.Copy(&issueResponses[i], issue.Model)
		copier.Copy(&issueResponses[i], issue)
	}

	metadata := responses.NewPaginationMetaData(options.Page, options.Limit, int(totalItem), issueResponses)
	data := responses.NewPaginatedResponse(metadata)

	return &data, nil
}

// Update implements service_interfaces.IssueService.
func (i *IssueServiceImpl) Update(id string, data requests.CreateOrUpdateIssueRequest, files []string, user *entities.User) (*responses.IssueResponse, *fiber.Error) {
	//validate input data
	if err := i.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadGateway, "Invalid data")
	}

	// check if issue name already exists
	existedIssue, _ := i.repository.FindOneBy("name = ? AND created_by_id = ? AND id <> ?", data.Name, user.Id, id)
	if existedIssue != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Issue already exists")
	}

	issue, _ := i.repository.FindByID(id)
	if issue == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Issue not found")
	}

	// copy data from request to issue
	copier.Copy(issue, data)

	if !utils.IsEmpty(&data.ProjectId) {
		var project entities.Project
		err := i.repository.GetDataSource().First(&project, "id = ?", data.ProjectId).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Project not found")
		}
		issue.Project = project
	} else {
		issue.ProjectId = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	if len(files) > 0 {
		issue.Files = sql.NullString{
			String: strings.Join(files, ","),
			Valid:  true,
		}
	}

	cerr := i.repository.Update(issue)
	if cerr != nil {
		log.Println("Error: ", cerr)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update this issue")
	}

	updatedIssue := &responses.IssueResponse{}
	copier.Copy(updatedIssue, issue)
	return updatedIssue, nil
}
