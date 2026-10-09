package transaction_test

import (
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
