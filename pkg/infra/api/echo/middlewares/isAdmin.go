package middlewares

import (
	"net/http"

	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/labstack/echo/v4"
)

func IsAdmin() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		fn := func(c echo.Context) error {
			token, err := getToken(c.Request().Header.Get("Authorization"))
			if err != nil {
				return c.JSON(http.StatusForbidden, &outputs.CustomError{
					Name:    AuthorizationError,
					Message: err.Error(),
				})
			}
			_, role, err := parseToken(token)
			if err != nil {
				return c.JSON(http.StatusForbidden, &outputs.CustomError{
					Name:    AuthorizationError,
					Message: err.Error(),
				})
			}
			if role != "admin" {
				return c.JSON(http.StatusMethodNotAllowed, &outputs.CustomError{
					Name:    AuthorizationError,
					Message: "Method not allowed",
				})
			}
			return next(c)
		}
		return fn
	}
}
