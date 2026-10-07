package main

import (
	"fmt"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

func main() {
	tx := transaction.Transaction{
		ID:     "tx-001",
		Amount: transaction.AmountCents(12549),
		Status: transaction.StatusPending,
	}

	fmt.Printf(
		"transaction=%s amount=%d status=%s\n",
		tx.ID,
		tx.Amount,
		tx.Status,
	)
}
