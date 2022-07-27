package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	helpers "github.com/O-Mercan/Product--Service-V2/helpers"
	models "github.com/O-Mercan/Product--Service-V2/models"
	"github.com/gorilla/mux"
)

var productStore = make(map[string]models.Product)
var id int = 0

// HTTP POST - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	helpers.CheckError(err)

	product.CreatedOn = time.Now()
	id++
	product.Id = id
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content- Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

// HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	for _, product := range productStore {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	helpers.CheckError(err)

	w.Header().Set("Content- Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

// HTTP Get - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	vars := mux.Vars(r) //request-URL içerisindeki variableları alıyor
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.Id == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content- Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP PUT - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var prodUpd models.Product
	err = json.NewDecoder(r.Body).Decode(&prodUpd)
	helpers.CheckError(err)

	if _, ok := productStore[key]; ok {
		prodUpd.Id, _ = strconv.Atoi(key)
		prodUpd.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodUpd
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusNoContent)
}

// HTTP Delete - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Değer bulunamadı : %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
