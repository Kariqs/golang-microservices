package handlers

import (
	"log"

	"github.com/Kariqs/mesh-art-gallery-api/models"
	"github.com/Kariqs/mesh-art-gallery-api/services"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product
	if err := ctx.BodyParser(&product); err != nil {
		log.Println(err)
		return services.SendErrorResponse(ctx, fiber.StatusBadRequest, "unable to parse request body")
	}

	var productExists models.Product
	result := services.FindProductByTag(product.Tag, &productExists)
	if result.RowsAffected > 0 {
		return services.SendErrorResponse(ctx, fiber.StatusConflict, "product with this tag already exist")
	}

	if err := services.CreateProduct(&product).Error; err != nil {
		services.SendErrorResponse(ctx, fiber.StatusInternalServerError, "unable to create product")
	}

	return services.SendJSONResponse(ctx, fiber.StatusOK, fiber.Map{
		"message": "product has been created successfully",
		"product": product,
	})
}

func GetProducts(ctx *fiber.Ctx) error {
	var products []models.Product
	if err := services.GetProducts(&products).Error; err != nil {
		return services.SendErrorResponse(ctx, fiber.StatusInternalServerError, "unable to fetch products")
	}
	return services.SendJSONResponse(ctx, fiber.StatusOK, fiber.Map{
		"products": products,
	})
}

func GetProductByTag(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")
	var product models.Product
	result := services.FindProductByTag(tag, &product)
	if result.RowsAffected < 1 {
		return services.SendErrorResponse(ctx, fiber.StatusNotFound, "product not foundF")
	}
	return services.SendJSONResponse(ctx, fiber.StatusOK, fiber.Map{
		"product": product,
	})
}
