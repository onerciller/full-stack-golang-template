package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/onerciller/fullstack-golang-template/internal/shared"

	"golang.org/x/crypto/bcrypt"
)

// Handler manages user-related HTTP endpoints
type Handler struct {
	shared.BaseHandler
	UserStore UserStore
}

// New creates a new user Handler instance
func NewHandler(base *shared.BaseHandler, userStore UserStore) *Handler {
	return &Handler{
		BaseHandler: *base,
		UserStore:   userStore,
	}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/auth/register", h.Register)
	router.Post("/auth/login", h.Login)
	router.Get("/api/v1/user", h.GetUser)
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "User registration details"
// @Success 201 {object} model.AuthResponse "Returns access and refresh tokens"
// @Failure 400 {object} apperror.AppError "Invalid request or user already exists"
// @Router /auth/register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return shared.ErrInvalidRequestBody.ToBadRequestAppError()
	}

	if err := h.validateRegistration(c, &req); err != nil {
		return err
	}

	user, err := h.createUser(c, &req)
	if err != nil {
		return err
	}

	tokens, err := h.generateAndStoreTokens(user.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(tokens)
}

// Login handles user authentication
// @Summary Login user
// @Description Authenticate a user and return access tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "User login credentials"
// @Success 200 {object} model.AuthResponse "Returns access and refresh tokens"
// @Failure 400 {object} apperror.AppError "Invalid request"
// @Failure 401 {object} apperror.AppError "Invalid credentials"
// @Router /auth/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return shared.ErrInvalidRequestBody.ToBadRequestAppError()
	}

	user, err := h.validateLogin(c, &req)
	if err != nil {
		return err
	}

	tokens, err := h.generateAndStoreTokens(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(tokens)
}

// GetUsers returns a list of all users
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.UsersResponse
// @Failure 401 {object} apperror.AppError "Unauthorized"
// @Failure 400 {object} apperror.AppError "Bad Request"
// @Router /api/v1/users [get]
func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.UserStore.FindAll(c.Context())
	if err != nil {
		return shared.ErrFailedToGetUsers.ToBadRequestAppError()
	}

	return c.JSON(UsersResponse{
		Users: users,
	})
}

func (h *Handler) validateRegistration(c *fiber.Ctx, req *RegisterRequest) error {
	existingUser, err := h.UserStore.FindByUsername(c.Context(), req.Username)
	if err == nil && existingUser != nil {
		return shared.ErrUserAlreadyExists.ToBadRequestAppError()
	}
	return nil
}

func (h *Handler) createUser(c *fiber.Ctx, req *RegisterRequest) (*UserEntity, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, shared.ErrFailedToHashPassword.ToBadRequestAppError()
	}

	user := &UserEntity{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := h.UserStore.Create(c.Context(), user); err != nil {
		return nil, shared.ErrFailedToCreateUser.ToBadRequestAppError()
	}

	return user, nil
}

func (h *Handler) validateLogin(c *fiber.Ctx, req *LoginRequest) (*UserEntity, error) {
	user, err := h.UserStore.FindByUsername(c.Context(), req.Username)
	if err != nil {
		return nil, shared.ErrInvalidCredentials.ToUnauthorizedAppError()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, shared.ErrInvalidCredentials.ToUnauthorizedAppError()
	}

	return user, nil
}

func (h *Handler) generateAndStoreTokens(userID uint) (*AuthResponse, error) {
	accessToken, err := h.JWT.GenerateAccessToken(userID)
	if err != nil {
		return nil, shared.ErrFailedToGenToken.ToBadRequestAppError()
	}

	return &AuthResponse{
		AccessToken: accessToken,
	}, nil
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
	return c.JSON(UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}
