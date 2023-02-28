package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tonatos/goal-tracker/internal/app"
	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"
	"github.com/tonatos/goal-tracker/test"
	"github.com/tonatos/goal-tracker/test/fixtures"
)

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}

func (suite *HandlersTestSuite) SetupSuite() {
	log.Println("Init handlers tests")

	suite.TestDBName = fmt.Sprintf("%s_test", utils.GetEnvWithDefault("POSTGRES_DB", "postgres"))

	// init test database
	test.InitTestDataBase(suite.TestDBName)

	// init app and test database connection
	suite.App = app.SetupInit()

	// setup actual connection and other conf
	suite.DB = database.DB
	suite.Api = "/api/v1/"

	// migrate
	app.MigrateDB(suite.DB)

	// fill data
	goals := fixtures.GoalFixturesFabric().Items
	suite.DB.Create(&goals)
	suite.DB.Create(&fixtures.ContributionFixturesFabric(&goals[0]).Items)
}

func (suite *HandlersTestSuite) TearDownSuite() {
	log.Println("Teardown handlers test")

	// close existing connection to *_test
	connection, _ := suite.DB.DB()
	connection.Close()

	// dropoff test database
	test.DropTestDataBase(suite.TestDBName)
}

func (suite *HandlersTestSuite) Request(method, target string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, fmt.Sprintf("%s%s", suite.Api, target), body)
	req.Header.Set("Content-type", "application/json")
	return req
}
