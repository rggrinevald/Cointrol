package repository

import (
	"time"

	"github.com/rggrinevald/cointrol/internal/domain/entity"
	domainRepo "github.com/rggrinevald/cointrol/internal/domain/repository"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domainRepo.TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *entity.Transaction) (error) {
	resultado := r.db.Create(transaction)
	return resultado.Error
}

func (r *transactionRepository) FindByID(id string) (*entity.Transaction, error) {
	var transaction entity.Transaction
	resultado := r.db.First(&transaction, "id = ?", id)
	return &transaction, resultado.Error
}

func (r *transactionRepository) FindByUserID(user_id string) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	resultado := r.db.Where("user_id = ?", user_id).Find(&transactions)
	return transactions, resultado.Error
}

func (r *transactionRepository) FindByCategoryID(category_id string) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	resultado := r.db.Where("category_id = ?", category_id).Find(&transactions)
	return transactions, resultado.Error
}

func (r *transactionRepository) FindByDate(date time.Time) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	resultado := r.db.Where("date = ?", date).Find(&transactions)
	return transactions, resultado.Error
}

func (r *transactionRepository) Update(transaction *entity.Transaction) (error) {
	resultado := r.db.Save(transaction)
	return resultado.Error
}

func (r *transactionRepository) Delete(id string) (error) {
	resultado := r.db.Delete(&entity.Transaction{}, "id = ?", id)
	return resultado.Error
}
