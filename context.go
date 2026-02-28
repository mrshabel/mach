package mach

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	ErrEmptyRequestBody = errors.New("request body is empty")
)

// Context adds helpful methods to the ongoing request
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter

	app *App

	// cached url data
	query        url.Values
	IsFormParsed bool
}

// reset prepares a context to be reused by a new request
func (ctx *Context) reset(w http.ResponseWriter, r *http.Request) {
	ctx.Response = w
	ctx.Request = r
	ctx.query = nil
	ctx.IsFormParsed = false
}

// Param gets a path parameter by name.
// For example, this returns the value of id from /users/{id}
func (c *Context) Param(name string) string {
	return c.Request.PathValue(name)
}

// Query returns a named query parameter
func (c *Context) Query(name string) string {
	// extract all query parameters once
	if c.query == nil {
		c.query = c.Request.URL.Query()
	}

	return c.query.Get(name)
}

// DefaultQuery gets query param with default value
func (c *Context) DefaultQuery(name, defaultValue string) string {
	val := c.Query(name)
	if val == "" {
		return defaultValue
	}

	return val
}

// Form gets a form value
func (c *Context) Form(name string) string {
	// parse form values only once. its values cached by default once parsed
	if !c.IsFormParsed {
		c.Request.ParseForm()
		c.IsFormParsed = true
	}

	return c.Request.FormValue(name)
}

// File gets an uploaded file by key name. The file header containing the file is returned
func (c *Context) File(name string) (*multipart.FileHeader, error) {
	_, header, err := c.Request.FormFile(name)
	return header, err
}

// Cookie gets a request cookie by name
func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.Request.Cookie(name)
}

// GetHeader retrieves a request header by key
func (c *Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

// Method returns the request method
func (c *Context) Method() string {
	return c.Request.Method
}

// Path retrieves the request path
func (c *Context) Path() string {
	return c.Request.URL.Path
}

// ClientIP returns the client IP address.
// Use this if you trust request headers passed to the server (ie: reverse proxy sits before server)
// else use c.Request.RemoteAddr()
func (c *Context) ClientIP() string {
	if forwarded := c.Request.Header.Get("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}
	if realIP := c.Request.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	ip := c.Request.RemoteAddr
	if parts := strings.Split(ip, ":"); len(parts) > 1 {
		ip = parts[0]
	}
	return ip
}

// Body reads the request body
func (c *Context) Body() ([]byte, error) {
	return io.ReadAll(c.Request.Body)
}

// Context returns the request original context from context.Context
func (c *Context) Context() context.Context {
	return c.Request.Context()
}

// binding methods

// DecodeJSON decodes a request body into a struct
func (c *Context) DecodeJSON(data interface{}) error {
	if c.Request.Body == nil {
		return ErrEmptyRequestBody
	}

	return json.NewDecoder(c.Request.Body).Decode(data)
}

// DecodeXML decodes a request body into a struct
func (c *Context) DecodeXML(data interface{}) error {
	if c.Request.Body == nil {
		return ErrEmptyRequestBody
	}

	return xml.NewDecoder(c.Request.Body).Decode(data)
}

// response methods

// JSON sends a JSON response
func (c *Context) JSON(status int, data interface{}) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(status)

	return json.NewEncoder(c.Response).Encode(data)
}

// XML sends an XML response
func (c *Context) XML(status int, data interface{}) error {
	c.SetHeader("Content-Type", "application/xml")
	c.Response.WriteHeader(status)

	return xml.NewEncoder(c.Response).Encode(data)
}

// Text sends a plain text response
func (c *Context) Text(status int, format string, values ...interface{}) error {
	c.SetHeader("Content-Type", "text/plain; charset=utf-8")
	c.Response.WriteHeader(status)

	_, err := fmt.Fprintf(c.Response, format, values...)

	return err
}
func (c *Context) HTML(status int, html string) error {
	c.SetHeader("Content-Type", "text/html; charset=utf-8")
	c.Response.WriteHeader(status)

	_, err := c.Response.Write([]byte(html))
	return err
}

// Data sends raw bytes
func (c *Context) Data(status int, contentType string, data []byte) error {
	c.SetHeader("Content-Type", contentType)
	c.Response.WriteHeader(status)

	_, err := c.Response.Write(data)
	return err
}

// NoContent sends a response with no body
func (c *Context) NoContent(status int) {
	c.Response.WriteHeader(status)
}

// SetHeader sets a response header
func (c *Context) SetHeader(key, value string) {
	c.Response.Header().Set(key, value)
}

// SetCookie sets a response cookie
func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Response, cookie)
}

// Redirect redirects to a URL
func (c *Context) Redirect(status int, url string) {
	http.Redirect(c.Response, c.Request, url, status)
}

// utilities

// SaveFile saves an uploaded file to the specified destination path.
func (c *Context) SaveFile(file *multipart.FileHeader, path string) error {
	// copy file to destination
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return err
}
