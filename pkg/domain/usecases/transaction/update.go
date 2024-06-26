package bank

type UpdateInput struct {
	Name        string
	Description string
	Nickname    string
	Balance     float64
}

type Update interface {
	Execute(id string, input UpdateInput) error
}
