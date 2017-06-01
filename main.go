package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikovacevic/commonwealth/logger"
)

func main() {
	r := NewRouter()
	http.ListenAndServe(":3000", r)
}

// NewRouter returns a gorilla/mux Router with registered routing based on
// var routes (see routes.go)
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Register routes defined in routes
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
