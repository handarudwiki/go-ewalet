package domain

import (
	"context"
	"time"

	"github.com/handarudwiki/golang-ewalet/dto"
)

type Transaction struct {
	ID                  int64     `db:"id"`
	SofNumber           string    `db:"sof_number"`
	DofNumber           string    `db:"dof_number"`
	Amount              float64   `db:"ammount"`
	TransactionType     string    `db:"transaction_type"`
	AccountID           int64     `db:"account_id"`
	TransactionDatetime time.Time `db:"transaction_datetime"`
}

type TransactionRepository interface {
	Insert(ctx context.Context, transaction *Transaction) error
}

type TransactionService interface {
	TransferInquiry(ctx context.Context, req dto.TransferInquiryReq) (dto.TransferInquiryRes, error)
	TransferExecute(ctx context.Context, req dto.TransferExecuteReq) error
}
