package postgres

import (
	domainRepo "github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database/postgres/repositories"
)

type PostgresRepositoryHandler struct {
	database *database
}

func NewPostgresRepositoryHandler(database *database) *PostgresRepositoryHandler {
	return &PostgresRepositoryHandler{
		database: database,
	}
}

func (p *PostgresRepositoryHandler) UserRepository() domainRepo.UserRepository {
	return *repositories.NewUserRepository(p.database.conn)
}
