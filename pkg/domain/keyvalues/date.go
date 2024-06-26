package keyvalues

import (
	"time"
)

type Date struct {
	value time.Time
}

func NewDate(date time.Time) (*Date, error) {
	return &Date{
		value: date,
	}, nil
}

func (d Date) Value() time.Time {
	return d.value
}
