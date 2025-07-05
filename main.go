package main

import (
	"time"

	"github.com/Kariqs/mesh-art-gallery-api/initializers"
	"github.com/Kariqs/mesh-art-gallery-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Mesh Art Gallery",
		CaseSensitive: true,
		ReadTimeout:   5 * time.Second,
		IdleTimeout:   10 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://yourdomain.com, http://localhost:4200",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")

	routes.RegisterProductRoutes(api.Group("/product"))

	app.Listen(":8080")
	app.ShutdownWithTimeout(30 * time.Second)
}
