package routes

import "github.com/gofiber/fiber/v2"

func RegisterAuthRoutes(router fiber.Router) {
	router.Post("/auth")
}
