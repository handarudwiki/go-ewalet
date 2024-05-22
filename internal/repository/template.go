package repository

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"gorm.io/gorm"
)

type templateRepository struct {
	db *gorm.DB
}

func NewTemplate(db *gorm.DB) domain.TemplateRepository {
	return &templateRepository{
		db: db,
	}
}

func (r *templateRepository) FindByCode(ctx context.Context, code string) (domain.Template, error) {
	var template domain.Template

	err := r.db.Where("code=?", code).First(&template).Error

	if err != nil {
		return domain.Template{}, err
	}

	return template, err
}
