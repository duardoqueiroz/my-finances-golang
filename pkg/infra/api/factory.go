package api

import (
	"errors"

	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
)

var (
	errInvalidServerInstance = errors.New("invalid server instance")
)

const (
	EchoInstance int = iota
)

func NewServerInstanceFactory(instance int, repositories database.RepositoryHandler) (Server, error) {
	switch instance {
	case EchoInstance:
		return echo.NewEchoServer(repositories), nil
	default:
		return nil, errInvalidServerInstance
	}
}
