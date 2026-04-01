package repository

import (
	"github.com/rggrinevald/cointrol/internal/domain/entity"
	domainRepo "github.com/rggrinevald/cointrol/internal/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domainRepo.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) (error) {
	resultado := r.db.Create(user)
	return resultado.Error
}

func (r *userRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	resultado := r.db.First(&user, "id = ?", id)
	return &user, resultado.Error
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	resultado := r.db.First(&user, "email = ?", email)
	return &user, resultado.Error
}

func (r *userRepository) Update(user *entity.User) (error) {
	resultado := r.db.Save(user)
	return resultado.Error
}

func (r *userRepository) Delete(id string) (error) {
	resultado := r.db.Delete(&entity.User{}, "id = ?", id)
	return resultado.Error
}