package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	ProductID    uint    `json:"product_id" gorm:"not null;index"`
	SaleQuantity int     `json:"sale_quantity" gorm:"not null;check:sale_quantity > 0"`
	Product      Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
