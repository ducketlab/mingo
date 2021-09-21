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

	// SetAuther Set auther
	SetAuther(auther Auther)
}