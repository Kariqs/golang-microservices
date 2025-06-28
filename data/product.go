package data

import (
	"fmt"
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
		Id:          2,
		Name:        "Espresso",
		Description: "Conc coffee.",
		Price:       100.99,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	}}

func GetProducts() Products {
	return productList
}

func GetProduct(id int) (*Product, error) {
	for _, prod := range productList {
		if prod.Id == id {
			return prod, nil
		}
	}
	return nil, fmt.Errorf("product with id %d does not exist", id)
}

func AddProduct(product Product) error {
	for _, prod := range productList {
		if prod.Id == product.Id {
			return fmt.Errorf("product with ID %d already exists", product.Id)
		}
	}
	productList = append(productList, &product)
	return nil
}

func UpdateProduct(id int, product Product) error {
	for _, prod := range productList {
		if prod.Id == id {
			prod.Name = product.Name
			prod.Description = product.Description
			prod.Price = product.Price
			prod.UpdatedAt = time.Now().UTC().String()
			return nil
		}
	}
	return fmt.Errorf("product with id %d does not exist", id)
}
