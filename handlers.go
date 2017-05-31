package main

import (
	"fmt"
	"net/http"
)

// Index handles "GET /index"
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Commonwealth</h1>")
}
