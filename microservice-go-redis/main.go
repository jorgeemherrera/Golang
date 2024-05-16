package main

import (
	"context"
	"fmt"
	"microservice-go-redis/application"
	"strings"
)

func main() {
	// new instance of our application using the constructor
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		// Verifica si el error es de tipo http.MaxBytesError
		if strings.Contains(err.Error(), "http: request body too large") {
			fmt.Println("Failed to start app due to MaxBytesError:", err)
		} else {
			// Si el error no es de tipo http.MaxBytesError, maneja otros casos de error
			fmt.Println("Failed to start app:", err)
		}
	}
}
