package application

import (
	domain "github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
)

// create payments use case implementing usecase interface
type UpdatePaymentMethodUsecaseUseCase struct {
	repository domain.PaymentRepository
}

// create payments use case constructor
func NewUpdatePaymentMethodUsecaseUseCase(repository domain.PaymentRepository) *UpdatePaymentMethodUsecaseUseCase {
	return &UpdatePaymentMethodUsecaseUseCase{
		repository: repository,
	}
}

// create payments use case implementation
func (u *UpdatePaymentMethodUsecaseUseCase) Do(input interface{}) (interface{}, error) {
	updatePaymentMethodDto := input.(*UpdatePaymentMethodDto)
	paymentMethod, err := domain.NewPaymentMethod(updatePaymentMethodDto.PaymentMethod.Type, updatePaymentMethodDto.PaymentMethod.Details)
	if err != nil {
		return nil, err
	}
	payment, err := u.repository.Find(updatePaymentMethodDto.PaymentId)
	if err != nil {
		return nil, err
	}
	payment.UpdatePaymentMethod(paymentMethod)
	payment, err = u.repository.Save(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil

}
