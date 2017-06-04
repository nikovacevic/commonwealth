package handlers

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
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
	w.Header().Set("Content-Type", "text/html")
	// TODO
}

// GETRegister GET /register
func GETRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

// POSTRegister POST /register
func POSTRegister(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("boltdb/session.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO Check already logged in

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: password,
	}

	/*
		// TODO Check for unique email
		err = db.View(func(tx *bolt.Tx) error {
			userBkt := tx.Bucket([]byte("user"))
			data := userBkt.Get([]byte(user.Email))
			if data != nil {
				return &emailTakenError{user.Email}
			}
			return nil
		})
		switch err := err.(type) {
		case nil:
			// Email is available, so continue
		case *emailTakenError:
			// TODO send back to form with values
			log.Printf(err.Error())
			http.Redirect(w, r, "/register", http.StatusSeeOther)
		default:
			log.Fatal(err)
		}
	*/

	err = db.Update(func(tx *bolt.Tx) error {
		userBkt := tx.Bucket([]byte("user"))
		user.ID, err = userBkt.NextSequence()
		data, er := json.Marshal(user)
		if er != nil {
			return er
		}
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, user.ID)
		er = userBkt.Put(b, data)
		if er != nil {
			return er
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	s := sessions.NewSession(w, user.ID, db)
	if s == nil {
		log.Println("auth.go\t110\tFailed to create session")
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
