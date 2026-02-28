package mach

import "net/http"

// Group is a route group with common named prefix
type Group struct {
	prefix      string
	middlewares []MiddlewareFunc

	app *App
}

// Use registers middlewares to the group
func (g *Group) Use(middlewares ...MiddlewareFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

// Group creates a sub-group. Global middlewares come first in the chain
func (g *Group) Group(prefix string, middlewares ...MiddlewareFunc) *Group {
	// copy all applicable middlewares
	return &Group{
		prefix:      g.prefix + prefix,
		middlewares: append(append([]MiddlewareFunc{}, g.middlewares...), middlewares...),
		app:         g.app,
	}
}

func (g *Group) handle(method, path string, handler HandlerFunc) {
	// compose full path with group prefix
	path = g.prefix + path

	if len(g.middlewares) == 0 {
		g.app.handle(method, path, handler)
		return
	}

	// wrap middlewares in reverse order so they execute in order of addition
	wrappedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// retrieve context from pool
		c := g.app.pool.Get().(*Context)
		c.reset(w, r)

		handler(c)
		g.app.pool.Put(c)
	})

	for idx := len(g.middlewares) - 1; idx >= 0; idx-- {
		wrappedHandler = g.middlewares[idx](wrappedHandler).(http.HandlerFunc)
	}

	// handle request with middlewares
	pattern := method + " " + path
	g.app.router.Handle(pattern, wrappedHandler)
}

// GET registers a GET route for the group
func (g *Group) GET(path string, handler HandlerFunc) {
	g.handle(http.MethodGet, path, handler)
}

// POST registers a POST route for the group
func (g *Group) POST(path string, handler HandlerFunc) {
	g.handle(http.MethodPost, path, handler)
}

// PATCH registers a PATCH route for the group
func (g *Group) PATCH(path string, handler HandlerFunc) {
	g.handle(http.MethodPatch, path, handler)
}

// PUT registers a PUT route for the group
func (g *Group) PUT(path string, handler HandlerFunc) {
	g.handle(http.MethodPut, path, handler)
}

// DELETE registers a DELETE route for the group
func (g *Group) DELETE(path string, handler HandlerFunc) {
	g.handle(http.MethodDelete, path, handler)
}

// HEAD registers a HEAD route for the group
func (g *Group) HEAD(path string, handler HandlerFunc) {
	g.handle(http.MethodHead, path, handler)
}

// OPTIONS registers a OPTIONS route for the group
func (g *Group) OPTIONS(path string, handler HandlerFunc) {
	g.handle(http.MethodOptions, path, handler)
}
