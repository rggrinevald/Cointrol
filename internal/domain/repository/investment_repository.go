package repository

import "github.com/rggrinevald/cointrol/internal/domain/entity"

type InvestmentRepository interface {
	Create(investment *entity.Investment) error
	FindByID(id string) (*entity.Investment, error)
	FindByUserID(id string) ([]*entity.Investment, error)
	FindByAccountID(id string) ([]*entity.Investment, error)
	Update(user *entity.Investment) error
	Delete(id string) error
}