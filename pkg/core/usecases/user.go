package usecases

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/usecases"
)

type UserUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{repo: userRepo}
}

func (u UserUseCase) Create(input usecases.CreateUserInput) (*usecases.CreateUserOutput, error) {
	user, err := entities.NewUser(input.Name, input.Email, input.CPF, input.Phone)
	if err != nil {
		return nil, err
	}
	err = u.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u UserUseCase) Update(id string, input usecases.UpdateUserInput) error {
	return nil
}

func (u UserUseCase) FindAll() ([]usecases.FindAllUserOutput, error) {
	return []usecases.FindAllUserOutput{}, nil
}

func (u UserUseCase) FindByID(id string) (*usecases.FindUserByIDOutput, error) {
	return nil, nil
}

func (u UserUseCase) Delete(id string) error {

	return nil
}
