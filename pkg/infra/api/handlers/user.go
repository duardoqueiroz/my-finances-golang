package handlers

import (
	"fmt"

	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{usecase: userUseCase}
}

func (u UserHandler) Create(c echo.Context) error {
	var createUserInput usecases.CreateUserInput
	err := c.Bind(&createUserInput)
	if err != nil {
		return c.JSON(400, fmt.Sprintf("Error binding data: %s", err))
	}

	output, err := u.usecase.Create(createUserInput)
	if err != nil {
		return c.JSON(500, fmt.Sprintf("Error creating user: %s", err))
	}

	return c.JSON(201, output)
}
