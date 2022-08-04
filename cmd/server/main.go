package main

import (
	"net/http"

	"github.com/O-Mercan/Product--Service-V2/internal/database"
	"github.com/O-Mercan/Product--Service-V2/internal/product"
	transportHTTP "github.com/O-Mercan/Product--Service-V2/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

// App - the struct with contains like pointers
// database connections
type App struct{}

func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting up our app")

	db, err := database.NewDatabase()
	if err != nil {
		log.Error("New Database Error")
		return err
	}

	if err = database.MigrateDB(db); err != nil {
		log.Error("Migrate Error")
		return err
	}

	productService := product.NewService(db)

	handler := transportHTTP.NewHandler(productService)
	handler.SetUpRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}
	log.Info("Server is running on port 8080")
	return nil

}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Error("Error starting API")
	}
}
