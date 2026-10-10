package main

import (
	"errors"
	"fmt"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("txflow failed:", err)
	}
}

func run() error {
	tx, err := transaction.New(
		"tx-001",
		transaction.AmountCents(12549),
	)
	if err != nil {
		switch {
		case errors.Is(
			err,
			transaction.ErrEmptyID,
		):
			return fmt.Errorf(
				"invalid transaction ID: %w",
				err,
			)

		case errors.Is(
			err,
			transaction.ErrInvalidAmount,
		):
			return fmt.Errorf(
				"invalid transaction amount: %w",
				err,
			)

		default:
			return fmt.Errorf(
				"create transaction: %w",
				err,
			)
		}
	}

	fmt.Printf(
		"before: transaction=%s amount=%d status=%s\n",
		tx.ID(),
		tx.Amount(),
		tx.Status(),
	)

	if err := tx.MarkProcessed(); err != nil {
		return fmt.Errorf(
			"process transaction %s: %w",
			tx.ID(),
			err,
		)
	}

	fmt.Printf(
		"after: transaction=%s amount=%d status=%s\n",
		tx.ID(),
		tx.Amount(),
		tx.Status(),
	)

	return nil
}
