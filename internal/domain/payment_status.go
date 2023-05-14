package domain

import "errors"

const (
	pending   = "PENDING"
	confirmed = "CONFIRMED"
	rejected  = "REJECTED"
)

var stateMapper = map[string]PaymentStatus{
	pending:   PendingPaymentStatus{},
	confirmed: ConfirmedPaymentStatus{},
	rejected:  RejectedPaymentStatus{},
}

type PaymentStatus interface {
	Confirm(payment *Payment) error
	Reject(payment *Payment) error
	Reset(payment *Payment) error
	IsConfirmed() bool
	IsRejected() bool
	IsPending() bool
	String() string
}

type PendingPaymentStatus struct{}

func (p PendingPaymentStatus) Confirm(payment *Payment) error {
	payment.Status = stateMapper[confirmed]
	return nil
}

func (p PendingPaymentStatus) Reject(payment *Payment) error {
	payment.Status = stateMapper[rejected]
	return nil
}

func (p PendingPaymentStatus) Reset(payment *Payment) error {
	return nil
}

func (p PendingPaymentStatus) IsConfirmed() bool {
	return false
}

func (p PendingPaymentStatus) IsRejected() bool {
	return false
}

func (p PendingPaymentStatus) IsPending() bool {
	return true
}

func (p PendingPaymentStatus) String() string {
	return pending
}

type ConfirmedPaymentStatus struct{}

func (c ConfirmedPaymentStatus) Confirm(payment *Payment) error {
	return errors.New("payment already confirmed")
}

func (c ConfirmedPaymentStatus) Reject(payment *Payment) error {
	return errors.New("payment already confirmed")
}

func (c ConfirmedPaymentStatus) Reset(payment *Payment) error {
	payment.Status = stateMapper[pending]
	return nil
}

func (c ConfirmedPaymentStatus) IsConfirmed() bool {
	return true
}

func (c ConfirmedPaymentStatus) IsRejected() bool {
	return false
}

func (c ConfirmedPaymentStatus) IsPending() bool {
	return false
}

func (c ConfirmedPaymentStatus) String() string {
	return confirmed
}

type RejectedPaymentStatus struct{}

func (r RejectedPaymentStatus) Confirm(payment *Payment) error {
	return errors.New("payment already rejected")
}

func (r RejectedPaymentStatus) Reject(payment *Payment) error {
	return errors.New("payment already rejected")
}

func (r RejectedPaymentStatus) Reset(payment *Payment) error {
	payment.Status = stateMapper[pending]
	return nil
}

func (r RejectedPaymentStatus) IsConfirmed() bool {
	return false
}

func (r RejectedPaymentStatus) IsRejected() bool {
	return true
}

func (r RejectedPaymentStatus) IsPending() bool {
	return false
}

func (r RejectedPaymentStatus) String() string {
	return rejected
}

func NewPaymentStatus(status string) (PaymentStatus, error) {
	paymentStatus, ok := stateMapper[status]
	if !ok {
		return nil, errors.New("invalid payment status")
	}
	return paymentStatus, nil
}
