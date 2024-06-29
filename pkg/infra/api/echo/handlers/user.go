package handlers

import (
	"fmt"
	"net/http"

	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecases.User
}

func NewUserHandler(userUseCase usecases.User) *UserHandler {
	return &UserHandler{usecase: userUseCase}
}

func (u UserHandler) Create(c echo.Context) error {
	var createUserInput inputs.CreateUserInput
	err := c.Bind(&createUserInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &outputs.CustomError{Name: "InvalidInput", Message: fmt.Sprintf("Error parsing input: %s", err)})
	}

	output, err := u.usecase.Create(createUserInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &outputs.CustomError{Name: "InternalError", Message: fmt.Sprintf("Error creating user: %s", err)})
	}

	return c.JSON(http.StatusCreated, output)
}
