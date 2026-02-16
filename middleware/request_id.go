package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/mrshabel/mach"
)

type contextKey string

const requestIDKey contextKey = "requestID"

// generateRequestID generates a 128-bit random id
func generateRequestID() string {
	var b [16]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return ""
	}

	return hex.EncodeToString(b[:])
}

// RequestID adds a unique request ID to each request
func RequestID() mach.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get("X-Request-ID")

			if id == "" {
				id = generateRequestID()
			}

			// set context value and header
			ctx := context.WithValue(r.Context(), requestIDKey, id)
			w.Header().Set("X-Request-ID", id)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}
