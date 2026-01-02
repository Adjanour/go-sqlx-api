package models

import "fmt"

// Error types
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// Common errors
func ErrNotFound(resource string) *AppError {
	return &AppError{
		Code:    404,
		Message: fmt.Sprintf("%s not found", resource),
	}
}

func ErrInvalidInput(msg string) *AppError {
	return &AppError{
		Code:    400,
		Message: msg,
	}
}

func ErrInternalServer(msg string) *AppError {
	return &AppError{
		Code:    500,
		Message: msg,
	}
}
