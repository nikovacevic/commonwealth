package handlers

import (
	"net/http"

	"github.com/nikovacevic/commonwealth/views"
)

var error404View = views.NewView("default", "views/404.gohtml")

// NotFound handles 404 errors
func (hdl *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	error404View.Render(w, nil)
}
