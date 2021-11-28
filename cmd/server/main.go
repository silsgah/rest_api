package main

import "fmt"

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up App")
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
