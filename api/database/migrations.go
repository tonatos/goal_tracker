package database

import (
	"log"

	"goal-tracker/api/models/table"

	"gorm.io/gorm"
)

func MigrateDB(connection *gorm.DB) error {
	connection.AutoMigrate(&table.Goal{})
	connection.AutoMigrate(&table.Contribution{})

	log.Println("DB migrate successful!!")
	return nil
}
