package routes

import (
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/usecases"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/api/echo/handlers"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/database"
	"github.com/labstack/echo/v4"
)

func LoadAuthRoutes(group *echo.Group, repoHandler database.RepositoryHandler) {
	userGroup := group.Group("/auth")

	userRepo := repoHandler.UserRepository()
	authUseCase := usecases.NewAuthUseCase(userRepo)
	authHandler := handlers.NewAuthHandler(authUseCase)

	userGroup.POST("/login", authHandler.Login)
	userGroup.POST("/logout", authHandler.Logout)
	userGroup.POST("/signup", authHandler.SignUp)
}
