package repositories

import (
	"personal-growth/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	BaseRepository[models.Payment]
}

type paymentRepository struct {
	BaseRepository[models.Payment]
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		BaseRepository: NewBaseRepository[models.Payment](db),
		db:             db,
	}
}
