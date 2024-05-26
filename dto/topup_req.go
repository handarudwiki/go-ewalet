package dto

type ToUpReg struct {
	Amount float64 `json:"amount"`
	UserID int64   `json:"-"`
}
