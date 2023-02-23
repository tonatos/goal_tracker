package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	db_model "goal-tracker/api/models/db"
	"goal-tracker/api/utils"
)

func GetGoalById(c *fiber.Ctx, db *gorm.DB, id string) (*db_model.Goal, *utils.HTTPError) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, &utils.HTTPError{
			Code:    400,
			Message: "Invalid goal id",
		}
	}

	var goal db_model.Goal
	query := db_model.Goal{ID: uint(intId)}
	err = db.First(&goal, query).Error

	if err == gorm.ErrRecordNotFound {
		return nil, &utils.HTTPError{
			Code:    404,
			Message: "Can't find goal with this id",
		}
	}
	return &goal, nil
}

func GetContributionById(c *fiber.Ctx, db *gorm.DB, goal_id string, id string) (
	*db_model.Contribution,
	*utils.HTTPError,
) {
	_, http_err := GetGoalById(c, db, goal_id)
	if http_err != nil {
		return nil, http_err
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, &utils.HTTPError{
			Code:    400,
			Message: "Invalid contrbution id",
		}
	}

	var contrbution db_model.Contribution
	query := db_model.Contribution{ID: uint(intId)}
	err = db.Joins("Goal").First(&contrbution, query).Error
	if err != nil {
		return nil, &utils.HTTPError{
			Code:    404,
			Message: "Can't find contrbution with this id",
		}
	}

	return &contrbution, nil
}
