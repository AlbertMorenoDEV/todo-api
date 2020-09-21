package route

import (
	"net/http"
)

// Route handles HTTP requests
type Route interface {
	Path() string
	Method() string
	Name() string
	HandlerFunc() http.HandlerFunc
}

type route struct {
	path        string
	method      string
	name        string
	handlerFunc http.HandlerFunc
}

// New creates a route
func New(path string, method string, name string, handlerFunc http.HandlerFunc) Route {
	return &route{
		path:        path,
		method:      method,
		name:        name,
		handlerFunc: handlerFunc,
	}
}

func (r *route) Path() string {
	return r.path
}

func (r *route) Method() string {
	return r.method
}

func (r *route) Name() string {
	return r.name
}

func (r *route) HandlerFunc() http.HandlerFunc {
	return r.handlerFunc
}
