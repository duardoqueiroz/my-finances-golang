package database

import (
	"errors"

	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database/postgres"
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

const (
	PostgresInstance int = iota
)

func NewSQLDatabaseFactory(instance int) (Connection, RepositoryHandler, error) {
	switch instance {
	case PostgresInstance:
		pgDb, err := postgres.NewPostgresDB()
		if err != nil {
			return nil, nil, err
		}
		repositoryHandler := postgres.NewPostgresRepositoryHandler(pgDb)
		return pgDb, repositoryHandler, nil
	default:
		return nil, nil, errInvalidSQLDatabaseInstance
	}
}
