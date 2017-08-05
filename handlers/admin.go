package handlers

import (
	"net/http"

	"github.com/nikovacevic/commonwealth/views"
)

var adminView = views.NewView("default", "views/admin/index.gohtml")

// GETAdmin GET /admin
func (hdl *Handler) GETAdmin(w http.ResponseWriter, r *http.Request) {
	adminView.Render(w, nil)
}
