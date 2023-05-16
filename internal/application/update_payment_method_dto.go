package application

type UpdatePaymentMethodDto struct {
	PaymentId  string `json:"payment_id"`
	MethodType string `json:"method_type"`
}
