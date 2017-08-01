package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // postgres driver
	"github.com/nikovacevic/commonwealth/auth"
	"github.com/nikovacevic/commonwealth/models"
)

type emailTakenError struct {
	email string
}

func (ete *emailTakenError) Error() string {
	return fmt.Sprintf("Email address %s has already been taken", ete.email)
}

// GETLogin GET /login
func (hdl *Handler) GETLogin(w http.ResponseWriter, r *http.Request) {
	loginView.Render(w, nil)
}

// POSTLogin POST /login
func (hdl *Handler) POSTLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	ctx := context.Background()
	row := hdl.db.QueryRowContext(ctx, "SELECT u.id, u.password_hash FROM users AS u WHERE email = $1;", email)
	var uid uint64
	var hash string
	if err := row.Scan(&uid, &hash); err != nil {
		log.Fatal(err)
	}

	if err := auth.CheckPassword([]byte(hash), []byte(password)); err != nil {
		log.Println("Passwords do not match")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	s := sess.NewSession(w, uid)
	if s == nil {
		log.Println("auth.go\t110\tFailed to create session")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

// GETRegister GET /register
func (hdl *Handler) GETRegister(w http.ResponseWriter, r *http.Request) {
	registerView.Render(w, nil)
}

// POSTRegister POST /register
func (hdl *Handler) POSTRegister(w http.ResponseWriter, r *http.Request) {
	// TODO Check if already logged in

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	organization := r.FormValue("organization")
	password := r.FormValue("password")

	hash, err := auth.HashPassword([]byte(password))
	if err != nil {
		log.Fatal(err)
	}

	user := &models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		Phone:        phone,
		Organization: organization,
		PasswordHash: string(hash),
	}

	user, err = userService.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	// TODO start session
	// s := sess.NewSession(w, user.ID)
	// if s == nil {
	// 	log.Println("auth.go\t110\tFailed to create session")
	// }

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
