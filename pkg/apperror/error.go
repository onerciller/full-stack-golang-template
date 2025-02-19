package apperror

import (
	"net/http"
)

// Error codes for HTTP responses
const (
	CodeBadRequest     = "BAD_REQUEST"
	CodeUnauthorized   = "UNAUTHORIZED"
	CodeForbidden      = "FORBIDDEN"
	CodeNotFound       = "NOT_FOUND"
	CodeConflict       = "CONFLICT"
	CodeInternalError  = "INTERNAL_ERROR"
	CodeValidationFail = "VALIDATION_FAILED"
)

// AppError represents an application error with HTTP status code and error details
// @Description Application error response
type AppError struct {
	// HTTP status code
	Status int `json:"status" example:"400"`
	// Error code
	Code string `json:"code" example:"BAD_REQUEST"`
	// Error message
	Message string `json:"message" example:"Invalid request parameters"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new AppError
func NewAppError(status int, code, message string) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

// HTTP error responses
func BadRequest(message string) *AppError {
	return NewAppError(http.StatusBadRequest, CodeBadRequest, message)
}

func Unauthorized(message string) *AppError {
	return NewAppError(http.StatusUnauthorized, CodeUnauthorized, message)
}

func Forbidden(message string) *AppError {
	return NewAppError(http.StatusForbidden, CodeForbidden, message)
}

func NotFound(message string) *AppError {
	return NewAppError(http.StatusNotFound, CodeNotFound, message)
}

func Conflict(message string) *AppError {
	return NewAppError(http.StatusConflict, CodeConflict, message)
}

func InternalError(message string) *AppError {
	return NewAppError(http.StatusInternalServerError, CodeInternalError, message)
}

func ValidationFailed(message string) *AppError {
	return NewAppError(http.StatusBadRequest, CodeValidationFail, message)
}
