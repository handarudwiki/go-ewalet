package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
)

type topuUSevice struct {
	notificationService   domain.NotificationService
	topUpRepository       domain.TopUpRepository
	midtransService       domain.MidtransService
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
}

func NewTopup(notificationService domain.NotificationService,
	topUpRepository domain.TopUpRepository,
	midtransService domain.MidtransService,
	accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository,
) domain.TopUpService {
	return &topuUSevice{
		topUpRepository:       topUpRepository,
		notificationService:   notificationService,
		midtransService:       midtransService,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (t topuUSevice) InitialieTopUp(ctx context.Context, req dto.ToUpReg) (dto.TopUpRes, error) {
	topUp := domain.TopUp{
		ID:     uuid.NewString(),
		UserID: req.UserID,
		Status: 0,
		Amount: req.Amount,
	}
	err := t.midtransService.GenerateSnapURL(ctx, &topUp)
	if err != nil {
		return dto.TopUpRes{}, err
	}

	return dto.TopUpRes{
		SnapURL: topUp.SnapURL,
	}, nil
}
func (t topuUSevice) CondfirmedTopUp(ctx context.Context, id string) error {
	topUp, err := t.topUpRepository.FindBYID(ctx, id)

	if err != nil {
		return err
	}
	if topUp == (domain.TopUp{}) {
		return errors.New("top up not found")
	}

	account, err := t.accountRepository.FindByUserID(ctx, topUp.UserID)
	if err != nil {
		return err
	}
	if account == (domain.Account{}) {
		return domain.ErrAccountNotFound
	}

	err = t.transactionRepository.Insert(ctx, &domain.Transaction{
		AccountID:           account.ID,
		SofNumber:           "00",
		DofNumber:           account.AccountNumber,
		TransactionType:     "C",
		Amount:              topUp.Amount,
		TransactionDatetime: time.Now(),
	})

	if err != nil {
		return err
	}

	account.Balance += topUp.Amount

	err = t.accountRepository.Update(ctx, &account)
	if err != nil {
		return err
	}

	data := map[string]string{
		"amount": fmt.Sprintf("%.2f", topUp.Amount),
	}

	err = t.notificationService.Insert(ctx, account.UserID, "TOPUP_SUCCESS", data)

	if err != nil {
		return err
	}

	return nil
}
