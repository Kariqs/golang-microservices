package data

import (
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
}

type Products []*Product

var productList = Products{
	&Product{
		Id:          1,
		Name:        "Latte",
		Description: "Coffee with milk.",
		Price:       100.99,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	}, &Product{
		Id:          1,
		Name:        "Espresso",
		Description: "Conc coffee.",
		Price:       100.99,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	}}

func GetProducts() Products {
	return productList
}

func AddProduct(product Product) {
	productList = append(productList, &product)
}

