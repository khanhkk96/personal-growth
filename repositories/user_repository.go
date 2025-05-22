package repositories

import (
	"personal-growth/db/entities"

	"gorm.io/gorm"
)

// UserRepository defines methods specific to the User model.
type UserRepository interface {
	BaseRepository[entities.User]
	FindByEmail(email string) (*entities.User, error)
}

// userRepository is the implementation of UserRepository.
type userRepository struct {
	BaseRepository[entities.User]
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[entities.User](db),
		db:             db,
	}
}

// FindByEmail finds a user by their email.
func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
