package repository

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"gorm.io/gorm"
)

type TopUpRepository struct {
	db *gorm.DB
}

func NewTopUp(db *gorm.DB) domain.TopUpRepository {
	return &TopUpRepository{
		db: db,
	}
}

func (t TopUpRepository) FindBYID(ctx context.Context, id string) (domain.TopUp, error) {
	var topUp domain.TopUp

	err := t.db.Where("id=?", id).First(&topUp).Error

	if err != nil {
		return domain.TopUp{}, err
	}
	return topUp, nil
}

func (t TopUpRepository) Insert(ctx context.Context, topup *domain.TopUp) error {
	err := t.db.Create(&topup).Error

	if err != nil {
		return err
	}

	return nil
}

func (t TopUpRepository) Update(ctx context.Context, topup *domain.TopUp) error {
	var data domain.TopUp

	err := t.db.Where("id=?", topup.ID).Error

	if err != nil {
		return err
	}

	data.Amount = topup.Amount
	data.SnapURL = topup.SnapURL
	data.Status = topup.Status

	err = t.db.Save(&data).Error

	if err != nil {
		return err
	}

	return nil
}
