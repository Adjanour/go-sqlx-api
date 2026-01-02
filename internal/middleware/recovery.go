package middleware

import (
	"log"
	"net/http"

	"github.com/Adjanour/go-sqlx-api/pkg/response"
)

// Recovery recovers from panics and returns a 500 error
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				response.Error(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
