package usecase

import (
	"errors"

	"github.com/rggrinevald/cointrol/internal/domain/entity"
	"github.com/rggrinevald/cointrol/internal/domain/repository"
)

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *userUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) CreateUser(user *entity.User)  error {
	usuarioExistente, _ := u.repo.FindByEmail(user.Email)

	if usuarioExistente != nil {
		return errors.New("Ja existe uma conta ativa com esse email!")
	}

	return u.repo.Create(user)
}

func (u *userUsecase) GetByID(id string) (*entity.User, error) {
	usuarioExistente, _ := u.repo.FindByID(id)

	if usuarioExistente == nil {
		return nil, errors.New("Usuario nao encontrado!")
	}

	return usuarioExistente, nil
}

func (u *userUsecase) UpdateUser(user *entity.User) error {
	usuarioExistente, _ := u.repo.FindByID(user.ID)

	if usuarioExistente == nil {
		return errors.New("Usuario nao encontrado!")
	}

	return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(user *entity.User) error {
	usuarioExistente, _ := u.repo.FindByID(user.ID)

	if usuarioExistente == nil {
		return errors.New("Usuario nao encontrado!")
	}

	return u.repo.Delete(user.ID)
}