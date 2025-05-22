package repositories

import (
	"personal-growth/db/entities"

	"gorm.io/gorm"
)

type IssueRepository interface {
	BaseRepository[entities.Issue]
}

type issueRepository struct {
	BaseRepository[entities.Issue]
	db *gorm.DB
}

func NewIssueRepository(db *gorm.DB) IssueRepository {
	return &issueRepository{
		BaseRepository: NewBaseRepository[entities.Issue](db),
		db:             db,
	}
}
