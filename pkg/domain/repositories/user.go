package repositories

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindByID(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Create(user *entities.User) (string, error)
	Update(id string, user *entities.User) error
	Delete(id string) error
}
