package handlers

import (
	"log"
	"net/http"

	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/views"
)

var indexView = views.NewView("default", "views/index.gohtml")

// GETIndex GET /
func (hdl *Handler) GETIndex(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var err error

	uid := sess.GetUserID(r)

	if uid > 0 {
		user, err = userService.ByID(uid)
		if err != nil {
			log.Fatal(err)
		}
	}

	indexView.Render(w, user)
}
