package bank

type FindByIdOutput struct {
	ID          string
	Name        string
	Description string
	Nickname    string
	Balance     float64
}

type FindById interface {
	Execute(id string) (*FindByIdOutput, error)
}
