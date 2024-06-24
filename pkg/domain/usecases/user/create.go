package user

type CreateUserInput struct {
	Name  string
	Email string
	CPF   string
	Phone string
}

type CreateUserOutput struct {
	ID string
}

type CreateUserUseCase interface {
	Execute(input CreateUserInput) (*CreateUserOutput, error)
}
