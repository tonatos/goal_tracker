package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tonatos/goal-tracker/internal/handlers"
)

func UploadRoutes(app *fiber.App) {
	routes := app.Group("/api/v1")
	routes.Post("/upload/", handlers.UploadImage)
}
