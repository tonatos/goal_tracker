package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tonatos/goal-tracker/internal/handlers/helpers"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

// @Summary      Upload Image
// @Description  Upload image for goal
// @Tags         Upload
// @Produce      json
// @Param	  	 image	path	int	true	"Image file"
// @Success 	 200 {string} status "ok"
// @Failure      400  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router /v1/upload/ [post]
func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return utils.NewError(c, 400, err)
	}

	url, err := helpers.SaveFile(c, file)
	if err != nil {
		return utils.NewError(c, 400, err)
	}

	return c.JSON(utils.JSONResult{
		Code:    200,
		Message: "success",
		Data: map[string]string{
			"image": url,
		},
	})
}
