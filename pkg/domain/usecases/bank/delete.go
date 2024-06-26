package bank

type Delete interface {
	Execute(id string) error
}
