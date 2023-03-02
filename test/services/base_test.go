package services

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tonatos/goal-tracker/pkg/database"
)

func TestServicesTestSuite(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}

func (suite *ServicesTestSuite) SetupSuite() {
	log.Println("Init services tests")

	var err error
	database.Redis, err = database.ConnectRedis(&database.ConfigRedis{
		ServerName: os.Getenv("REDIS_HOST"),
		Port:       os.Getenv("REDIS_PORT"),
		Password:   "",
		DB:         1,
	})
	if err != nil {
		panic(err.Error())
	}
}

func (suite *ServicesTestSuite) TearDownSuite() {
	log.Println("Teardown handlers test")

	// clear redis keys
	database.Redis.FlushDB()
}
