package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/tonatos/goal-tracker/internal/app/routes"
	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

func initDatabase() {
	config := database.PostgresConfig{
		ServerName: os.Getenv("POSTGRES_HOST"),
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		DB: func() string {
			if utils.GetCurrentEnv() == "test" {
				return fmt.Sprintf("%s_test", utils.GetEnvWithDefault("POSTGRES_DB", "postgres"))
			} else {
				return os.Getenv("POSTGRES_DB")
			}
		}(),
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

	// initialize database
	initDatabase()

	// init redis
	initRedis()

	// initialize routes
	initRoutes(app)

	return app
}
