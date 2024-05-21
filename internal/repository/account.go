package repository

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) domain.AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (a accountRepository) FindByUserID(ctx context.Context, id int64) (domain.Account, error) {
	var account domain.Account

	err := a.db.Where("user_id=?", id).First(&account).Error
	if err != nil {
		return domain.Account{}, err
	}

	return account, nil
}

func (a accountRepository) FindByAccountNumber(ctx context.Context, accNumber string) (domain.Account, error) {
	var account domain.Account
	err := a.db.Where("account_number=?", accNumber).First(&account).Error
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (a accountRepository) Update(ctx context.Context, account *domain.Account) error {
	var data domain.Account

	err := a.db.Where("id=?", account.ID).First(&data).Error

	if err != nil {
		return err
	}
	data.Balance = account.Balance
	err = a.db.Save(&data).Error
	if err != nil {
		return err
	}

	return nil
}
