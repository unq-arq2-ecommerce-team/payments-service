package domain

type PaymentMethodType string

const (
	CreaditCard PaymentMethodType = "credit_card"
	MercadoPago PaymentMethodType = "mercado_pago"
	Cash        PaymentMethodType = "cash"
)

type PaymentMethod interface {
	Execute(payment *Payment) error
}

type paymentMethod struct {
	Type PaymentMethodType
}

func NewPaymentMethod(t PaymentMethodType) PaymentMethod {
	return &paymentMethod{
		Type: t,
	}
}

func (p *paymentMethod) Execute(payment *Payment) error {
	return nil
}
