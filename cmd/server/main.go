package main

import (
	"fmt"
	"github.com/O-Mercan/Product--Service-V2/internal/database"
	"github.com/O-Mercan/Product--Service-V2/internal/product"
	transportHTTP "github.com/O-Mercan/Product--Service-V2/internal/transport/http"
	"net/http"
)

// App - the struct with contains like pointers
// database connections
type App struct{}

func (a *App) Run() error {
	fmt.Println("Setting up our app")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	if err = database.MigrateDB(db); err != nil {
		return err
	}

	productService := product.NewService(db)

	handler := transportHTTP.NewHandler(productService)
	handler.SetUpRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	fmt.Println("Server is running on port 8080")
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting API")
	}
}
