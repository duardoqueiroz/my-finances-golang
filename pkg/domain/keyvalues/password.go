package keyvalues

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	if !isPasswordSecure(value) {
		return nil, fmt.Errorf("password must have at least 8 characters, one lowercase letter, one uppercase letter, one number and one special character")
	}

	hash, err := hashPassword(value)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	return &Password{value: hash}, nil
}

func NewExistentPassword(value string) *Password {
	return &Password{value: value}
}

func hashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p Password) Value() string {
	return p.value
}

func (p Password) Compare(pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(pwd))
	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func isPasswordSecure(pwd string) bool {
	secure := true
	tests := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, pwd)
		if !t {
			secure = false
			break
		}
	}
	return secure
}
