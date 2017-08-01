package handlers

import (
	"database/sql"
	"log"

	"github.com/nikovacevic/commonwealth/services"
	"github.com/nikovacevic/commonwealth/sessions"
)

// Handler receives all app-specific HandlerFunctions, embedding the app's
// database connection pool and templates.
type Handler struct {
	db *sql.DB
}

var hdl *Handler
var sess *sessions.SessionHandler

func init() {
	// TODO toggle SSL disable with config
	// TODO grab DB credentials from config
	db, err := sql.Open("postgres", "postgres://niko:@localhost/commonwealth?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	hdl = &Handler{db: db}

	userService = services.NewUser(hdl.db)

	sess = sessions.GetSessionHandler()
}

// GetHandler returns the initialized Handler
func GetHandler() *Handler {
	return hdl
}
