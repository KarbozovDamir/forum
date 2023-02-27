package models

// ViewData - struct to template html page
type ViewData struct {
	Result               string
	Authorised           bool
	ID                   int
	Title                string
	Time                 string
	Threads              []Thread // ctreate posts and output posts
	LikedThreads         []Thread
	LikedComments        []Thread
	Error                []byte
	CountOfPosts         int
	CountOfLikedThreads  int
	CountOfLikedComments int
	Me                   bool
	Image                string
}
