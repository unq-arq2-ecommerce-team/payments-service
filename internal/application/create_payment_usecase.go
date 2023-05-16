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
	payer := domain.NewPayer(paymentDto.Payer.ID, paymentDto.Payer.Name, paymentDto.Payer.Email)
	paymentMethod, err := domain.NewPaymentMethod(paymentDto.MethodType)
	if err != nil {
		return nil, err
	}

	payment := domain.NewPayment(money, paymentMethod, payer, paymentDto.OrderID)

	payment, err = u.repository.Save(payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
