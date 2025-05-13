package services

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/models"

	"github.com/gofiber/fiber/v2"
)

type ProjectService interface {
	Add(data requests.CreateOrUpdateProjectRequest, user *models.User) (*responses.ProjectResponse, *fiber.Error)
	Update(id string, data requests.CreateOrUpdateProjectRequest, user *models.User) (*responses.ProjectResponse, *fiber.Error)
	Delete(id string, user *models.User) (*responses.ProjectResponse, *fiber.Error)
	Detail(id string) (*responses.ProjectResponse, *fiber.Error)
	List(options requests.ProjectFilters, user *models.User) (*responses.ProjectPageResponse, *fiber.Error)
}
