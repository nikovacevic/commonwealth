package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // postgres driver
	"github.com/nikovacevic/commonwealth/auth"
	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/views"
)

type emailTakenError struct {
	email string
}

func (ete *emailTakenError) Error() string {
	return fmt.Sprintf("Email address %s has already been taken", ete.email)
}

var loginView = views.NewView("default", "views/auth/login.gohtml")
var registerView = views.NewView("default", "views/auth/register.gohtml")

// GETLogin GET /login
func (hdl *Handler) GETLogin(w http.ResponseWriter, r *http.Request) {
	loginView.Render(w, nil)
}

// POSTLogin POST /login
func (hdl *Handler) POSTLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Find user for given email
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

	user := models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		Phone:        phone,
		Organization: organization,
		PasswordHash: string(hash),
	}

	ctx := context.Background()
	stmt, err := hdl.db.PrepareContext(ctx, "INSERT INTO users (first_name, last_name, email, phone, organization, password_hash) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.ExecContext(ctx, user.FirstName, user.LastName, user.Email, user.Phone, user.Organization, user.PasswordHash)
	if err != nil {
		log.Fatal(err)
	}
	uid, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	user.ID = uint64(uid)

	s := sess.NewSession(w, user.ID)
	if s == nil {
		log.Println("auth.go\t110\tFailed to create session")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
