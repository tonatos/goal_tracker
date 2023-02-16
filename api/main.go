package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"goal-tracker/api/database"
	"goal-tracker/api/handlers"
)

func dbPortFronEnv() int64 {
	db_port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 0, 64)
	if db_port == 0 || err == nil {
		db_port = 5432
	}
	log.Println("DB port is", db_port)
	return db_port
}

func initDatabase() {
	config := database.Config{
		ServerName: os.Getenv("POSTGRES_HOST"),
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		DB:         os.Getenv("POSTGRES_DB"),
		Port:       dbPortFronEnv(),
	}

	connectionString := database.GetConnectionString(config)
	connection, err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.MigrateDB(connection)
}

func initRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		return c.SendString("GoalTracker API v1")
	})
	v1.Get("/goal", handlers.GetGoals)
	v1.Get("/goal/:id", handlers.GetGoal)
	v1.Post("/goal", handlers.CreateGoal)
	v1.Put("/goal/:id", handlers.UpdateGoal)
	v1.Delete("/goal/:id", handlers.DeleteGoal)

	v1.Get("/contribution", handlers.GetContributions)
	v1.Post("/contribution", handlers.CreateContribution)
	v1.Put("/contribution/:id", handlers.UpdateContribution)
	v1.Delete("/contribution/:id", handlers.DeleteContribution)
}

func main() {
	app := fiber.New()

	// initialize database
	initDatabase()

	// initialize routes
	initRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
