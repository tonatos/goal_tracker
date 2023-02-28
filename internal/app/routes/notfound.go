package routes

import (
	"github.com/gofiber/fiber/v2"
)

func NotFoundRoutes(app *fiber.App) {
	routes := app.Group("/api/v1")

	routes.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}
