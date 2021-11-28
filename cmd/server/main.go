package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/Tutorial/rest_api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up App")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed to setup server")
		return err
	}
	return nil
}
func main() {
	fmt.Println("GO REST")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our Rest api")
		fmt.Println(err)
	}
}
