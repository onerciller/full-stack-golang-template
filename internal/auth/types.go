package auth

import (
	"context"

	"github.com/onerciller/fullstack-golang-template/internal/shared"
)

// UserRepository defines the interface for user-related database operations
type UserStore interface {
	FindByUsername(ctx context.Context, username string) (*UserEntity, error)
	FindByEmail(ctx context.Context, email string) (*UserEntity, error)
	FindByID(ctx context.Context, id uint) (*UserEntity, error)
	FindAll(ctx context.Context) ([]*UserEntity, error)
	Create(ctx context.Context, user *UserEntity) error
	Update(ctx context.Context, user *UserEntity) error
	Delete(ctx context.Context, id uint) error
}

// User represents a user in the system
// @Description User entity with basic information
type UserEntity struct {
	shared.BaseEntity
	// Username of the user
	Username string
	// Email address of the user
	Email string
	// Password (hashed, not exposed in JSON)
	Password string
}

func (u *UserEntity) TableName() string {
	return "users"
}

// LoginRequest represents user login credentials
// @Description User login request
type LoginRequest struct {
	// Username of the user
	Username string `json:"username" validate:"required,min=3"`
	// User's password (min 6 characters)
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterRequest represents user registration data
// @Description User registration request
type RegisterRequest struct {
	// Username (min 3 characters)
	Username string `json:"username" validate:"required,min=3"`
	// Email address of the user
	Email string `json:"email" validate:"required,email"`
	// Password (min 6 characters)
	Password string `json:"password" validate:"required,min=6"`
}

// AuthResponse represents authentication tokens response
// @Description Authentication response containing access and refresh tokens
type AuthResponse struct {
	// JWT access token
	AccessToken string `json:"access_token"`
}

// UsersResponse represents a list of users response
// @Description Response containing a list of users
type UsersResponse struct {
	// List of users
	Users []*UserEntity `json:"users"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
