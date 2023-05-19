package domain

import "errors"

var (
	ErrInvalidHolderName = errors.New("invalid holder name")
)

type HolderName struct {
	Name string
}

func NewHolderName(name string) (*HolderName, error) {
	if len(name) < 2 || len(name) > 60 {
		return nil, ErrInvalidHolderName
	}
	return &HolderName{Name: name}, nil
}
