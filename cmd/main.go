package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/KarbozovDamir/forum/internal/data"
	"github.com/KarbozovDamir/forum/web/router"
	_ "github.com/mattn/go-sqlite3"
)

var Ok bool

func main() {
	data.Init()
	defer data.CloseDB()
	mux := http.NewServeMux()
	//handle static assets
	files := http.FileServer(http.Dir("web"))
	mux.Handle("/static/", http.StripPrefix("/web/static/styles.css", files))
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

// func main() {
// 	// create handlers
// 	handlers, err := router.NewMainHandler()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	server := new(Server)
// 	if err := server.Run("4000", handlers.InitRoutes()); err != nil {
// 		log.Printf("ERROR: %s\n", err)

// 	}
// }

// type Server struct {
// 	httpServer *http.Server
// }

// // starting up the server
// func (s *Server) Run(port string, handler http.Handler) error {
// 	s.httpServer = &http.Server{
// 		Addr:         ":" + port,
// 		Handler:      handler,
// 		ReadTimeout:  5 * time.Second, // return cont.Done() if up time limit
// 		WriteTimeout: 5 * time.Second, // return cont.Done() if up time limit
// 	}

// 	log.Printf("Server runs on http://localhost%s\n", s.httpServer.Addr)
// 	err := s.httpServer.ListenAndServe()
// 	return fmt.Errorf("Server.Run: %w", err)
// }
