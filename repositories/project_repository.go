package repositories

import (
	"personal-growth/db/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	BaseRepository[models.Project]
}

type projectRepository struct {
	BaseRepository[models.Project]
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		BaseRepository: NewBaseRepository[models.Project](db),
		db:             db,
	}
}
