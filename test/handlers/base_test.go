package handlers

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tonatos/goal-tracker/internal/app"
	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"
	"github.com/tonatos/goal-tracker/test"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}

func (suite *HandlersTestSuite) SetupSuite() {
	log.Println("Init handlers tests")

	base_db_connection, _ := gorm.Open(postgres.Open(database.GetConnectionString(database.PostgresConfig{
		ServerName: os.Getenv("POSTGRES_HOST"),
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		DB:         "",
		Port:       5432,
	})))

	suite.TestDBName = fmt.Sprintf("%s_test", utils.GetEnvWithDefault("POSTGRES_DB", "postgres"))

	// init test database
	test.InitTestDataBase(base_db_connection, suite.TestDBName)
	connection, _ := base_db_connection.DB()
	connection.Close()

	// init app and databases
	suite.App = app.SetupInit()

	// setup base structure config
	suite.DB = database.DB

	app.MigrateDB(suite.DB)
}

func (suite *HandlersTestSuite) TearDownSuite() {
	log.Println("Teardown handlers test")

	connection, _ := suite.DB.DB()
	connection.Close()

	base_db_connection, _ := gorm.Open(postgres.Open(database.GetConnectionString(database.PostgresConfig{
		ServerName: os.Getenv("POSTGRES_HOST"),
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		DB:         "",
		Port:       5432,
	})))

	// dropoff test database
	test.DropTestDataBase(base_db_connection, suite.TestDBName)
}
