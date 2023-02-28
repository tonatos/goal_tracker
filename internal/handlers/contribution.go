package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tonatos/goal-tracker/internal/handlers/helpers"
	"github.com/tonatos/goal-tracker/internal/models/request"
	"github.com/tonatos/goal-tracker/internal/models/response"
	"github.com/tonatos/goal-tracker/internal/models/table"
	"github.com/tonatos/goal-tracker/pkg/database"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

// @Summary      Contribution List
// @Description  Get contributions list
// @Tags         Contribution
// @Produce      json
// @Param	  	 goal	path	int	true	"Goal ID for contribution"
// @Success 	 200  {array}  utils.JSONResult{data=[]response.ResponesContribution}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution [get]
func GetContributions(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if err != nil {
		return utils.NewError(c, err.Code, err)
	}

	var contributions []response.ResponesContribution
	query := table.Contribution{Goal: *goal}
	database.DB.Model(&table.Contribution{}).Find(&contributions, query)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    contributions,
	})
}

// @Summary      Contribution Create
// @Description  Create contribuitions
// @Tags         Contribution
// @Produce      json
// @Param	  	 goal	path	int	true	"Goal ID for contribution"
// @Param        contribuition	body	request.RequestCreateContribution true "Contribution object for create"
// @Success 	 200  {object}  utils.JSONResult{data=response.ResponesContribution}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution [post]
func CreateContribution(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if err != nil {
		return utils.NewError(c, err.Code, err)
	}

	json := new(request.RequestCreateContribution)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	newContriibution := table.Contribution{
		GoalID: goal.ID,
		Amount: json.Amount,
	}

	if err := database.DB.Create(&newContriibution).Error; err != nil {
		return utils.NewError(c, 200, err)
	}

	var result response.ResponesContribution
	database.DB.Model(&table.Contribution{}).Joins("Goal").First(&result, newContriibution)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    result,
	})
}

// @Summary      Contribution Update
// @Description  Update contributions by id
// @Tags         Contribution
// @Produce      json
// @Param	  	 goal	path	int	true	"Goal ID for contribution"
// @Param	  	 id	path	int	true	"Contribution ID"
// @Param        contribution	body	request.RequestUpdateContribution true "Contribution`s fields for update"
// @Success 	 200  {object}  utils.JSONResult{data=response.ResponesContribution}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution/:id [put]
func UpdateContribution(c *fiber.Ctx) error {
	contribution, http_err := helpers.GetContributionById(
		c, database.DB, c.Params("goal"), c.Params("id"),
	)
	if http_err != nil {
		return utils.NewError(c, http_err.Code, http_err)
	}

	data := new(request.RequestUpdateContribution)
	if err := c.BodyParser(data); err != nil {
		return utils.NewError(c, 400, err)
	}
	database.DB.Model(&contribution).Updates(utils.StructToMap(data))

	var result response.ResponesContribution
	database.DB.Model(&table.Contribution{}).Joins("Goal").First(&result, contribution)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    result,
	})
}

// @Summary      Contribution Delete
// @Description  Delete contribution by id
// @Tags         Contribution
// @Produce      json
// @Param	  	 goal	path	int	true	"Goal ID for contribution"
// @Param	  	 id	path	int	true	"Contribution ID"
// @Success 	 200 {string} status "ok"
// @Failure      400  {object}  utils.HTTPError
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution/:id [delete]
func DeleteContribution(c *fiber.Ctx) error {
	contribution, http_err := helpers.GetContributionById(
		c, database.DB, c.Params("goal"), c.Params("id"),
	)
	if http_err != nil {
		return utils.NewError(c, http_err.Code, http_err)
	}
	database.DB.Delete(&contribution)
	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
	})
}
