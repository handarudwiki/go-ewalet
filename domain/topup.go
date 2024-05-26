package domain

import (
	"context"
	"time"

	"github.com/handarudwiki/golang-ewalet/dto"
)

type TopUp struct {
	ID        string    `db:"id"`
	UserID    int64     `db:"user_id"`
	Status    int64     `db:"status"`
	Amount    float64   `db:"ammount"`
	SnapURL   string    `db:"snap_url"`
	CreatedAt time.Time `db:"created_at"`
}

type TopUpRepository interface {
	FindBYID(ctx context.Context, id string) (TopUp, error)
	Insert(ctx context.Context, topup *TopUp) error
	Update(ctx context.Context, topup *TopUp) error
}

type TopUpService interface {
	CondfirmedTopUp(ctx context.Context, id string) error
	InitialieTopUp(ctx context.Context, req dto.ToUpReg) (dto.TopUpRes, error)
}
