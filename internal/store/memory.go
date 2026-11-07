package store

import (
	"sync"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

type MemoryStore struct {
	mu           sync.Mutex
	transactions map[string]transaction.Transaction
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		transactions: make(map[string]transaction.Transaction),
	}
}

func (s *MemoryStore) Save(tx transaction.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.transactions[tx.ID()] = tx
}

func (s *MemoryStore) Get(id string) (transaction.Transaction, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx, ok := s.transactions[id]

	return tx, ok
}
