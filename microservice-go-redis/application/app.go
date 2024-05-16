package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

// Constructor
func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

// (a *App) --> this on javascript
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
