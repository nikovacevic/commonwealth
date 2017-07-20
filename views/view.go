package views

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// View defines a render-able set of template files and a layout to execute.
type View struct {
	Template *template.Template
	Layout   string
}

// NewView creates a new View, given the names of the requisite files.
func NewView(layout string, files ...string) *View {
	// Always parse all layouts
	layouts, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	files = append(files, layouts...)

	// Always parse all components
	components, err := filepath.Glob("views/components/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	files = append(files, components...)

	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	return &View{
		Template: tpl,
		Layout:   layout,
	}
}

// Render writes a View to an HTTP response.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// ServeHTTP renders a view to an HTTP response given a request, allowing Views
// to implement the http.Handler interface.
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}
