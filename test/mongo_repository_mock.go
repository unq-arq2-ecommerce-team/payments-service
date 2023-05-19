package test

import (
	"errors"

	"github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
)

type MockMongoRepository struct {
	SaveFn   func(payment *domain.Payment) (*domain.Payment, error)
	FindFn   func(id string) (*domain.Payment, error)
	UpdateFn func(payment *domain.Payment) (*domain.Payment, error)
}

func NewMockMongoRepository() *MockMongoRepository {
	return &MockMongoRepository{}
}

func (r *MockMongoRepository) Save(payment *domain.Payment) (*domain.Payment, error) {
	if r.SaveFn != nil {
		return r.SaveFn(payment)
	}
	return nil, errors.New("Save mock not implemented")
}

func (r *MockMongoRepository) Find(id string) (*domain.Payment, error) {
	if r.FindFn != nil {
		return r.FindFn(id)
	}
	return nil, errors.New("Find mock not implemented")
}

func (r *MockMongoRepository) Update(payment *domain.Payment) (*domain.Payment, error) {
	if r.UpdateFn != nil {
		return r.UpdateFn(payment)
	}
	return nil, errors.New("Update mock not implemented")
}

func (r *MockMongoRepository) WhenSaveDo(callback func(payment *domain.Payment) (*domain.Payment, error)) {
	r.SaveFn = callback
}

func (r *MockMongoRepository) WhenFindDo(callback func(id string) (*domain.Payment, error)) {
	r.FindFn = callback
}

func (r *MockMongoRepository) WhenUpdateDo(callback func(payment *domain.Payment) (*domain.Payment, error)) {
	r.UpdateFn = callback
}
