package domain

import "errors"

var (
	ErrInvalidExpirationDate = errors.New("invalid expiration date")
)

type ExpirationDate struct {
	Value string
}

func NewExpirationDate(value string) (*ExpirationDate, error) {
	if len(value) != 7 {
		return nil, ErrInvalidExpirationDate
	}
	return &ExpirationDate{Value: value}, nil

}
