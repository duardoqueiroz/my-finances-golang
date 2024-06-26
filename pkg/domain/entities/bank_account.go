package entities

import (
	"fmt"

	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/keyvalues"
	"github.com/google/uuid"
)

type BankAccount struct {
	id          string
	name        *keyvalues.BankAccountName
	balance     float64
	nickname    *keyvalues.NickName
	description *keyvalues.Description
}

func NewBankAccount(name, nickname, description string, balance float64) (*BankAccount, error) {
	nameValue, err := keyvalues.NewBankAccountName(name)
	if err != nil {
		return nil, fmt.Errorf("error creating bank account: %w", err)
	}

	nicknameValue, err := keyvalues.NewNickName(nickname)
	if err != nil {
		return nil, fmt.Errorf("error creating bank account: %w", err)
	}

	descriptionValue, err := keyvalues.NewDescription(description)
	if err != nil {
		return nil, fmt.Errorf("error creating bank account: %w", err)
	}

	id := uuid.New().String()
	return &BankAccount{
		id:          id,
		name:        nameValue,
		balance:     balance,
		nickname:    nicknameValue,
		description: descriptionValue,
	}, nil
}

func (b BankAccount) Name() string {
	return b.name.Value()
}

func (b BankAccount) Nickname() string {
	return b.nickname.Value()
}

func (b BankAccount) Description() string {
	return b.description.Value()
}

func (b BankAccount) Balance() string {
	return b.Balance()
}

func (b BankAccount) Id() string {
	return b.id
}
