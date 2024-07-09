package usecases

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/entities"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/security"
)

type Auth struct {
	userRepository repositories.UserRepository
}

func NewAuthUseCase(userRepo repositories.UserRepository) *Auth {
	return &Auth{userRepository: userRepo}
}

func (a *Auth) Login(input inputs.Login) (*outputs.Login, error) {
	user, err := a.userRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if err := user.Password().Compare(input.Password); err != nil {
		return nil, err
	}

	claims := &security.UserClaims{
		Id:   user.ID(),
		Role: user.Role(),
	}
	token, expires, err := security.NewAccessToken(*claims)
	if err != nil {
		return nil, err
	}

	return &outputs.Login{
		Id:   user.ID(),
		Name: user.Name(),
		Role: user.Role(),
		Token: &outputs.Token{
			Value:   token,
			Expires: expires,
		},
	}, nil
}

func (a Auth) SignUp(input inputs.SignUp) (*outputs.SignUp, error) {
	user, err := entities.NewUser(input.Name, input.Email, input.CPF, input.Phone, input.Password)
	if err != nil {
		return nil, err
	}
	id, err := a.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	claims := security.UserClaims{
		Id:   id,
		Role: user.Role(),
	}
	token, expires, err := security.NewAccessToken(claims)
	if err != nil {
		return nil, err
	}

	return &outputs.SignUp{ID: id, Token: &outputs.Token{
		Value:   token,
		Expires: expires,
	}}, nil
}
