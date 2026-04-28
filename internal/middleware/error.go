package middleware

import (
	"go-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	// fiber error
	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(fiber.Map{
			"status": e.Code,
			"error":  e.Message,
		})
	}

	// custom error
	if apiErr, ok := err.(*utils.ApiError); ok {
		return c.Status(apiErr.StatusCode).JSON(fiber.Map{
			"status": apiErr.StatusCode,
			"error":  apiErr.Message,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"status": 500,
		"error":  err.Error(),
	})
}
