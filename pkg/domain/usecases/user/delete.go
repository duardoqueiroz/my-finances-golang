package user

type DeleteUserUseCase interface {
	Execute(id string) error
}
