package handlers

import (
	"github.com/gofiber/fiber/v2"

	"goal-tracker/api/database"
	"goal-tracker/api/handlers/helpers"
	db_model "goal-tracker/api/models/db"
	response_model "goal-tracker/api/models/response"
	"goal-tracker/api/services/auto_ru"
	"goal-tracker/api/utils"
)

// @Summary      Goals List
// @Description  Get goals list
// @Tags         Goals
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=[]db_model.Goal}
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal [get]
func GetGoal(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("id"))
	if err != nil {
		return err
	}

	ar := auto_ru.AutoruInit(goal.GoalAmount)
	ads_count, _ := ar.CountAds()
	catalog_link, _ := ar.GetCatalogLink()

	goal_response := response_model.Goal{
		Goal:        goal,
		CatalogUrl:  catalog_link,
		AdsByAmount: ads_count,
	}

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    goal_response,
	})
}

// @Summary      Goal Item
// @Description  Get goal by :id
// @Tags         Goals
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=db_model.Goal}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:id [get]
func GetGoals(c *fiber.Ctx) error {
	var goals []db_model.Goal
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
// @Success 	 200  {object}  utils.JSONResult{data=db_model.Goal}
// @Failure      400  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal [post]
func CreateGoal(c *fiber.Ctx) error {
	json := new(db_model.Goal)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	newGoal := db_model.Goal{
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
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("id"))
	if err != nil {
		return err
	}

	json := new(db_model.Goal)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	database.DB.Model(&goal).Updates(json)
	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    goal,
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
	goal, http_err := helpers.GetGoalById(c, database.DB, c.Params("id"))
	if http_err != nil {
		return utils.NewError(c, http_err.Code, http_err)
	}
	database.DB.Delete(&goal)
	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
	})
}
