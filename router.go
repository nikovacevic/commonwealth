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
		index,
	},
	Route{
		"GET",
		"/login",
		login,
	},
}

// NewRouter returns a new gorilla/mux Router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Register routes defined in routes
	for _, route := range routes {
		r.HandleFunc(route.Path, Log(route.HandlerFunc))
	}
	// Register route for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	// Register special assets
	r.HandleFunc("/favicon.ico", http.NotFound)
	return r
}
