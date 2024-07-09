package usecases

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
)

type UserUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{repo: userRepo}
}

func (u UserUseCase) Update(id string, input inputs.UpdateUserInput) error {
	user, err := entities.NewUser(input.Name, input.Email, input.CPF, input.Phone, input.Password)
	if err != nil {
		return err
	}
	err = u.repo.Update(id, user)
	return err
}

func (u UserUseCase) FindAll() ([]outputs.FindAllUserOutput, error) {
	return []outputs.FindAllUserOutput{}, nil
}

func (u UserUseCase) FindByID(id string) (*outputs.FindUserByIDOutput, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &outputs.FindUserByIDOutput{
		ID:    user.ID(),
		Name:  user.Name(),
		Email: user.Email(),
		CPF:   user.CPF(),
		Phone: user.Phone(),
		Role:  user.Role(),
	}, nil
}

func (u UserUseCase) Delete(id string) error {
	if _, err := u.repo.FindByID(id); err != nil {
		return err
	}
	err := u.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
