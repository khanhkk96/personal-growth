package repositories

import (
	"personal-growth/db/entities"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	BaseRepository[entities.Project]
}

type projectRepository struct {
	BaseRepository[entities.Project]
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		BaseRepository: NewBaseRepository[entities.Project](db),
		db:             db,
	}
}
