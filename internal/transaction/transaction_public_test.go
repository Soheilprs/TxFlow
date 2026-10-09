package transaction_test

import (
	"errors"
	"testing"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

func TestPublicTransactionAPI(t *testing.T) {
	tx, err := transaction.New(
		"tx-public",
		transaction.AmountCents(5000),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// id, amount, status := transaction.Details(tx)

	if tx.ID() != "tx-public" {
		t.Fatalf(
			"expected ID %q, got %q",
			"tx-public",
			tx.ID(),
		)
	}

	if tx.Amount() != transaction.AmountCents(5000) {
		t.Fatalf(
			"expected amount %d, got %d",
			5000,
			tx.Amount(),
		)
	}

	if tx.Status() != transaction.StatusPending {
		t.Fatalf(
			"expected status %q, got %q",
			transaction.StatusPending,
			tx.ID(),
		)
	}
}

func TestPublicTransactionLifecycle(t *testing.T) {
	tx, err := transaction.New(
		"tx-public",
		transaction.AmountCents(5000),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !tx.IsPending() {
		t.Fatal("expected transaction to be pending")
	}

	if err := tx.MarkProcessed(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tx.Status() != transaction.StatusProcessed {
		t.Fatalf(
			"expected status %q, got %q",
			transaction.StatusProcessed,
			tx.Status(),
		)
	}

	if !tx.Status().IsFinal() {
		t.Fatal("expected processed transaction to be final")
	}
}

func TestPublicErrors(t *testing.T) {
	_, err := transaction.New(
		"",
		transaction.AmountCents(5000),
	)

	if !errors.Is(
		err,
		transaction.ErrEmptyID,
	) {
		t.Fatalf(
			"expected ErrEmptyID, got %v",
			err,
		)
	}
}
