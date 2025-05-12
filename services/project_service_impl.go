package services

import (
	"log"
	"personal-growth/common/enums"
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
	project.CreatedById = user.Id

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
	project, _ := p.repository.FindByID(id)
	if project == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Project not found")
	}

	if user.Role != enums.ADMIN && user.Id != project.CreatedById {
		return nil, fiber.NewError(fiber.StatusForbidden, "You do not have permission to delete this project")
	}

	derr := p.repository.Remove(id)
	if derr != nil {
		log.Println("Error: ", derr)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete this project")
	}

	projectDetail := &responses.ProjectResponse{}
	copier.Copy(projectDetail, project)
	return projectDetail, nil
}

// Detail implements ProjectService.
func (p *ProjectServiceImpl) Detail(id string) (*responses.ProjectResponse, *fiber.Error) {
	project, _ := p.repository.FindByID(id)
	if project == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Project not found")
	}

	projectDetail := &responses.ProjectResponse{}
	copier.Copy(projectDetail, project)
	return projectDetail, nil
}

// List implements ProjectService.
func (p *ProjectServiceImpl) List(options requests.GetProjectOptions) (*responses.ProjectPageResponse, *fiber.Error) {
	panic("unimplemented")
}

// Update implements ProjectService.
func (p *ProjectServiceImpl) Update(id string, data requests.CreateOrUpdateProjectRequest, user *models.User) (*responses.ProjectResponse, *fiber.Error) {
	//validate input data
	if err := p.validate.Struct(data); err != nil {
		return nil, fiber.NewError(fiber.StatusBadGateway, "Invalid data")
	}

	// check if project name already exists
	existedProject, _ := p.repository.FindOneBy("name = ? AND created_by_id = ? AND id = ?", data.Name, user.Id, id)
	if existedProject != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Project already exists")
	}

	project, _ := p.repository.FindByID(id)
	if project == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Project not found")
	}

	// copy data from request to project
	copier.Copy(project, data)

	cerr := p.repository.Update(project)
	if cerr != nil {
		log.Println("Error: ", cerr)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update this project")
	}

	updatedProject := &responses.ProjectResponse{}
	copier.Copy(updatedProject, project)
	return updatedProject, nil
}
