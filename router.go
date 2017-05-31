package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is an HTTP method and request path with a HandlerFunc to run
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Routes
type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
}

// NewRouter returns a new gorilla/mux Router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.HandleFunc(route.Path, Log(route.HandlerFunc))
	}
	return r
}
