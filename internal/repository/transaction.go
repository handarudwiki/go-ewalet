package repository

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t transactionRepository) Insert(ctx context.Context, transaction *domain.Transaction) error {
	err := t.db.Create(&transaction).Error

	if err != nil {
		return err
	}

	return nil
}
