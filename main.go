package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikovacevic/commonwealth/handlers"
	"github.com/nikovacevic/commonwealth/logger"
)

// Route is an HTTP method and request path with a HandlerFunc to run
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Routes
type Routes []Route

func main() {
	h := handlers.GetHandler()
	r := NewRouter(h)
	http.ListenAndServe(":3000", r)
}

// NewRouter returns a gorilla/mux Router with registered routing based on
// var routes (see routes.go)
func NewRouter(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	// Register routes
	routes := Routes{
		// Index
		Route{
			"GET",
			"/",
			h.GETIndex,
		},
		// Auth
		Route{
			"GET",
			"/login",
			h.GETLogin,
		},
		Route{
			"POST",
			"/login",
			h.POSTLogin,
		},
		Route{
			"GET",
			"/register",
			h.GETRegister,
		},
		Route{
			"POST",
			"/register",
			h.POSTRegister,
		},
	}

	for _, route := range routes {
		fn := logger.LogRequest(route.HandlerFunc)
		r.HandleFunc(route.Path, fn).Methods(route.Method)
	}

	// Register route for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Register special assets
	r.HandleFunc("/favicon.ico", http.NotFound)

	return r
}
