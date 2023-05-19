package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaymentMethod(t *testing.T) {
	t.Run("Credit card method", func(t *testing.T) {
		_, err := NewPaymentMethod("creditCard", map[string]interface{}{
			"card_number":     "12341234",
			"expiration_date": "01/27",
			"cvv":             "12",
			"holder_name":     "W",
		})
		require.Error(t, err)

		_, err = NewPaymentMethod("creditCard", map[string]interface{}{
			"card_number":     "1234123412341234",
			"expiration_date": "01",
			"cvv":             "12",
			"holder_name":     "W",
		})
		require.Error(t, err)
		_, err = NewPaymentMethod("creditCard", map[string]interface{}{
			"card_number":     "1234123412341234",
			"expiration_date": "01/2027",
			"cvv":             "12",
			"holder_name":     "W",
		})
		require.Error(t, err)

		_, err = NewPaymentMethod("creditCard", map[string]interface{}{
			"card_number":     "1234123412341234",
			"expiration_date": "01/2027",
			"cvv":             "123",
			"holder_name":     "W",
		})
		require.Error(t, err)

		method, err := NewPaymentMethod("creditCard", map[string]interface{}{
			"card_number":     "1234123412341234",
			"expiration_date": "01/2027",
			"cvv":             "123",
			"holder_name":     "Walter White",
		})
		require.NoError(t, err)

		require.Equal(t, "creditCard", method.Type())
	})
}
