package keyvalues

import (
	"fmt"
	"net/mail"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	email, err := mail.ParseAddress(value)
	if err != nil {
		return nil, fmt.Errorf("invalid email")
	}

	return &Email{value: email.Address}, nil
}

func (e Email) Value() string {
	return e.value
}
