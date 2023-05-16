package application

import (
	domain "github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
)

// confirm payment usecase
type ConfirmPaymentUseCase struct {
	PaymentRepository domain.PaymentRepository
}

// confirm payment usecase constructor
func NewConfirmPaymentUseCase(paymentRepository domain.PaymentRepository) *ConfirmPaymentUseCase {
	return &ConfirmPaymentUseCase{
		PaymentRepository: paymentRepository,
	}
}

// confirm payment usecase implementation
func (u *ConfirmPaymentUseCase) Do(input interface{}) (interface{}, error) {
	paymentId := input.(string)
	payment, err := u.PaymentRepository.Find(paymentId)
	if err != nil {
		return nil, err
	}
	payment.Confirm()
	payment, err = u.PaymentRepository.Save(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
