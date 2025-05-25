package repositories

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// BaseRepository defines common methods for any model.
type BaseRepository[T any] interface {
	Create(entity *T) error
	FindByID(id interface{}) (*T, error)
	FindOneBy(query interface{}, args ...interface{}) (*T, error)
	FindAll(cons ...interface{}) ([]T, error)
	FindMany(query interface{}, args ...interface{}) ([]T, error)
	FindMapMany(cons map[string]interface{}) ([]T, error)
	Update(entity *T) error
	Delete(id interface{}) error
	Remove(id interface{}) error
	GetDataSource() *gorm.DB
}

// baseRepository is the default implementation of BaseRepository.
type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db}
}

func (r *baseRepository[T]) GetDataSource() *gorm.DB {
	return r.db
}

func (r *baseRepository[T]) Create(entity *T) error {
	return r.db.Create(&entity).Error
}

func (r *baseRepository[T]) FindByID(id interface{}) (*T, error) {
	var entity T
	if err := r.db.First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *baseRepository[T]) FindOneBy(query interface{}, args ...interface{}) (*T, error) {
	var entity T

	if err := r.db.Where(query, args...).First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &entity, nil
}

func (r *baseRepository[T]) FindAll(cons ...interface{}) ([]T, error) {
	var models []T
	if err := r.db.Find(&models, cons).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r *baseRepository[T]) FindMany(query interface{}, args ...interface{}) ([]T, error) {
	var models []T
	if err := r.db.Where(query, args).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r *baseRepository[T]) FindMapMany(cons map[string]interface{}) ([]T, error) {
	var models []T
	if err := r.db.Where(cons).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r *baseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepository[T]) Delete(id interface{}) error {
	var entity T
	return r.db.Unscoped().Delete(&entity, "id = ?", id).Error
}

func (r *baseRepository[T]) Remove(id interface{}) error {
	var entity T
	return r.db.Delete(&entity, "id = ?", id).Error
}
