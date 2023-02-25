package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/handlers/helpers"
	"github.com/tonatos/goal-tracker/pkg/models/request"
	"github.com/tonatos/goal-tracker/pkg/models/response"
	"github.com/tonatos/goal-tracker/pkg/models/table"
	"github.com/tonatos/goal-tracker/pkg/services/auto_ru"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

// @Summary      Goals Item
// @Description  Get goal by :id
// @Tags         Goals
// @Produce      json
// @Param	  	 id	path	int	true	"Goal ID"
// @Success 	 200  {array}  utils.JSONResult{data=[]response.ResponesGoal}
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal [get]
func GetGoal(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("id"))
	if err != nil {
		return err
	}

	// @todo: считать накопления, а не передавать сумму
	accumulatedAmount := goal.GoalAmount

	ar := auto_ru.AutoruInit(accumulatedAmount)
	ads_count, _ := ar.CountAds()
	catalog_link, _ := ar.GetCatalogLink()

	goal_response := response.ResponesGoal{
		Goal:              goal,
		CatalogUrl:        catalog_link,
		AdsByAmount:       ads_count,
		AccumulatedAmount: accumulatedAmount,
		DaysUntilBang:     int(goal.TargetDate.Sub(time.Now()).Hours() / 24),
	}

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    goal_response,
	})
}

// @Summary      Goal List
// @Description  Get goals list
// @Tags         Goals
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=response.ResponesGoal}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:id [get]
func GetGoals(c *fiber.Ctx) error {
	var goals []table.Goal
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
// @Param        goal	body	request.RequestCreateGoal true "Goal object for create"
// @Success 	 200  {object}  utils.JSONResult{data=response.ResponesGoal}
// @Failure      400  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal [post]
func CreateGoal(c *fiber.Ctx) error {
	data := new(request.RequestCreateGoal)
	if err := c.BodyParser(data); err != nil {
		return utils.NewError(c, 400, err)
	}

	newGoal := table.Goal{
		Name:       data.Name,
		Slug:       utils.Slugificator(data.Name),
		GoalAmount: data.GoalAmount,
		TargetDate: data.TargetDate,
	}

	err := database.DB.Create(&newGoal).Error
	if err != nil {
		return utils.NewError(c, 400, err)
	}

	var result response.ResponesGoal
	database.DB.Model(&table.Goal{}).First(&result.Goal, newGoal)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    result,
	})
}

// @Summary      Goal Update
// @Description  Update goal by id
// @Tags         Goals
// @Accept       json
// @Produce      json
// @Param	  	 id	path	int	true	"Goal ID"
// @Param        goal	body	request.RequestUpdateGoal true "Goal`s fields for update"
// @Success 	 200  {object}  utils.JSONResult{data=response.ResponesGoal}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:id [put]
func UpdateGoal(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("id"))
	if err != nil {
		return err
	}

	data := new(request.RequestUpdateGoal)
	if err := c.BodyParser(data); err != nil {
		return utils.NewError(c, 400, err)
	}

	database.DB.Model(&goal).Updates(utils.StructToMap(data))

	var result response.ResponesGoal
	database.DB.Model(&table.Goal{}).First(&result.Goal, goal)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    result,
	})
}

// @Summary      Goal Delete
// @Description  Delete goal by id
// @Tags         Goals
// @Produce      json
// @Param	  	 id	path	int	true	"Goal ID"
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
