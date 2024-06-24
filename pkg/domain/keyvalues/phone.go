package keyvalues

import (
	"fmt"
	"regexp"

	"github.com/nyaruka/phonenumbers"
)

type Phone struct {
	value string
}

func NewPhone(value string) (*Phone, error) {
	p, err := phonenumbers.Parse(value, "BR")
	if err != nil {
		return nil, fmt.Errorf("invalid phone")
	}

	phone := regexp.MustCompile(`[()-]`).ReplaceAllString(phonenumbers.Format(p, phonenumbers.E164), "")

	return &Phone{value: phone}, nil
}

func (p Phone) Value() string {
	return p.value
}
