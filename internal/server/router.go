package server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"go-api/internal/middleware"
	"go-api/internal/modules/account"
	"go-api/internal/modules/room"
)

func New(db *sql.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.GlobalErrorHandler,
	})

	v1 := app.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API v1 base route",
		})
	})

	room.RegisterRoutes(v1, db)
	account.RegisterRoutes(v1, db)

	return app
}
