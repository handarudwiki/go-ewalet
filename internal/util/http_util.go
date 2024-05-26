package util

import (
	"errors"

	"github.com/handarudwiki/golang-ewalet/domain"
)

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	case errors.Is(err, domain.ErrUsernameTakeb):
		return 400
	case errors.Is(err, domain.ErrOTPInvalid):
		return 400
	case errors.Is(err, domain.ErrPinInvalid):
		return 400
	default:
		return 500
	}
}
