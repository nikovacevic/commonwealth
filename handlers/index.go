package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/nikovacevic/commonwealth/sessions"
)

// GETIndex GET /
func GETIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Open session DB
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

	user := sessions.GetUser(w, r, bdb, db)

	tpl.ExecuteTemplate(w, "index.gohtml", user)
}
