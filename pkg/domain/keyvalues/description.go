package keyvalues

import "fmt"

type Description struct {
	value string
}

func NewDescription(value string) (*Description, error) {
	if len(value) > 300 {
		return nil, fmt.Errorf("description must have at most 300 characters")
	}
	return &Description{value: value}, nil
}

func (n Description) Value() string {
	return n.value
}
