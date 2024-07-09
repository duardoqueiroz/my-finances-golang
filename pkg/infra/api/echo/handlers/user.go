package handlers

import (
	"net/http"

	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
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

func (u UserHandler) Update(c echo.Context) error {
	var input inputs.UpdateUserInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid data")
	}
	id := c.Param("id")
	err = u.usecase.Update(id, input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &outputs.CustomError{
			Name:    "Error updating data",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "User updated")
}

func (uh UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	err := uh.usecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &outputs.CustomError{
			Name:    "User cant be deleted",
			Message: err.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}
