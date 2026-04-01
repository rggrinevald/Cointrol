package repository

import (
	"github.com/rggrinevald/cointrol/internal/domain/entity"
	domainRepo "github.com/rggrinevald/cointrol/internal/domain/repository"
	"gorm.io/gorm"
)

type investmentRepository struct {
	db *gorm.DB
}

func NewInvestmentRepository(db *gorm.DB) domainRepo.InvestmentRepository {
	return &investmentRepository{db: db}
}

func (r *investmentRepository) Create(investment *entity.Investment) (error) {
	resultado := r.db.Create(investment)
	return resultado.Error
}

func (r *investmentRepository) FindByID(id string) (*entity.Investment, error) {
	var investment entity.Investment
	resultado := r.db.First(&investment, "id = ?", id)
	return &investment, resultado.Error
}

func (r *investmentRepository) FindByUserID(user_id string) ([]*entity.Investment, error) {
	var investments []*entity.Investment
	resultado := r.db.Where("user_id = ?", user_id).Find(&investments)
	return investments, resultado.Error
}

func (r *investmentRepository) FindByAccountID(account_id string) ([]*entity.Investment, error) {
	var investments []*entity.Investment
	resultado := r.db.Where("account_id = ?", account_id).Find(&investments)
	return investments, resultado.Error
}

func (r *investmentRepository) Update(investment *entity.Investment) (error) {
	resultado := r.db.Save(investment)
	return resultado.Error
}

func (r *investmentRepository) Delete(id string) (error) {
	resultado := r.db.Delete(&entity.Investment{}, "id = ?", id)
	return resultado.Error
}
