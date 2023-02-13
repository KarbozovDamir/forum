package router

import (
	"errors"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	data "github.com/KarbozovDamir/forum/internal/data"
	utils "github.com/KarbozovDamir/forum/internal/utils"
)

//ViewData - struct to template html page
type ViewData struct {
	Result               string
	Authorised           bool
	ID                   int
	Title                string
	Time                 string
	Data                 []data.Thread
	Data2                []data.Thread
	Data3                []data.Thread
	Error                []byte
	CountOfPosts         int
	CountOfLikedThreads  int
	CountOfLikedComments int
	Me                   bool
	Image                string
}

// Parser for templates - value to template html page
var Parser ViewData

// DefaultHandler - Default Request Handler1
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("/web/templates/login.html")
	if r.URL.Path[1:] == "register" {
		if Parser.Authorised == true {
			http.Redirect(w, r, "/articles", code)
		}
		tmpl, _ = template.ParseFiles("web/templates/register.html")
		if r.Method == "GET" {
			tmpl.Execute(w, Parser)
		}
	} else if r.URL.Path[1:] == "login" || r.URL.Path == "/" {
		if Parser.Authorised == true {
			http.Redirect(w, r, "/articles", code)
		}
		tmpl, _ = template.ParseFiles("web/templates/login.html")
		if r.Method == "GET" {
			tmpl.Execute(w, Parser)
		} else {
			r.ParseForm()
			_, ok := r.Form["checkbox"]
			if len(r.Form) == 3 && !ok {
				utils.AddUser(r.Form["username"][0], r.Form["mail"][0], r.Form["psw"][0])
			}
			user, err := data.GetUserByUsername(r.Form["username"][0])
			if err != nil {
				user, err = data.GetUserByMail(r.Form["username"][0])
			}
			if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Psw), []byte(r.Form["psw"][0])) != nil {
				Parser.Result = "Sorry the username or password is not correct\n"
				defer Reset(&Parser)
			} else {
				user.CreateAndSetSession(w, r)
				Parser.Authorised = true
				Parser.ID = user.ID
				http.Redirect(w, r, "/articles", code)
			}
			tmpl.Execute(w, Parser)
		}
	} else {
		ErrorHandler(w, r, errors.New("no such page"), 2)
	}
}
