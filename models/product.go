package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"name" gorm:"not null;size:255"`
	Tag      string  `json:"tag" gorm:"size:100"`
	Price    float64 `json:"price" gorm:"not null;check:price > 0"`
	Quantity int     `json:"quantity" gorm:"not null;default:0"`
	Sales    []Sale  `json:"sales,omitempty" gorm:"foreignKey:ProductID"`
}