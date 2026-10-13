package processor

import (
	"fmt"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

type TransactionStore interface {
	Get(id string) (transaction.Transaction, bool)
	Save(tx transaction.Transaction)
}

type Processor struct {
	store TransactionStore
}

func New(store TransactionStore) *Processor {
	return &Processor{
		store: store,
	}
}

func (p *Processor) Process(
	id string,
) (transaction.Transaction, error) {
	tx, ok := p.store.Get(id)
	if !ok {
		return transaction.Transaction{},
			fmt.Errorf(
				"load transaction %s: %w",
				id,
				ErrTransactionNotFound,
			)
	}

	if err := tx.MarkProcessed(); err != nil {
		return transaction.Transaction{},
			fmt.Errorf(
				"process transaction %s: %w",
				id,
				err,
			)
	}

	p.store.Save(tx)

	return tx, nil
}
