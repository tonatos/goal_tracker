package helpers

import (
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tonatos/goal-tracker/pkg/utils"
)

func SaveFile(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	image := fmt.Sprintf(
		"%s.%s",
		strings.Replace(uuid.New().String(), "-", "", -1),
		strings.Split(file.Filename, ".")[1],
	)
	imagePath := fmt.Sprintf("public/images/%s", image)
	err := c.SaveFile(file, fmt.Sprintf("%s/%s", utils.GetBaseDir(), imagePath))
	if err != nil {
		log.Println("Image save error --> ", err)
		return "", err
	}
	return imagePath, nil
}
