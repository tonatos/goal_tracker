package database

import (
	"log"

	"goal-tracker/api/models/db"

	"gorm.io/gorm"
)

func MigrateDB(connection *gorm.DB) error {
	connection.AutoMigrate(&db.Goal{})
	connection.AutoMigrate(&db.Contribution{})

	log.Println("DB migrate successful!!")
	return nil
}
