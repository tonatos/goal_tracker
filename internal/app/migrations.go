package app

import (
	"log"

	"github.com/tonatos/goal-tracker/internal/models/table"

	"gorm.io/gorm"
)

func MigrateDB(connection *gorm.DB) error {
	connection.AutoMigrate(&table.Goal{})
	connection.AutoMigrate(&table.Contribution{})

	log.Println("DB migrate successful!!")
	return nil
}
