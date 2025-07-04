package services

import "github.com/gofiber/fiber/v2"

func SendErrorResponse(ctx *fiber.Ctx, errorCode int, errorMessage string) error {
	return ctx.Status(errorCode).JSON(fiber.Map{
		"error": errorMessage,
	})
}

func SendJSONResponse(ctx *fiber.Ctx, httpCode int, response fiber.Map) error {
	return ctx.Status(httpCode).JSON(response)
}
