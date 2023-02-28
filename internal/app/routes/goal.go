package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tonatos/goal-tracker/internal/handlers"
)

func GoalRoutes(app *fiber.App) {
	routes := app.Group("/api/v1")

	routes.Get("/goal", handlers.GetGoals)
	routes.Get("/goal/:id", handlers.GetGoal)
	routes.Post("/goal", handlers.CreateGoal)
	routes.Put("/goal/:id", handlers.UpdateGoal)
	routes.Delete("/goal/:id", handlers.DeleteGoal)
}
