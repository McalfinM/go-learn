package server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"go-api/internal/modules/room"
)

func New(db *sql.DB) *fiber.App {
	app := fiber.New()

	v1 := app.Group("/v1")

	// base route
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API v1 base route",
		})
	})

	// modules
	room.RegisterRoutes(v1, db)

	return app
}
