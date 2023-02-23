package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"goal-tracker/api/database"
	"goal-tracker/api/handlers/helpers"
	db_model "goal-tracker/api/models/db"
	"goal-tracker/api/utils"
)

type ContributionRespones struct {
	Id        uint      `json:"id"`
	Amount    float32   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// @Summary      Contribution List
// @Description  Get contributions list
// @Tags         Contribution
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=[]ContributionRespones}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution [get]
func GetContributions(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if err != nil {
		return utils.NewError(c, err.Code, err)
	}

	var contributions []ContributionRespones
	query := db_model.Contribution{Goal: *goal}
	database.DB.Model(&db_model.Contribution{}).Find(&contributions, query)

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
// @Success 	 200  {object}  utils.JSONResult{data=db_model.Contribution}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution [post]
func CreateContribution(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if err != nil {
		return utils.NewError(c, err.Code, err)
	}

	json := new(db_model.Contribution)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	newContriibution := db_model.Contribution{
		GoalID: goal.ID,
		Amount: json.Amount,
	}

	if err := database.DB.Create(&newContriibution).Error; err != nil {
		return utils.NewError(c, 200, err)
	}

	database.DB.Joins("Goal").First(&newContriibution, newContriibution)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    newContriibution,
	})
}

// @Summary      Contribution Update
// @Description  Update contributions by id
// @Tags         Contribution
// @Produce      json
// @Success 	 200  {object}  utils.JSONResult{data=db_model.Contribution}
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

	json := new(db_model.Contribution)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}
	database.DB.Model(&contribution).Updates(json)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    contribution,
	})
}

// @Summary      Contribution Delete
// @Description  Delete contribution by id
// @Tags         Contribution
// @Produce      json
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
