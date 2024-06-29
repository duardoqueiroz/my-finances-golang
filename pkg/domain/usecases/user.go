package usecases

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
)

type User interface {
	Create(input inputs.CreateUserInput) (*outputs.CreateUserOutput, error)
	Update(id string, input inputs.UpdateUserInput) error
	FindAll() ([]outputs.FindAllUserOutput, error)
	FindByID(id string) (*outputs.FindUserByIDOutput, error)
	Delete(id string) error
}
