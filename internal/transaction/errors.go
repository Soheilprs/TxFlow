package transaction

import "errors"

var (
	ErrEmptyID = errors.New(
		"transaction ID must not be empty",
	)

	ErrInvalidAmount = errors.New(
		"transaction amount must be greater than zero",
	)

	ErrFinalState = errors.New(
		"transaction is already in a final state",
	)
)
