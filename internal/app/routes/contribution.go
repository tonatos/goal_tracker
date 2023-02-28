package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tonatos/goal-tracker/internal/handlers"
)

func ContributionRoutes(app *fiber.App) {
	routes := app.Group("/api/v1")

	routes.Get("/goal/:goal/contribution/", handlers.GetContributions)
	routes.Post("/goal/:goal/contribution/", handlers.CreateContribution)
	routes.Put("/goal/:goal/contribution/:id", handlers.UpdateContribution)
	routes.Delete("/goal/:goal/contribution/:id", handlers.DeleteContribution)
}
