package usecases

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
)

type Auth interface {
	Login(input inputs.Login) (*outputs.Login, error)
	SignUp(input inputs.SignUp) (*outputs.SignUp, error)
}
