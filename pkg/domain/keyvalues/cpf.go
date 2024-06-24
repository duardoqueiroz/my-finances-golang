package keyvalues

import (
	"fmt"
	"regexp"

	"github.com/klassmann/cpfcnpj"
)

type CPF struct {
	value string
}

func NewCPF(value string) (*CPF, error) {
	cpf := cpfcnpj.NewCPF(value)
	if !cpf.IsValid() {
		return nil, fmt.Errorf("invalid cpf")
	}

	cpfValue := regexp.MustCompile(`[.-]`).ReplaceAllString(cpf.String(), "")

	return &CPF{value: cpfValue}, nil
}

func (c CPF) Value() string {
	return c.value
}
