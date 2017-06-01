package handlers

import (
	"fmt"
	"net/http"

	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/sessions"
)

// GETLogin GET /login
func GETLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

// POSTLogin POST /login
func POSTLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// TODO
}

// GETRegister GET /register
func GETRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println("GETRegister")
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

// POSTRegister POST /register
func POSTRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println("POSTRegister")

	// TODO Check already logged in

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		ID:           uint(len(sessions.DBUser) + 1),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: password,
	}

	fmt.Println(user)

	sessions.DBUser[user.ID] = user
	sessions.NewSession(w, user.ID)

	// http.Redirect(w, r, "/", http.StatusSeeOther)
	// return

	view := struct {
		User models.User
	}{
		user,
	}

	tpl.ExecuteTemplate(w, "register.gohtml", view)
}
