package bank

type CreateInput struct {
	Name        string
	Description string
	Nickname    string
	Balance     float64
}

type CreateOutput struct {
	ID string
}

type Create interface {
	Execute(input CreateInput) (*CreateOutput, error)
}
