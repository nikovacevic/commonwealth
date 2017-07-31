package handlers

import (
	"log"
	"net/http"

	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/views"
)

var adminIndexView = views.NewView("default", "views/admin/index.gohtml")
var adminUsersView = views.NewView("default", "views/admin/users.gohtml")

// GETAdminIndex GET /admin
func (hdl *Handler) GETAdminIndex(w http.ResponseWriter, r *http.Request) {
	adminIndexView.Render(w, nil)
}

// GETAdminUsers GET /admin/users
func (hdl *Handler) GETAdminUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := hdl.db.Query("SELECT id, first_name, last_name, email, organization FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Organization,
		)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	adminUsersView.Render(w, struct {
		Users []*models.User
	}{
		users,
	})
}
