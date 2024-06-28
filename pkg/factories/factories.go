package factories

import "github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"

type RepositoryFactory interface {
	UserRepository() repositories.UserRepository
}

type repositoryFactory struct {
	userRepository repositories.UserRepository
}

func NewRepositoryFactory(userRepository repositories.UserRepository) RepositoryFactory {
	return &repositoryFactory{
		userRepository: userRepository,
	}
}

func (r *repositoryFactory) UserRepository() repositories.UserRepository {
	return r.userRepository
}
