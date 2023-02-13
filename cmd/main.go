package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KarbozovDamir/forum/internal/data"
	"github.com/KarbozovDamir/forumweb/router"
	_ "github.com/mattn/go-sqlite3"
)

var Ok bool

func main() {
	data.Init()
	defer data.CloseDB()
	mux := http.NewServeMux()
	//handle static assets
	files := http.FileServer(http.Dir("web"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mapper := map[string]func(http.ResponseWriter, *http.Request){
		"/":                    router.DefaultHandler,
		"/ajax":                router.AjaxHandler,
		"/logout":              router.LogOut,
		"/id/":                 router.Profile,
		"/about":               router.About,
		"/rules":               router.Rules,
		"/restore":             router.Restore,
		"/thread/create":       router.CreateThread,
		"/thread/comment/":     router.Comment,
		"/thread/":             router.Post,
		"/error":               router.Error,
		"/stats":               router.Stats,
		"/articles":            router.Articles,
		"/updateProfileImage/": router.UpdateAva,
	}
	for pattern, handler := range mapper {
		mux.HandleFunc(pattern, handler)
	}

	fmt.Println("Server is listening...")
	server := &http.Server{
		Addr:           "0.0.0.0:8181",
		Handler:        mux,
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
