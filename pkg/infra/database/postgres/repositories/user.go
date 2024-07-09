package repositories

import (
	"errors"
	"fmt"

	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database/postgres/dtos"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database/postgres/queries"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type User struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) *User {
	return &User{conn: conn}
}

func (u User) FindAll() ([]entities.User, error) {
	return []entities.User{}, nil
}

func (u User) FindByID(id string) (*entities.User, error) {
	dest := dtos.User().Select().ById()
	err := u.conn.Get(&dest, queries.User().Select().ById(), id)
	if err != nil {
		return nil, u.handleErrors(err)
	}
	user := dest.ToDomain()
	return user, nil
}

func (u User) FindByEmail(email string) (*entities.User, error) {
	dest := dtos.User().Select().ByEmail()
	err := u.conn.Get(&dest, queries.User().Select().ByEmail(), email)
	if err != nil {
		return nil, u.handleErrors(err)
	}
	user := dest.ToDomain()
	return user, nil
}

func (u User) Create(user *entities.User) (string, error) {
	_, err := u.conn.Exec("INSERT INTO users (id, name, email, cpf, phone, password, role) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.ID(), user.Name(), user.Email(), user.CPF(), user.Phone(), user.PasswordHash(), user.Role())
	if err != nil {
		return "", u.handleErrors(err)
	}
	return user.ID(), nil
}

func (u User) Update(id string, user *entities.User) error {
	return nil
}

func (u User) Delete(id string) error {
	return nil
}

func (u User) handleErrors(err error) error {
	if err == nil {
		return nil
	}
	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code.Name() {
		case "unique_violation":
			switch pgErr.Constraint {
			case "users_email_key":
				return errors.New("user email already exisits")
			case "users_phone_key":
				return errors.New("user phone already exisits")
			case "users_cpf_key":
				return errors.New("user cpf already exisits")
			default:
				return errors.New(pgErr.Detail)
			}
		}
	}
	return errors.Unwrap(fmt.Errorf("user unknown error: %w", err))
}
