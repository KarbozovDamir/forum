package router

import (
	"errors"
	"net/http"
	"text/template"
)

//About - info about forum
func About(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("/web/templates/about.html")
	if r.URL.Path[6:] != "" {
		ErrorHandler(w, r, errors.New("no such page"), 2)
	}
	tmpl.Execute(w, Parser)
}
