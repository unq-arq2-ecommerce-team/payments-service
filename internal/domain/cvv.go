package domain

import "errors"

var (
	ErrInvalidCVV = errors.New("invalid cvv")
)

type CVV struct {
	Code string
}

func NewCVV(code string) (*CVV, error) {
	if len(code) != 3 {
		return nil, ErrInvalidCVV
	}
	return &CVV{Code: code}, nil
}
