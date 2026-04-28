package room

import (
	"go-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) CreateRoom(c *fiber.Ctx) error {
	var room Room

	if err := c.BodyParser(&room); err != nil {
		return utils.Error(c, utils.NewApiError(400, err.Error()))
	}

	if err := h.Service.CreateRoom(&room); err != nil {
		if err != nil {
			return err
		}
	}

	return utils.JSON(c, 201, room)
}

func (h *Handler) GetRooms(c *fiber.Ctx) error {
	name := c.Query("name")

	filters := map[string]interface{}{}

	if name != "" {
		filters["name"] = name
	}

	rooms, err := h.Service.GetRooms(filters)
	if err != nil {
		if err != nil {
			return err
		}
	}

	return utils.JSON(c, 200, rooms)
}

func (h *Handler) UpdateRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	var room Room

	if err := c.BodyParser(&room); err != nil {
		if err != nil {
			return err
		}
	}

	room.RoomID = id

	if err := h.Service.UpdateRoom(&room); err != nil {
		if err != nil {
			return err
		}
	}

	return utils.JSON(c, 200, room)
}

func (h *Handler) DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.Service.DeleteRoom(id); err != nil {
		if err != nil {
			return err
		}
	}

	return utils.JSON(c, 200, fiber.Map{
		"message": "deleted",
	})
}
