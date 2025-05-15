package repositories

import (
	"personal-growth/models"

	"gorm.io/gorm"
)

type IssueRepository interface {
	BaseRepository[models.Issue]
}

type issueRepository struct {
	BaseRepository[models.Issue]
	db *gorm.DB
}

func NewIssueRepository(db *gorm.DB) IssueRepository {
	return &issueRepository{
		BaseRepository: NewBaseRepository[models.Issue](db),
		db:             db,
	}
}
