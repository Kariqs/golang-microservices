package routes

import (
	"github.com/Kariqs/mesh-art-gallery-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(router fiber.Router) {
	router.Post("/", handlers.CreateProduct)
	router.Get("/", handlers.GetProducts)
	router.Get("/:tag", handlers.GetProductByTag)
	router.Put("/:tag", handlers.UpdateProduct)
	router.Delete("/:tag", handlers.DeleteProductHandler)
}
