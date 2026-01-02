package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Adjanour/go-sqlx-api/internal/models"
	"github.com/Adjanour/go-sqlx-api/pkg/response"
)

// HandleUsers handles requests to /api/v1/users
func (h *Handlers) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listUsers(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// HandleUser handles requests to /api/v1/users/{id}
func (h *Handlers) HandleUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")
	if path == "" {
		response.Error(w, http.StatusBadRequest, "User ID is required")
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getUser(w, r, id)
	case http.MethodPut:
		h.updateUser(w, r, id)
	case http.MethodDelete:
		h.deleteUser(w, r, id)
	default:
		response.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// listUsers retrieves all users
func (h *Handlers) listUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	query := `SELECT id, username, email, created_at, updated_at FROM users ORDER BY id`
	if err := h.db.Select(&users, query); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	response.Success(w, http.StatusOK, users)
}

// createUser creates a new user
func (h *Handlers) createUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	query := `
		INSERT INTO users (username, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, created_at, updated_at
	`

	now := time.Now()
	err := h.db.QueryRowx(query, req.Username, req.Email, now, now).StructScan(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			response.Error(w, http.StatusConflict, "User with this username or email already exists")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	response.Success(w, http.StatusCreated, user)
}

// getUser retrieves a user by ID
func (h *Handlers) getUser(w http.ResponseWriter, r *http.Request, id int) {
	var user models.User

	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1`
	err := h.db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Error(w, http.StatusNotFound, "User not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	response.Success(w, http.StatusOK, user)
}

// updateUser updates a user
func (h *Handlers) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Build dynamic update query
	updates := []string{}
	args := []interface{}{}
	argPos := 1

	if req.Username != "" {
		updates = append(updates, "username = $"+strconv.Itoa(argPos))
		args = append(args, req.Username)
		argPos++
	}

	if req.Email != "" {
		updates = append(updates, "email = $"+strconv.Itoa(argPos))
		args = append(args, req.Email)
		argPos++
	}

	updates = append(updates, "updated_at = $"+strconv.Itoa(argPos))
	args = append(args, time.Now())
	argPos++

	args = append(args, id)

	query := `
		UPDATE users
		SET ` + strings.Join(updates, ", ") + `
		WHERE id = $` + strconv.Itoa(argPos) + `
		RETURNING id, username, email, created_at, updated_at
	`

	var user models.User
	err := h.db.QueryRowx(query, args...).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Error(w, http.StatusNotFound, "User not found")
			return
		}
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			response.Error(w, http.StatusConflict, "User with this username or email already exists")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	response.Success(w, http.StatusOK, user)
}

// deleteUser deletes a user
func (h *Handlers) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	query := `DELETE FROM users WHERE id = $1`
	result, err := h.db.Exec(query, id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	if rowsAffected == 0 {
		response.Error(w, http.StatusNotFound, "User not found")
		return
	}

	response.Success(w, http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}
