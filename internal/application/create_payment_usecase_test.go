package application

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	domain "github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
	test "github.com/unq-arq2-ecommerce-team/payments-service/test"
)

func TestCreatePaymentUsecase(t *testing.T) {
	repository := test.NewMockMongoRepository()
	usecase := NewCreatePaymentUseCase(repository)

	t.Run("should create payment", func(t *testing.T) {
		var resPayment *domain.Payment
		repository.WhenSaveDo(func(payment *domain.Payment) (*domain.Payment, error) {
			resPayment = payment
			return payment, nil
		})

		result, err := usecase.Do(&CreatePaymentDto{
			Amount:   100,
			Currency: "USD",
			PaymentMethod: PaymentMethodDto{
				Type: "creditCard",
				Details: map[string]interface{}{
					"card_number":     "1234123412341234",
					"expiration_date": "01/2027",
					"cvv":             "123",
					"holder_name":     "Walter White",
				},
			},
			CustomerId: "123",
			OrderID:    "123",
		})
		require.NoError(t, err)
		require.Equal(t, resPayment, result)

	})

	t.Run("should return error", func(t *testing.T) {
		repository.WhenSaveDo(func(payment *domain.Payment) (*domain.Payment, error) {
			return nil, errors.New("error")
		})

		result, err := usecase.Do(&CreatePaymentDto{
			Amount:   100,
			Currency: "USD",
			PaymentMethod: PaymentMethodDto{
				Type: "creditCard",
				Details: map[string]interface{}{
					"card_number":     "1234123412341234",
					"expiration_date": "01/2027",
					"cvv":             "123",
					"holder_name":     "Walter White",
				},
			},
			CustomerId: "123",
			OrderID:    "123",
		})
		require.Error(t, err)
		require.Equal(t, result, nil)

	})

}
