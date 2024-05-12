package main

//inyeccion de dependencias
import (
	"context"

	"github.com/jorgeemherrera/Golang/database"
	"github.com/jorgeemherrera/Golang/internal/repository"
	"github.com/jorgeemherrera/Golang/internal/service"
	"github.com/jorgeemherrera/Golang/settings"
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
		),
		fx.Invoke(),
	)

	app.Run()
}
