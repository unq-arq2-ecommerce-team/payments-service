package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPayment(t *testing.T) {
	confirmPendingPayment(t)
	rejectPendingPayment(t)
	updatePaymentMethodToRejectedPayment(t)
	executePayment(t)
}

func createPayment(statusString string) *Payment {
	status, _ := NewPaymentStatus(statusString)
	method, _ := NewPaymentMethod("creditCard", map[string]interface{}{
		"card_number":     "1234123412341234",
		"expiration_date": "01/2027",
		"cvv":             "123",
		"holder_name":     "Walter White",
	})
	payment := NewPayment(NewMoney(100, "USD"), method, "123", "123")
	payment.Status = status
	return payment

}
func confirmPendingPayment(t *testing.T) {
	payment := createPayment(pending)
	require.Equal(t, pending, payment.Status.String())
	err := payment.Confirm()
	require.NoError(t, err)
	require.Equal(t, confirmed, payment.Status.String())
}

func rejectPendingPayment(t *testing.T) {
	payment := createPayment(pending)
	require.Equal(t, pending, payment.Status.String())
	err := payment.Reject()
	require.NoError(t, err)
	require.Equal(t, rejected, payment.Status.String())
}

func updatePaymentMethodToRejectedPayment(t *testing.T) {
	payment := createPayment(rejected)
	require.Equal(t, rejected, payment.Status.String())
	newPaymentMethod := &CreditCardMethod{
		CardNumber: &CardNumber{
			Number: "2222333344445555",
		},
		ExpirationDate: &ExpirationDate{
			Value: "01/2027",
		},
		CVV: &CVV{
			Code: "123",
		},
		HolderName: &HolderName{
			Name: "Walter White",
		},
	}
	payment.UpdatePaymentMethod(
		newPaymentMethod,
	)
	require.Equal(t, newPaymentMethod, payment.Method)
	require.Equal(t, pending, payment.Status.String())
}

func executePayment(t *testing.T) {
	payment := createPayment(pending)
	err := payment.Execute()
	require.NoError(t, err)
}
