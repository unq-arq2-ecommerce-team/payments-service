package application

// payment dto
type CreatePaymentDto struct {
	ID         string    `json:"id"`
	Amount     float64   `json:"amount"`
	Currency   string    `json:"currency"`
	MethodType string    `json:"method_type"`
	Payer      *PayerDto `json:"payer"`
	OrderID    string    `json:"order_id"`
}
