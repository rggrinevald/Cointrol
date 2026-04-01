package repository

import (
	"github.com/rggrinevald/cointrol/internal/domain/entity"
	domainRepo "github.com/rggrinevald/cointrol/internal/domain/repository"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) domainRepo.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *entity.Account) (error) {
	resultado := r.db.Create(account)
	return resultado.Error
}

func (r *accountRepository) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	resultado := r.db.First(&account, "id = ?", id)
	return &account, resultado.Error
}

func (r *accountRepository) FindByUserID(user_id string) ([]*entity.Account, error) {
	var accounts []*entity.Account
	resultado := r.db.Where("user_id = ?", user_id).Find(&accounts)
	return accounts, resultado.Error
}

func (r *accountRepository) Update(account *entity.Account) (error) {
	resultado := r.db.Save(account)
	return resultado.Error
}

func (r *accountRepository) Delete(id string) (error) {
	resultado := r.db.Delete(&entity.Account{}, "id = ?", id)
	return resultado.Error
}
