package postgres

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) *UserRepository {
	return &UserRepository{conn: conn}
}

func (u UserRepository) FindAll() ([]entities.User, error) {
	return []entities.User{}, nil
}

func (u UserRepository) FindByID(id string) (*entities.User, error) {
	return &entities.User{}, nil
}

func (u UserRepository) Create(user *entities.User) error {
	return nil
}

func (u UserRepository) Update(id string, user *entities.User) error {
	return nil
}

func (u UserRepository) Delete(id string) error {
	return nil
}
