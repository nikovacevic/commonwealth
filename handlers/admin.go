package handlers

import (
	"net/http"
)

// GETAdmin GET /admin
func (hdl *Handler) GETAdmin(w http.ResponseWriter, r *http.Request) {
	adminView.Render(w, nil)
}
