package services

import (
	"log"
	"personal-growth/data/requests"
	"personal-growth/data/responses"
	"personal-growth/models"
	"personal-growth/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type ProjectServiceImpl struct {
	repository repositories.ProjectRepository
	validate   *validator.Validate
}

func NewProjectServiceImpl(repository repositories.ProjectRepository, validate *validator.Validate) ProjectService {
	return &ProjectServiceImpl{
		repository: repository,
		validate:   validate,
	}
}

// Add implements ProjectService.
func (p *ProjectServiceImpl) Add(data requests.CreateOrUpdateProjectRequest, user *models.User) (*responses.ProjectResponse, *fiber.Error) {
	//validate input data
	if err := p.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadGateway, "Invalid data")
	}

	// check if project name already exists
	project, _ := p.repository.FindOneBy("name = ? AND created_by_id = ?", data.Name, user.Id)
	if project != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Project already exists")
	}

	project = &models.Project{}
	// copy data from request to project
	copier.Copy(project, data)
	project.CreatedById = user.Id.String()

	cerr := p.repository.Create(project)
	if cerr != nil {
		log.Println("Error: ", cerr)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to add new project")
	}

	newProject := &responses.ProjectResponse{}
	copier.Copy(newProject, project)
	return newProject, nil
}

// Delete implements ProjectService.
func (p *ProjectServiceImpl) Delete(id string, user *models.User) (*responses.ProjectResponse, *fiber.Error) {
	panic("unimplemented")
}

// Detail implements ProjectService.
func (p *ProjectServiceImpl) Detail(id string) (*responses.ProjectResponse, *fiber.Error) {
	panic("unimplemented")
}

// List implements ProjectService.
func (p *ProjectServiceImpl) List(options requests.GetProjectOptions) (*responses.ProjectPageResponse, *fiber.Error) {
	panic("unimplemented")
}

// Update implements ProjectService.
func (p *ProjectServiceImpl) Update(id string, data requests.CreateOrUpdateProjectRequest) (*responses.ProjectResponse, *fiber.Error) {
	panic("unimplemented")
}
