package dtos

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
)

type userDtoManager struct{}
type userDtoSelectManager struct{}

type userDto struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
	Cpf      string `db:"cpf"`
	Role     string `db:"role"`
}

func User() *userDtoManager {
	return &userDtoManager{}
}

func (u *userDto) ToDomain() *entities.User {
	user := entities.NewExistentUser(u.Id, u.Name, u.Email, u.Cpf, u.Phone, u.Role, u.Password)
	return user
}

func (*userDtoManager) Select() *userDtoSelectManager {
	return &userDtoSelectManager{}
}

func (userDtoSelectManager) ById() userDto {
	return userDto{}
}

func (userDtoSelectManager) ByEmail() userDto {
	return userDto{}
}

func (userDtoManager) Update(id string, user *entities.User) []interface{} {
	return []interface{}{
		user.Name(),
		user.Email(),
		user.Phone(),
		user.CPF(),
		user.PasswordHash(),
		user.Role(),
		id,
	}
}

func (userDtoManager) Create(user *entities.User) []interface{} {
	return []interface{}{
		user.ID(),
		user.Name(),
		user.Email(),
		user.CPF(),
		user.Phone(),
		user.PasswordHash(),
		user.Role(),
	}
}
