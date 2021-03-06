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

// NewRouter creates a gorilla/mux Router, registers routes, and returns it.
func NewRouter(h *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	// Register routes
	routes := Routes{
		// Index
		Route{"GET", "/", h.GETIndex},
		// Auth
		Route{"GET", "/login", h.GETLogin},
		Route{"POST", "/login", h.POSTLogin},
		Route{"GET", "/register", h.GETRegister},
		Route{"POST", "/register", h.POSTRegister},
		// Admin
		Route{"GET", "/admin", h.GETAdmin},
		// Users
		Route{"GET", "/users", h.GETUsers},
		// Products
		Route{"GET", "/products", h.Products},
		Route{"GET", "/products/{id:[0-9]+}", h.ViewProduct},
		Route{"GET", "/products/create", h.CreateProduct},
		Route{"POST", "/products/create", h.PostProduct},
		Route{"GET", "/products/{id:[0-9]+}/update", h.UpdateProduct},
		Route{"POST", "/products/{id:[0-9]+}/update", h.PatchProduct},
	}

	for _, route := range routes {
		fn := logger.LogRequest(route.HandlerFunc)
		r.HandleFunc(route.Path, fn).Methods(route.Method)
	}

	// Register error routes
	r.NotFoundHandler = logger.LogError("404", http.HandlerFunc(h.NotFound))

	// Register route for static files
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))

	// Register special assets
	r.HandleFunc("/favicon.ico", http.NotFound)

	return r
}
