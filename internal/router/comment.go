package router

import (
	"net/http"
	"strconv"

	data "github.com/KarbozovDamir/forum/data"
)

//Comment - Allow authorised users to comment Threads
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
