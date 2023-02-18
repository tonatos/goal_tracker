package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"goal-tracker/api/database"
	"goal-tracker/api/utils"
)

func GetGoalById(c *fiber.Ctx, db *gorm.DB, id string) (*database.Goal, *utils.HTTPError) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, &utils.HTTPError{
			Code:    400,
			Message: "Invalid id",
		}
	}

	var goal database.Goal
	query := database.Goal{ID: uint(intId)}
	err = db.First(&goal, query).Error

	if err == gorm.ErrRecordNotFound {
		return nil, &utils.HTTPError{
			Code:    404,
			Message: "Can't find goal with this id",
		}
	}
	return &goal, nil
}
