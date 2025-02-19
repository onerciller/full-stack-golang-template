package model

import "github.com/onerciller/fullstack-golang-template/internal/entity"

// LoginRequest represents user login credentials
// @Description User login request
type LoginRequest struct {
	// Username of the user
	Username string `json:"username" validate:"required,min=3" example:"johndoe"`
	// User's password (min 6 characters)
	Password string `json:"password" validate:"required,min=6" example:"password123"`
}

// RegisterRequest represents user registration data
// @Description User registration request
type RegisterRequest struct {
	// Username (min 3 characters)
	Username string `json:"username" validate:"required,min=3" example:"johndoe"`
	// Email address of the user
	Email string `json:"email" validate:"required,email" example:"john@example.com"`
	// Password (min 6 characters)
	Password string `json:"password" validate:"required,min=6" example:"password123"`
}

// AuthResponse represents authentication tokens response
// @Description Authentication response containing access and refresh tokens
type AuthResponse struct {
	// JWT access token
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	// JWT refresh token
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// UsersResponse represents a list of users response
// @Description Response containing a list of users
type UsersResponse struct {
	// List of users
	Users []*entity.User `json:"users"`
}
