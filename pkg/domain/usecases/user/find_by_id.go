package user

type FindUserByIDOutput struct {
	ID    string
	Name  string
	Email string
	CPF   string
	Phone string
}

type FindUserByIDUseCase interface {
	Execute(id string) (*FindUserByIDOutput, error)
}
