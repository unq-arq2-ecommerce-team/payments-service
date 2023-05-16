package application

type UpdatePaymentMethodDto struct {
	PaymentId     string           `json:"payment_id"`
	PaymentMethod PaymentMethodDto `json:"method"`
}
