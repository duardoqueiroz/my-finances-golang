package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/duardoqueiroz/my-finances-golang/pkg/application/inputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/application/outputs"
	"github.com/duardoqueiroz/my-finances-golang/pkg/domain/usecases"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase usecases.Auth
}

func NewAuthHandler(usecase usecases.Auth) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (a AuthHandler) SignUp(c echo.Context) error {
	var createUserInput inputs.SignUp
	err := c.Bind(&createUserInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &outputs.CustomError{Name: "InvalidInput", Message: fmt.Sprintf("Error parsing input: %s", err)})
	}

	output, err := a.usecase.SignUp(createUserInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &outputs.CustomError{Name: "InternalError", Message: fmt.Sprintf("Error creating user: %s", err)})
	}

	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   output.Token.Value,
		Expires: time.Unix(output.Token.Expires, 0),
	})
	return c.JSON(http.StatusCreated, output)
}

func (ah *AuthHandler) Login(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err == nil {
		if cookie.Value != "" {
			return c.JSON(http.StatusForbidden, "user already logged in")
		}
	}

	var loginInput inputs.Login
	err = c.Bind(&loginInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid body")
	}

	output, err := ah.usecase.Login(loginInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid credentials")
	}

	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   output.Token.Value,
		Expires: time.Unix(output.Token.Expires, 0),
	})
	return c.JSON(http.StatusAccepted, output)
}

func (ah *AuthHandler) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:  "token",
		Value: "",
	})
	return c.JSON(http.StatusOK, "user logged out")
}
