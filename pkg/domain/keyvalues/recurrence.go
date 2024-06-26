package keyvalues

import "fmt"

type Recurrence struct {
	value int
}

func NewRecurrence(value int) (*Recurrence, error) {
	if value < 0 {
		return nil, fmt.Errorf("recurrence must be greater than 1")
	}
	return &Recurrence{
		value: value,
	}, nil
}

func (r Recurrence) Value() int {
	return r.value
}
