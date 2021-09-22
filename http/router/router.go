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

	// EnableApiRoot Route/expose the service routing table
	EnableApiRoot()
}