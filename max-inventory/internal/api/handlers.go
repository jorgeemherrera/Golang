package api

import (
	"log"
	"net/http"

	"github.com/jorgeemherrera/Golang/encryption"
	"github.com/jorgeemherrera/Golang/internal/api/dtos"
	"github.com/jorgeemherrera/Golang/internal/service"
	"github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:message`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()

	params := dtos.RegisterUser{}

	//Obtener los parametros del request
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid Request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.srv.RegisterUser(ctx, params.Email, params.Name, params.Password)

	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "User already exists"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}
	return c.JSON(http.StatusCreated, nil)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid Request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	user, err := a.srv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}
	//TODO: crear jwt token

	token, err := encryption.SignedLoginToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}
	// retornar jwt token
	return c.JSON(http.StatusOK, map[string]string{"success": token})
}
