package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"goal-tracker/api/database"
	"goal-tracker/api/handlers/helpers"
	"goal-tracker/api/utils"
)

// @Summary      Goals List
// @Description  Get goals list
// @Tags         Goals
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=[]database.Goal}
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal [get]
func GetGoal(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if err != nil {
		return err
	}

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    goal,
	})
}

// @Summary      Goal Item
// @Description  Get goal by :id
// @Tags         Goals
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=database.Goal}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:id [get]
func GetGoals(c *fiber.Ctx) error {
	var goals []database.Goal
	database.DB.Find(&goals)
	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    goals,
	})
}

// @Summary      Goal Create
// @Description  Create goal
// @Tags         Goals
// @Accept       json
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=database.Goal}
// @Failure      400  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal [post]
func CreateGoal(c *fiber.Ctx) error {
	json := new(database.Goal)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	newGoal := database.Goal{
		Name:       json.Name,
		Slug:       utils.Slugificator(json.Name),
		GoalAmount: json.GoalAmount,
		TargetDate: json.TargetDate,
	}

	err := database.DB.Create(&newGoal).Error
	if err != nil {
		return utils.NewError(c, 400, err)
	}

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    newGoal,
	})
}

// @Summary      Goal Update
// @Description  Update goal by id
// @Tags         Goals
// @Accept       json
// @Produce      json
// @Success 	 200 {string} status "ok"
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:id [put]
func UpdateGoal(c *fiber.Ctx) error {
	json := new(database.Goal)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewError(c, 400, utils.HTTPError{
			Message: "Invalid id",
		})
	}

	var found database.Goal
	query := database.Goal{ID: uint(id)}
	err = database.DB.Preload("Goal").First(&found, query).Error

	if err == gorm.ErrRecordNotFound {
		return utils.NewError(c, 400, utils.HTTPError{
			Message: "Can't find goal with this id",
		})
	}

	database.DB.Model(&found).Updates(json)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    found,
	})
}

// @Summary      Goal Delete
// @Description  Delete goal by id
// @Tags         Goals
// @Produce      json
// @Success 	 200 {string} status "ok"
// @Failure      400  {object}  utils.HTTPError
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:id [delete]
func DeleteGoal(c *fiber.Ctx) error {
	goal, http_err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if http_err != nil {
		return utils.NewError(c, http_err.Code, http_err)
	}
	database.DB.Delete(&goal)
	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
	})
}
