package main

import (
	"fmt"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

func main() {
	tx, err := transaction.New(
		"tx-001",
		transaction.AmountCents(12549),
	)
	if err != nil {
		fmt.Println("failed to create transaction:", err)
		return
	}

	fmt.Printf(
		"transaction=%s amount=%d status=%s\n",
		tx.ID,
		tx.Amount,
		tx.Status,
	)
}
