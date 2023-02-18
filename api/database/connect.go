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

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
	Port       int64
}

func GetConnectionString(config Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Yekaterinburg",
		config.ServerName,
		config.User,
		config.Password,
		config.DB,
		strconv.FormatInt(config.Port, 10),
	)
}

func Connect(connectionString string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
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

func MigrateDB(connection *gorm.DB) error {
	connection.AutoMigrate(&Goal{})
	connection.AutoMigrate(&Contribution{})
	log.Println("DB migrate successful!!")
	return nil
}

// DB gorm connector
var DB *gorm.DB
