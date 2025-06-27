package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kariqs/golang-microservices/data"
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
	jsonData, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshal response.", http.StatusInternalServerError)
	}
	fmt.Fprintf(rw,"%s", string(jsonData))
}
