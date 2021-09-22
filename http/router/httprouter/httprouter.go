package httprouter

import (
	"github.com/ducketlab/mingo/http/context"
	"github.com/ducketlab/mingo/http/response"
	"github.com/ducketlab/mingo/http/router"
	httppb "github.com/ducketlab/mingo/pb/http"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type httpRouter struct {
	r *httprouter.Router

	middlewareChain  []router.Middleware
	entrySet         *entrySet
	auther           router.Auther
	mergedHandler    http.Handler
	notFound         http.Handler
	authEnable       bool
	permissionEnable bool
	allow            []string
}

func New() router.Router {
	r := &httpRouter{
		notFound: http.HandlerFunc(http.NotFound),
		r: &httprouter.Router{
			RedirectTrailingSlash:  true,
			RedirectFixedPath:      true,
			HandleMethodNotAllowed: true,
		},
	}

	return r
}

func (r *httpRouter) Use(m router.Middleware) {

	r.middlewareChain = append(r.middlewareChain, m)
	r.mergedHandler = r.r

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		r.mergedHandler = r.middlewareChain[i].Handler(r.mergedHandler)
	}
}

func (r *httpRouter) Handle(method string, path string, h http.HandlerFunc) httppb.EntryDecorator {

	e := &entry{
		Entry: &httppb.Entry{
			Path:             path,
			Method:           method,
			AuthEnable:       r.authEnable,
			PermissionEnable: r.permissionEnable,
			Allow:            r.allow,
			Labels:           map[string]string{},
		},
		h: h,
	}

	r.add(e)

	return e.Entry
}

func (r *httpRouter) add(e *entry) {
	var handler http.Handler

	handler = e.h
	r.addHandler(e.Method, e.Path, handler)
	r.addEntry(e)
}

func (r *httpRouter) addHandler(method string, path string, h http.Handler) {
	wrapper := func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rc := context.GetContext(req)

		var (
			authInfo interface{}
			entry    *entry
		)

		entry = r.findEntry(method, path)

		if entry == nil {
			r.notFound.ServeHTTP(w, req)
			return
		}

		rc.Entry = entry.Entry

		if r.auther != nil {
			ai, err := r.auther.Auth(req, *entry.Entry)

			if err != nil {
				response.Failed(w, err)
				return
			}

			authInfo = ai
		}

		rc.AuthInfo = authInfo
		rc.Ps = ps
		req = context.WithContext(req, rc)

		h.ServeHTTP(w, req)

		if r.auther != nil {
			r.auther.ResponseHook(w, req, *entry.Entry)
		}
	}

	r.r.Handle(method, path, wrapper)
}

func (r *httpRouter) addEntry(e *entry) {
	if err := r.entrySet.AddEntry(e); err != nil {
		panic(err)
	}
}

func (r *httpRouter) findEntry(method string, path string) *entry {
	handler, paras, _ := r.r.Lookup(method, path)

	if handler == nil {
		return nil
	}

	targetPath := path

	for _, kv := range paras {
		targetPath = strings.Replace(targetPath, ":"+kv.Key, kv.Value, 1)
	}

	return r.entrySet.FindEntry(targetPath, method)
}

func (r *httpRouter) SetAuther(auther router.Auther) {
	r.auther = auther
}

func (r *httpRouter) Auth(isEnable bool) {
	r.authEnable = isEnable
}

func (r *httpRouter) Permission(isEnable bool) {
	r.permissionEnable = isEnable
}

func (r *httpRouter) EnableApiRoot() {
	r.Handle("GET", "/", r.apiRoot)
}

func (r *httpRouter) apiRoot(w http.ResponseWriter, req *http.Request) {
	response.Success(w, r.entrySet.EntrySet())
}
