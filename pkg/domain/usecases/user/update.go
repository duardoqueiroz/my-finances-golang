package user

type UpdateUserInput struct {
	Name  string
	Email string
	CPF   string
	Phone string
}

type UpdateUserUseCase interface {
	Execute(id string, input UpdateUserInput) error
}
