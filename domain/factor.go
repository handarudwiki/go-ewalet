package domain

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/dto"
)

type Factor struct {
	ID     int64  `db:"id"`
	UserID int64  `db:"user_id"`
	PIN    string `db:"pin"`
}

type FactoryRepository interface {
	FindByUserID(ctx context.Context, userID int64) (Factor, error)
}

type FactoryService interface {
	ValidatePIN(ctx context.Context, req dto.ValidatePinReq) error
}
