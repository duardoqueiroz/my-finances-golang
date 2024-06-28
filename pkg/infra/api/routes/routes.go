package routes

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/factories"
	"github.com/labstack/echo/v4"
)

func LoadRoutes(repoFactory factories.RepositoryFactory) *echo.Echo {
	router := echo.New()

	apiGroup := router.Group("/api")

	loadUserRoutes(apiGroup, repoFactory.UserRepository())
	return router
}
