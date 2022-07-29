package main

import (
	"log"
	"net/http"

	handlers "github.com/O-Mercan/Product--Service-V2/handlers"
	models "github.com/O-Mercan/Product--Service-V2/models"
	repo "github.com/O-Mercan/Product--Service-V2/repositories"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	log.Println("Server starting...")

	product := models.Product{
		Name:        "Samsung Z Fold",
		Category:    "Tech",
		Summary:     "asdasd",
		Description: "asdsdas",
		Price:       6000,
	}
	repo.InsertProduct(product)

	//repo.GetProducts()
	//models.GetProductByID(2)

	/* product := models.Product{
		Id:          3,
		Name:        "İphone 6",
		Category:    "Tech",
		Summary:     "asdasd",
		Description: "asdsdas",
		Price:       2000,
	}
	models.UpdateProduct(product) */
	//models.DeleteProductByID(4)
	//models.GetProducts()

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
