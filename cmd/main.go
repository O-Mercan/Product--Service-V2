package main

import (
	"log"
	"net/http"

	handlers "github.com/O-Mercan/Product--Service-V2/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter()
	r.HandleFunc("/api/prooducts", handlers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/prooducts/{id}", handlers.GetProductHandler).Methods("GET")
	r.HandleFunc("/api/prooducts/{id}", handlers.PostProductHandler).Methods("POST")
	r.HandleFunc("/api/prooducts", handlers.PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/prooducts", handlers.DeleteProductHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

	log.Println("Server ending...")

}
