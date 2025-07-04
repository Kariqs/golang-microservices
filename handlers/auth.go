package handlers

import "github.com/gofiber/fiber/v2"

func AuthHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"handler": "Auth",
	})
}
