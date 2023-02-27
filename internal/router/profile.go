package router

import (
	"errors"
	"net/http"
	"strconv"
	"text/template"

	data "github.com/KarbozovDamir/forum/data"
)

// Profile - Profile page of user
func Profile(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("web/templates/profile.html")
	x, err2 := strconv.Atoi(r.URL.Path[4:])
	user, err := data.GetUserByID(x)
	if err2 != nil {
		user, err = data.GetUserByUsername(r.URL.Path[4:])
	}
	if err != nil {
		ErrorHandler(w, r, err, 3)
	} else {
		Parser.Title = user.Username
		Parser.Result = user.Mail
		if Parser.Authorised {
			Parser.Threads = data.GetAllUserCreatedPosts(user.ID)
			Parser.LikedThreads = data.GetAllUserLikedThread(user.ID)
			Parser.LikedComments = data.GetAllUserLikedComments(user.ID)
		}
		Parser.CountOfPosts = len(Parser.Threads)
		Parser.CountOfLikedThreads = len(Parser.LikedThreads)
		Parser.CountOfLikedComments = len(Parser.LikedComments)
		if Parser.ID == user.ID {
			Parser.Me = true
		}
		Parser.Image = "User" + strconv.Itoa(user.ID)
		err = data.FindImage(Parser.Image)
		if err != nil {
			Parser.Image = "/static/images/default-avatar.jpg"
		} else {
			Parser.Image = "/static/images/User" + strconv.Itoa(user.ID)
		}
		tmpl.Execute(w, Parser)
		Reset(&Parser)
	}
}

// UpdateAva - Updating our avatar image
func UpdateAva(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	x, err2 := strconv.Atoi(r.URL.Path[20:])
	user, err := data.GetUserByID(x)
	if err != nil || err2 != nil || r.Method == "GET" {
		ErrorHandler(w, r, err, 2)
		return
	}
	if user.ID != Parser.ID {
		ErrorHandler(w, r, errors.New("No permission"), 4)
		return
	}
	Parser.Image = "User" + strconv.Itoa(user.ID)
	err = data.AddImage(Parser.Image, 1, user.ID, r)
	if err != nil {
		ErrorHandler(w, r, err, 5)
		return
	}
	http.Redirect(w, r, "/id/"+strconv.Itoa(user.ID), code)
}
