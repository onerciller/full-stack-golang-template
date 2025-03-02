package shared

import "github.com/onerciller/fullstack-golang-template/pkg/apperror"

// DomainErrorCode represents a domain-specific error code
type AppErrorCode string

// Domain error codes
const (
	// User domain errors
	ErrUserNotFound         AppErrorCode = "USER_NOT_FOUND"
	ErrUserAlreadyExists    AppErrorCode = "USER_ALREADY_EXISTS"
	ErrInvalidCredentials   AppErrorCode = "INVALID_CREDENTIALS"
	ErrFailedToCreateUser   AppErrorCode = "FAILED_TO_CREATE_USER"
	ErrFailedToHashPassword AppErrorCode = "FAILED_TO_HASH_PASSWORD"
	ErrFailedToGenToken     AppErrorCode = "FAILED_TO_GEN_TOKEN"
	ErrFailedToStoreToken   AppErrorCode = "FAILED_TO_STORE_TOKEN"
	// Authentication domain errors
	ErrInvalidRequestBody      AppErrorCode = "INVALID_REQUEST_BODY"
	ErrInvalidToken            AppErrorCode = "INVALID_TOKEN"
	ErrMissingAuthHeader       AppErrorCode = "MISSING_AUTH_HEADER"
	ErrInvalidAuthHeaderFormat AppErrorCode = "INVALID_AUTH_HEADER_FORMAT"
	ErrFailedToUpdateUser      AppErrorCode = "FAILED_TO_UPDATE_USER"
	ErrFailedToGetUsers        AppErrorCode = "FAILED_TO_GET_USERS"
	// Sorting center domain errors
)

// GetMessage returns a human-readable message for the DomainErrorCode
func (code AppErrorCode) GetMessage() string {
	switch code {
	// User related messages
	case ErrUserNotFound:
		return "User not found"
	case ErrUserAlreadyExists:
		return "User already exists"
	case ErrInvalidCredentials:
		return "Invalid credentials"
	case ErrInvalidToken:
		return "Invalid token"
	case ErrMissingAuthHeader:
		return "Missing authorization header"
	case ErrInvalidAuthHeaderFormat:
		return "Invalid authorization header format"
	case ErrInvalidRequestBody:
		return "Invalid request body"
	case ErrFailedToCreateUser:
		return "Failed to create user"
	case ErrFailedToHashPassword:
		return "Failed to hash password"
	case ErrFailedToGenToken:
		return "Failed to generate token"
	case ErrFailedToStoreToken:
		return "Failed to store token"
	case ErrFailedToUpdateUser:
		return "Failed to update user"
	case ErrFailedToGetUsers:
		return "Failed to get users"
	default:
		return "Unknown error occurred"
	}
}

// ToNotFoundAppError converts a DomainErrorCode to an AppError
func (code AppErrorCode) ToNotFoundAppError() *apperror.AppError {
	return apperror.NotFound(code.GetMessage())
}

// ToBadRequestAppError converts a DomainErrorCode to an AppError
func (code AppErrorCode) ToBadRequestAppError() *apperror.AppError {
	return apperror.BadRequest(code.GetMessage())
}

// ToUnauthorizedAppError converts a DomainErrorCode to an AppError
func (code AppErrorCode) ToUnauthorizedAppError() *apperror.AppError {
	return apperror.Unauthorized(code.GetMessage())
}
