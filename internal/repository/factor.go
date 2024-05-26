package repository

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"gorm.io/gorm"
)

type factorRepository struct {
	db *gorm.DB
}

func NewFactor(db *gorm.DB) domain.FactoryRepository {
	return &factorRepository{
		db: db,
	}
}

func (r factorRepository) FindByUserID(ctx context.Context, userID int64) (domain.Factor, error) {
	var factor domain.Factor

	err := r.db.Where("user_id=?", userID).First(&factor).Error

	if err != nil {
		return domain.Factor{}, err
	}

	return factor, nil
}
