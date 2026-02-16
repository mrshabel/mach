package mach

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// responseWriter extends the http response writer to capture additional details
type responseWriter struct {
	http.ResponseWriter

	status int
	size   int
	// write header only once
	isHeaderWritten bool
}

func (rw *responseWriter) WriteHeader(status int) {
	if rw.isHeaderWritten {
		return
	}

	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
	rw.isHeaderWritten = true
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	// write status header if not done
	if !rw.isHeaderWritten {
		rw.WriteHeader(http.StatusOK)
	}

	size, err := rw.ResponseWriter.Write(data)
	rw.size += size

	return size, err
}

func Logger() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			// call next handler in chain
			next.ServeHTTP(rw, r)

			// log details. [method] /path address status duration size
			log.Printf("[%s] %s %s - %d (%v) %d bytes", r.Method, r.URL.Path, r.RemoteAddr,
				rw.status, time.Since(start), rw.size)
		})
	}
}

func Recovery() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// log stack trace
					log.Printf("PANIC: %v\n%s", err, debug.Stack())
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func CORS(allowOrigins []string) MiddlewareFunc {
	// build allowed origin list
	origins := make(map[string]struct{}, len(allowOrigins))
	for _, org := range allowOrigins {
		origins[org] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// validate that origin is whitelisted
			origin := r.Header.Get("Origin")

			if _, ok := origins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				// avoid caching origin
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}

			// handle preflight request
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
