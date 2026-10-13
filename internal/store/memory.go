package store

import "github.com/Soheilprs/TxFlow/internal/transaction"

type MemoryStore struct {
	transactions map[string]transaction.Transaction
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		transactions: make(map[string]transaction.Transaction),
	}
}

func (s *MemoryStore) Save(tx transaction.Transaction) {
	s.transactions[tx.ID()] = tx
}

func (s *MemoryStore) Get(id string) (transaction.Transaction, bool) {
	tx, ok := s.transactions[id]

	return tx, ok
}
