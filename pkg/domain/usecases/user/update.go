package user

type UpdateInput struct {
	Name  string
	Email string
	CPF   string
	Phone string
}

type Update interface {
	Execute(id string, input UpdateInput) error
}
