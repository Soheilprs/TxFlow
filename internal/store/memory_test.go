package store

import (
	"testing"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

func TestMemoryStoreSaveAndGet(t *testing.T) {
	s := NewMemoryStore()

	tx, err := transaction.New(
		"tx-001",
		transaction.AmountCents(12549),
	)
	if err != nil {
		t.Fatalf(
			"create transaction: %v",
			err,
		)
	}

	s.Save(tx)

	stored, ok := s.Get("tx-001")
	if !ok {
		t.Fatal(
			"expected transaction to exist",
		)
	}

	if stored.ID() != tx.ID() {
		t.Fatalf(
			"expected ID %q, got %q",
			tx.ID(),
			stored.ID(),
		)
	}
}

func TestMemoryStoreGetMissingTransaction(
	t *testing.T,
) {
	s := NewMemoryStore()

	_, ok := s.Get("missing")

	if ok {
		t.Fatal(
			"expected transaction not to exist",
		)
	}
}

func TestMemoryStoreStoresTransactionValue(
	t *testing.T,
) {
	s := NewMemoryStore()

	tx, err := transaction.New(
		"tx-001",
		transaction.AmountCents(12549),
	)
	if err != nil {
		t.Fatalf(
			"create transaction: %v",
			err,
		)
	}

	s.Save(tx)

	if err := tx.MarkProcessed(); err != nil {
		t.Fatalf(
			"process transaction: %v",
			err,
		)
	}

	stored, ok := s.Get("tx-001")
	if !ok {
		t.Fatal(
			"expected transaction to exist",
		)
	}

	if stored.Status() != transaction.StatusPending {
		t.Fatalf(
			"expected stored status %q, got %q",
			transaction.StatusPending,
			stored.Status(),
		)
	}
}
