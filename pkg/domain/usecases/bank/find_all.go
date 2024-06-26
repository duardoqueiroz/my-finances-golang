package bank

type FindAllOutput struct {
	ID          string
	Name        string
	Description string
	Nickname    string
	Balance     float64
}

type FindAll interface {
	Execute() (*FindAllOutput, error)
}
