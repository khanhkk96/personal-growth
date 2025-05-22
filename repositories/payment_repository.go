package repositories

import (
	"personal-growth/db/entities"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	BaseRepository[entities.Payment]
}

type paymentRepository struct {
	BaseRepository[entities.Payment]
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		BaseRepository: NewBaseRepository[entities.Payment](db),
		db:             db,
	}
}
