package factories

import (
	"fmt"

	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database/repositories/postgres"
)

func GetRepositories(driver string, database database.Connection) (repositories.UserRepository, error) {
	switch driver {
	case "postgres":
		return postgres.NewUserRepository(database.Connection()), nil
	default:
		return nil, fmt.Errorf("%s driver not found", driver)
	}
}
