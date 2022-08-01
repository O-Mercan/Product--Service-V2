package http

import (
	"encoding/json"
	"fmt"
	"github.com/O-Mercan/Product--Service-V2/internal/product"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	fmt.Println("Routes are setting")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/products/{id}", h.GetProduct).Methods("GET")
	h.Router.HandleFunc("/api/products", h.PostProduct).Methods("POST")

	h.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(writer, Response{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}
	pr, err := h.Service.GetProduct(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving Product By ID", err)
		return
	}

	if err := sendOkResponse(w, pr); err != nil {
		panic(err)
	}
}

func (h *Handler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var pr product.Product
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}
	pr, err := h.Service.PostProduct(pr)
	if err != nil {
		sendErrorResponse(w, "Failed to create new product", err)
	}

	if err := sendOkResponse(w, pr); err != nil {
		panic(err)
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
		panic(err)
	}
}
