package main

import (
	"errors"
	"fmt"

	"github.com/Soheilprs/TxFlow/internal/transaction"
)

type transactionInput struct {
	id     string
	amount transaction.AmountCents
}

func main() {
	if err := run(); err != nil {
		fmt.Println("txflow failed:", err)
	}
}

func run() error {
	inputs := []transactionInput{
		{
			id:     "tx-001",
			amount: transaction.AmountCents(12549),
		},
		{
			id:     "tx-002",
			amount: transaction.AmountCents(5000),
		},
		{
			id:     "tx-003",
			amount: transaction.AmountCents(9999),
		},
	}

	transactions := make(
		[]transaction.Transaction,
		0,
		len(inputs),
	)

	for _, input := range inputs {
		tx, err := transaction.New(
			input.id,
			input.amount,
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
					"create transaction %s: %w",
					input.id,
					err,
				)
			}
		}

		transactions = append(
			transactions,
			tx,
		)
	}

	for i := range transactions {
		tx := &transactions[i]

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
	}

	return nil
}
