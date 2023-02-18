package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"goal-tracker/api/database"
	"goal-tracker/api/handlers/helpers"
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
	query := database.Contribution{Goal: *goal}
	database.DB.Model(&database.Contribution{}).Find(&contributions, query)

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
// @Success 	 200  {object}  utils.JSONResult{data=database.Contribution}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution [post]
func CreateContribution(c *fiber.Ctx) error {
	goal, err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if err != nil {
		return utils.NewError(c, err.Code, err)
	}

	json := new(database.Contribution)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	newContriibution := database.Contribution{
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
// @Success 	 200  {object}  utils.JSONResult{data=database.Contribution}
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/goal/:goal/contribution/:id [put]
func UpdateContribution(c *fiber.Ctx) error {
	_, http_err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if http_err != nil {
		return utils.NewError(c, http_err.Code, http_err)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewError(c, 400, utils.HTTPError{
			Message: "Invalid id",
		})
	}

	json := new(database.Contribution)
	if err := c.BodyParser(json); err != nil {
		return utils.NewError(c, 400, err)
	}

	var found database.Contribution
	query := database.Contribution{ID: uint(id)}

	if err := database.DB.Joins("Goal").First(&found, query).Error; err == gorm.ErrRecordNotFound {
		return utils.NewError(c, 400, utils.HTTPError{
			Message: "Can't find contribution with this id",
		})
	}
	database.DB.Model(&found).Updates(json)

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data:    found,
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
	_, http_err := helpers.GetGoalById(c, database.DB, c.Params("goal"))
	if http_err != nil {
		return utils.NewError(c, http_err.Code, http_err)
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewError(c, 400, utils.HTTPError{
			Message: "Invalid id",
		})
	}

	var found database.Contribution
	query := database.Contribution{ID: uint(id)}

	if err := database.DB.First(&found, query).Error; err == gorm.ErrRecordNotFound {
		return utils.NewError(c, 400, utils.HTTPError{
			Message: "Can't find contribution with this id",
		})
	}

	database.DB.Delete(&found)
	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
	})
}
