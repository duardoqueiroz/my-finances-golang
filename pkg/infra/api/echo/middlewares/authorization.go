package middlewares

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/repositories"
	"github.com/labstack/echo/v4"
)

var (
	AuthorizationError = "Authorization error"
)

func UserAuthorizer(authEnforcer *casbin.Enforcer, userRepo repositories.UserRepository) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		fn := func(c echo.Context) error {
			var id string
			role := "anonymous"
			token, err := getToken(c.Request().Header.Get("Authorization"))
			if err == nil {
				id, role, err = parseToken(token)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, &outputs.CustomError{
						Name:    AuthorizationError,
						Message: err.Error(),
					})
				}
			}

			if role != "anonymous" {
				_, err := userRepo.FindByID(id)
				if err != nil {
					return c.JSON(http.StatusForbidden, &outputs.CustomError{
						Name:    AuthorizationError,
						Message: "Method not allowed",
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
