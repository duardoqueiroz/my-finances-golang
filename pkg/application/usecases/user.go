package usecases

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/security"
)

type UserUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{repo: userRepo}
}

func (u UserUseCase) Create(input inputs.CreateUserInput) (*outputs.CreateUserOutput, error) {
	user, err := entities.NewUser(input.Name, input.Email, input.CPF, input.Phone, input.Password)
	if err != nil {
		return nil, err
	}
	id, err := u.repo.Create(user)
	if err != nil {
		return nil, err
	}
	claims := security.UserClaims{
		Id:   id,
		Role: user.Role(),
	}
	jwtToken, err := security.NewAccessToken(claims)
	if err != nil {
		return nil, err
	}
	return &outputs.CreateUserOutput{ID: id, Token: jwtToken}, nil
}

func (u UserUseCase) Update(id string, input inputs.UpdateUserInput) error {
	return nil
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

	return nil
}
