package handlers

import "github.com/gofiber/fiber/v2"

func GetGoal(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func GetGoals(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func CreateGoal(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func UpdateGoal(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}

func DeleteGoal(c *fiber.Ctx) error {
	return c.JSON(c.App().Stack())
}
