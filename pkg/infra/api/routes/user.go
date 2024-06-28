package routes

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/core/usecases"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/handlers"
	"github.com/labstack/echo/v4"
)

func loadUserRoutes(group *echo.Group, userRepo repositories.UserRepository) {
	userGroup := group.Group("/users")

	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	userGroup.POST("", userHandler.Create)
}
