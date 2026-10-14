package main

import (
	"errors"
	"fmt"

	"github.com/Soheilprs/TxFlow/internal/processor"
	"github.com/Soheilprs/TxFlow/internal/store"
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
	memoryStore := store.NewMemoryStore()

	txProcessor := processor.New(
		memoryStore,
	)

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

	for _, input := range inputs {
		tx, err := txProcessor.Create(
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

		fmt.Printf(
			"before: transaction=%s amount=%d status=%s\n",
			tx.ID(),
			tx.Amount(),
			tx.Status(),
		)

		processedTx, err := txProcessor.Process(
			tx.ID(),
		)
		if err != nil {
			return fmt.Errorf(
				"process transaction %s: %w",
				tx.ID(),
				err,
			)
		}

		fmt.Printf(
			"after: transaction=%s amount=%d status=%s\n",
			processedTx.ID(),
			processedTx.Amount(),
			processedTx.Status(),
		)
	}

	return nil
}
