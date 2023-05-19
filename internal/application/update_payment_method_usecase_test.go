package application

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	domain "github.com/unq-arq2-ecommerce-team/payments-service/internal/domain"
	test "github.com/unq-arq2-ecommerce-team/payments-service/test"
)

func TestUpdatePaymentMethodUsecase(t *testing.T) {
	repository := test.NewMockMongoRepository()
	usecase := NewUpdatePaymentMethodUsecaseUseCase(repository)

	t.Run("should change payment method", func(t *testing.T) {
		method, _ := domain.NewPaymentMethod("creditCard", map[string]interface{}{
			"card_number":     "1234123412341234",
			"expiration_date": "01/2027",
			"cvv":             "123",
			"holder_name":     "Walter White",
		})
		resPayment := domain.NewPayment(domain.NewMoney(100, "USD"), method, "123", "123")
		repository.WhenSaveDo(func(payment *domain.Payment) (*domain.Payment, error) {
			return payment, nil
		})

		repository.WhenFindDo(func(id string) (*domain.Payment, error) {
			return resPayment, nil
		})

		result, err := usecase.Do(&UpdatePaymentMethodDto{
			PaymentId: "123",
			PaymentMethod: PaymentMethodDto{
				Type: "creditCard",
				Details: map[string]interface{}{
					"card_number":     "2222333344445555",
					"expiration_date": "01/2027",
					"cvv":             "123",
					"holder_name":     "Walter White",
				},
			},
		})
		require.NoError(t, err)
		require.NotEqual(t, result.(*domain.Payment).Method, method)

	})

	t.Run("should return error", func(t *testing.T) {
		method, _ := domain.NewPaymentMethod("creditCard", map[string]interface{}{
			"card_number":     "1234123412341234",
			"expiration_date": "01/2027",
			"cvv":             "123",
			"holder_name":     "Walter White",
		})
		resPayment := domain.NewPayment(domain.NewMoney(100, "USD"), method, "123", "123")
		repository.WhenSaveDo(func(payment *domain.Payment) (*domain.Payment, error) {
			return nil, errors.New("error")
		})

		repository.WhenFindDo(func(id string) (*domain.Payment, error) {
			return resPayment, nil
		})

		result, err := usecase.Do(&UpdatePaymentMethodDto{
			PaymentId: "123",
			PaymentMethod: PaymentMethodDto{
				Type: "creditCard",
				Details: map[string]interface{}{
					"card_number":     "2222333344445555",
					"expiration_date": "01/2027",
					"cvv":             "123",
					"holder_name":     "Walter White",
				},
			},
		})
		require.Error(t, err)
		require.Equal(t, result, nil)

	})

}
