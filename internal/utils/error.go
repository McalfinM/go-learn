package utils

import "github.com/gofiber/fiber/v2"

type ApiError struct {
	StatusCode int
	Message    string
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(code int, message string) *ApiError {
	return &ApiError{
		StatusCode: code,
		Message:    message,
	}
}

func Error(c *fiber.Ctx, err error) error {
	if apiErr, ok := err.(*ApiError); ok {
		return c.Status(apiErr.StatusCode).JSON(fiber.Map{
			"status": apiErr.StatusCode,
			"error":  apiErr.Message,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"status": 500,
		"error":  "Internal Server Error",
	})
}
