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
