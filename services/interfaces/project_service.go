package service_interfaces

import (
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/db/entities"

	"github.com/gofiber/fiber/v2"
)

type ProjectService interface {
	Add(data requests.CreateOrUpdateProjectRequest, user *entities.User) (*responses.ProjectResponse, *fiber.Error)
	Update(id string, data requests.CreateOrUpdateProjectRequest, user *entities.User) (*responses.ProjectResponse, *fiber.Error)
	Delete(id string, user *entities.User) (*responses.ProjectResponse, *fiber.Error)
	Detail(id string) (*responses.ProjectResponse, *fiber.Error)
	List(options requests.ProjectFilters, user *entities.User) responses.ProjectPageResponse
}
