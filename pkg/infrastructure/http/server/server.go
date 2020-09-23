package server

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/http/logging"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/http/server/route"
	"github.com/gorilla/mux"
	"net/http"
)

// Server that handles the HTTP API
type Server interface {
	Router() http.Handler
	AddRoute(route route.Route) error
}

type server struct {
	router *mux.Router
}

// New creates a server with all the dependencies
func New() Server {
	return &server{
		mux.NewRouter().StrictSlash(true),
	}
}

// Router returns the http handler
func (s *server) Router() http.Handler {
	return s.router
}

// AddRoute add a new route to the server
func (s *server) AddRoute(route route.Route) error {
	var handler http.Handler

	handler = route.HandlerFunc()
	handler = logging.Logger(handler, route.Name())

	s.router.
		Methods(route.Method()).
		Path(route.Path()).
		Name(route.Name()).
		Handler(handler)

	return nil
}
