package account

import (
	"go-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{UserService: s}
}

// REGISTER
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return utils.NewApiError(400, err.Error())
	}

	if err := h.UserService.Register(body.Email, body.Password); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Register Success",
	})
}

// LOGIN
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return utils.NewApiError(400, err.Error())
	}

	token, err := h.UserService.Login(body.Email, body.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) UploadKTP(c *fiber.Ctx) error {
	file, err := c.FormFile("ktp")
	if err != nil {
		return fiber.NewError(400, "file required")
	}

	userUUID := c.Locals("user_uuid").(string)

	url, err := h.UserService.UploadKTP(userUUID, file)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"url": url,
	})
}
