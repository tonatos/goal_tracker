package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// NewError example
func NewError(ctx *fiber.Ctx, status int, err error) error {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	log.Printf("API Error [%d]: %s", status, err.Error())
	return ctx.Status(status).JSON(er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func (e HTTPError) Error() string {
	return e.Message
}

type JSONResult struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
