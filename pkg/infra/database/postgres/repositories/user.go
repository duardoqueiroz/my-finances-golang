package repositories

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
	"github.com/jmoiron/sqlx"
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
	return &entities.User{}, nil
}

func (u User) Create(user *entities.User) (string, error) {
	_, err := u.conn.Exec("INSERT INTO users (id, name, email, cpf, phone, password, role) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.ID(), user.Name(), user.Email(), user.CPF(), user.Phone(), user.Password(), user.Role())
	if err != nil {
		return "", err
	}
	return user.ID(), nil
}

func (u User) Update(id string, user *entities.User) error {
	return nil
}

func (u User) Delete(id string) error {
	return nil
}
