package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type PostgresConfig struct {
	ServerName string
	User       string
	Password   string
	DB         string
	Port       int64
}

func GetPostgresPort(port string) int64 {
	db_port, err := strconv.ParseInt(port, 0, 64)
	if db_port == 0 || err != nil {
		return 5432
	}
	return db_port
}

func GetGormLogLevel(env string) logger.LogLevel {
	if env == "prod" || env == "test" {
		return logger.Error
	} else if env == "stage" {
		return logger.Warn
	} else {
		return logger.Info
	}
}

func GetConnectionString(config PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s %s port=%s sslmode=disable TimeZone=Asia/Yekaterinburg",
		config.ServerName,
		config.User,
		config.Password,
		func(dbname string) string {
			if dbname != "" {
				return fmt.Sprintf("dbname=%s", config.DB)
			} else {
				return ""
			}
		}(config.DB),
		strconv.FormatInt(config.Port, 10),
	)
}

func ConnectPostgres(connectionString string, env string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,          // Slow SQL threshold
			LogLevel:                  GetGormLogLevel(env), // Log level
			IgnoreRecordNotFoundError: true,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                 // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: false,
		},
	})

	if err != nil {
		return nil, err
	}
	log.Println("Connection was successful!!")
	return db, nil
}

// DB gorm connector
var DB *gorm.DB
