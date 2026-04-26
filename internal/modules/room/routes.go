package room

import (
	"database/sql"

	"go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, db *sql.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	rooms := r.Group("/rooms")

	// public
	rooms.Get("/", handler.GetRooms)

	// protected
	rooms.Use(middleware.AuthMiddleware())

	rooms.Post("/", handler.CreateRoom)
	rooms.Put("/:id", handler.UpdateRoom)
	rooms.Delete("/:id", handler.DeleteRoom)
}
