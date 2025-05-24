package service_interfaces

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/db/models"

	"github.com/gofiber/fiber/v2"
)

type IssueService interface {
	Add(data requests.CreateOrUpdateIssueRequest, files []string, user *models.User) (*responses.IssueResponse, *fiber.Error)
	Update(id string, data requests.CreateOrUpdateIssueRequest, files []string, user *models.User) (*responses.IssueResponse, *fiber.Error)
	Delete(id string, user *models.User) (*responses.IssueResponse, *fiber.Error)
	Detail(id string) (*responses.IssueResponse, *fiber.Error)
	List(options requests.IssueFilters, user *models.User) responses.IssuePageResponse
}
