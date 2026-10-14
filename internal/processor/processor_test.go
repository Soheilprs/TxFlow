package processor

import (
	"errors"
	"testing"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

type fakeStore struct {
	transactions map[string]transaction.Transaction
}

func newFakeStore() *fakeStore {
	return &fakeStore{
		transactions: make(
			map[string]transaction.Transaction,
		),
	}
}

func (f *fakeStore) Get(
	id string,
) (transaction.Transaction, bool) {
	tx, ok := f.transactions[id]

	return tx, ok
}

func (f *fakeStore) Save(
	tx transaction.Transaction,
) {
	f.transactions[tx.ID()] = tx
}

func TestProcessorProcessesTransaction(
	t *testing.T,
) {
	fake := newFakeStore()

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

	fake.Save(tx)

	p := New(fake)

	processed, err := p.Process("tx-001")
	if err != nil {
		t.Fatalf(
			"process transaction: %v",
			err,
		)
	}

	if processed.Status() != transaction.StatusProcessed {
		t.Fatalf(
			"expected status %q, got %q",
			transaction.StatusProcessed,
			processed.Status(),
		)
	}

	stored, ok := fake.Get("tx-001")
	if !ok {
		t.Fatal(
			"expected transaction to exist",
		)
	}

	if stored.Status() != transaction.StatusProcessed {
		t.Fatalf(
			"expected stored status %q, got %q",
			transaction.StatusProcessed,
			stored.Status(),
		)
	}
}

func TestProcessorRejectsMissingTransaction(
	t *testing.T,
) {
	fake := newFakeStore()

	p := New(fake)

	_, err := p.Process("missing")

	if !errors.Is(
		err,
		ErrTransactionNotFound,
	) {
		t.Fatalf(
			"expected ErrTransactionNotFound, got %v",
			err,
		)
	}
}

func TestProcessorCreatesTransaction(
	t *testing.T,
) {
	fake := newFakeStore()

	p := New(fake)

	tx, err := p.Create(
		"tx-001",
		transaction.AmountCents(12549),
	)
	if err != nil {
		t.Fatalf(
			"create transaction: %v",
			err,
		)
	}

	if tx.ID() != "tx-001" {
		t.Fatalf(
			"expected ID %q, got %q",
			"tx-001",
			tx.ID(),
		)
	}

	stored, ok := fake.Get("tx-001")
	if !ok {
		t.Fatal(
			"expected transaction to be stored",
		)
	}

	if stored.ID() != tx.ID() {
		t.Fatalf(
			"expected stored ID %q, got %q",
			tx.ID(),
			stored.ID(),
		)
	}
}

func TestProcessorRejectsInvalidAmount(
	t *testing.T,
) {
	fake := newFakeStore()

	p := New(fake)

	_, err := p.Create(
		"tx-001",
		transaction.AmountCents(0),
	)

	if !errors.Is(
		err,
		transaction.ErrInvalidAmount,
	) {
		t.Fatalf(
			"expected ErrInvalidAmount, got %v",
			err,
		)
	}

	_, ok := fake.Get("tx-001")
	if ok {
		t.Fatal(
			"expected invalid transaction not to be stored",
		)
	}
}
