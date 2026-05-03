package order

import (
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	BookingService *BookingService
}

func NewUserHandler(s *BookingService) *BookingHandler {
	return &BookingHandler{BookingService: s}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var body struct {
		RoomUUID      string `json:"room_uuid"`
		BookingType   string `json:"booking_type"`
		DurationHours *int   `json:"duration_hours"`

		GuestName string `json:"guest_name"`
		GuestKTP  string `json:"guest_ktp"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	userUUID := c.Locals("user_uuid").(string)

	booking := &Booking{
		BookingType:   body.BookingType,
		DurationHours: body.DurationHours,
	}

	err := h.BookingService.CreateBooking(userUUID, body.RoomUUID, booking, body.GuestName, body.GuestKTP)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "booking created",
	})
}
