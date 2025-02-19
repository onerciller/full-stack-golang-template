package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func WithHealthCheck(path string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Path() == path {
			return ctx.JSON("ok")
		}
		return ctx.Next()
	}
}

func WithLogger(logger *zap.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Info("request", zap.String("path", ctx.Path()))
		return ctx.Next()
	}
}

func WithCORS() fiber.Handler {
	return cors.New()
}
