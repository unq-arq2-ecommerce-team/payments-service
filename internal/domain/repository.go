package domain

type PaymentRepository interface {
	Store(payment *Payment) error
	Find(id string) (*Payment, error)
	Confirm(id string) error
	Reject(id string) error
	Reset(id string) error
}
