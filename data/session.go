package data

import (
	"errors"
	"net/http"
	"time"
)

type Session struct {
	UserID int
	UUID   string
	Time   time.Time
}

// DeleteByUUID - Easy way
func DeleteByUUID(s string) {
	Db.Exec("delete from Session where Uuid = $1", s)
}

// CheckCookie - takes cookie from browser and
func CheckCookie(r *http.Request) (ID int, err2 error) {
	tmpSession := Session{}
	cookie, err := r.Cookie("Cookie")
	if err == nil {
		err = Db.QueryRow("select * from Session where Uuid = $1", cookie.Value).Scan(&tmpSession.UserID, &tmpSession.UUID, &tmpSession.Time)
		if time.Now().After(tmpSession.Time) {
			err = errors.New("Time is not correct")
		}
	}
	err2 = err
	ID = tmpSession.UserID
	return
}
