package domain

type PaymentRepository interface {
	Save(payment *Payment) (*Payment, error)
	Find(id string) (*Payment, error)
	Update(payment *Payment) (*Payment, error)
}
