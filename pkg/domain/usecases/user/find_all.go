package user

type FindAllUserOutput struct {
	ID    string
	Name  string
	Email string
	CPF   string
	Phone string
}

type FindAllUserUseCase interface {
	Execute() ([]FindAllUserOutput, error)
}
