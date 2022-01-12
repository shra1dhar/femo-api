package main

import (
	transportHTTP "femo-go/internal/transport/http"
	"fmt"
	"net/http"
)

// Main structure to contain server thingy like pointer to DB
type App struct {}

func (a *App) Run() error {
	fmt.Println("Running")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":9009",handler.Router); err != nil {
		fmt.Println("Error occured while instantiating")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Running")

	app := App{}
	if err := app.Run(); err != nil {

	}
}