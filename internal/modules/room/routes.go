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

	// admin only
	rooms.Post("/", middleware.RequireRole("admin"), handler.CreateRoom)
	rooms.Put("/:id", middleware.RequireRole("admin"), handler.UpdateRoom)
	rooms.Delete("/:id", middleware.RequireRole("admin"), handler.DeleteRoom)

	// customer + admin
	// rooms.Get("/:id", middleware.RequireRole("admin", "customer"), handler.ge)
}
