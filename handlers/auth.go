package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	_ "github.com/lib/pq" // postgres driver
	"github.com/nikovacevic/commonwealth/auth"
	"github.com/nikovacevic/commonwealth/models"
	"github.com/nikovacevic/commonwealth/sessions"
)

type emailTakenError struct {
	email string
}

func (ete *emailTakenError) Error() string {
	return fmt.Sprintf("Email address %s has already been taken", ete.email)
}

// GETLogin GET /login
func GETLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

// POSTLogin POST /login
func POSTLogin(w http.ResponseWriter, r *http.Request) {
	// BoltDB for Session persistence
	bdb, err := bolt.Open("boltdb/session.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer bdb.Close()

	// Postgres for non-Session persistence
	db, err := sql.Open("postgres", "postgres://niko:@localhost/commonwealth?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	email := r.FormValue("email")
	password := r.FormValue("password")

	ctx := context.Background()
	// Retrieve User from Session
	row := db.QueryRowContext(ctx, "SELECT u.id, u.password_hash FROM users AS u WHERE email = $1;", email)
	var uid uint64
	var hash string
	if err = row.Scan(&uid, &hash); err != nil {
		log.Fatal(err)
	}

	err = auth.CheckPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("Passwords do not match")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	s := sessions.NewSession(w, uid, bdb)
	if s == nil {
		log.Println("auth.go\t110\tFailed to create session")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

// GETRegister GET /register
func GETRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

// POSTRegister POST /register
func POSTRegister(w http.ResponseWriter, r *http.Request) {
	// BoltDB for Session persistence
	bdb, err := bolt.Open("boltdb/session.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer bdb.Close()

	// Postgres for non-Session persistence
	db, err := sql.Open("postgres", "postgres://niko:@localhost/commonwealth?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
	stmt, err := db.PrepareContext(ctx, "INSERT INTO users (first_name, last_name, email, phone, organization, password_hash) VALUES ($1, $2, $3, $4, $5, $6)")
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

	s := sessions.NewSession(w, user.ID, bdb)
	if s == nil {
		log.Println("auth.go\t110\tFailed to create session")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
