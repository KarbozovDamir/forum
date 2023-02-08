package router

import (
	"errors"
	"net/http"
	"text/template"
)

//Rules - Our rules on forum
func Rules(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	if r.URL.Path[6:] != "" {
		ErrorHandler(w, r, errors.New("no such page"), 2)
	}
	tmpl, _ := template.ParseFiles("/web/templatess/rules.html")
	tmpl.Execute(w, Parser)
}
