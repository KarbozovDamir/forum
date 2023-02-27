package models

//Thread - struct for Thread
type Thread struct {
	Title      string
	UserID     int
	Likes      int
	Dislikes   int
	ThreadID   int
	ToThreadID int
	Date       string
	Content    string
	Category   string
	Username   string
	Image      string
	Liked      int
}

//ThreadStats - struct for statistic Thread
type ThreadStats struct {
	FromUserID int
	ToThreadID int
	Value      int
}
