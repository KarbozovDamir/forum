package model

type PostRate struct {
	PostID, Score int64
}

type CommentRate struct {
	CommentID, Score int64
}

type UserRate struct {
	UserID, Score int64
}

type PostRating struct {
	Positive       bool
	PostID, UserID int64
}

type CommentRating struct {
	Positive          bool
	CommentID, UserID int64
}

type UserRating struct {
	Positive         bool
	CriticID, UserID int64
}
