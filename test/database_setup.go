package test

import (
	"fmt"

	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ClearConnectionToPostgres() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(
			database.GetConnectionString(
				database.PostgresConfig{
					ServerName: utils.GetEnvWithDefault("POSTGRES_HOST", "localhost"),
					User:       utils.GetEnvWithDefault("POSTGRES_USER", "postgres"),
					Password:   utils.GetEnvWithDefault("POSTGRES_PASSWORD", "postgres"),
					Port:       database.GetPostgresPort(utils.GetEnvWithDefault("POSTGRES_PORT", "5432")),
					DB:         "",
				},
			),
		),
	)

	if err != nil {
		panic(err)
	}
	return db
}

func InitTestDataBase(dbName string) {
	DropTestDataBase(dbName)

	db := ClearConnectionToPostgres()
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)).Error; err != nil {
		panic(err)
	}

	connection, _ := db.DB()
	connection.Close()
}

func DropTestDataBase(dbName string) {
	db := ClearConnectionToPostgres()
	if err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)).Error; err != nil {
		panic(err)
	}

	connection, _ := db.DB()
	connection.Close()
}
