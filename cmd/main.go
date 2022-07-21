package main

import (
	"log"
	"net/http"

	. "../handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter()
	r.HandleFunc("/api/prooducts", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/prooducts/{id}", GetProductHandler).Methods("GET")
	r.HandleFunc("/api/prooducts/{id}", PostProductHandler).Methods("POST")
	r.HandleFunc("/api/prooducts", PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/prooducts", DeleteProductHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	server.ListenAndServe()

	log.Println("Server ending...")

}
