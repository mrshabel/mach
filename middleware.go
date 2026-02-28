package mach

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
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

type CORSConfig struct {
	AllowOrigins      []string
	AllowMethods      []string
	AllowHeaders      []string
	ExposeHeaders     []string
	AllowCredentials  bool
	MaxAge            int
	PreflightContinue bool
}

func CORS(allowOrigins []string) MiddlewareFunc {
	return CORSWithConfig(CORSConfig{
		AllowOrigins: allowOrigins,
	})
}

func CORSWithConfig(config CORSConfig) MiddlewareFunc {
	allowAll := false
	origins := make(map[string]struct{}, len(config.AllowOrigins))

	// validate origin
	for _, origin := range config.AllowOrigins {
		if origin == "*" {
			allowAll = true
			break
		}
		origins[origin] = struct{}{}
	}

	defaultMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD"}
	defaultHeaders := []string{"Content-Type", "Authorization"}

	allowMethods := config.AllowMethods
	if len(allowMethods) == 0 {
		allowMethods = defaultMethods
	}

	allowHeaders := config.AllowHeaders
	if len(allowHeaders) == 0 {
		allowHeaders = defaultHeaders
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			if allowAll {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else if _, ok := origins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
			}

			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if len(config.ExposeHeaders) > 0 {
				w.Header().Set("Access-Control-Expose-Headers", strings.Join(config.ExposeHeaders, ", "))
			}

			w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowHeaders, ", "))

			if config.MaxAge > 0 {
				w.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", config.MaxAge))
			}

			if r.Method == "OPTIONS" {
				if config.PreflightContinue {
					next.ServeHTTP(w, r)
				} else {
					w.WriteHeader(http.StatusNoContent)
				}
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
