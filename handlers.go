package main

import (
	"net/http"
)

// Index handles "GET /index"
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
