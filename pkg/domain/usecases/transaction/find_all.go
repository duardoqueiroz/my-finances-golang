package bank

import "time"

type FindAllOutput struct {
	ID                string
	Value             float64
	Date              time.Time
	Description       string
	PaymentMethod     string
	PaymentConditions string
	Installments      int
	Recurrence        int
}

type FindAll interface {
	Execute() (*FindAllOutput, error)
}
