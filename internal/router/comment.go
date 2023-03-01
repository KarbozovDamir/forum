package router

import (
	"net/http"
	"strconv"

	data "github.com/KarbozovDamir/forum/data"
)

const code = 302 // redirect

// ErrorHandler - Allow us to handle errors in forum
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error, status int) {
	if status == 0 {
		if !Parser.Authorised {
			http.Redirect(w, r, "/login", code)
		}
	}
	if err != nil {
		switch status {
		case 1:
			Parser.Error = []byte("400: Bad Request => No such thread")
		case 2:
			Parser.Error = []byte("404: Not Found => No such page")
		case 3:
			Parser.Error = []byte("400: Bad Request => No such user")
		case 4:
			Parser.Error = []byte("403: Forbidden => You don't have permission")
		case 5:
			Parser.Error = []byte("500: Internal Server Error")
		}
		http.Redirect(w, r, "/error", code) // Found error
	}
}

// Comment - Allow authorised users to comment Threads
func Comment(w http.ResponseWriter, r *http.Request) {
	Authorised(r)
	ErrorHandler(w, r, nil, 0)
	x, err := strconv.Atoi(r.URL.Path[8:])
	ErrorHandler(w, r, err, 1)
	CurThread, err := data.GetThreadByID(x)
	ErrorHandler(w, r, err, 1)
	user, _ := data.GetUserByID(Parser.ID)
	data.CreateCommentToThread(r, Parser.ID, CurThread.ThreadID, user.Username)
}
