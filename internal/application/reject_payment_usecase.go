package application

import (
	domain "github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
)

// reject payment usecase
type RejectPaymentUseCase struct {
	PaymentRepository domain.PaymentRepository
}

// reject payment usecase constructor
func NewRejectPaymentUseCase(paymentRepository domain.PaymentRepository) *RejectPaymentUseCase {
	return &RejectPaymentUseCase{
		PaymentRepository: paymentRepository,
	}
}

// reject payment usecase implementation
func (u *RejectPaymentUseCase) Do(input interface{}) (interface{}, error) {
	paymentId := input.(string)
	payment, err := u.PaymentRepository.Find(paymentId)
	if err != nil {
		return nil, err
	}
	payment.Reject()
	payment, err = u.PaymentRepository.Save(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
