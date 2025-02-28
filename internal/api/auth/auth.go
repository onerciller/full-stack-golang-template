package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onerciller/fullstack-golang-template/internal/api/base"
	"github.com/onerciller/fullstack-golang-template/internal/entity"
	"github.com/onerciller/fullstack-golang-template/internal/errors"
	"github.com/onerciller/fullstack-golang-template/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Handler manages user-related HTTP endpoints
type Handler struct {
	base.Handler
}

// New creates a new user Handler instance
func New(base base.Handler) *Handler {
	return &Handler{
		Handler: base,
	}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/register", h.Register)
	router.Post("/login", h.Login)
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
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.ErrInvalidRequestBody.ToBadRequestAppError()
	}

	if err := h.validateRegistration(c, &req); err != nil {
		return err
	}

	user, err := h.createUser(c, &req)
	if err != nil {
		return err
	}

	tokens, err := h.generateAndStoreTokens(c, user.ID)
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
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.ErrInvalidRequestBody.ToBadRequestAppError()
	}

	user, err := h.validateLogin(c, &req)
	if err != nil {
		return err
	}

	tokens, err := h.generateAndStoreTokens(c, user.ID)
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
		return errors.ErrFailedToGetUsers.ToBadRequestAppError()
	}

	return c.JSON(model.UsersResponse{
		Users: users,
	})
}

func (h *Handler) validateRegistration(c *fiber.Ctx, req *model.RegisterRequest) error {
	existingUser, err := h.UserStore.FindByUsername(c.Context(), req.Username)
	if err == nil && existingUser != nil {
		return errors.ErrUserAlreadyExists.ToBadRequestAppError()
	}
	return nil
}

func (h *Handler) createUser(c *fiber.Ctx, req *model.RegisterRequest) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrFailedToHashPassword.ToBadRequestAppError()
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := h.UserStore.Create(c.Context(), user); err != nil {
		return nil, errors.ErrFailedToCreateUser.ToBadRequestAppError()
	}

	return user, nil
}

func (h *Handler) validateLogin(c *fiber.Ctx, req *model.LoginRequest) (*entity.User, error) {
	user, err := h.UserStore.FindByUsername(c.Context(), req.Username)
	if err != nil {
		return nil, errors.ErrInvalidCredentials.ToUnauthorizedAppError()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.ErrInvalidCredentials.ToUnauthorizedAppError()
	}

	return user, nil
}

func (h *Handler) generateAndStoreTokens(c *fiber.Ctx, userID uint) (*model.AuthResponse, error) {
	accessToken, err := h.JWT.GenerateAccessToken(userID)
	if err != nil {
		return nil, errors.ErrFailedToGenToken.ToBadRequestAppError()
	}

	return &model.AuthResponse{
		AccessToken: accessToken,
	}, nil
}
