package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/O-Mercan/Product--Service-V2/internal/product"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Router  *mux.Router
	Service *product.Service
}

func NewHandler(service *product.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetUpRoutes() {
	log.Info("Routes are setting")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET")
	h.Router.HandleFunc("/api/products", h.GetProducts).Methods("GET")
	h.Router.HandleFunc("/api/products", h.PostProduct).Methods("POST")
	h.Router.HandleFunc("/api/products/{id}", h.PutProduct).Methods("PUT")
	h.Router.HandleFunc("/api/products/{id}", h.DeleteProduct).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(writer, Response{Message: "I am alive!"}); err != nil {
			log.Panic(err)
		}
	})
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID, GetProduct handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)

	}
	pr, err := h.Service.GetProduct(uint(i))
	if err != nil {
		log.Error("Error retrieving Product By ID. GetProduct handler")
		sendErrorResponse(w, "Error retrieving Product By ID", err)
		return
	}

	if err := sendOkResponse(w, pr); err != nil {
		log.Panic(err)
	}
}

func (h *Handler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var pr product.Product
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		log.Error("Error retrieving Product By ID. PostProduct handler")
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	pr, err := h.Service.PostProduct(pr)
	if err != nil {
		log.Error("Error retrieving Product By ID. PostProduct handler")
		sendErrorResponse(w, "Failed to create new product", err)
	}

	if err := sendOkResponse(w, pr); err != nil {
		log.Panic(err)
	}
}

//Get products
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	pr, err := h.Service.GetProducts()
	if err != nil {
		log.Error("Failed to get products, GetProducts handler")
		sendErrorResponse(w, "Failed to get products", err)
		return
	}

	if err := sendOkResponse(w, pr); err != nil {
		log.Panic(err)
	}
}

//Put products
func (h *Handler) PutProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("Unable to parse UINT from ID. PutProduct Handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}
	var pr product.Product
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		log.Error("Faile to decode json body. PutProduct Handler")
		sendErrorResponse(w, "Faile to decode json body", err)
	}

	product, err := h.Service.PutProduct(uint(i), pr)
	if err != nil {
		log.Error("Failed to update a product. PutProduct Handler")
		sendErrorResponse(w, "Failed to update a product", err)
	}

	if err := sendOkResponse(w, product); err != nil {
		log.Panic(err)
	}
}

//Delete product
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error("convert Error, DeleteProduct handler")
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}

	pr, err := h.Service.DeleteProduct(uint(i))
	if err != nil {
		log.WithFields(log.Fields{
			"Function": "DeleteProduct",
		}).Error("ID doesn't exist")
		//log.Error("DeleteProduct handler")
		sendErrorResponse(w, "Failed to delete a product", err)
	}

	if err := sendOkResponse(w, pr); err != nil {
		log.Panic(err)
	}
}

type Response struct {
	Message string
	Error   string
}

func sendOkResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		log.Panic(err)
	}
}
