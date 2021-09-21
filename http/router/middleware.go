package router

import "net/http"

// Middleware router
type Middleware interface {
	Handler(handler http.Handler) http.Handler
}

// MiddlewareFunc is an adapter to allow
type MiddlewareFunc func(handler http.Handler) http.Handler

// Handler wrapper fo function
func (h MiddlewareFunc) Handler(next http.Handler) http.Handler {
	return h(next)
}
