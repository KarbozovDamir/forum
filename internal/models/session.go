package models

import "time"

type Session struct {
	UserID int
	UUID   string
	Time   time.Time
}
