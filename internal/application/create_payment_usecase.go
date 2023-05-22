package application

import (
	domain "github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
)

// create payments use case implementing usecase interface
type CreatePaymentUseCase struct {
	repository domain.PaymentRepository
}

// create payments use case constructor
func NewCreatePaymentUseCase(repository domain.PaymentRepository) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{
		repository: repository,
	}
}

// create payments use case implementation
func (u *CreatePaymentUseCase) Do(input interface{}) (interface{}, error) {
	paymentDto := input.(*CreatePaymentDto)
	money := domain.NewMoney(paymentDto.Amount, paymentDto.Currency)
	paymentMethod, err := domain.NewPaymentMethod(paymentDto.PaymentMethod.Type, paymentDto.PaymentMethod.Details)
	if err != nil {
		return nil, err
	}

	payment := domain.NewPayment(money, paymentMethod, paymentDto.CustomerId, paymentDto.OrderID)

	payment, err = u.repository.Save(payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
