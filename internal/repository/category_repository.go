package repository

import (
	"github.com/rggrinevald/cointrol/internal/domain/entity"
	domainRepo "github.com/rggrinevald/cointrol/internal/domain/repository"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domainRepo.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *entity.Category) (error) {
	resultado := r.db.Create(category)
	return resultado.Error
}

func (r *categoryRepository) FindByID(id string) (*entity.Category, error) {
	var category entity.Category
	resultado := r.db.First(&category, "id = ?", id)
	return &category, resultado.Error
}

func (r *categoryRepository) FindByUserID(user_id string) ([]*entity.Category, error) {
	var categorys []*entity.Category
	resultado := r.db.Where("user_id = ?", user_id).Find(&categorys)
	return categorys, resultado.Error
}

func (r *categoryRepository) Update(category *entity.Category) (error) {
	resultado := r.db.Save(category)
	return resultado.Error
}

func (r *categoryRepository) Delete(id string) (error) {
	resultado := r.db.Delete(&entity.Category{}, "id = ?", id)
	return resultado.Error
}
