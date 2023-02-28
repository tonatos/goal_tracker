package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type HandlersTestSuite struct {
	suite.Suite
	App        *fiber.App
	DB         *gorm.DB
	TestDBName string
}
