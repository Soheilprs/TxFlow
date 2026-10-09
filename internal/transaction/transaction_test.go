package transaction

import (
	"errors"
	"testing"
)

func TestTransactionZeroValue(t *testing.T) {
	var tx Transaction

	if tx.id != "" {
		t.Fatalf("expected empty ID, got %q", tx.id)
	}

	if tx.amount != 0 {
		t.Fatalf("expected zero amount, got %d", tx.amount)
	}

	if tx.status != "" {
		t.Fatalf("expected empty status, got %q", tx.status)
	}
}

func TestNew(t *testing.T) {
	tx, err := New(
		"tx-001",
		AmountCents(12549),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tx.id != "tx-001" {
		t.Fatalf(
			"expected ID %q, got %q",
			"tx-001",
			tx.id,
		)
	}

	if tx.amount != AmountCents(12549) {
		t.Fatalf(
			"expected amount %d, got %d",
			12549,
			tx.amount,
		)
	}

	if tx.status != StatusPending {
		t.Fatalf(
			"expected status %q, got %q",
			StatusPending,
			tx.status,
		)
	}
}

func TestDetails(t *testing.T) {
	tx, err := New(
		"tx-001",
		AmountCents(12549),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// id, amount, status := Details(tx)

	if tx.ID() != "tx-001" {
		t.Fatalf(
			"expected ID %q, got %q",
			"tx-001",
			tx.ID(),
		)
	}

	if tx.Amount() != AmountCents(12549) {
		t.Fatalf(
			"expected amount %d, got %d",
			12549,
			tx.Amount(),
		)
	}

	if tx.Status() != StatusPending {
		t.Fatalf(
			"expected status %q, got %q",
			StatusPending,
			tx.Status(),
		)
	}
}

func TestStatusIsFinal(t *testing.T) {
	if StatusPending.IsFinal() {
		t.Fatal("expected pending status not to be final")
	}

	if !StatusProcessed.IsFinal() {
		t.Fatal("expected processed status to be final")
	}

	if !StatusFailed.IsFinal() {
		t.Fatal("expected failed status to be final")
	}
}

func TestMarkProcessed(t *testing.T) {
	tx, err := New(
		"tx-001",
		AmountCents(12549),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = tx.MarkProcessed()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tx.Status() != StatusProcessed {
		t.Fatalf(
			"expected status %q, got %q",
			StatusProcessed,
			tx.Status(),
		)
	}
}

func TestMarkFailed(t *testing.T) {
	tx, err := New(
		"tx-001",
		AmountCents(12549),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = tx.MarkFailed()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tx.Status() != StatusFailed {
		t.Fatalf(
			"expected status %q, got %q",
			StatusFailed,
			tx.Status(),
		)
	}
}

func TestNewRejectsEmptyID(t *testing.T) {
	_, err := New(
		"",
		AmountCents(12549),
	)

	if !errors.Is(err, ErrEmptyID) {
		t.Fatalf(
			"expected ErrEmptyID, got %v",
			err,
		)
	}
}

func TestNewRejectsNonPositiveAmount(
	t *testing.T,
) {
	_, err := New(
		"tx-001",
		AmountCents(0),
	)

	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf(
			"expected ErrInvalidAmount, got %v",
			err,
		)
	}
}

func TestMarkProcessedRejectsFinalTransaction(
	t *testing.T,
) {
	tx, err := New(
		"tx-001",
		AmountCents(12549),
	)
	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if err := tx.MarkProcessed(); err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	err = tx.MarkProcessed()

	if !errors.Is(err, ErrFinalState) {
		t.Fatalf(
			"expected ErrFinalState, got %v",
			err,
		)
	}
}
