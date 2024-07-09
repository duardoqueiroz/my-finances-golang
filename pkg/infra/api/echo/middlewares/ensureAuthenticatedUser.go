package middlewares

import (
	"fmt"
	"net/http"

	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/labstack/echo/v4"
)

func EnsureAuthenticatedUser() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		fn := func(c echo.Context) error {
			userId := c.Param("id")
			token := c.Request().Header.Get("Authorization")
			id, role, err := parseToken(token)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, &outputs.CustomError{
					Name:    AuthorizationError,
					Message: err.Error(),
				})
			}
			fmt.Println(id)
			fmt.Println(role)
			if id != userId && role != "admin" {
				return c.JSON(http.StatusMethodNotAllowed, &outputs.CustomError{
					Name:    AuthorizationError,
					Message: "Unauthorized resource",
				})
			}
			return next(c)
		}
		return fn
	}
}
