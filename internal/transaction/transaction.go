package transaction

type Status string
type AmountCents int64

const (
	StatusPending   Status = "pending"
	StatusProcessed Status = "processed"
	StatusFailed    Status = "failed"
)

type Transaction struct {
	ID     string
	Amount AmountCents
	Status Status
}
