package model

type Post struct {
	PostID, UID    int
	Title, Content string
	Tags           []string
}
