package handlers

import (
	"net/http"

	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/usecases"
	"github.com/duardoqueiroz/my-finances-golang/pkg/infra/security"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecases.User
}

func NewUserHandler(userUseCase usecases.User) *UserHandler {
	return &UserHandler{usecase: userUseCase}
}

func (u UserHandler) FindMe(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	claims, err := security.ParseAccessToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error trying to parse access token")
	}
	user, err := u.usecase.FindByID(claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error trying to find user")
	}

	return c.JSON(http.StatusOK, user)
}
