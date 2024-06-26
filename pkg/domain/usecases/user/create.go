package user

type CreateInput struct {
	Name  string
	Email string
	CPF   string
	Phone string
}

type CreateOutput struct {
	ID string
}

type Create interface {
	Execute(input CreateInput) (*CreateOutput, error)
}
