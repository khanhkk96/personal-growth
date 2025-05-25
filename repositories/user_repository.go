package repositories

import (
	"personal-growth/db/models"

	"gorm.io/gorm"
)

// UserRepository defines methods specific to the User model.
type UserRepository interface {
	BaseRepository[models.User]
	FindByEmail(email string) (*models.User, error)
}

// userRepository is the implementation of UserRepository.
type userRepository struct {
	BaseRepository[models.User]
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[models.User](db),
		db:             db,
	}
}

// FindByEmail finds a user by their email.
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
