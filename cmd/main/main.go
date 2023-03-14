package main

import (
	"fmt"
	"log"

	"github.com/tonatos/goal-tracker/internal/app"
	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"

	_ "github.com/tonatos/goal-tracker/docs"
)

// @title           GoalTracker API
// @version         0.1
// @description     API Service of simple app for tracking your widescale goals

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	application := app.SetupInit()

	app.MigrateDB(database.DB)

	log.Fatal(
		application.Listen(
			fmt.Sprintf(":%s", utils.GetEnvWithDefault("APP_PORT", "8000")),
		),
	)
}
