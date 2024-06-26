package entities

import (
	"fmt"

	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/keyvalues"
	"github.com/google/uuid"
)

type User struct {
	id    string
	name  *keyvalues.Name
	email *keyvalues.Email
	cpf   *keyvalues.CPF
	phone *keyvalues.Phone
}

func NewUser(name, email, cpf, phone string) (*User, error) {
	nameValue, err := keyvalues.NewName(name)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	emailValue, err := keyvalues.NewEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	cpfValue, err := keyvalues.NewCPF(cpf)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	phoneValue, err := keyvalues.NewPhone(phone)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	id := uuid.New().String()

	return &User{
		id:    id,
		name:  nameValue,
		email: emailValue,
		cpf:   cpfValue,
		phone: phoneValue,
	}, nil
}

func (u User) ID() string {
	return u.id
}

func (u User) Name() string {
	return u.name.Value()
}

func (u User) Email() string {
	return u.email.Value()
}

func (u User) Phone() string {
	return u.phone.Value()
}

func (u User) CPF() string {
	return u.cpf.Value()
}
