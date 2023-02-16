package database

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connection was successful!!")
	return db, nil
}

func MigrateDB(connection *gorm.DB) error {
	connection.AutoMigrate(&Goal{})
	connection.AutoMigrate(&Сontribution{})
	log.Println("DB migrate successful!!")
	return nil
}
