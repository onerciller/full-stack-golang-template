package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/onerciller/fullstack-golang-template/internal/api/base"
	"github.com/onerciller/fullstack-golang-template/internal/model"
)

type Handler struct {
	base.Handler
}

func New(base base.Handler) *Handler {
	return &Handler{base}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Get("/user", h.GetUser)
}

// GetUser returns the current user
// @Summary Get current user
// @Description Get the current user's details
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "Current user details"
// @Failure 401 {object} apperror.AppError "Unauthorized"
// @Failure 500 {object} apperror.AppError "Internal server error"
// @Security BearerAuth
// @Router /api/v1/user [get]
func (h *Handler) GetUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	user, err := h.UserStore.FindByID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(model.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}
