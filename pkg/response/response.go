package response

import (
	"encoding/json"
	"net/http"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// JSON writes a JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		// If encoding fails, headers are already sent - nothing we can do
		// In production, this should be logged for monitoring
		_ = err
	}
}

// Success writes a success response
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	JSON(w, statusCode, Response{
		Success: true,
		Data:    data,
	})
}

// Error writes an error response
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, Response{
		Success: false,
		Error:   message,
	})
}
