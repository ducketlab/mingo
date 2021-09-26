package httprouter

import (
	"github.com/ducketlab/mingo/http/router"
	httppb "github.com/ducketlab/mingo/pb/http"
	"net/http"
	"path"
)

type subRouter struct {
	basePath          string
	resourceName      string
	resourceBasePath  string
	root              *httpRouter
	middlewareChain   []router.Middleware
	authEnable        bool
	permissionEnable  bool
	allow             []string
	requiredNamespace bool
}

func newSubRouter(basePath string, root *httpRouter) *subRouter {
	return &subRouter{
		basePath:          basePath,
		root:              root,
		authEnable:        root.authEnable,
		permissionEnable:  root.permissionEnable,
		requiredNamespace: root.requiredNamespace,
	}
}

func (r *subRouter) Use(m router.Middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *subRouter) With(m ...router.Middleware) router.SubRouter {
	sr := &subRouter{
		basePath:        r.basePath,
		root:            r.root,
		middlewareChain: r.middlewareChain,
	}

	sr.middlewareChain = append(sr.middlewareChain, m...)
	return sr
}

func (r *subRouter) Auth(isEnable bool) {
	r.authEnable = isEnable
}

func (r *subRouter) Permission(isEnable bool) {
	r.permissionEnable = isEnable
}

func (r *subRouter) RequiredNamespace(isEnable bool) {
	r.requiredNamespace = isEnable
}

func (r *subRouter) Handle(method string, path string, h http.HandlerFunc) httppb.EntryDecorator {
	e := &entry{
		Entry: &httppb.Entry{
			Resource:         r.resourceName,
			Method:           method,
			Path:             path,
			AuthEnable:       r.authEnable,
			PermissionEnable: r.permissionEnable,
			Allow:            r.allow,
		},
		h: h,
	}

	r.add(e)
	return e.Entry
}

func (r *subRouter) add(e *entry) {
	mergedHandler := r.combineHandler(e)
	e.Path = r.calculateAbsolutePath(e.Path)
	r.root.addHandler(e.Method, e.Path, mergedHandler)
	r.root.addEntry(e)
}

func (r *subRouter) combineHandler(e *entry) http.Handler {
	var mergedHandler http.Handler
	mergedHandler = http.HandlerFunc(e.h)
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i].Handler(mergedHandler)
	}
	return mergedHandler
}

func (r *subRouter) ResourceRouter(resourceName string) router.ResourceRouter {
	return &subRouter{
		resourceName:      resourceName,
		resourceBasePath:  r.basePath,
		basePath:          r.basePath,
		root:              r.root,
		authEnable:        r.authEnable,
		permissionEnable:  r.permissionEnable,
		allow:             r.allow,
		requiredNamespace: r.requiredNamespace,
	}
}

func (r *subRouter) BasePath(path string) {
	r.basePath = r.resourceBasePath + "/" + path
}

func (r *subRouter) calculateAbsolutePath(relativePath string) string {

	if relativePath == "" {
		return r.basePath
	}

	finalPath := path.Join(r.basePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}

	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
