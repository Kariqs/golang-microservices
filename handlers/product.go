package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Kariqs/golang-microservices/data"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (p *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()

	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(lp)
	if err != nil {
		http.Error(rw, "unable to encode response", http.StatusInternalServerError)
		return
	}
}

func (p *ProductHandler) GetProductById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "unable to extract id from the request", http.StatusBadRequest)
		return
	}

	product, err := data.GetProduct(id)
	if err != nil {
		p.l.Println(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(product)
	if err != nil {
		http.Error(rw, "unable to encode response", http.StatusInternalServerError)
		return
	}
}

func (p *ProductHandler) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var product data.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		p.l.Println(err)
		http.Error(rw, "unable to decode request body", http.StatusInternalServerError)
		return
	}
	err := data.AddProduct(product)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(product)
	if err != nil {
		http.Error(rw, "unable to marshal response.", http.StatusInternalServerError)
		return
	}
	response := map[string]any{
		"message":        "Product created successfully.",
		"createdProduct": jsonData,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}
func (p *ProductHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "unable to extract id from the request", http.StatusBadRequest)
		return
	}

	var newProduct data.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		p.l.Println(err)
		http.Error(rw, "unable to decode request body", http.StatusInternalServerError)
		return
	}

	err = data.UpdateProduct(id, newProduct)
	if err != nil {
		p.l.Println(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]any{"message": "Product updated successfully."})
}
