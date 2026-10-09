package transaction

import "testing"

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

func TestDetails(t *testing.T) {
	tx, err := New(
		"tx-001",
		AmountCents(12549),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	id, amount, status := Details(tx)

	if id != "tx-001" {
		t.Fatalf(
			"expected ID %q, got %q",
			"tx-001",
			id,
		)
	}

	if amount != AmountCents(12549) {
		t.Fatalf(
			"expected amount %d, got %d",
			12549,
			amount,
		)
	}

	if status != StatusPending {
		t.Fatalf(
			"expected status %q, got %q",
			StatusPending,
			status,
		)
	}
}
