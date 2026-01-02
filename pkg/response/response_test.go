package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		data       interface{}
	}{
		{
			name:       "Success response",
			statusCode: http.StatusOK,
			data:       map[string]string{"message": "success"},
		},
		{
			name:       "Error response",
			statusCode: http.StatusBadRequest,
			data:       map[string]string{"error": "bad request"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			JSON(w, tt.statusCode, tt.data)

			if w.Code != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, w.Code)
			}

			contentType := w.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("expected Content-Type application/json, got %s", contentType)
			}
		})
	}
}

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{"message": "test"}

	Success(w, http.StatusOK, data)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var resp Response
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !resp.Success {
		t.Error("expected success to be true")
	}

	if resp.Error != "" {
		t.Errorf("expected no error, got %s", resp.Error)
	}
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	errorMsg := "test error"

	Error(w, http.StatusBadRequest, errorMsg)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var resp Response
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Success {
		t.Error("expected success to be false")
	}

	if resp.Error != errorMsg {
		t.Errorf("expected error %s, got %s", errorMsg, resp.Error)
	}
}
