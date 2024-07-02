package api

import (
	"errors"

	"github.com/casbin/casbin"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/logger"
)

var (
	errInvalidServerInstance = errors.New("invalid server instance")
)

const (
	EchoInstance int = iota
)

func NewServerInstanceFactory(instance int, repositories database.RepositoryHandler, logger logger.Logger, enforcer *casbin.Enforcer) (Server, error) {
	switch instance {
	case EchoInstance:
		return echo.NewEchoServer(repositories, enforcer, logger), nil
	default:
		return nil, errInvalidServerInstance
	}
}
