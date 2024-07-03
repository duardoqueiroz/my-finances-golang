package dtos

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
)

type userDtoManager struct{}
type userDtoSelectManager struct{}

type userDto struct {
	Id    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Phone string `db:"phone"`
	Cpf   string `db:"cpf"`
	Role  string `db:"role"`
}

func User() *userDtoManager {
	return &userDtoManager{}
}

func (u *userDto) ToDomain() *entities.User {
	user := entities.NewExistentUser(u.Id, u.Name, u.Email, u.Cpf, u.Phone, u.Role)
	return user
}

func (*userDtoManager) Select() *userDtoSelectManager {
	return &userDtoSelectManager{}
}

func (userDtoSelectManager) ById() userDto {
	return userDto{}
}
