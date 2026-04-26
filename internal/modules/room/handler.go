package room

import (
	"go-api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
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
		return utils.Error(c, err)
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
		return utils.Error(c, err)
	}

	return utils.JSON(c, 200, rooms)
}

func (h *Handler) UpdateRoom(c *gin.Context) {
	id := c.Param("id")

	var room Room

	if err := c.ShouldBindJSON(&room); err != nil {
		utils.Error(c, utils.NewApiError(400, err.Error()))
		return
	}

	room.RoomID = id

	if err := h.Service.UpdateRoom(&room); err != nil {
		utils.Error(c, err)
		return
	}

	utils.JSON(c, 200, room)
}
