# Mach - A Lightweight Go Web Framework

[![Go Reference](https://pkg.go.dev/badge/github.com/mrshabel/mach.svg)](https://pkg.go.dev/github.com/mrshabel/mach)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrshabel/mach)](https://goreportcard.com/report/github.com/mrshabel/mach)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/mrshabel/mach.svg)](https://github.com/mrshabel/mach/releases)

Mach is a minimalist web framework for Go, inspired by Python's Bottle. My motivation for building this was to go deeper into the internals of the frameworks I've been using to build HTTP backend services. Rather than reinventing the wheel with a custom radix tree router, I leverage Go 1.22's enhanced `net/http` server, is incredibly fast with minimal allocations. The framework also uses context pooling to reduce allocations in the hot path. I hope you find this tool useful.

## Features

- **Middleware chain** - Global and route-specific middlewares
- **Route groups** - Organize routes with common prefixes and middleware
- **Flexible configuration** - Functional options for app and server setup
- **Graceful shutdown** - Built-in support for clean server termination
- **Zero dependencies** - Core framework uses only the standard library

## Quick Start

### Installation

```bash
go get github.com/mrshabel/mach
```

### Basic Usage

```go
package main

import (
    "github.com/mrshabel/mach"
)

func main() {
    // create app with default middleware (logger + recovery)
    app := mach.Default()

    // simple route
    app.GET("/", func(c *mach.Context) {
        c.Text(200, "Hello, Mach!")
    })

    // path parameters
    app.GET("/users/{id}", func(c *mach.Context) {
        id := c.Param("id")
        c.JSON(200, map[string]string{"user_id": id})
    })

    // JSON binding
    app.POST("/users", func(c *mach.Context) {
        var user struct {
            Name  string `json:"name"`
            Email string `json:"email"`
        }

        if err := c.BindJSON(&user); err != nil {
            c.JSON(400, map[string]string{"error": err.Error()})
            return
        }

        c.JSON(201, user)
    })

    app.Run(":8080")
}
```

## Examples

### Route Groups

```go
app := mach.Default()

// public routes
app.GET("/", homeHandler)
app.GET("/about", aboutHandler)

// api group with CORS
api := app.Group("/api", mach.CORS("*"))
{
    // v1 endpoints
    v1 := api.Group("/v1")
    {
        v1.GET("/users", listUsers)
        v1.POST("/users", createUser)
        v1.GET("/users/{id}", getUser)
    }
}

// admin group with authentication
admin := app.Group("/admin", mach.BasicAuth("admin", "secret"))
{
    admin.GET("/dashboard", dashboardHandler)
    admin.GET("/settings", settingsHandler)
}
```

### Custom Middleware

```go
// custom logger middleware
func CustomLogger() mach.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()

            next.ServeHTTP(w, r)

            log.Printf("%s %s (%v)", r.Method, r.URL.Path, time.Since(start))
        })
    }
}

app := mach.New()
app.Use(CustomLogger())
app.Use(mach.Recovery())
```

### Configuration Options

```go
// app with options
app := mach.New(
    mach.WithLogger(),
    mach.WithRecovery(),
    mach.WithDebug(),
)

// Server with options
app.Run(":8080",
    mach.WithReadTimeout(10*time.Second),
    mach.WithWriteTimeout(10*time.Second),
    mach.WithGracefulShutdown(30*time.Second),
)
```

### Static Files

```go
app := mach.Default()

// serve static files from ./public
app.Static("/static/", "./public")

// Access at http://localhost:8000/static/css/style.css
```

## Context API

### Request Data

```go
app.GET("/example", func(c *mach.Context) {
    // path parameters
    id := c.Param("id")

    // query parameters
    page := c.Query("page")
    pageOrDefault := c.DefaultQuery("page", "1")

    // form data
    name := c.Form("name")

    // headers
    token := c.GetHeader("Authorization")

    // cookies
    cookie, err := c.Cookie("session")

    // client IP
    ip := c.ClientIP()

    // Request body
    body, err := c.Body()
})
```

### Response Methods

```go
// JSON response
c.JSON(200, map[string]string{"status": "ok"})

// text response
c.Text(200, "Hello, %s", name)

// HTML response
c.HTML(200, "<h1>Welcome</h1>")

// XML response
c.XML(200, data)

// Raw data
c.Data(200, "application/octet-stream", bytes)

// no content
c.NoContent(204)

// redirect
c.Redirect(302, "/login")
```

### Binding

```go
// JSON decoding
var user User
if err := c.DecodeJSON(&user); err != nil {
    c.JSON(400, map[string]string{"error": err.Error()})
    return
}

// XML decoding
var config Config
if err := c.DecodeXML(&config); err != nil {
    c.JSON(400, map[string]string{"error": err.Error()})
    return
}
```

## Built-in Middleware

- **Logger()** - Request logging with timing
- **Recovery()** - Panic recovery with stack traces
- **CORS(origin)** - CORS headers configuration
- **RequestID()** - Unique request ID generation

## Design Philosophy

Mach follows these principles:

1. **Leverage the standard library** - Use Go 1.22's enhanced routing instead of custom implementations
2. **Performance matters** - Context pooling and efficient middleware chaining
3. **Simplicity over features** - Clean API inspired by Bottle's minimalism
4. **Standard Go patterns** - No magic, just idiomatic Go code
5. **Zero core dependencies** - Framework core uses only stdlib

## Technical Details

### Routing

Mach uses Go 1.22+'s `net/http.ServeMux` pattern matching:

```go
// method-specific routes
GET /users/{id}
POST /users
PUT /users/{id}

// wildcard patterns
GET /files/{path...}
```

Pattern matching is handled by the standard library with [precedence rules](https://go.dev/blog/routing-enhancements) that ensure the most specific pattern wins.

### Context Pooling

To minimize allocations, Context objects are pooled using `sync.Pool`:

```go
type App struct {
    pool sync.Pool
}

func (app *App) handle(method, path string, handler HandlerFunc) {
    app.router.HandleFunc(method+" "+path, func(w http.ResponseWriter, r *http.Request) {
        c := app.pool.Get().(*Context)
        c.reset(w, r)
        handler(c)
		// return to pool
        app.pool.Put(c)
    })
}
```

This reduces GC pressure in high-throughput scenarios.

### Middleware Chain

Middleware uses the standard `func(http.Handler) http.Handler` pattern:

```go
func (app *App) buildHandler() http.Handler {
    handler := http.Handler(app.router)

    // apply middleware in reverse order
    for i := len(app.middleware) - 1; i >= 0; i-- {
        handler = app.middleware[i](handler)
    }

    return handler
}
```

This creates a chain where the first middleware added wraps all others, enabling proper before/after request handling.

## Documentation

Full documentation is coming soon.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) for details.

## Acknowledgments

Inspired by [Bottle](https://bottlepy.org/) - Python's minimalist web framework.
