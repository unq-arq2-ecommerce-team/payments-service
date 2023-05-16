package domain

import "errors"

const (
	CreaditCard = "creditCard"
	MercadoPago = "mercadoPago"
	Cash        = "cash"
)

var methodMapper = map[string]func(map[string]interface{}) (PaymentMethod, error){
	CreaditCard: NewCashMethod,
	MercadoPago: NewMercadoPagoMethod,
	Cash:        NewCashMethod,
}

type PaymentMethod interface {
	Execute(payment *Payment) error
	Type() string
}

func NewPaymentMethod(methodType string, details map[string]interface{}) (PaymentMethod, error) {
	methodFactory, ok := methodMapper[methodType]
	if !ok {
		return nil, errors.New("invalid payment method")
	}
	return methodFactory(details)

}

type MercadoPagoMethod struct {
	OperationId string
}

func (m *MercadoPagoMethod) Type() string {
	return MercadoPago
}

func (m *MercadoPagoMethod) Execute(payment *Payment) error {
	return nil
}

func NewMercadoPagoMethod(details map[string]interface{}) (PaymentMethod, error) {
	operationId, ok := details["operation_id"].(string)
	if !ok {
		return nil, errors.New("invalid operation_id")
	}
	return &MercadoPagoMethod{
		OperationId: operationId,
	}, nil
}

type CreditCardMethod struct {
	CardNumber     string
	ExpirationDate string
	CVV            string
	HolderName     string
}

func (c *CreditCardMethod) Type() string {
	return CreaditCard
}

func (c *CreditCardMethod) Execute(payment *Payment) error {
	return nil
}

func NewCreditCardMethod(details map[string]interface{}) (PaymentMethod, error) {
	cardNumber, ok := details["card_number"].(string)
	if !ok {
		return nil, errors.New("invalid card_number")
	}
	expirationDate, ok := details["expiration_date"].(string)
	if !ok {
		return nil, errors.New("invalid expiration_date")
	}
	cvv, ok := details["cvv"].(string)
	if !ok {
		return nil, errors.New("invalid cvv")
	}
	holderName, ok := details["holder_name"].(string)
	if !ok {
		return nil, errors.New("invalid holder_name")
	}
	return &CreditCardMethod{
		CardNumber:     cardNumber,
		ExpirationDate: expirationDate,
		CVV:            cvv,
		HolderName:     holderName,
	}, nil
}

type CashMethod struct{}

func (c *CashMethod) Type() string {
	return Cash
}

func (c *CashMethod) Execute(payment *Payment) error {
	return nil
}

func NewCashMethod(details map[string]interface{}) (PaymentMethod, error) {
	return &CashMethod{}, nil
}
