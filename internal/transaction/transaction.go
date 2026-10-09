package transaction

import "errors"

type Status string
type AmountCents int64

const (
	StatusPending   Status = "pending"
	StatusProcessed Status = "processed"
	StatusFailed    Status = "failed"
)

type Transaction struct {
	id     string
	amount AmountCents
	status Status
}

func New(id string, amount AmountCents) (Transaction, error) {
	if id == "" {
		return Transaction{}, errors.New("transaction ID must not be empty")
	}

	if amount <= 0 {
		return Transaction{}, errors.New("transaction amount must be greater than zero")
	}

	tx := Transaction{
		id:     id,
		amount: amount,
		status: StatusPending,
	}

	return tx, nil
}

func (t Transaction) ID() string {
	return t.id
}

func (t Transaction) Amount() AmountCents {
	return t.amount
}

func (t Transaction) Status() Status {
	return t.status
}

func (t Transaction) IsPending() bool {
	return t.status == StatusPending
}

func (s Status) IsFinal() bool {
	return s == StatusProcessed || s == StatusFailed
}
