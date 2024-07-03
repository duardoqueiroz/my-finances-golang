package keyvalues

import (
	"fmt"
	"regexp"
	"strings"

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

	phone := regexp.MustCompile(`[()-]`).ReplaceAllString(phonenumbers.Format(p, phonenumbers.NATIONAL), "")
	phone = strings.ReplaceAll(phone, " ", "")
	return &Phone{value: phone}, nil
}

func NewExistentPhone(value string) *Phone {
	return &Phone{value: value}
}

func (p Phone) Value() string {
	return p.value
}
