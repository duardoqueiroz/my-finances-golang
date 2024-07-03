package entities

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/keyvalues"
	"github.com/google/uuid"
)

type User struct {
	id       string
	name     *keyvalues.Name
	email    *keyvalues.Email
	cpf      *keyvalues.CPF
	phone    *keyvalues.Phone
	password *keyvalues.Password
	role     *keyvalues.Role
}

func NewUser(name, email, cpf, phone, password string) (*User, error) {
	nameValue, err := keyvalues.NewName(name)
	if err != nil {
		return nil, err
	}

	emailValue, err := keyvalues.NewEmail(email)
	if err != nil {
		return nil, err
	}

	cpfValue, err := keyvalues.NewCPF(cpf)
	if err != nil {
		return nil, err
	}

	phoneValue, err := keyvalues.NewPhone(phone)
	if err != nil {
		return nil, err
	}

	passwordValue, err := keyvalues.NewPassword(password)
	if err != nil {
		return nil, err
	}

	// Default value for user role: member
	roleValue, err := keyvalues.NewRole(keyvalues.MemberRole)
	if err != nil {
		return nil, err
	}

	return &User{
		id:       uuid.New().String(),
		name:     nameValue,
		email:    emailValue,
		cpf:      cpfValue,
		phone:    phoneValue,
		password: passwordValue,
		role:     roleValue,
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

func (u User) CPF() string {
	return u.cpf.Value()
}

func (u User) Phone() string {
	return u.phone.Value()
}

func (u User) Password() string {
	return u.password.Value()
}

func (u User) Role() string {
	return u.role.Value()
}
