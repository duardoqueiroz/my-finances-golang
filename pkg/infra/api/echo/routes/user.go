package routes

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/usecases"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo/handlers"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo/middlewares"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/labstack/echo/v4"
)

func LoadUserRoutes(group *echo.Group, repoHandler database.RepositoryHandler) {
	userGroup := group.Group("/users")

	userRepo := repoHandler.UserRepository()
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	userGroup.GET("/:id", userHandler.FindById, middlewares.EnsureAuthenticatedUser())
	userGroup.GET("", userHandler.FindAll, middlewares.IsAdmin())
	userGroup.PUT("/:id", userHandler.Update, middlewares.EnsureAuthenticatedUser())
	userGroup.DELETE("/:id", userHandler.Delete, middlewares.EnsureAuthenticatedUser())
}
