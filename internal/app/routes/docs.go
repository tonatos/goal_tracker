package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func DocsRoutes(app *fiber.App) {
	routes := app.Group("/api/v1")

	routes.Get("/docs/*", swagger.New(swagger.Config{
		DeepLinking:  true,
		DocExpansion: "list",
	}))
}
