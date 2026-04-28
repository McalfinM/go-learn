package account

import (
	"database/sql"
	"go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, db *sql.DB) {
	// accounts := r.Group("/accounts")

	// accounts.Get("/", handler.GetAccounts)
	// accounts.Post("/", middleware.RequireRole("admin"), handler.CreateAccount)
	// accounts.Put("/:id", middleware.RequireRole("admin"), handler.UpdateAccount)
	// accounts.Delete("/:id", middleware.RequireRole("admin"), handler.DeleteAccount)

	repo := NewRoleRepository(db)
	service := NewRoleService(repo)
	handler := NewRoleHandler(service)

	roles := r.Group("/roles")

	roles.Get("/", handler.GetRoles)
	roles.Post("/", middleware.RequireRole("admin"), handler.CreateRole)
	roles.Delete("/:id", middleware.RequireRole("admin"), handler.DeleteRole)
}
