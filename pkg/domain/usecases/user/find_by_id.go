package user

type FindByIDOutput struct {
	ID    string
	Name  string
	Email string
	CPF   string
	Phone string
}

type FindByID interface {
	Execute(id string) (*FindByIDOutput, error)
}
