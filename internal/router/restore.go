package router

import (
	"errors"
	"net/http"
	"text/template"

	data "github.com/KarbozovDamir/forum/data"
)

// Restore password
func Restore(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	if r.URL.Path[8:] != "" {
		ErrorHandler(w, r, errors.New("no such page"), 2)
	}
	Parser.Result = ""
	tmpl, _ := template.ParseFiles("web/templates/restore.html")
	if r.Method == "GET" {
		tmpl.Execute(w, Parser)
	} else {
		r.ParseForm()
		capsule, err := data.GetUserByUsername(r.Form["username"][0])
		service := data.UserService{User: &capsule}
		if err != nil || service.User.Code != r.Form["code"][0] {
			Parser.Result = "Sorry wrong code or username"
		} else {
			Parser.Result = "Your password successfully changed"
			service.Update(w, r, r.Form["psw"][0])
			if Parser.Authorised == true {
				http.Redirect(w, r, "/logout", code)
			} else {
				http.Redirect(w, r, "/", code)
			}
		}
		tmpl.Execute(w, Parser)
	}
	Reset(&Parser)
}
