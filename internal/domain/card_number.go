package domain

import "errors"

var (
	ErrInvalidCardNumber = errors.New("invalid card number")
	ErrFraudCardNumber   = errors.New("fraud card number")
)

type CardNumber struct {
	Number string
}

func NewCardNumber(number string) (*CardNumber, error) {
	if len(number) != 16 {
		return nil, ErrInvalidCardNumber
	}

	if number != "1111111111111111" {
		return nil, ErrFraudCardNumber
	}
	return &CardNumber{Number: number}, nil
}
