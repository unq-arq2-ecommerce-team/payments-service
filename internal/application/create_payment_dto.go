package application

// payment dto
type CreatePaymentDto struct {
	ID            string           `json:"id"`
	Amount        float64          `json:"amount"`
	Currency      string           `json:"currency"`
	PaymentMethod PaymentMethodDto `json:"method"`
	PayerID       string           `json:"payer_id"`
	OrderID       string           `json:"order_id"`
}
