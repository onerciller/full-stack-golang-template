package httpserver

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/onerciller/fullstack-golang-template/pkg/apperror"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		// Handle fiber's built-in errors
		if e, ok := err.(*fiber.Error); ok {
			return ctx.Status(e.Code).JSON(apperror.NewAppError(
				e.Code,
				apperror.CodeInternalError,
				e.Message,
			))
		}

		// Handle our custom AppError
		if appErr, ok := err.(*apperror.AppError); ok {
			return ctx.Status(appErr.Status).JSON(appErr)
		}

		// Handle standard errors
		return ctx.Status(http.StatusInternalServerError).JSON(apperror.InternalError(err.Error()))
	}
}
