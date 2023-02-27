package main

import (
	"fmt"
	"net/http"

	data "github.com/KarbozovDamir/forum/data"
	"github.com/KarbozovDamir/forum/internal/router"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	data.Init()
	// defer data.CloseDB()
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

	fmt.Println("Server is listening on port: 8181")
	server := &http.Server{
		Addr:    "127.0.0.1:8181",
		Handler: mux,
	}
	server.ListenAndServe()
}
