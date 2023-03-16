package app

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/tonatos/goal-tracker/internal/app/routes"
	"github.com/tonatos/goal-tracker/internal/services/auto_ru"
	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

func initDatabase() {
	config := database.PostgresConfig{
		ServerName: utils.GetEnvWithDefault("POSTGRES_HOST", "localhost"),
		User:       utils.GetEnvWithDefault("POSTGRES_USER", "postgres"),
		Password:   utils.GetEnvWithDefault("POSTGRES_PASSWORD", "postgres"),
		Port:       database.GetPostgresPort(utils.GetEnvWithDefault("POSTGRES_PORT", "5432")),
		DB: func() string {
			if utils.GetCurrentEnv() == "test" {
				return fmt.Sprintf("%s_test", utils.GetEnvWithDefault("POSTGRES_DB", "postgres"))
			} else {
				return os.Getenv("POSTGRES_DB")
			}
		}(),
	}

	connectionString := database.GetConnectionString(config)

	var err error
	database.DB, err = database.ConnectPostgres(connectionString, utils.GetCurrentEnv())
	if err != nil {
		panic(err.Error())
	}
}

func initRedis() {
	config := database.ConfigRedis{
		ServerName: os.Getenv("REDIS_HOST"),
		Port:       os.Getenv("REDIS_PORT"),
		Password:   "",
		DB:         0,
	}

	var err error
	database.Redis, err = database.ConnectRedis(&config)
	if err != nil {
		panic(err.Error())
	}
}

func initRoutes(app *fiber.App) {
	routes.GoalRoutes(app)
	routes.ContributionRoutes(app)
	routes.DocsRoutes(app)
	routes.NotFoundRoutes(app)
}

func SetupInit() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: true,
	}))

	// initialize database
	initDatabase()

	// init redis
	initRedis()

	// initialize routes
	initRoutes(app)

	auto_ru.AutoruInit()

	return app
}
