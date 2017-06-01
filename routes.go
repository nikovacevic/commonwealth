package main

import (
	"net/http"

	"github.com/nikovacevic/commonwealth/handlers"
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
	// Index
	Route{
		"GET",
		"/",
		handlers.GETIndex,
	},
	// Auth
	Route{
		"GET",
		"/login",
		handlers.GETLogin,
	},
	Route{
		"POST",
		"/login",
		handlers.POSTLogin,
	},
	Route{
		"GET",
		"/register",
		handlers.GETRegister,
	},
	Route{
		"POST",
		"/register",
		handlers.POSTRegister,
	},
}
