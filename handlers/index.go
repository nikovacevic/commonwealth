package handlers

import (
	"net/http"

	"github.com/nikovacevic/commonwealth/sessions"
)

// GETIndex GET /
func GETIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	user := sessions.GetUser(w, r)

	tpl.ExecuteTemplate(w, "index.gohtml", user)
}
