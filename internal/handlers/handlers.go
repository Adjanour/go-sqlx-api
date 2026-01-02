package handlers

import (
	"net/http"

	"github.com/Adjanour/go-sqlx-api/pkg/response"
	"github.com/jmoiron/sqlx"
)

// Handlers contains all HTTP handlers
type Handlers struct {
	db *sqlx.DB
}

// NewHandlers creates a new Handlers instance
func NewHandlers(db *sqlx.DB) *Handlers {
	return &Handlers{db: db}
}

// HealthCheck returns the health status of the API
func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Check database connection
	if err := h.db.Ping(); err != nil {
		response.Error(w, http.StatusServiceUnavailable, "Database connection failed")
		return
	}

	response.Success(w, http.StatusOK, map[string]string{
		"status":   "healthy",
		"database": "connected",
	})
}
