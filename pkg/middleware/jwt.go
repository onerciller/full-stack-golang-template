package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/onerciller/fullstack-golang-template/internal/errors"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
)

// JWTAuth creates a JWT authentication middleware
func JWTAuth(jwtService jwt.Jwt) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return errors.ErrMissingAuthHeader.ToUnauthorizedAppError()
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return errors.ErrInvalidAuthHeaderFormat.ToUnauthorizedAppError()
		}

		// Extract the token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		userID, err := jwtService.ValidateToken(token)
		if err != nil {
			return errors.ErrInvalidToken.ToUnauthorizedAppError()
		}

		// Store user ID in context for later use
		c.Locals("userID", userID)

		return c.Next()
	}
}
