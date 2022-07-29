package main

import (
	"log"
	"net/http"

	handlers "github.com/O-Mercan/Product--Service-V2/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Server starting...")

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/products", handlers.GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", handlers.GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products", handlers.PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products", handlers.PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/products", handlers.DeleteProductHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

	log.Println("Server ending...")

}
