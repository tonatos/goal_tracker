package test

import (
	"fmt"

	"gorm.io/gorm"
)

func InitTestDataBase(db *gorm.DB, dbName string) {
	DropTestDataBase(db, dbName)

	if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)).Error; err != nil {
		panic(err)
	}
}

func DropTestDataBase(db *gorm.DB, dbName string) {
	if err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)).Error; err != nil {
		panic(err)
	}
}
