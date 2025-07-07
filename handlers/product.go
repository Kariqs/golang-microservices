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

func UpdateProduct(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")
	var productInfo models.Product

	if err := ctx.BodyParser(&productInfo); err != nil {
		log.Println(err)
		return services.SendErrorResponse(ctx, fiber.StatusBadRequest, "unable to parse request body")
	}

	var product models.Product
	result := services.FindProductByTag(tag, &product)
	if result.RowsAffected < 1 {
		return services.SendErrorResponse(ctx, fiber.StatusNotFound, "product not found")
	}

	if result = services.UpdateProduct(tag, &productInfo); result.Error != nil {
		return services.SendErrorResponse(ctx, fiber.StatusInternalServerError, "unable to update product information")
	}

	var updatedProduct models.Product
	services.FindProductByTag(tag, &updatedProduct)

	return services.SendJSONResponse(ctx, fiber.StatusOK, fiber.Map{
		"message": "Product updated successfully.",
		"product": updatedProduct,
	})
}

func DeleteProductHandler(ctx *fiber.Ctx) error {
	tag := ctx.Params("tag")

	result := services.DeleteProduct(tag)
	if result.Error != nil {
		log.Println(result.Error)
		return services.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Unable to delete product.")
	}
	if result.RowsAffected < 0 {
		return services.SendErrorResponse(ctx, fiber.StatusNotFound, "Product not found.")
	}

	return services.SendJSONResponse(ctx, fiber.StatusOK, fiber.Map{
		"message": "Product deleted successfully.",
	})
}
