package models

import "time"

// User represents a user in the system
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

// Validate validates the CreateUserRequest
func (r *CreateUserRequest) Validate() error {
	if r.Username == "" {
		return ErrInvalidInput("username is required")
	}
	if r.Email == "" {
		return ErrInvalidInput("email is required")
	}
	return nil
}

// Validate validates the UpdateUserRequest
func (r *UpdateUserRequest) Validate() error {
	if r.Username == "" && r.Email == "" {
		return ErrInvalidInput("at least one field must be provided")
	}
	return nil
}
