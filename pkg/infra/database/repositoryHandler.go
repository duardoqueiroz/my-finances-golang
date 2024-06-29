package database

import "github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"

type RepositoryHandler interface {
	UserRepository() repositories.UserRepository
}
