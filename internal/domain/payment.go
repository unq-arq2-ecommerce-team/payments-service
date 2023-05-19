package domain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID           string
	Money        *Money
	Method       PaymentMethod
	Status       PaymentStatus
	PayerID      string
	CreationDate time.Time
	OrderID      string
}

func NewPayment(money *Money, method PaymentMethod, payerID string, orderID string) *Payment {
	id := uuid.New().String()
	payment := &Payment{
		ID:           id,
		Money:        money,
		Method:       method,
		Status:       stateMapper[confirmed],
		PayerID:      payerID,
		CreationDate: time.Now(),
		OrderID:      orderID,
	}
	return payment
}

func (p *Payment) Confirm() error {
	return p.Status.Confirm(p)
}

func (p *Payment) Reject() error {
	return p.Status.Reject(p)
}

func (p *Payment) Reset() error {
	return p.Status.Reset(p)
}

func (p *Payment) Execute() error {
	return p.Method.Execute(p)
}

func (p *Payment) UpdatePaymentMethod(method PaymentMethod) {
	p.Method = method
	p.Reset()
}
