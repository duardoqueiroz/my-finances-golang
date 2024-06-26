package keyvalues

import "fmt"

type TransactionValue struct {
	value float64
}

func NewTransactionValue(value float64) (*TransactionValue, error) {
	if value < 0.01 {
		return nil, fmt.Errorf("transaction value must be greater than 0.01")
	}
	return &TransactionValue{
		value: value,
	}, nil
}

func (t TransactionValue) Value() float64 {
	return t.value
}
