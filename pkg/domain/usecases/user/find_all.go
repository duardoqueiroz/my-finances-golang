package user

type FindAllOutput struct {
	ID    string
	Name  string
	Email string
	CPF   string
	Phone string
}

type FindAll interface {
	Execute() ([]FindAllOutput, error)
}
