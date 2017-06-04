package handlers

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/nikovacevic/commonwealth/sessions"
)

// GETIndex GET /
func GETIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Open session DB
	db, err := bolt.Open("boltdb/session.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := sessions.GetUser(w, r, db)

	tpl.ExecuteTemplate(w, "index.gohtml", user)
}
