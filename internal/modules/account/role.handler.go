package account

import (
	"go-api/internal/utils"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	Service *RoleService
}

func NewRoleHandler(s *RoleService) *RoleHandler {
	return &RoleHandler{Service: s}
}

func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var role Role

	if err := c.BodyParser(&role); err != nil {
		return utils.NewApiError(400, err.Error())
	}

	if err := h.Service.Create(&role); err != nil {
		return err
	}

	return c.JSON(role)
}

func (h *RoleHandler) GetRoles(c *fiber.Ctx) error {
	roles, err := h.Service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(roles)
}

func (h *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "invalid id")
	}

	return h.Service.Delete(id)
}
