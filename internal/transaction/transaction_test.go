package transaction

import "testing"

func TestTransactionZeroValue(t *testing.T) {
	var tx Transaction

	if tx.ID != "" {
		t.Fatalf("expected empty ID, got %q", tx.ID)
	}

	if tx.Amount != 0 {
		t.Fatalf("expected zero amount, got %d", tx.Amount)
	}

	if tx.Status != "" {
		t.Fatalf("expected empty status, got %q", tx.Status)
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

	if tx.ID != "tx-001" {
		t.Fatalf(
			"expected ID %q, got %q",
			"tx-001",
			tx.ID,
		)
	}

	if tx.Amount != AmountCents(12549) {
		t.Fatalf(
			"expected amount %d, got %d",
			12549,
			tx.Amount,
		)
	}

	if tx.Status != StatusPending {
		t.Fatalf(
			"expected status %q, got %q",
			StatusPending,
			tx.Status,
		)
	}
}

func TestNewRejectsEmptyID(t *testing.T) {
	_, err := New(
		"",
		AmountCents(12549),
	)
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestNewRejectsNonPositiveAmount(t *testing.T) {
	_, err := New(
		"tx-001",
		AmountCents(0),
	)
	if err == nil {
		t.Fatal("expected an error")
	}
}
