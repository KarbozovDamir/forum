package router

import (
	"log"
	"net/http"
	"text/template"
)

// SetAndDelete - Sets and clears two variables
func SetAndDelete(s *string, t *[]byte) {
	*s = string(*t)
	*t = make([]byte, 0)
}

// Reset - Reset all values after templates except main attributes
func Reset(x *ViewData) {
	*x = ViewData{Authorised: x.Authorised, ID: x.ID}
}

// Error - html page to handle errors
func Error(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	defer Reset(&Parser)
	SetAndDelete(&Parser.Result, &Parser.Error)
	if string(Parser.Result[:]) == "" {
		http.Redirect(w, r, "/articles", code)
	}
	tmpl, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		log.Println("err: ", err.Error())
		return
	}
	tmpl.Execute(w, Parser)
}
