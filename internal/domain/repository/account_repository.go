package repository

import "github.com/rggrinevald/cointrol/internal/domain/entity"

type AccountRepository interface {
	Create(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	FindByUserID(id string) ([]*entity.Account, error)
	Update(user *entity.Account) error
	Delete(id string) error
}