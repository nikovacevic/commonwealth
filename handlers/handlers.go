package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/nikovacevic/commonwealth/sessions"
)

// Handler receives all app-specific HandlerFunctions, embedding the app's
// database connection pool and templates.
type Handler struct {
	db  *sql.DB
	tpl *template.Template
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
	tpl := template.Must(template.ParseGlob("templates/*"))

	hdl = &Handler{
		db:  db,
		tpl: tpl,
	}

	sess = sessions.GetSessionHandler()
}

// GetHandler returns the initialized Handler
func GetHandler() *Handler {
	return hdl
}

// Render writes a template--with or without data--to a HTTP ResponseWriter
func (hdl *Handler) Render(w http.ResponseWriter, name string, data interface{}) {
	// TODO Allow this to be given as argument
	// Set headers
	w.Header().Set("Content-Type", "text/html")
	// Execute template
	hdl.tpl.ExecuteTemplate(w, name, data)
}
