package main

import (
	"femo-go/internal/database"
	transportHTTP "femo-go/internal/transport/http"
	"fmt"
	"net/http"
)

// Main structure to contain server thingy like pointer to DB
type App struct{}

func (a *App) Run() error {
	fmt.Println("Initializing the app")

	_, err := database.NewDatabase();
	if err != nil {
		fmt.Println(err)
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":9009", handler.Router); err != nil {
		fmt.Println("Error occurred while instantiating")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Running")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
