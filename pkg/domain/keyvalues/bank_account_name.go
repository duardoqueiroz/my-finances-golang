package keyvalues

import "fmt"

type BankAccountName struct {
	value string
}

func NewBankAccountName(value string) (*BankAccountName, error) {
	if len(value) < 2 {
		return nil, fmt.Errorf("bank account name must have at least 2 characters")
	}

	if len(value) > 100 {
		return nil, fmt.Errorf("bank account name must have at most 100 characters")
	}
	return &BankAccountName{value: value}, nil
}

func (n BankAccountName) Value() string {
	return n.value
}
