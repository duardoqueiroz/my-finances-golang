package keyvalues

import "fmt"

type NickName struct {
	value string
}

func NewNickName(value string) (*NickName, error) {
	if len(value) < 2 {
		return nil, fmt.Errorf("nickname must have at least 2 characters")
	}

	if len(value) > 30 {
		return nil, fmt.Errorf("nickname must have at most 30 characters")
	}
	return &NickName{value: value}, nil
}

func (n NickName) Value() string {
	return n.value
}
