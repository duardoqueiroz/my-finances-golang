package user

type Delete interface {
	Execute(id string) error
}
