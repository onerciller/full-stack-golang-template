package middleware

import (
	"strings"

	"github.com/onerciller/fullstack-golang-template/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtService *jwt.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "auth header is missing",
			})
		}

		// Check if the Authorization header has the Bearer scheme
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid auth header format",
			})
		}

		// Extract the token
		tokenString := parts[1]

		// Validate the token
		userID, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Set the user ID in the context for later use
		c.Locals("userID", userID)

		return c.Next()
	}
}
