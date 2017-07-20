package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/views"
)

var indexView = views.NewView("default", "views/index.gohtml")

// GETIndex GET /
func (hdl *Handler) GETIndex(w http.ResponseWriter, r *http.Request) {
	var user *models.User

	uid := sess.GetUserID(r)

	if uid > 0 {
		// TODO Something like this
		// user, err := userService.GetUserById(uid)

		// Find user for given email
		user = &models.User{ID: uid}
		ctx := context.Background()
		row := hdl.db.QueryRowContext(ctx, "SELECT u.first_name, u.last_name, u.email, u.phone, u.organization FROM users AS u WHERE u.id = $1;", uid)
		if err := row.Scan(&(user.FirstName), &(user.LastName), &(user.Email), &(user.Phone), &(user.Organization)); err != nil {
			log.Fatal(err)
		}
	}

	indexView.Render(w, user)
}
