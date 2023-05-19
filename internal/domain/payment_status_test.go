package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaymentStatus(t *testing.T) {
	t.Run("Confirm", func(t *testing.T) {
		confirmPendingStatus(t)
		confirmConfirmedStatus(t)
		confirmRejectedStatus(t)
	})
	t.Run("Reject", func(t *testing.T) {
		rejectPendingStatus(t)
		rejectConfirmedStatus(t)
		rejectRejectedStatus(t)
	})
	t.Run("Reset", func(t *testing.T) {
		resetPendingStatus(t)
		resetConfirmedStatus(t)
		resetRejectedStatus(t)
	})
}

func confirmPendingStatus(t *testing.T) {
	status, err := NewPaymentStatus(pending)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, pending, payment.Status.String())
	err = status.Confirm(payment)
	require.NoError(t, err)
	require.Equal(t, confirmed, payment.Status.String())
}
func confirmConfirmedStatus(t *testing.T) {
	status, err := NewPaymentStatus(confirmed)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, confirmed, payment.Status.String())
	err = status.Confirm(payment)
	require.NoError(t, err)
	require.Equal(t, confirmed, payment.Status.String())
}
func confirmRejectedStatus(t *testing.T) {
	status, err := NewPaymentStatus(rejected)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, rejected, payment.Status.String())
	err = status.Confirm(payment)
	require.NoError(t, err)
	require.Equal(t, confirmed, payment.Status.String())
}

func rejectPendingStatus(t *testing.T) {
	status, err := NewPaymentStatus(pending)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, pending, payment.Status.String())
	err = status.Reject(payment)
	require.NoError(t, err)
	require.Equal(t, rejected, payment.Status.String())
}

func rejectConfirmedStatus(t *testing.T) {
	status, err := NewPaymentStatus(confirmed)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, confirmed, payment.Status.String())
	err = status.Reject(payment)
	require.Error(t, err)
	require.Equal(t, confirmed, payment.Status.String())
}

func rejectRejectedStatus(t *testing.T) {
	status, err := NewPaymentStatus(rejected)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, rejected, payment.Status.String())
	err = status.Reject(payment)
	require.NoError(t, err)
	require.Equal(t, rejected, payment.Status.String())
}

func resetPendingStatus(t *testing.T) {
	status, err := NewPaymentStatus(pending)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, pending, payment.Status.String())
	err = status.Reset(payment)
	require.NoError(t, err)
	require.Equal(t, pending, payment.Status.String())
}

func resetConfirmedStatus(t *testing.T) {
	status, err := NewPaymentStatus(confirmed)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, confirmed, payment.Status.String())
	err = status.Reset(payment)
	require.NoError(t, err)
	require.Equal(t, pending, payment.Status.String())
}

func resetRejectedStatus(t *testing.T) {
	status, err := NewPaymentStatus(rejected)
	require.NoError(t, err)
	payment := &Payment{
		Status: status,
	}
	require.Equal(t, rejected, payment.Status.String())
	err = status.Reset(payment)
	require.NoError(t, err)
	require.Equal(t, pending, payment.Status.String())
}
