package services

import (
	"github.com/Kariqs/mesh-art-gallery-api/initializers"
	"github.com/Kariqs/mesh-art-gallery-api/models"
	"gorm.io/gorm"
)

func CreateProduct(product models.Product) *gorm.DB {
	return initializers.DB.Create(&product)
}

func GetProducts(products *[]models.Product) *gorm.DB {
	return initializers.DB.Find(products)
}

func FindProductByTag(product, productExists models.Product) *gorm.DB {
	return initializers.DB.Where("tag=?", product.Tag).First(&productExists)
}
