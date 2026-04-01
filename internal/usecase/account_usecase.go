package usecase

import (
	"errors"

	"github.com/rggrinevald/cointrol/internal/domain/entity"
	"github.com/rggrinevald/cointrol/internal/domain/repository"
)

type accountUsecase struct {
	repo repository.AccountRepository
}

func NewAccountUsecase(repo repository.AccountRepository) *accountUsecase {
	return &accountUsecase{repo: repo}
}

func (a *accountUsecase) CreateAccount(account *entity.Account) error {
	accounts, _ := a.repo.FindByUserID(account.UserID)
	for _, acc := range accounts {
		if acc.Bank == account.Bank {
			return errors.New("Voce ja possui uma conta para esse banco")
		}
	}
	return a.repo.Create(account)
}

func (a *accountUsecase) GetByID(id string) (*entity.Account, error) {
	contaExistente, _ := a.repo.FindByID(id)

	if contaExistente == nil {
		return nil, errors.New("Conta nao encontrada")
	}

	return contaExistente, nil
}

func (a *accountUsecase) UpdateAccount(account *entity.Account) error {
	contaExistente, _ := a.repo.FindByID(account.ID)

	if contaExistente == nil {
		return errors.New("Conta nao encontrada")
	}

	return a.repo.Update(account)
}

func (a *accountUsecase) DeleteAccount(account *entity.Account) error {
	contaExistente, _ := a.repo.FindByID(account.ID)

	if contaExistente == nil {
		return errors.New("Conta nao encontrada")
	}

	return a.repo.Delete(account.ID)
}