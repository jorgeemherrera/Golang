package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jorgeemherrera/Golang/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	srv           service.Service
	dataValidator *validator.Validate
}

func New(srv service.Service) *API {
	return &API{
		srv:           srv,
		dataValidator: validator.New(),
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.RegisterRoutes(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500"},
		AllowMethods: []string{echo.POST},
		AllowHeaders: []string{echo.HeaderContentType},
	}))
	return e.Start(address)
}
