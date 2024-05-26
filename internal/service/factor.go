package service

import (
	"context"

	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
	"golang.org/x/crypto/bcrypt"
)

type factorService struct {
	factorRepository domain.FactoryRepository
}

func NewFactor(factorRepository domain.FactoryRepository) domain.FactoryService {
	return &factorService{
		factorRepository: factorRepository,
	}
}

func (f factorService) ValidatePIN(ctx context.Context, req dto.ValidatePinReq) error {
	factor, err := f.factorRepository.FindByUserID(ctx, req.UserID)

	if err != nil {
		return domain.ErrPinInvalid
	}

	if factor == (domain.Factor{}) {
		return domain.ErrPinInvalid
	}

	err = bcrypt.CompareHashAndPassword([]byte(factor.PIN), []byte(req.PIN))

	if err != nil {
		return domain.ErrPinInvalid
	}

	return nil
}
