package bank

import "time"

type CreateInput struct {
	Value             float64
	Date              time.Time
	Description       string
	PaymentMethod     string
	PaymentConditions string
	Installments      int
	Recurrence        int
}

type CreateOutput struct {
	ID string
}

type Create interface {
	Execute(input CreateInput) (*CreateOutput, error)
}
