package router

import (
	httppb "github.com/ducketlab/mingo/pb/http"
	"net/http"
)

// Auther Authenticator used by the protected route
// Title The title is used for identification
// The record of the record is used for authentication of permissions
type Auther interface {
	Auth(req *http.Request, entry httppb.Entry) (authInfo interface{}, err error)
	ResponseHook(w http.ResponseWriter, r *http.Request, entry httppb.Entry)
}

// AuthFunc is an adapter to allow
type AuthFunc func(header http.Header) (authInfo interface{}, err error)

// Auth calls auth
func (f AuthFunc) Auth(h http.Header) (authInfo interface{}, err error) {
	return f(h)
}