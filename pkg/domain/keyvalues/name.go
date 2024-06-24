package keyvalues

import "fmt"

type Name struct {
	value string
}

func NewName(value string) (*Name, error) {
	if len(value) < 2 {
		return nil, fmt.Errorf("name must have at least 2 characters")
	}

	if len(value) > 100 {
		return nil, fmt.Errorf("name must have at most 100 characters")
	}
	return &Name{value: value}, nil
}

func (n Name) Value() string {
	return n.value
}
