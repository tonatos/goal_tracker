package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"goal-tracker/api/database"
	_ "goal-tracker/api/docs"
	"goal-tracker/api/handlers"
	"goal-tracker/api/utils"
)

func getEnv() string {
	var env string
	switch os.Getenv("ENV") {
	case "prod", "stage":
		env = "prod"
	default:
		env = "dev"
	}
	return env
}

func initDatabase() {
	config := database.Config{
		ServerName: os.Getenv("POSTGRES_HOST"),
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		DB:         os.Getenv("POSTGRES_DB"),
		Port: func() int64 {
			db_port, err := strconv.ParseInt(
				utils.GetEnvWithDefault("POSTGRES_PORT", "5432"), 0, 64,
			)
			if db_port == 0 || err != nil {
				return 5432
			}
			return db_port
		}(),
	}

	connectionString := database.GetConnectionString(config)

	var err error
	database.DB, err = database.Connect(connectionString, getEnv())
	if err != nil {
		panic(err.Error())
	}

	database.MigrateDB(database.DB)
}

func initRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/goal", handlers.GetGoals)
	v1.Get("/goal/:id", handlers.GetGoal)
	v1.Post("/goal", handlers.CreateGoal)
	v1.Put("/goal/:id", handlers.UpdateGoal)
	v1.Delete("/goal/:id", handlers.DeleteGoal)

	v1.Get("/goal/:goal/contribution/", handlers.GetContributions)
	v1.Post("/goal/:goal/contribution/", handlers.CreateContribution)
	v1.Put("/goal/:goal/contribution/:id", handlers.UpdateContribution)
	v1.Delete("/goal/:goal/contribution/:id", handlers.DeleteContribution)

	v1.Get("/docs/*", swagger.New(swagger.Config{
		DeepLinking:  true,
		DocExpansion: "list",
	}))

	v1.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}

// @title           GoalTracker API
// @version         0.1
// @description     API Service of simple app for tracking your widescale goals

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	app := fiber.New()

	// initialize database
	initDatabase()

	// initialize routes
	initRoutes(app)
	log.Fatal(
		app.Listen(
			fmt.Sprintf(":%s", utils.GetEnvWithDefault("APP_PORT", "8000")),
		),
	)
}
