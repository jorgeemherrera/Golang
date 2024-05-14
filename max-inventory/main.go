package main

//inyeccion de dependencias
import (
	"context"
	"fmt"

	"github.com/jorgeemherrera/Golang/database"
	"github.com/jorgeemherrera/Golang/internal/api"
	"github.com/jorgeemherrera/Golang/internal/repository"
	"github.com/jorgeemherrera/Golang/internal/service"
	"github.com/jorgeemherrera/Golang/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
