package repository

import (
	"time"

	"github.com/rggrinevald/cointrol/internal/domain/entity"
)

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	FindByID(id string) (*entity.Transaction, error)
	FindByUserID(id string) ([]*entity.Transaction, error)
	FindByCategoryID(id string) ([]*entity.Transaction, error)
	FindByDate(date time.Time) ([]*entity.Transaction, error)
	Update(user *entity.Transaction) error
	Delete(id string) error
}