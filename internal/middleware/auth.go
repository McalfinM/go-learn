package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("your_secret")

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")

		if auth == "" {
			return fiber.NewError(401, "Unauthorized")
		}

		tokenString := strings.Replace(auth, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			return fiber.NewError(401, "Invalid token")
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Locals("user_uuid", claims["user_uuid"])
		c.Locals("role_id", claims["role_id"])

		return c.Next()
	}
}

func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")

		if auth == "" {
			return fiber.NewError(401, "Unauthorized")
		}

		tokenString := strings.Replace(auth, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			return fiber.NewError(401, "Invalid token")
		}

		claims := token.Claims.(jwt.MapClaims)
		userRole := claims["role_name"].(string)

		for _, r := range roles {
			if r == userRole {
				c.Locals("user_uuid", claims)
				return c.Next()
			}
		}

		return fiber.NewError(403, "Forbidden")
	}
}
