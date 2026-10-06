package main

import (
	"fmt"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

func main() {
	tx := transaction.Transaction{
		ID:          "tx-001",
		AmountCents: 12549,
		Status:      transaction.StatusPending,
	}

	fmt.Printf(
		"transaction=%s amount=%d status=%s\n",
		tx.ID,
		tx.AmountCents,
		tx.Status,
	)
}
