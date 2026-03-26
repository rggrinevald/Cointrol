package repository

import "github.com/rggrinevald/cointrol/internal/domain/entity"

type CategoryRepository interface {
	Create(category *entity.Category) error
	FindByID(id string) (*entity.Category, error)
	FindByUserID(id string) ([]*entity.Category, error)
	Update(category *entity.Category) error
	Delete(id string) error
}