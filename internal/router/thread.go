package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"text/template"

	data "github.com/KarbozovDamir/forum/internal/data"
)

//StatsTH - struct of Thread statistic
type StatsTH struct {
	ThreadID string
	Value    int
	Likes    string
	Dislikes string
	Liked    string
}

//Authorised - Is user authorised ?
func Authorised(r *http.Request) {
	id, err := data.CheckCookie(r)
	if err == nil {
		Parser.Authorised = true
		Parser.ID = id
	} else {
		Parser.Authorised = false
	}
}

//CreateThread - Html page to Create Thread
func CreateThread(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	ErrorHandler(w, r, nil, 0)
	user, _ := data.GetUserByID(Parser.ID)
	tmpl, _ := template.ParseFiles("web/templates/create_thread.html")
	if r.Method == "POST" {
		num, err := data.CreateTH(r, Parser.ID, user.Username)
		if err != nil {
			ErrorHandler(w, r, err, 5)
		}
		http.Redirect(w, r, "/thread/"+strconv.Itoa(num), code)

	}
	tmpl.Execute(w, Parser)
}

//Post - Page to Thread
func Post(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	tmpl, _ := template.ParseFiles("web/templates/thread.html")
	if r.Method == "POST" {
		Comment(w, r)
		http.Redirect(w, r, r.URL.Path, code)
	}
	x, err := strconv.Atoi(r.URL.Path[8:])
	if err != nil {
		ErrorHandler(w, r, err, 2)
	} else {
		CurThread, err := data.GetThreadByID(x)
		if err != nil {
			ErrorHandler(w, r, err, 1)
		} else {
			Parser.Title = CurThread.Title
			user, _ := data.GetUserByID(CurThread.UserID)
			Parser.Result = user.Username
			Parser.Time = CurThread.Date
			Parser.Data = data.GetAllToThreadByID(CurThread.ThreadID, Parser.ID)
			Parser.Image = "Thread" + strconv.Itoa(CurThread.ThreadID)
			if data.FindImage(Parser.Image) != nil {
				Parser.Image = ""
			}
			tmpl.Execute(w, Parser)
			Reset(&Parser)
		}
	}

}

// Stats - Ajax handler to online statistic
func Stats(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	if Parser.Authorised == false {
		return
	}
	ErrorHandler(w, r, nil, 0)
	if r.Method == "GET" {
		ErrorHandler(w, r, errors.New("page for ajax"), 2)
	} else {
		var x StatsTH
		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		ID, err := strconv.Atoi(x.ThreadID)
		ErrorHandler(w, r, err, 1)
		data.AddNewValueToThread(Parser.ID, ID, x.Value)
		Thread, err := data.GetThreadByID(ID)
		x.Likes = strconv.Itoa(Thread.Likes)
		x.Dislikes = strconv.Itoa(Thread.Dislikes)
		x.Liked = strconv.Itoa(data.CheckUserLikedThread(Parser.ID, ID))
		a, err := json.Marshal(x)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)
	}
}

//Articles - List of articles
func Articles(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	Parser.Data = data.GetAll(Parser.ID)
	tmpl, _ := template.ParseFiles("web/templates/articles.html")
	tmpl.Execute(w, Parser)
}

//DeleteThread - Need Update After
func DeleteThread() {}

//EditThread - New Feature
func EditThread() {}
