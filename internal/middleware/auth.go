package middleware

import (
	"strings"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("your_secret")

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")

		if auth == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		tokenString := strings.Replace(auth, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		return c.Next()
	}
}
