package domain

import "errors"

const (
	CreaditCard = "creditCard"
	MercadoPago = "mercadoPago"
	Cash        = "cash"
)

var methodMapper = map[string]PaymentMethod{
	CreaditCard: &paymentMethod{
		Type: CreaditCard,
	},
	MercadoPago: &paymentMethod{
		Type: MercadoPago,
	},
	Cash: &paymentMethod{
		Type: Cash,
	},
}

type PaymentMethod interface {
	Execute(payment *Payment) error
}

type paymentMethod struct {
	Type string
}

func NewPaymentMethod(t string) (PaymentMethod, error) {
	method, ok := methodMapper[t]
	if !ok {
		return nil, errors.New("invalid payment method")
	}
	return method, nil

}

func (p *paymentMethod) Execute(payment *Payment) error {
	return nil
}
