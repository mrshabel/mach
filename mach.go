// Package mach provides a lightweight web framework for Go.
//
// Mach is built on Go 1.22's enhanced net/http router with zero dependencies.
// It provides a simple, intuitive API for building web applications while
// leveraging the standard library's performance and reliability.
//
// Example usage:
//
//	app := mach.Default()
//
//	app.GET("/", func(c *mach.Context) {
//	    c.JSON(200, map[string]string{"message": "Hello, Mach!"})
//	})
//
//	app.Run(":8080")
//
// Features:
//   - Go 1.22+ native routing with method matching and path parameters
//   - Standard http.Handler middleware pattern
//   - Route groups for organization
//   - Zero dependencies
package mach

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// App is the main application instance
type App struct {
	router      *http.ServeMux
	middlewares []MiddlewareFunc
	pool        sync.Pool
	logger      *log.Logger
	// whether to run the app in debug mode or not
	debug bool
}

// HandlerFunc is the handler signature
type HandlerFunc func(c *Context)

// MiddlewareFunc is the middleware signature
type MiddlewareFunc func(http.Handler) http.Handler

// Option configures the app
type Option func(*App)

// RunOption configures the server
type RunOption func(*serverConfig)

type serverConfig struct {
	readTimeout     time.Duration
	writeTimeout    time.Duration
	gracefulTimeout time.Duration
}

// New instantiates a new app instance
func New(opts ...Option) *App {
	app := &App{
		router: http.NewServeMux(),
		logger: log.New(os.Stdout, "[mach] ", log.LstdFlags),
	}

	// setup context pool
	app.pool.New = func() interface{} {
		return &Context{app: app}
	}

	// apply server configuration
	for _, opt := range opts {
		opt(app)
	}

	return app
}

// Default instantiates an app with common settings
func Default() *App {
	app := New()
	app.Use(Logger())
	app.Use(Recovery())

	return app
}

// app configuration

// WithLogger adds logger middleware
func WithLogger() Option {
	return func(app *App) {
		app.Use(Logger())
	}
}

// WithRecovery adds recovery middleware
func WithRecovery() Option {
	return func(app *App) {
		app.Use(Recovery())
	}
}

// WithDebug enables debug mode
func WithDebug() Option {
	return func(app *App) {
		app.debug = true
	}
}

// app core methods

// ServeHTTP implements the handler for serving each request
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := app.buildHandler()
	handler.ServeHTTP(w, r)
}

// buildHandler creates the middleware chain
func (app *App) buildHandler() http.Handler {
	handler := http.Handler(app.router)

	// chain middleware in reverse order so they execute in order which they were added
	for idx := len(app.middlewares) - 1; idx >= 0; idx-- {
		handler = app.middlewares[idx](handler)
	}

	return handler
}

// Use adds a global middleware to the application
func (app *App) Use(middlewares ...MiddlewareFunc) {
	app.middlewares = append(app.middlewares, middlewares...)
}

// internal routing method to fit the http.Handler method signature
func (app *App) handle(method, path string, handler HandlerFunc) {
	// std lib pattern is space delimiter of method and path
	pattern := method + " " + path

	// pick request context from pool, process and drop it back
	app.router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := app.pool.Get().(*Context)
		c.reset(w, r)
		handler(c)
		app.pool.Put(c)
	})
}

// routing methods

// Route registers a handler for the given method and path
func (app *App) Route(method, path string, handler HandlerFunc) {
	app.handle(method, path, handler)
}

// GET registers a GET route
func (app *App) GET(path string, handler HandlerFunc) {
	app.handle(http.MethodGet, path, handler)
}

// POST registers a POST route
func (app *App) POST(path string, handler HandlerFunc) {
	app.handle(http.MethodPost, path, handler)
}

// PATCH registers a PATCH route
func (app *App) PATCH(path string, handler HandlerFunc) {
	app.handle(http.MethodPatch, path, handler)
}

// PUT registers a PUT route
func (app *App) PUT(path string, handler HandlerFunc) {
	app.handle(http.MethodPut, path, handler)
}

// DELETE registers a DELETE route
func (app *App) DELETE(path string, handler HandlerFunc) {
	app.handle(http.MethodDelete, path, handler)
}

// HEAD registers a HEAD route
func (app *App) HEAD(path string, handler HandlerFunc) {
	app.handle(http.MethodHead, path, handler)
}

// OPTIONS registers a OPTIONS route
func (app *App) OPTIONS(path string, handler HandlerFunc) {
	app.handle(http.MethodOptions, path, handler)
}

func (app *App) Static(prefix, dir string) {
	// get file server
	fileServer := http.FileServer(http.Dir(dir))
	app.router.Handle(fmt.Sprintf("GET %s{path...}", prefix), http.StripPrefix(prefix, fileServer))
}

// Group creates a route group with common prefix and middleware
func (app *App) Group(prefix string, middlewares ...MiddlewareFunc) *Group {
	return &Group{
		prefix:      prefix,
		middlewares: middlewares,
		app:         app,
	}
}

// run options

// WithReadTimeout sets the server read timeout
func WithReadTimeout(duration time.Duration) RunOption {
	return func(cfg *serverConfig) {
		cfg.readTimeout = duration
	}
}

// WithReadTimeout sets the server write timeout
func WithWriteTimeout(duration time.Duration) RunOption {
	return func(cfg *serverConfig) {
		cfg.writeTimeout = duration
	}
}

// WithReadTimeout sets the server graceful shutdown timeout
func WithGracefulShutdown(timeout time.Duration) RunOption {
	return func(cfg *serverConfig) {
		cfg.gracefulTimeout = timeout
	}
}

// server methods

// Run starts the http server
func (app *App) Run(addr string, opts ...RunOption) error {
	cfg := &serverConfig{
		readTimeout:     30 * time.Second,
		writeTimeout:    30 * time.Second,
		gracefulTimeout: 0,
	}

	// apply provided options
	for _, opt := range opts {
		opt(cfg)
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      app,
		ReadTimeout:  cfg.readTimeout,
		WriteTimeout: cfg.writeTimeout,
	}

	if cfg.gracefulTimeout > 0 {
		return app.runWithGracefulShutdown(server, cfg.gracefulTimeout)
	}

	app.logger.Printf("Server started on %s", addr)
	return server.ListenAndServe()
}

// RunTLS starts the HTTPS server
func (app *App) RunTLS(addr, certFile, keyFile string, opts ...RunOption) error {
	cfg := &serverConfig{
		readTimeout:  30 * time.Second,
		writeTimeout: 30 * time.Second,
	}

	// apply provided options
	for _, opt := range opts {
		opt(cfg)
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      app,
		ReadTimeout:  cfg.readTimeout,
		WriteTimeout: cfg.writeTimeout,
	}

	app.logger.Printf("Server started on %s", addr)
	return server.ListenAndServeTLS(certFile, keyFile)
}

func (app *App) runWithGracefulShutdown(server *http.Server, timeout time.Duration) error {
	errChan := make(chan error, 1)

	// run non blocking server
	go func(addr string) {
		app.logger.Printf("Server started on %s with graceful shutdown", server.Addr)
		// send error reports
		errChan <- server.ListenAndServe()
	}(server.Addr)

	// wait for shutdown or os interrupts
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	// server error
	case err := <-errChan:
		return err
	// os interrupt
	case <-quitChan:
		app.logger.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return err
		}

		app.logger.Println("Server shutdown complete")
		return nil
	}
}
