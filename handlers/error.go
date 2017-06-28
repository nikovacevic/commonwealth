package handlers

import (
	"net/http"
)

// NotFound handles 404 errors
func (hdl *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	hdl.Render(w, "404.gohtml", nil)
}
