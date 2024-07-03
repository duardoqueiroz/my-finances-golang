package middlewares

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/security"
	"github.com/labstack/echo/v4"
)

var (
	AuthorizationError = "Authorization error"
)

func UserAuthorizer(authEnforcer *casbin.Enforcer, userRepo repositories.UserRepository) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		fn := func(c echo.Context) error {
			role := "anonymous"
			token := c.Request().Header.Get("Authorization")
			var id string
			var err error
			if token != "" {
				id, role, err = parseToken(token)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, &outputs.CustomError{
						Name:    AuthorizationError,
						Message: err.Error(),
					})
				}
			}

			if role != "anonymous" {
				user, err := userRepo.FindByID(id)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, &outputs.CustomError{
						Name:    AuthorizationError,
						Message: err.Error(),
					})
				}
				if user == nil {
					return c.JSON(http.StatusForbidden, &outputs.CustomError{
						Name:    AuthorizationError,
						Message: "Method not allowed for anonymous user",
					})
				}
			}

			result, err := authEnforcer.EnforceSafe(role, c.Request().URL.Path, c.Request().Method)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, &outputs.CustomError{
					Name:    AuthorizationError,
					Message: err.Error(),
				})
			}

			if !result {
				return c.JSON(http.StatusForbidden, &outputs.CustomError{
					Name:    AuthorizationError,
					Message: "Method not allowed",
				})
			}

			return next(c)
		}
		return fn
	}
}

func parseToken(token string) (string, string, error) {
	parsedToken, err := security.ParseAccessToken(token)
	if err != nil {
		return "", "", fmt.Errorf("error parsing token: %w", err)
	}
	return parsedToken.Id, parsedToken.Role, nil
}
