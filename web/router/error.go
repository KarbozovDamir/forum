package router

import (
	"log"
	"net/http"
	"text/template"
)

const code = 302

//ErrorHandler - Allow us to handle errors in forum
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error, status int) {
	if status == 0 {
		if !Parser.Authorised {
			http.Redirect(w, r, "/login", code)
		}
	}
	if err != nil {
		switch status {
		case 1:
			Parser.Error = []byte("No such thread")
		case 2:
			Parser.Error = []byte("No such page")
		case 3:
			Parser.Error = []byte("No such user")
		case 4:
			Parser.Error = []byte("You don't have permission")
		case 5:
			Parser.Error = []byte("Internal error has been occured")
		}
		http.Redirect(w, r, "/error", code) //Found error
	}
}

//SetAndDelete - Sets and clears two variables
func SetAndDelete(s *string, t *[]byte) {
	*s = string(*t)
	*t = make([]byte, 0)
}

//Reset - Reset all values after templates except main attributes
func Reset(x *ViewData) {
	*x = ViewData{Authorised: x.Authorised, ID: x.ID}
}

//Error - html page to handle errors
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
