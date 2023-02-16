package handlers

import "github.com/gofiber/fiber/v2"

func GetContributions(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func CreateContribution(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func UpdateContribution(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func DeleteContribution(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}
