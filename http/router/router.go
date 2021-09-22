package router

import (
	httppb "github.com/ducketlab/mingo/pb/http"
	"net/http"
)

type Router interface {
	// Use Add middleware
	Use(m Middleware)

	// Handle Add an authentication protected route
	Handle(method string, path string, h http.HandlerFunc) httppb.EntryDecorator

	// Auth Whether to enable user authentication
	Auth(isEnable bool)

	// Permission Whether to enable user permission verification
	Permission(isEnable bool)

	// SetAuther Set auther
	SetAuther(auther Auther)

	// RequiredNamespace namespace
	RequiredNamespace(isEnable bool)

	// EnableApiRoot Route/expose the service routing table
	EnableApiRoot()

	// SubRouter base path
	SubRouter(basePath string) SubRouter
}

type ResourceRouter interface {
	SubRouter
	BasePath(path string)
}

type SubRouter interface {
	// Auth Whether to enable user authentication
	Auth(isEnable bool)

	// Permission Whether to enable user permission verification
	Permission(isEnable bool)

	// RequiredNamespace namespace
	RequiredNamespace(isEnable bool)

	// Use Add middleware
	Use(m Middleware)

	// With handler
	With(m ...Middleware) SubRouter

	// Handle Add an authentication protected route
	Handle(method string, path string, h http.HandlerFunc) httppb.EntryDecorator
}